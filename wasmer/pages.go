package wasmer

// #include <wasmer_wasm.h>
import "C"

type Pages C.wasm_memory_pages_t

const WasmPageSize = uint(0x10000)
const WasmMaxPages = uint(0x10000)
const WasmMinPages = uint(0x100)

func (self *Pages) ToUint32() uint32 {
	return uint32(C.wasm_memory_pages_t(*self))
}

func (self *Pages) ToBytes() uint {
	return uint(self.ToUint32()) * WasmPageSize
}
