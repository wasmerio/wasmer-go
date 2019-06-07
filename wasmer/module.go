package wasmer

import (
	"io/ioutil"
	"unsafe"
)

// ReadBytes reads a `.wasm` file and returns its content as an array of bytes.
func ReadBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Validate validates a sequence of bytes that is supposed to represent a valid
// WebAssembly module.
func Validate(bytes []byte) bool {
	return true == cWasmerValidate((*cUchar)(unsafe.Pointer(&bytes[0])), cUint(len(bytes)))
}

// ModuleError represents any kind of errors related to a WebAssembly
// module.
type ModuleError struct {
	// Error message.
	message string
}

// NewModuleError constructs a new `ModuleError`.
func NewModuleError(message string) *ModuleError {
	return &ModuleError{message}
}

// `ModuleError` is an actual error. The `Error` function returns the
// error message.
func (error *ModuleError) Error() string {
	return error.message
}

// ExportDescriptor represents an export descriptor of a WebAssembly
// module. It is different of an export of a WebAssembly instance. An
// export descriptor only has a name and a kind/type.
type ExportDescriptor struct {
	// The export name.
	Name string

	// The export kind/type.
	Kind ExportKind
}

// ExportKind represents an export descriptor kind/type.
type ExportKind int

const (
	// ExportKindFunction represents an export descriptor of kind function.
	ExportKindFunction = ExportKind(cWasmFunction)

	// ExportKindGlobal represents an export descriptor of kind global.
	ExportKindGlobal = ExportKind(cWasmGlobal)

	// ExportKindMemory represents an export descriptor of kind memory.
	ExportKindMemory = ExportKind(cWasmMemory)

	// ExportKindTable represents an export descriptor of kind table.
	ExportKindTable = ExportKind(cWasmTable)
)

// Module represents a WebAssembly module.
type Module struct {
	module  *cWasmerModuleT
	Exports []ExportDescriptor
}

// Compile compiles a WebAssembly module from bytes.
func Compile(bytes []byte) (Module, error) {
	var module *cWasmerModuleT

	var compileResult = cWasmerCompile(
		&module,
		(*cUchar)(unsafe.Pointer(&bytes[0])),
		cUint(len(bytes)),
	)

	var emptyModule = Module{module: nil}

	if compileResult != cWasmerOk {
		return emptyModule, NewModuleError("Failed to compile the module.")
	}

	var exports = moduleExports(module)

	return Module{module, exports}, nil
}

func moduleExports(module *cWasmerModuleT) []ExportDescriptor {
	var exportDescriptors *cWasmerExportDescriptorsT
	cWasmerExportDescriptors(module, &exportDescriptors)
	defer cWasmerExportDescriptorsDestroy(exportDescriptors)

	var numberOfExportDescriptors = int(cWasmerExportDescriptorsLen(exportDescriptors))
	var exports = make([]ExportDescriptor, numberOfExportDescriptors)

	for nth := 0; nth < numberOfExportDescriptors; nth++ {
		var exportDescriptor = cWasmerExportDescriptorsGet(exportDescriptors, cInt(nth))
		var exportKind = cWasmerExportDescriptorKind(exportDescriptor)
		var wasmExportName = cWasmerExportDescriptorName(exportDescriptor)
		var exportName = cGoStringN((*cChar)(unsafe.Pointer(wasmExportName.bytes)), (cInt)(wasmExportName.bytes_len))

		exports[nth] = ExportDescriptor{
			Name: exportName,
			Kind: ExportKind(exportKind),
		}
	}

	return exports
}

// Instantiate creates a new instance of the WebAssembly module.
func (module *Module) Instantiate() (Instance, error) {
	return module.InstantiateWithImports(NewImports())
}

// InstantiateWithImports creates a new instance with imports of the WebAssembly module.
func (module *Module) InstantiateWithImports(imports *Imports) (Instance, error) {
	return newInstanceWithImports(
		imports,
		func(wasmImportsCPointer *cWasmerImportT, numberOfImports int) (*cWasmerInstanceT, error) {
			var instance *cWasmerInstanceT

			var instantiateResult = cWasmerModuleInstantiate(
				module.module,
				&instance,
				wasmImportsCPointer,
				cInt(numberOfImports),
			)

			if instantiateResult != cWasmerOk {
				return nil, NewModuleError("Failed to instantiate the module.")
			}

			return instance, nil
		},
	)
}

// Serialize serializes the current module into a sequence of
// bytes. Those bytes can be deserialized into a module with
// `DeserializeModule`.
func (module *Module) Serialize() ([]byte, error) {
	var serializedModule *cWasmerSerializedModuleT
	var serializeResult = cWasmerModuleSerialize(&serializedModule, module.module)
	defer cWasmerSerializedModuleDestroy(serializedModule)

	if serializeResult != cWasmerOk {
		return nil, NewModuleError("Failed to serialize the module.")
	}

	return cWasmerSerializedModuleBytes(serializedModule), nil
}

// DeserializeModule deserializes a sequence of bytes into a
// module. Ideally, those bytes must come from `Module.Serialize`.
func DeserializeModule(serializedModuleBytes []byte) (Module, error) {
	var emptyModule = Module{module: nil}

	if len(serializedModuleBytes) < 1 {
		return emptyModule, NewModuleError("Serialized module bytes are empty.")
	}

	var serializedModule *cWasmerSerializedModuleT
	var deserializeBytesResult = cWasmerSerializedModuleFromBytes(
		&serializedModule,
		(*cUint8T)(unsafe.Pointer(&serializedModuleBytes[0])),
		cInt(len(serializedModuleBytes)),
	)
	defer cWasmerSerializedModuleDestroy(serializedModule)

	if deserializeBytesResult != cWasmerOk {
		return emptyModule, NewModuleError("Failed to reconstitute the serialized module from the given bytes.")
	}

	var module *cWasmerModuleT
	var deserializeResult = cWasmerModuleDeserialize(&module, serializedModule)

	if deserializeResult != cWasmerOk {
		return emptyModule, NewModuleError("Failed to deserialize the module.")
	}

	var exports = moduleExports(module)

	return Module{module, exports}, nil
}

// Close closes/frees a `Module`.
func (module *Module) Close() {
	if module.module != nil {
		cWasmerModuleDestroy(module.module)
	}
}
