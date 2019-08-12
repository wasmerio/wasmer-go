package wasmer

import "C"
import "unsafe"

func WasmerSvmRegisterSet(instanceContext InstanceContext, regIndex uint, buf []byte) {
	cWasmerSvmRegisterSet(instanceContext.context, regIndex, buf)
 	// TODO:
	// ensure that `res` is `cWasmerOK`
	// _res := cWasmerSvmRegisterGet(nil, instanceContext, regIndex)
}

func WasmerSvmRegisterGet(instanceContext InstanceContext, regIndex int32) []byte {
	// TODO:
	// ensure that `res` is `cWasmerOK`
	var registerData unsafe.Pointer

	cWasmerSvmRegisterGet(&registerData, instanceContext.context, (C.uint)(regIndex))

	// TODO: do we need to call `C.free` on `registerData` ??
	return C.GoBytes(registerData, 8)
}
