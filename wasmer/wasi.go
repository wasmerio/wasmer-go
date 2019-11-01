package wasmer

import (
	"unsafe"
)

// An entry that can be passed to `NewWasiImportObject`.
// Preopens a file for the WASI module but renames it to the given name
type MapDirEntry struct {
	alias    string
	hostPath string
}

// NewDefaultWasiImportObject constructs a new `ImportObject`
// with WASI host imports.
// To specify WASI program arguments, environment variables,
// preopened directories, and more, see `NewWasiImportObject`
func NewDefaultWasiImportObject() *ImportObject {
	var inner = cNewWasmerDefaultWasiImportObject()

	return &ImportObject{inner}
}

// Creates an `ImportObject` with the default WASI imports.
// Specify arguments (the first is the program name),
// environment variables ("envvar=value"), preopened directories
// (host file paths), and mapped directories (host file paths with an
// alias, see `MapDirEntry`)
func NewWasiImportObject(args []string, envs []string,
	preopenedDirs []string, mappedDirs []MapDirEntry) *ImportObject {
	var argBytes = []cWasmerByteArray{}
	for _, arg := range args {
		argBytes = append(argBytes, cGoStringToWasmerByteArray(arg))
	}
	var envBytes = []cWasmerByteArray{}
	for _, env := range envs {
		envBytes = append(envBytes, cGoStringToWasmerByteArray(env))
	}
	var poDirBytes = []cWasmerByteArray{}
	for _, poDir := range preopenedDirs {
		poDirBytes = append(poDirBytes, cGoStringToWasmerByteArray(poDir))
	}
	var mappedDirBytes = []cWasmerWasiMapDirEntryT{}
	for _, mappedDir := range mappedDirs {
		var wasiMappedDir = cAliasAndHostPathToWasiDirEntry(mappedDir.alias, mappedDir.hostPath)
		mappedDirBytes = append(mappedDirBytes, wasiMappedDir)
	}

	var inner = cNewWasmerWasiImportObject((*cWasmerByteArray)(unsafe.Pointer(&argBytes)), len(argBytes),
		(*cWasmerByteArray)(unsafe.Pointer(&envBytes)), len(envBytes),
		(*cWasmerByteArray)(unsafe.Pointer(&poDirBytes)), len(poDirBytes),
		(*cWasmerWasiMapDirEntryT)(unsafe.Pointer(&mappedDirBytes)), len(mappedDirBytes),
	)

	return &ImportObject{inner}
}
