package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type TableType struct {
	_inner   *C.wasm_tabletype_t
	_ownedBy interface{}
}

func newTableType(pointer *C.wasm_tabletype_t, ownedBy interface{}) *TableType {
	tableType := &TableType{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(tableType, func(tableType *TableType) {
			C.wasm_tabletype_delete(tableType.inner())
		})
	}

	return tableType
}

func NewTableType(valueType *ValueType, limits *Limits) *TableType {
	pointer := C.wasm_tabletype_new(valueType.inner(), limits.inner())

	return newTableType(pointer, nil)
}

func (self *TableType) inner() *C.wasm_tabletype_t {
	return self._inner
}

func (self *TableType) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *TableType) ValueType() *ValueType {
	pointer := C.wasm_tabletype_element(self.inner())

	runtime.KeepAlive(self)

	return newValueType(pointer, self.ownedBy())
}

func (self *TableType) Limits() *Limits {
	limits := newLimits(C.wasm_tabletype_limits(self.inner()), self.ownedBy())

	runtime.KeepAlive(self)

	return limits
}

func (self *TableType) IntoExternType() *ExternType {
	pointer := C.wasm_tabletype_as_externtype_const(self.inner())

	return newExternType(pointer, self.ownedBy())
}
