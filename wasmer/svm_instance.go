package wasmer

import "unsafe"

type SvmInstanceConfig struct {
	addr           []byte
	maxPages       uint
	maxPagesSlices uint
	nodeDataPtr    unsafe.Pointer
}

func NewSvmInstanceConfig(addr []byte, maxPages uint, maxPagesSlices uint, nodeDataPtr unsafe.Pointer) *SvmInstanceConfig {
	return &SvmInstanceConfig{addr, maxPages, maxPagesSlices, nodeDataPtr}
}

func NewSvmInstance(
	module *Module,
	imports *Imports,
	config *SvmInstanceConfig,
) (Instance, error) {
	return NewInstanceWithModuleAndImportObject(module, imports,
		func(wasmImportsCPointer *cWasmerImportT, numberOfImports uint) (*cWasmerImportObjectT, error) {
			var importObject *cWasmerImportObjectT

			importResult := cWasmerSvmImportObject(&importObject, wasmImportsCPointer, numberOfImports, config)
			if importResult != cWasmerOk {
				return nil, buildInstantiateError()
			}

			return importObject, nil
		})
}
