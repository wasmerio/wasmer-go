package wasmer

// #include <wasmer_wasm.h>
//
// extern wasm_trap_t* function_trampoline(
//   void *environment,
//   /* const */ wasm_val_vec_t* arguments,
//   wasm_val_vec_t* results
// );
//
// typedef void (*wasm_func_callback_env_finalizer_t)(void* environment);
//
// extern void function_environment_finalizer(void *environment);
import "C"
import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

type Function struct {
	_inner      *C.wasm_func_t
	_ownedBy    interface{}
	environment *FunctionEnvironment
	lazyNative  func(...interface{}) (interface{}, error)
}

func newFunction(pointer *C.wasm_func_t, environment *FunctionEnvironment, ownedBy interface{}) *Function {
	function := &Function{
		_inner:      pointer,
		_ownedBy:    ownedBy,
		environment: environment,
		lazyNative:  nil,
	}

	if ownedBy == nil {
		runtime.SetFinalizer(function, func(function *Function) {
			if function.environment != nil {
				hostFunctionStore.remove(function.environment.hostFunctionStoreIndex)
			}

			C.wasm_func_delete(function.inner())
		})
	}

	return function
}

func NewFunction(store *Store, ty *FunctionType, function func([]Value) ([]Value, error)) *Function {
	hostFunction := &hostFunction{
		store:    store,
		function: function,
	}
	environment := &FunctionEnvironment{
		hostFunctionStoreIndex: hostFunctionStore.store(hostFunction),
	}
	pointer := C.wasm_func_new_with_env(
		store.inner(),
		ty.inner(),
		(C.wasm_func_callback_t)(C.function_trampoline),
		unsafe.Pointer(environment),
		(C.wasm_func_callback_env_finalizer_t)(C.function_environment_finalizer),
	)

	runtime.KeepAlive(environment)

	return newFunction(pointer, environment, nil)
}

//export function_trampoline
func function_trampoline(env unsafe.Pointer, args *C.wasm_val_vec_t, res *C.wasm_val_vec_t) *C.wasm_trap_t {
	environment := (*FunctionEnvironment)(env)
	hostFunction, err := hostFunctionStore.load(environment.hostFunctionStoreIndex)

	if err != nil {
		panic(err)
	}

	arguments := toValueList(args)
	results, err := (hostFunction.function)(arguments)

	if err != nil {
		trap := NewTrap(hostFunction.store, err.Error())

		runtime.KeepAlive(trap)

		return trap.inner()
	}

	toValueVec(results, res)

	return nil
}

func (self *Function) inner() *C.wasm_func_t {
	return self._inner
}

func (self *Function) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *Function) IntoExtern() *Extern {
	pointer := C.wasm_func_as_extern(self.inner())

	return newExtern(pointer, self.ownedBy())
}

func (self *Function) Type() *FunctionType {
	ty := C.wasm_func_type(self.inner())

	runtime.KeepAlive(self)

	return newFunctionType(ty, self.ownedBy())
}

func (self *Function) ParameterArity() uint {
	return uint(C.wasm_func_param_arity(self.inner()))
}

func (self *Function) ResultArity() uint {
	return uint(C.wasm_func_result_arity(self.inner()))
}

func (self *Function) Call(parameters ...interface{}) (interface{}, error) {
	return self.Native()(parameters...)
}

func (self *Function) Native() func(...interface{}) (interface{}, error) {
	if self.lazyNative != nil {
		return self.lazyNative
	}

	self.lazyNative = func(receivedParameters ...interface{}) (interface{}, error) {
		ty := self.Type()
		expectedParameters := ty.Params()
		numberOfReceivedParameters := len(receivedParameters)
		numberOfExpectedParameters := len(expectedParameters)
		diff := numberOfExpectedParameters - numberOfReceivedParameters

		if diff > 0 {
			return nil, newErrorWith(fmt.Sprintf("Missing %d argument(s) when calling the function; Expected %d argument(s), received %d", diff, numberOfExpectedParameters, numberOfReceivedParameters))
		} else if diff < 0 {
			return nil, newErrorWith(fmt.Sprintf("Given %d extra argument(s) when calling the function; Expected %d argument(s), received %d", -diff, numberOfExpectedParameters, numberOfReceivedParameters))
		}

		allArguments := make([]C.wasm_val_t, numberOfReceivedParameters)

		for nth, receivedParameter := range receivedParameters {
			argument, err := fromGoValue(receivedParameter, expectedParameters[nth].Kind())

			if err != nil {
				return nil, newErrorWith(fmt.Sprintf("Argument %d of the function must of of type `%s`, cannot cast value to this type.", nth+1, err))
			}

			allArguments[nth] = argument
		}

		results := C.wasm_val_vec_t{}
		C.wasm_val_vec_new_uninitialized(&results, C.size_t(len(ty.Results())))
		defer C.wasm_val_vec_delete(&results)

		arguments := C.wasm_val_vec_t{}
		defer C.wasm_val_vec_delete(&arguments)

		if numberOfReceivedParameters > 0 {
			C.wasm_val_vec_new(&arguments, C.size_t(numberOfReceivedParameters), (*C.wasm_val_t)(unsafe.Pointer(&allArguments[0])))
		}

		trap := C.wasm_func_call(self.inner(), &arguments, &results)

		runtime.KeepAlive(arguments)
		runtime.KeepAlive(results)

		if trap != nil {
			return nil, newErrorFromTrap(trap)
		}

		switch results.size {
		case 0:
			return nil, nil
		case 1:
			return toGoValue(results.data), nil
		default:
			numberOfValues := int(results.size)
			allResults := make([]interface{}, numberOfValues)
			firstValue := unsafe.Pointer(results.data)
			sizeOfValuePointer := unsafe.Sizeof(C.wasm_val_t{})

			var currentValuePointer *C.wasm_val_t

			for nth := 0; nth < numberOfValues; nth++ {
				currentValuePointer = (*C.wasm_val_t)(unsafe.Pointer(uintptr(firstValue) + uintptr(nth)*sizeOfValuePointer))
				value := toGoValue(currentValuePointer)
				allResults[nth] = value
			}

			return allResults, nil
		}
	}

	return self.lazyNative
}

type FunctionEnvironment struct {
	store                  *Store
	hostFunctionStoreIndex uint
}

//export function_environment_finalizer
func function_environment_finalizer() {}

type hostFunction struct {
	function func([]Value) ([]Value, error)
	store    *Store
}

type hostFunctions struct {
	sync.RWMutex
	functions map[uint]*hostFunction
}

func (self *hostFunctions) load(index uint) (*hostFunction, error) {
	hostFunction, exists := self.functions[index]

	if exists && hostFunction != nil {
		return hostFunction, nil
	}

	return nil, newErrorWith(fmt.Sprintf("Host function `%d` does not exist", index))
}

func (self *hostFunctions) store(hostFunction *hostFunction) uint {
	self.Lock()
	// By default, the index is the size of the store.
	index := uint(len(self.functions))

	for nth, hostFunc := range self.functions {
		// Find the first empty slot in the store.
		if hostFunc == nil {
			// Use that empty slot for the index.
			index = nth
			break
		}
	}

	self.functions[index] = hostFunction
	self.Unlock()

	return index
}

func (self *hostFunctions) remove(index uint) {
	self.Lock()
	self.functions[index] = nil
	self.Unlock()
}

var hostFunctionStore = hostFunctions{
	functions: make(map[uint]*hostFunction),
}
