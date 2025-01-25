package wasmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawFunction(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (type $sum_t (func (param i32 i32) (result i32)))
			  (func $sum_f (type $sum_t) (param $x i32) (param $y i32) (result i32)
			    local.get $x
			    local.get $y
			    i32.add)
			  (export "sum" (func $sum_f)))
		`),
	)
	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	sum, err := instance.Exports.GetRawFunction("sum")
	assert.NoError(t, err)
	assert.Equal(t, sum.ParameterArity(), uint(2))
	assert.Equal(t, sum.ResultArity(), uint(1))

	result, err := sum.Call(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, result, int32(3))
}

func TestFunctionNative(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (type $sum_t (func (param i32 i32) (result i32)))
			  (func $sum_f (type $sum_t) (param $x i32) (param $y i32) (result i32)
			    local.get $x
			    local.get $y
			    i32.add)
			  (export "sum" (func $sum_f)))
		`),
	)
	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	sum, err := instance.Exports.GetFunction("sum")
	assert.NoError(t, err)

	result, err := sum(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, result, int32(3))
}

func TestFunctionCallReturnZeroValue(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (type $test_t (func (param i32 i32)))
			  (func $test_f (type $test_t) (param $x i32) (param $y i32))
			  (export "test" (func $test_f)))
		`),
	)
	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	test, err := instance.Exports.GetFunction("test")
	assert.NoError(t, err)

	result, err := test(1, 2)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestFunctionCallReturnMultipleValues(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (type $swap_t (func (param i32 i64) (result i64 i32)))
			  (func $swap_f (type $swap_t) (param $x i32) (param $y i64) (result i64 i32)
			    local.get $y
			    local.get $x)
			  (export "swap" (func $swap_f)))
		`),
	)
	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	swap, err := instance.Exports.GetFunction("swap")
	assert.NoError(t, err)

	results, err := swap(41, 42)
	assert.NoError(t, err)
	assert.Equal(t, results, []interface{}{int64(42), int32(41)})
}

func TestFunctionSum(t *testing.T) {
	instance := testGetInstance(t)

	f, _ := instance.Exports.GetFunction("sum")
	result, err := f(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, result, int32(3))
}

func TestFunctionArity0(t *testing.T) {
	instance := testGetInstance(t)

	f, _ := instance.Exports.GetFunction("arity_0")
	result, err := f()
	assert.NoError(t, err)
	assert.Equal(t, result, int32(42))
}

func TestFunctionI32I32(t *testing.T) {
	instance := testGetInstance(t)

	f, _ := instance.Exports.GetFunction("i32_i32")
	result, err := f(7)
	assert.NoError(t, err)
	assert.Equal(t, result, int32(7))

	result, _ = f(int8(7))
	assert.Equal(t, result, int32(7))

	result, _ = f(uint8(7))
	assert.Equal(t, result, int32(7))

	result, _ = f(int16(7))
	assert.Equal(t, result, int32(7))

	result, _ = f(uint16(7))
	assert.Equal(t, result, int32(7))

	result, _ = f(int32(7))
	assert.Equal(t, result, int32(7))

	result, _ = f(int(7))
	assert.Equal(t, result, int32(7))

	result, _ = f(uint(7))
	assert.Equal(t, result, int32(7))
}

func TestFunctionI64I64(t *testing.T) {
	instance := testGetInstance(t)

	f, _ := instance.Exports.GetFunction("i64_i64")
	result, err := f(7)
	assert.NoError(t, err)
	assert.Equal(t, result, int64(7))

	result, _ = f(int8(7))
	assert.Equal(t, result, int64(7))

	result, _ = f(uint8(7))
	assert.Equal(t, result, int64(7))

	result, _ = f(int16(7))
	assert.Equal(t, result, int64(7))

	result, _ = f(uint16(7))
	assert.Equal(t, result, int64(7))

	result, _ = f(int32(7))
	assert.Equal(t, result, int64(7))

	result, _ = f(int64(7))
	assert.Equal(t, result, int64(7))

	result, _ = f(int(7))
	assert.Equal(t, result, int64(7))

	result, _ = f(uint(7))
	assert.Equal(t, result, int64(7))
}

