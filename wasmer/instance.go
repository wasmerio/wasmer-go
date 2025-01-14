package wasmer

// #include <wasmer.h>
import "C"
import "runtime"

type Instance struct {
	CPtrBase[*C.wasm_instance_t]
	Exports *Exports

	// without this, imported functions may be freed before execution of an exported function is complete.
	imports *ImportObject
}

// NewInstance instantiates a new Instance.
//
// It takes two arguments, the Module and an ImportObject.
//
// Note:️ Instantiating a module may return TrapError if the module's
// start function traps.
//
//	wasmBytes := []byte(`...`)
//	engine := wasmer.NewEngine()
//	store := wasmer.NewStore(engine)
//	module, err := wasmer.NewModule(store, wasmBytes)
//	importObject := wasmer.NewImportObject()
//	instance, err := wasmer.NewInstance(module, importObject)
func NewInstance(module *Module, imports *ImportObject) (*Instance, error) {
	var traps *C.wasm_trap_t
	externs, err := imports.intoInner(module)

	if err != nil {
		return nil, err
	}

	var instance *C.wasm_instance_t

	err2 := maybeNewErrorFromWasmer(func() bool {
		instance = C.wasm_instance_new(
			module.store.inner(),
			module.inner(),
			externs,
			&traps,
		)

		return traps == nil && instance == nil
	})

	if err2 != nil {
		return nil, err2
	}

	if traps != nil {
		return nil, newErrorFromTrap(traps)
	}

	self := &Instance{
		CPtrBase: mkPtr(instance),
		Exports:  newExports(instance, module),
		imports:  imports,
	}

	runtime.SetFinalizer(self, func(self *Instance) {
		self.Close()
	})

	return self, nil
}

func (self *Instance) inner() *C.wasm_instance_t {
	return self.ptr()
}

// GetRemainingPoints exposes wasm meterings remaining gas or points
func (self *Instance) GetRemainingPoints() uint64 {
	return uint64(C.wasmer_metering_get_remaining_points(self.ptr()))
}

// GetRemainingPoints a bool to determine if the engine has been shutdown from meter exhaustion
func (self *Instance) MeteringPointsExhausted() bool {
	return bool(C.wasmer_metering_points_are_exhausted(self.ptr()))
}

// SetRemainingPoints imposes a new gas limit on the wasm engine
func (self *Instance) SetRemainingPoints(newLimit uint64) {
	C.wasmer_metering_set_remaining_points(self.ptr(), C.uint64_t(newLimit))
}

// ReleaseFn is a function to release resources
// This function is parameterized to make sure that the correct
// argument is passed to the function.
type ReleaseFn[T any] func(T)

func keepAlive[T any](value T) {
	runtime.KeepAlive(value)
}

// GetFunctionSafe performs the same job as instance.Exports.GetFunction
// but it returns a ReleaseFn to release the resources.
//
// The reason for this is that the following code *IS NOT SAFE*:
//
//	instance, err := NewInstance(module, NewImportObject())
//	fn, err := instance.Exports.GetFunction("function_name")
//	// No further usages of instance after this line.
//	fn()
//
// The problem is that when fn() is called, it will issue CGo call to the
// function.  When that happens, the instance may be garbage collected
// because there are no longer any uses of instance.  The returned function
// however *does* depended on the instance; we must ensure that instance
// remains live as long long as the calls to *any* instance functions may happen.
// One way to do that is to add the following line *immediately* after the
// instance intialization:
//
//	defer runtime.KeepAlive(instance)
//
// Another mechanism is to use GetFunctionSafe which returns a ReleaseFn
// which should be called if this function returns non-error value.
// This usage is preferred because it is less error-prone (i.e. it
// returns 3 values that should all be handled appropriately).
func (self *Instance) GetFunctionSafe(name string) (NativeFunction, ReleaseFn[*Instance], error) {
	fn, err := self.Exports.GetFunction(name)
	if err != nil {
		return nil, nil, err
	}
	return fn, keepAlive, nil
}

// Force to close the Instance.
//
// A runtime finalizer is registered on the Instance, but it is
// possible to force the destruction of the Instance by calling Close
// manually.
func (self *Instance) Close() {
	runtime.SetFinalizer(self, nil)
	C.wasm_instance_delete(self.inner())
	self.Exports.Close()
}
