package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type FunctionType struct {
	_inner   *C.wasm_functype_t
	_ownedBy interface{}
}

func NewFunctionType(params []*ValueType, results []*ValueType) *FunctionType {
	paramsAsValueTypeVec := toValueTypeVec(params)
	resultsAsValueTypeVec := toValueTypeVec(results)

	pointer := C.wasm_functype_new(&paramsAsValueTypeVec, &resultsAsValueTypeVec)

	return newFunctionType(pointer, nil)
}

func newFunctionType(pointer *C.wasm_functype_t, ownedBy interface{}) *FunctionType {
	functionType := &FunctionType{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(functionType, func(functionType *FunctionType) {
			C.wasm_functype_delete(functionType.inner())
		})
	}

	return functionType
}

func (self *FunctionType) inner() *C.wasm_functype_t {
	return self._inner
}

func (self *FunctionType) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *FunctionType) Params() []*ValueType {
	return toValueTypeList(C.wasm_functype_params(self.inner()), self.ownedBy())
}

func (self *FunctionType) Results() []*ValueType {
	return toValueTypeList(C.wasm_functype_results(self.inner()), self.ownedBy())
}

func (self *FunctionType) IntoExternType() *ExternType {
	pointer := C.wasm_functype_as_externtype_const(self.inner())

	return newExternType(pointer, self.ownedBy())
}
