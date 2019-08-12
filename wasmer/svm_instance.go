package wasmer

import "unsafe"

type SvmInstanceConfig struct {
	Addr           []byte
	State          []byte
	MaxPages       uint
	MaxPagesSlices uint
	NodeDataPtr    unsafe.Pointer
}

func NewSvmInstanceConfig(addr []byte, state []byte, maxPages uint, maxPagesSlices uint, nodeDataPtr unsafe.Pointer) *SvmInstanceConfig {
	return &SvmInstanceConfig{Addr: addr, State: state, MaxPages: maxPages, MaxPagesSlices: maxPagesSlices, NodeDataPtr: nodeDataPtr}
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
