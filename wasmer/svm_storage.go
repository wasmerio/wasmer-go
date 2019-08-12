package wasmer

import "C"
import "unsafe"

func WasmerSvmContextRegisterSet(context *cWasmerInstanceContextT, regIndex uint, buf []byte) {
	cWasmerSvmRegisterSet(context, regIndex, buf)
 	// TODO:
	// ensure that `res` is `cWasmerOK`
	// _res := cWasmerSvmRegisterGet(nil, instanceContext, regIndex)
}

func WasmerSvmContextRegisterGet(context *cWasmerInstanceContextT, regIndex int32) []byte {
	// TODO:
	// ensure that `res` is `cWasmerOK`
	var registerData unsafe.Pointer

	cWasmerSvmRegisterGet(&registerData, context, (C.uint)(regIndex))

	// TODO: do we need to call `C.free` on `registerData` ??
	return C.GoBytes(registerData, 8)
}

func WasmerSvmRegisterSet(instanceContext InstanceContext, regIndex uint, buf []byte) {
	WasmerSvmContextRegisterSet(instanceContext.context, regIndex, buf)
}

func WasmerSvmRegisterGet(instanceContext InstanceContext, regIndex int32) []byte {
	return WasmerSvmContextRegisterGet(instanceContext.context, regIndex)
}
