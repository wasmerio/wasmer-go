package wasmer

import "fmt"
import "unsafe"

// ExportedFunctionError represents any kind of errors related to a
// WebAssembly exported function. It is returned by `Instance`
// functions only.
type ExportedFunctionError struct {
	functionName string
	message      string
}

// NewExportedFunctionError constructs a new `ExportedFunctionError`,
// where `functionName` is the name of the exported function, and
// `message` is the error message. If the error message contains `%s`,
// then this parameter will be replaced by `functionName`.
func NewExportedFunctionError(functionName string, message string) *ExportedFunctionError {
	return &ExportedFunctionError{functionName, message}
}

// ExportedFunctionError is an actual error. The `Error` function
// returns the error message.
func (error *ExportedFunctionError) Error() string {
	return fmt.Sprintf(error.message, error.functionName)
}

type ExportFunctionT = func(...interface{}) (Value, error)

type ExportedFunction struct {
	name            string
	inputsArity     cUint32T
	outputsArity    cUint32T
	inputsSignature []cWasmerValueTag
}

type Exports struct {
	functions map[string]ExportFunctionT

	memory *Memory
}

func instanceExports(instance *cWasmerInstanceT) (*Exports, error) {
	var err error
	var memory *Memory
	var functions = make(map[string]ExportFunctionT)

	var wasmExports *cWasmerExportsT

	cWasmerInstanceExports(instance, &wasmExports)
	defer cWasmerExportsDestroy(wasmExports)

	var numberOfExports = int(cWasmerExportsLen(wasmExports))

	for nth := 0; nth < numberOfExports; nth++ {
		var wasmExport = cWasmerExportsGet(wasmExports, cInt(nth))
		var wasmExportKind = cWasmerExportKind(wasmExport)

		switch wasmExportKind {
		case cWasmMemory:
			memory, err = handleMemoryExport(wasmExport)
			if err != nil {
				return nil, err
			}
		case cWasmFunction:
			f, err := handleFunctionExport(wasmExport)

			if err != nil {
				return nil, err
			}

			functions[f.name] = wrapInstanceExportedFunction(instance, f)
		}
	}

	return &Exports{functions, memory}, nil
}

func wrapInstanceExportedFunction(instance *cWasmerInstanceT, f *ExportedFunction) ExportFunctionT {
	return func(arguments ...interface{}) (Value, error) {
		wasmInputs, err := prepareInstanceCall(f.name, int(f.inputsArity), f.inputsSignature, arguments...)

		if err != nil {
			return I32(0), err
		}

		var wasmOutputs = make([]cWasmerValueT, f.outputsArity)
		callResult := doInstanceCall(instance, f.inputsArity, f.outputsArity, f.name, wasmInputs, wasmOutputs)
		if callResult != cWasmerOk {
			return I32(0), NewExportedFunctionError(f.name, "Failed to call the `%s` exported function.")
		}

		if f.outputsArity > 0 {
			return castWasmerOutput(wasmOutputs), nil
		} else {
			return void(), nil
		}
	}
}

func doInstanceCall(instance *cWasmerInstanceT, wasmFunctionInputsArity cUint32T, wasmFunctionOutputsArity cUint32T, exportedFunctionName string, wasmInputs []cWasmerValueT, wasmOutputs []cWasmerValueT) cWasmerResultT {
	var wasmInputsCPointer *cWasmerValueT

	if wasmFunctionInputsArity > 0 {
		wasmInputsCPointer = (*cWasmerValueT)(unsafe.Pointer(&wasmInputs[0]))
	} else {
		wasmInputsCPointer = (*cWasmerValueT)(unsafe.Pointer(&wasmInputs))
	}

	var wasmFunctionName = cCString(exportedFunctionName)
	defer cFree(unsafe.Pointer(wasmFunctionName))

	var wasmOutputsCPointer *cWasmerValueT

	if wasmFunctionOutputsArity > 0 {
		wasmOutputsCPointer = (*cWasmerValueT)(unsafe.Pointer(&wasmOutputs[0]))
	} else {
		wasmOutputsCPointer = (*cWasmerValueT)(unsafe.Pointer(&wasmOutputs))
	}

	return cWasmerInstanceCall(
		instance,
		wasmFunctionName,
		wasmInputsCPointer,
		wasmFunctionInputsArity,
		wasmOutputsCPointer,
		wasmFunctionOutputsArity,
	)
}

