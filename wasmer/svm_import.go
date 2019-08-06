package wasmer

import "unsafe"

func (instanceContext *InstanceContext) NodeData() unsafe.Pointer {
	return cWasmerSvmInstanceContextNodeDataGet(instanceContext.context)
}
