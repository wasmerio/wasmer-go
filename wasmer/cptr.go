package wasmer

// CPtrBase is a based struct for a C pointer.
// It is intended to be embedded into any structure
// that stores a C pointer.
// While this based is parameterized on type any, using any
// type other than *C.xxx pointer should be considered an undefined behavior.
type CPtrBase[T comparable] struct {
	_ptr        T
	maybeStack  // stack of the creating goroutine (if enabled via memcheck)
	maybeNil[T] // indicates if initial value of _ptr was nil (if enabled via memcheck)
}

// release returns the C pointer stored in this base and clears the finalizer.
func (b *CPtrBase[T]) release() T {
	var zero T
	v := b.ptr()
	b._ptr = zero
	b.ClearFinalizer()
	return v
}

func (b *CPtrBase[T]) SetFinalizer(free func(v T)) {
	doSetFinalizer(b, free)
}

func (b *CPtrBase[T]) ClearFinalizer() {
	doClearFinalizer(b)
}

func mkPtr[T comparable](ptr T) CPtrBase[T] {
	return doMkPtr(ptr)
}
