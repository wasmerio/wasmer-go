package wasmer

// #include <wasmer_wasm.h>
import "C"
import (
	"runtime"
	"unsafe"
)

type ValueKind C.wasm_valkind_t

const (
	I32     = ValueKind(C.WASM_I32)
	I64     = ValueKind(C.WASM_I64)
	F32     = ValueKind(C.WASM_F32)
	F64     = ValueKind(C.WASM_F64)
	AnyRef  = ValueKind(C.WASM_ANYREF)
	FuncRef = ValueKind(C.WASM_FUNCREF)
)

func (self ValueKind) String() string {
	switch self {
	case I32:
		return "i32"
	case I64:
		return "i64"
	case F32:
		return "f32"
	case F64:
		return "f64"
	case AnyRef:
		return "anyref"
	case FuncRef:
		return "funcref"
	}
	panic("Unknown value kind") // unreachable
}

func (self ValueKind) IsNumber() bool {
	return bool(C.wasm_valkind_is_num(C.wasm_valkind_t(self)))
}

func (self ValueKind) IsReference() bool {
	return bool(C.wasm_valkind_is_ref(C.wasm_valkind_t(self)))
}

type ValueType struct {
	_inner   *C.wasm_valtype_t
	_ownedBy interface{}
}

func NewValueType(kind ValueKind) *ValueType {
	pointer := C.wasm_valtype_new(C.wasm_valkind_t(kind))

	return newValueType(pointer, nil)
}

func newValueType(pointer *C.wasm_valtype_t, ownedBy interface{}) *ValueType {
	valueType := &ValueType{_inner: pointer, _ownedBy: ownedBy}

	if ownedBy == nil {
		runtime.SetFinalizer(valueType, func(valueType *ValueType) {
			C.wasm_valtype_delete(valueType.inner())
		})
	}

	return valueType
}

func (self *ValueType) inner() *C.wasm_valtype_t {
	return self._inner
}

func (self *ValueType) Kind() ValueKind {
	kind := ValueKind(C.wasm_valtype_kind(self.inner()))

	runtime.KeepAlive(self)

	return kind
}

func toValueTypeVec(valueTypes []*ValueType) C.wasm_valtype_vec_t {
	vec := C.wasm_valtype_vec_t{}
	C.wasm_valtype_vec_new_uninitialized(&vec, C.size_t(len(valueTypes)))

	firstValueTypePointer := unsafe.Pointer(vec.data)

	for nth, valueType := range valueTypes {
		pointer := C.wasm_valtype_new(C.wasm_valtype_kind(valueType.inner()))
		*(**C.wasm_valtype_t)(unsafe.Pointer(uintptr(firstValueTypePointer) + unsafe.Sizeof(pointer)*uintptr(nth))) = pointer
	}

	runtime.KeepAlive(valueTypes)

	return vec
}

func toValueTypeList(valueTypes *C.wasm_valtype_vec_t, ownedBy interface{}) []*ValueType {
	numberOfValueTypes := int(valueTypes.size)
	list := make([]*ValueType, numberOfValueTypes)
	firstValueType := unsafe.Pointer(valueTypes.data)
	sizeOfValueTypePointer := unsafe.Sizeof(firstValueType)

	var currentValueTypePointer *C.wasm_valtype_t

	for nth := 0; nth < numberOfValueTypes; nth++ {
		currentValueTypePointer = *(**C.wasm_valtype_t)(unsafe.Pointer(uintptr(firstValueType) + uintptr(nth)*sizeOfValueTypePointer))
		valueType := newValueType(currentValueTypePointer, ownedBy)
		list[nth] = valueType
	}

	return list
}