func TestFunctionF32F32(t *testing.T) {
	instance := testGetInstance(t)

	f, _ := instance.Exports.GetFunction("f32_f32")
	result, err := f(float32(7.42))
	assert.NoError(t, err)
	assert.Equal(t, result, float32(7.42))
}

func TestFunctionF64F64(t *testing.T) {
	instance := testGetInstance(t)

	f, _ := instance.Exports.GetFunction("f64_f64")
	result, err := f(7.42)
	assert.NoError(t, err)
	assert.Equal(t, result, float64(7.42))

	result, _ = f(float64(7.42))
	assert.Equal(t, result, float64(7.42))
}

func TestFunctionI32I64F32F64F64(t *testing.T) {
	instance := testGetInstance(t)

	f, _ := instance.Exports.GetFunction("i32_i64_f32_f64_f64")
	result, err := f(1, 2, float32(3.4), 5.6)
	assert.NoError(t, err)
	assert.Equal(t, float64(int(result.(float64)*10000000))/10000000, 1+2+3.4+5.6)
}

func TestFunctionBoolCastedtoI32(t *testing.T) {
	instance := testGetInstance(t)

	f, _ := instance.Exports.GetFunction("bool_casted_to_i32")
	result, err := f()
	assert.NoError(t, err)
	assert.Equal(t, result, int32(1))
}

func TestHostFunction(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (import "math" "sum" (func $sum (param i32 i32) (result i32)))
			  (func (export "add_one") (param $x i32) (result i32)
			    local.get $x
			    i32.const 1
			    call $sum))
		`),
	)
	assert.NoError(t, err)

	function := NewFunction(
		store,
		NewFunctionType(NewValueTypes(I32, I32), NewValueTypes(I32)),
		func(args []Value) ([]Value, error) {
			x := args[0].I32()
			y := args[1].I32()

			return []Value{NewI32(x + y)}, nil
		},
	)

	importObject := NewImportObject()
	importObject.Register(
		"math",
		map[string]IntoExtern{
			"sum": function,
		},
	)

	instance, err := NewInstance(module, importObject)
	assert.NoError(t, err)

	addOne, err := instance.Exports.GetFunction("add_one")
	assert.NoError(t, err)

	result, err := addOne(41)
	assert.NoError(t, err)
	assert.Equal(t, result, int32(42))
}

func TestHostFunction_WithI64(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (import "math" "sum" (func $sum (param i64 i64) (result i64)))
			  (func (export "add_one") (param $x i64) (result i64)
			    local.get $x
			    i64.const 1
			    call $sum))
		`),
	)
	assert.NoError(t, err)

	function := NewFunction(
		store,
		NewFunctionType(NewValueTypes(I64, I64), NewValueTypes(I64)),
		func(args []Value) ([]Value, error) {
			x := args[0].I64()
			y := args[1].I64()

			return []Value{NewI64(x + y)}, nil
		},
	)

	importObject := NewImportObject()
	importObject.Register(
		"math",
		map[string]IntoExtern{
			"sum": function,
		},
	)

	instance, err := NewInstance(module, importObject)
	assert.NoError(t, err)

	addOne, err := instance.Exports.GetFunction("add_one")
	assert.NoError(t, err)

	result, err := addOne(41)
	assert.NoError(t, err)
	assert.Equal(t, result, int64(42))
}

