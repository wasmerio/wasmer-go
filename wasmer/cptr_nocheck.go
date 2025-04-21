//go:build !memcheck

package wasmer

import "runtime"

func doMkPtr[T comparable](ptr T) CPtrBase[T] {
	return CPtrBase[T]{_ptr: ptr}
}

// maybeStack is a no-op implementation for storing
// ptr creator stack.
type maybeStack struct{}

// maybeNil is a no-op implementation for indicating
// if initial value of _ptr was nil.
type maybeNil[T comparable] struct{}

// ptr returns the C pointer stored in this base.
func (b *CPtrBase[T]) ptr() T {
	return b._ptr
}

func doSetFinalizer[T comparable](b *CPtrBase[T], free func(v T)) {
	runtime.SetFinalizer(b, func(b *CPtrBase[T]) {
		free(b.ptr())
	})
}

func doClearFinalizer[T comparable](b *CPtrBase[T]) {
	runtime.SetFinalizer(b, nil)
}