func prepareInstanceCall(exportedFunctionName string, numberOfExpectedArguments int, wasmFunctionInputSignature []cWasmerValueTag, arguments ...interface{}) ([]cWasmerValueT, error) {
	var numberOfGivenArguments = len(arguments)
	var diff = numberOfExpectedArguments - numberOfGivenArguments

	if diff > 0 {
		return nil, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Missing %d argument(s) when calling the `%%s` exported function; Expect %d argument(s), given %d.", diff, numberOfExpectedArguments, numberOfGivenArguments))
	} else if diff < 0 {
		return nil, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Given %d extra argument(s) when calling the `%%s` exported function; Expect %d argument(s), given %d.", -diff, numberOfExpectedArguments, numberOfGivenArguments))
	}

	var wasmInputs = make([]cWasmerValueT, numberOfGivenArguments)

	for nth, value := range arguments {
		var wasmInputType = wasmFunctionInputSignature[nth]

		switch wasmInputType {
		case cWasmI32:
			wasmInputs[nth].tag = cWasmI32
			wasmerValue, err := castArgToWasmerI32(exportedFunctionName, nth, value)

			if err != nil {
				return nil, err
			}

			var pointer = (*int32)(unsafe.Pointer(&wasmInputs[nth].value))
			*pointer = wasmerValue

		case cWasmI64:
			wasmInputs[nth].tag = cWasmI64
			wasmerValue, err := castArgToWasmerI64(exportedFunctionName, nth, value)

			if err != nil {
				return nil, err
			}

			var pointer = (*int64)(unsafe.Pointer(&wasmInputs[nth].value))
			*pointer = wasmerValue
		case cWasmF32:
			wasmInputs[nth].tag = cWasmF32
			wasmerValue, err := castArgToWasmerF32(exportedFunctionName, nth, value)

			if err != nil {
				return nil, err
			}

			var pointer = (*float32)(unsafe.Pointer(&wasmInputs[nth].value))
			*pointer = wasmerValue

		case cWasmF64:
			wasmInputs[nth].tag = cWasmF64
			wasmerValue, err := castArgToWasmerF64(exportedFunctionName, nth, value)

			if err != nil {
				return nil, err
			}

			var pointer = (*float64)(unsafe.Pointer(&wasmInputs[nth].value))
			*pointer = wasmerValue
		default:
			panic(fmt.Sprintf("Invalid arguments type when calling the `%s` exported function.", exportedFunctionName))
		}
	}

	return wasmInputs, nil
}

func castArgToWasmerI32(exportedFunctionName string, nth int, value interface{}) (int32, error) {
	switch value.(type) {
	case int8:
		return int32(value.(int8)), nil
	case uint8:
		return int32(value.(uint8)), nil
	case int16:
		return int32(value.(int16)), nil
	case uint16:
		return int32(value.(uint16)), nil
	case int32:
		return int32(value.(int32)), nil
	case int:
		return int32(value.(int)), nil
	case uint:
		return int32(value.(uint)), nil
	case Value:
		var value = value.(Value)

		if value.GetType() != TypeI32 {
			return 0, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i32`, cannot cast given value to this type.", nth+1))
		}

		return value.ToI32(), nil
	default:
		return 0, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i32`, cannot cast given value to this type.", nth+1))
	}
}

func castArgToWasmerI64(exportedFunctionName string, nth int, value interface{}) (int64, error) {
	switch value.(type) {
	case int8:
		return int64(value.(int8)), nil
	case uint8:
		return int64(value.(uint8)), nil
	case int16:
		return int64(value.(int16)), nil
	case uint16:
		return int64(value.(uint16)), nil
	case int32:
		return int64(value.(int32)), nil
	case uint32:
		return int64(value.(uint32)), nil
	case int64:
		return int64(value.(int64)), nil
	case int:
		return int64(value.(int)), nil
	case uint:
		return int64(value.(uint)), nil
	case Value:
		var value = value.(Value)

		if value.GetType() != TypeI64 {
			return 0, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i64`, cannot cast given value to this type.", nth+1))
		}

		return value.ToI64(), nil
	default:
		return 0, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `i64`, cannot cast given value to this type.", nth+1))
	}
}

func castArgToWasmerF32(exportedFunctionName string, nth int, value interface{}) (float32, error) {
	switch value.(type) {
	case float32:
		return value.(float32), nil
	case Value:
		var value = value.(Value)

		if value.GetType() != TypeF32 {
			return 0, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f32`, cannot cast given value to this type.", nth+1))
		}

		return value.ToF32(), nil
	default:
		return 0, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f32`, cannot cast given value to this type.", nth+1))
	}
}

func castArgToWasmerF64(exportedFunctionName string, nth int, value interface{}) (float64, error) {
	switch value.(type) {
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return value.(float64), nil
	case Value:
		var value = value.(Value)

		if value.GetType() != TypeF64 {
			return 0, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f64`, cannot cast given value to this type.", nth+1))
		}

		return value.ToF64(), nil
	default:
		return 0, NewExportedFunctionError(exportedFunctionName, fmt.Sprintf("Argument #%d of the `%%s` exported function must be of type `f64`, cannot cast given value to this type.", nth+1))
	}
}

