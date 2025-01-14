//go:build memcheck

package wasmer

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

func doMkPtr[T comparable](ptr T) CPtrBase[T] {
	var zero T
	b := CPtrBase[T]{_ptr: ptr}
	b._stack = stack(3)
	b.maybeNil._nil = ptr == zero
	b.maybeNil._zero = zero
	return b
}

type maybeNil[T comparable] struct {
	_nil  bool
	_zero T
}

// maybeStack stores the stack of the goroutine.
type maybeStack struct {
	_stack []byte
}

// ptr returns the C pointer stored in this base.
func (b *CPtrBase[T]) ptr() T {
	if !b.maybeNil._nil && b._ptr == b.maybeNil._zero {
		panic(fmt.Errorf("attempt to use a released object; ptr was allocated by\n%s", b._stack))
	}
	runtime.GC()
	return b._ptr
}

func doSetFinalizer[T comparable](b *CPtrBase[T], free func(v T)) {
	finalizers.Store(uintptr(unsafe.Pointer(b)), stack(3))

	runtime.SetFinalizer(b, func(b *CPtrBase[T]) {
		ptr := b._ptr
		b._ptr = b.maybeNil._zero
		b.maybeNil._nil = false
		free(ptr)
		finalizers.Delete(uintptr(unsafe.Pointer(b)))
	})
}

func doClearFinalizer[T comparable](b *CPtrBase[T]) {
	runtime.SetFinalizer(b, nil)
	finalizers.Delete(uintptr(unsafe.Pointer(b)))
}

var finalizers sync.Map // uintptr -> []byte (the stack of the caller which created the finalizer)

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGABRT, syscall.SIGSEGV, syscall.SIGBUS, syscall.SIGINT)

	go func() {
		s := <-c
		fmt.Fprintf(os.Stderr, "Signal %s received\n", s)

		allStacks := make([]byte, 1<<18)
		sl := runtime.Stack(allStacks, true)
		if sl == len(allStacks) {
			// Try again, once, with a bigger buffer
			allStacks = make([]byte, 1<<20)
			sl = runtime.Stack(allStacks, true)
		}
		allStacks = allStacks[:sl]

		fmt.Fprintf(os.Stderr, "\n%s\n\n", allStacks)
		fmt.Fprintf(os.Stderr, "Pending Finalizers:\n")

		n := 0
		finalizers.Range(func(k, v interface{}) bool {
			n++
			stack := v.([]byte)
			fmt.Fprintf(os.Stderr, "\n--finilizer %d: %s\n\n", n, stack)
			return true
		})
		os.Exit(1)
	}()
}

// stackCache seems to be a bit of an overkill for debug build, but
// some of the tests (e.g. TestSumLoop) allocate many pointers; and
// the time and memory used for that adds up.
const stackCacheMaxSize = 1 << 9

var stackCache sync.Map // uintptr -> []byte
var stackCacheSize atomic.Int32

func stack(skip int) []byte {
	var buf bytes.Buffer
	const maxDepth = 32
	pc := make([]uintptr, maxDepth)
	// Skip the requested number of frames + the frames for this function
	n := runtime.Callers(skip+1, pc)
	if n == 0 {
		return nil
	}
	pc = pc[:n]

	if v, ok := stackCache.Load(pc[0]); ok {
		return v.([]byte)
	}

	// Convert program counters to frames
	frames := runtime.CallersFrames(pc)
	for {
		frame, more := frames.Next()
		// Format the frame as "file:line function"
		fmt.Fprintf(&buf, "%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}

	s := buf.Bytes()
	stackCache.Store(pc[0], s)
	if stackCacheSize.Add(1) > stackCacheMaxSize {
		stackCache.Range(func(k, v interface{}) bool {
			stackCache.Delete(k)
			return stackCacheSize.Add(-1) > (stackCacheMaxSize >> 1)
		})
	}
	return s
}