func TestHostFunctionWithEnv(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (import "math" "sum" (func $sum (param i32 i32) (result i32)))
			  (func (export "add_one") (param $x i32) (result i32)
			    local.get $x
			    i32.const 1
			    call $sum))
		`),
	)
	assert.NoError(t, err)

	type MyEnvironment struct {
		instance  *Instance
		theAnswer int32
	}

	environment := &MyEnvironment{
		instance:  nil,
		theAnswer: 42,
	}

	function := NewFunctionWithEnvironment(
		store,
		NewFunctionType(NewValueTypes(I32, I32), NewValueTypes(I32)),
		environment,
		func(environment interface{}, args []Value) ([]Value, error) {
			env := environment.(*MyEnvironment)
			assert.NotNil(t, env.instance)

			x := args[0].I32()
			y := args[1].I32()

			return []Value{NewI32(env.theAnswer + x + y)}, nil
		},
	)

	importObject := NewImportObject()
	importObject.Register(
		"math",
		map[string]IntoExtern{
			"sum": function,
		},
	)

	instance, err := NewInstance(module, importObject)
	assert.NoError(t, err)

	environment.instance = instance

	addOne, err := instance.Exports.GetFunction("add_one")
	assert.NoError(t, err)

	result, err := addOne(7)
	assert.NoError(t, err)
	assert.Equal(t, result, int32(50))
}

func TestHostFunctionStore(t *testing.T) {
	f := &hostFunction{
		store: NewStore(NewEngine()),
		function: func(args []Value) ([]Value, error) {
			return []Value{}, nil
		},
	}

	store := hostFunctions{
		functions: make(map[uint]*hostFunction),
	}
	_, err := store.load(0)
	assert.Error(t, err, "Host function `0` does not exist")

	indexA := store.store(f)
	indexB := store.store(f)
	indexC := store.store(f)
	assert.Equal(t, indexA, uint(0))
	assert.Equal(t, indexB, uint(1))
	assert.Equal(t, indexC, uint(2))

	store.remove(indexB)
	_, err = store.load(indexB)
	assert.Error(t, err, "Host function `1` does not exist")

	indexD := store.store(f)
	assert.Equal(t, indexD, indexB)
}

/*
func TestFunctionTrap(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (func (export "trap") (result i32) (unreachable) (i32.const 1)))
		`),
	)
	assert.NoError(t, err)

	instance, err := NewInstance(module, NewImportObject())
	assert.NoError(t, err)

	trap, err := instance.Exports.GetFunction("trap")
	assert.NoError(t, err)

	_, err = trap()

	assert.Error(t, err)
	assert.IsType(t, &TrapError{}, err)

	trapError := err.(*TrapError)
	trapOrigin := trapError.Origin()
	assert.NotNil(t, trapOrigin)
	assert.Equal(t, uint32(0), trapOrigin.FunctionIndex())
	assert.Equal(t, uint(1), trapOrigin.FunctionOffset())
	assert.Equal(t, uint(34), trapOrigin.ModuleOffset())

	trapTrace := trapError.Trace()
	assert.NotNil(t, trapTrace)
	assert.Len(t, trapTrace, 1)
	assert.Equal(t, trapOrigin.FunctionIndex(), trapTrace[0].FunctionIndex())
	assert.Equal(t, trapOrigin.FunctionOffset(), trapTrace[0].FunctionOffset())
	assert.Equal(t, trapOrigin.ModuleOffset(), trapTrace[0].ModuleOffset())
}
*/

type myError struct {
	message string
}

func (self *myError) Error() string {
	return self.message
}

func TestHostFunctionTrap(t *testing.T) {
	engine := NewEngine()
	store := NewStore(engine)
	module, err := NewModule(
		store,
		[]byte(`
			(module
			  (import "math" "sum" (func $sum (param i32 i32) (result i32)))
			  (func (export "add_one") (param $x i32) (result i32)
			    local.get $x
			    i32.const 1
			    call $sum))
		`),
	)
	assert.NoError(t, err)

	function := NewFunction(
		store,
		NewFunctionType(NewValueTypes(I32, I32), NewValueTypes(I32)),
		func(args []Value) ([]Value, error) {
			// Go Trap, go!
			return nil, &myError{message: "oops"}
		},
	)

	importObject := NewImportObject()
	importObject.Register(
		"math",
		map[string]IntoExtern{
			"sum": function,
		},
	)

	instance, err := NewInstance(module, importObject)
	assert.NoError(t, err)

	addOne, err := instance.Exports.GetFunction("add_one")
	assert.NoError(t, err)

	_, err = addOne(41)
	assert.IsType(t, err, &TrapError{})
	assert.Error(t, err, "oops")
}
