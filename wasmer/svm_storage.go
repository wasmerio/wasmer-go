package wasmer

import "C"
import "unsafe"

func WasmerSvmContextRegisterSet(context *cWasmerInstanceContextT, regBits int32, regIndex int32, buf []byte) {
	cWasmerSvmRegisterSet(context, int(regBits), int(regIndex), buf)

 	// TODO:
	// ensure that `res` is `cWasmerOK`
}

func WasmerSvmContextRegisterGet(context *cWasmerInstanceContextT, regBits int32, regIndex int32) []byte {
	// TODO:
	// ensure that `res` is `cWasmerOK`
	var registerData unsafe.Pointer

	cWasmerSvmRegisterGet(&registerData, context, int(regBits), int(regIndex))

	// TODO: do we need to call `C.free` on `registerData` ??
	regBytes := int(regBits / 8)
	return C.GoBytes(registerData, C.int(regBytes))
}

func WasmerSvmRegisterSet(instanceContext InstanceContext, regBits int32, regIndex int32, buf []byte) {
	WasmerSvmContextRegisterSet(instanceContext.context, regBits, regIndex, buf)
}

func WasmerSvmRegisterGet(instanceContext InstanceContext, regBits int32, regIndex int32) []byte {
	return WasmerSvmContextRegisterGet(instanceContext.context, regBits, regIndex)
}
