package wasmer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueKindToString(t *testing.T) {
	assert.Equal(t, I32.String(), "i32")
	assert.Equal(t, I64.String(), "i64")
	assert.Equal(t, F32.String(), "f32")
	assert.Equal(t, F64.String(), "f64")
	assert.Equal(t, AnyRef.String(), "anyref")
	assert.Equal(t, FuncRef.String(), "funcref")
}

func TestValueKindIsNumber(t *testing.T) {
	assert.Equal(t, I32.IsNumber(), true)
	assert.Equal(t, I64.IsNumber(), true)
	assert.Equal(t, F32.IsNumber(), true)
	assert.Equal(t, F64.IsNumber(), true)
	assert.Equal(t, AnyRef.IsNumber(), false)
	assert.Equal(t, FuncRef.IsNumber(), false)
}

func TestValueKindIsReference(t *testing.T) {
	assert.Equal(t, I32.IsReference(), false)
	assert.Equal(t, I64.IsReference(), false)
	assert.Equal(t, F32.IsReference(), false)
	assert.Equal(t, F64.IsReference(), false)
	assert.Equal(t, AnyRef.IsReference(), true)
	assert.Equal(t, FuncRef.IsReference(), true)
}

func TestValueTypeKind(t *testing.T) {
	assert.Equal(t, NewValueType(I32).Kind(), I32)
	assert.Equal(t, NewValueType(I64).Kind(), I64)
	assert.Equal(t, NewValueType(F32).Kind(), F32)
	assert.Equal(t, NewValueType(F64).Kind(), F64)
	assert.Equal(t, NewValueType(AnyRef).Kind(), AnyRef)
	assert.Equal(t, NewValueType(FuncRef).Kind(), FuncRef)
}

func TestValueTypeToVecToList(t *testing.T) {
	valueTypeList := []*ValueType{
		NewValueType(I32),
		NewValueType(I64),
		NewValueType(F32),
	}
	valueTypeVec := toValueTypeVec(valueTypeList)

	assert.Equal(t, int(valueTypeVec.size), 3)
	assert.Equal(t, toValueTypeList(&valueTypeVec, nil), valueTypeList)

	//C.wasm_valtype_delete(&valueTypeVec)
}
