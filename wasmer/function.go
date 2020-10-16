package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type Function struct {
	_inner     *C.wasm_func_t
	_ownedBy   interface{}
	lazyNative func(...interface{}) (interface{}, error)
}

func newFunction(pointer *C.wasm_func_t, ownedBy interface{}) *Function {
	function := &Function{_inner: pointer, _ownedBy: ownedBy, lazyNative: nil}

	if ownedBy == nil {
		runtime.SetFinalizer(function, func(function *Function) {
			C.wasm_func_delete(function.inner())
		})
	}

	return function
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

		var arguments C.wasm_val_vec_t
		defer C.wasm_val_vec_delete(&arguments)

		if numberOfReceivedParameters > 0 {
			C.wasm_val_vec_new(&arguments, C.size_t(numberOfReceivedParameters), (*C.wasm_val_t)(unsafe.Pointer(&allArguments[0])))
		} else {
			C.wasm_val_vec_new_empty(&arguments)
		}

		trap := C.wasm_func_call(self.inner(), &arguments, &results)

		runtime.KeepAlive(arguments)
		runtime.KeepAlive(results)

		if trap != nil {
			panic("trapped!")
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
