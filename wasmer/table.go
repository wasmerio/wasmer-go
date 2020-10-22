package wasmer

// #include <wasmer_wasm.h>
import "C"
import "runtime"

type TableSize C.wasm_table_size_t

func (self *TableSize) ToUint32() uint32 {
	return uint32(C.wasm_table_size_t(*self))
}

type Table struct {
	_inner   *C.wasm_table_t
	_ownedBy interface{}
}

func newTable(pointer *C.wasm_table_t, ownedBy interface{}) *Table {
	table := &Table{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(table, func(table *Table) {
			C.wasm_table_delete(table.inner())
		})
	}

	return table
}

func (self *Table) inner() *C.wasm_table_t {
	return self._inner
}

func (self *Table) ownedBy() interface{} {
	if self._ownedBy == nil {
		return self
	}

	return self._ownedBy
}

func (self *Table) Size() TableSize {
	return TableSize(C.wasm_table_size(self.inner()))
}

func (self *Table) IntoExtern() *Extern {
	pointer := C.wasm_table_as_extern(self.inner())

	return newExtern(pointer, self.ownedBy())
}