func castWasmerOutput(wasmOutputs []cWasmerValueT) Value {
	var result = wasmOutputs[0]

	switch result.tag {
	case cWasmI32:
		pointer := (*int32)(unsafe.Pointer(&result.value))

		return I32(*pointer)
	case cWasmI64:
		pointer := (*int64)(unsafe.Pointer(&result.value))

		return I64(*pointer)
	case cWasmF32:
		pointer := (*float32)(unsafe.Pointer(&result.value))

		return F32(*pointer)
	case cWasmF64:
		pointer := (*float64)(unsafe.Pointer(&result.value))

		return F64(*pointer)
	default:
		panic("unreachable")
	}
}

func handleMemoryExport(wasmExport *cWasmerExportT) (*Memory, error) {
	var wasmMemory *cWasmerMemoryT

	if cWasmerExportToMemory(wasmExport, &wasmMemory) != cWasmerOk {
		return nil, NewInstanceError("Failed to extract the exported memory.")
	}

	memory := newMemory(wasmMemory)
	return &memory, nil
}

func handleFunctionExport(wasmExport *cWasmerExportT) (*ExportedFunction, error) {
	var wasmFunctionInputsArity cUint32T

	var wasmExportName = cWasmerExportName(wasmExport)
	var wasmFunction = cWasmerExportToFunc(wasmExport)

	var exportedFunctionName = cGoStringN((*cChar)(unsafe.Pointer(wasmExportName.bytes)), (cInt)(wasmExportName.bytes_len))

	if cWasmerExportFuncParamsArity(wasmFunction, &wasmFunctionInputsArity) != cWasmerOk {
		return nil, NewExportedFunctionError(exportedFunctionName, "Failed to read the input arity of the `%s` exported function.")
	}

	var wasmFunctionInputSignatures = make([]cWasmerValueTag, int(wasmFunctionInputsArity))

	if wasmFunctionInputsArity > 0 {
		wasmFunctionInputSignaturesCPointer := (*cWasmerValueTag)(unsafe.Pointer(&wasmFunctionInputSignatures[0]))

		if cWasmerExportFuncParams(wasmFunction, wasmFunctionInputSignaturesCPointer, wasmFunctionInputsArity) != cWasmerOk {
			return nil, NewExportedFunctionError(exportedFunctionName, "Failed to read the signature of the `%s` exported function.")
		}
	}

	var wasmFunctionOutputsArity cUint32T

	if cWasmerExportFuncResultsArity(wasmFunction, &wasmFunctionOutputsArity) != cWasmerOk {
		return nil, NewExportedFunctionError(exportedFunctionName, "Failed to read the output arity of the `%s` exported function.")
	}

	return &ExportedFunction{
		name:            exportedFunctionName,
		inputsArity:     wasmFunctionInputsArity,
		outputsArity:    wasmFunctionOutputsArity,
		inputsSignature: wasmFunctionInputSignatures,
	}, nil
}
