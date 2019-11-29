package wasmer

import (
	"unsafe"
)

// MapDirEntry is an entry that can be passed to `NewWasiImportObject`.
// Preopens a file for the WASI module but renames it to the given name
type MapDirEntry struct {
	alias    string
	hostPath string
}

// NewDefaultWasiImportObject constructs a new `ImportObject`
// with WASI host imports.
//
// To specify WASI program arguments, environment variables,
// preopened directories, and more, see `NewWasiImportObject`
func NewDefaultWasiImportObject() *ImportObject {
	var inner = cNewWasmerDefaultWasiImportObject()

	return &ImportObject{inner}
}

// NewWasiImportObject creates an `ImportObject` with the default WASI imports.
// Specify arguments (the first is the program name),
// environment variables ("envvar=value"), preopened directories
// (host file paths), and mapped directories (host file paths with an
// alias, see `MapDirEntry`)
func NewWasiImportObject(
	arguments []string,
	environmentVariables []string,
	preopenedDirs []string,
	mappedDirs []MapDirEntry,
) *ImportObject {
	var argumentsBytes = []cWasmerByteArray{}

	for _, argument := range arguments {
		argumentsBytes = append(argumentsBytes, cGoStringToWasmerByteArray(argument))
	}

	var environmentVariablesBytes = []cWasmerByteArray{}

	for _, env := range environmentVariables {
		environmentVariablesBytes = append(environmentVariablesBytes, cGoStringToWasmerByteArray(env))
	}

	var preopenedDirsBytes = []cWasmerByteArray{}

	for _, preopenedDir := range preopenedDirs {
		preopenedDirsBytes = append(preopenedDirsBytes, cGoStringToWasmerByteArray(preopenedDir))
	}
	var mappedDirsBytes = []cWasmerWasiMapDirEntryT{}

	for _, mappedDir := range mappedDirs {
		var wasiMappedDir = cAliasAndHostPathToWasiDirEntry(mappedDir.alias, mappedDir.hostPath)
		mappedDirsBytes = append(mappedDirsBytes, wasiMappedDir)
	}

	var inner = cNewWasmerWasiImportObject(
		(*cWasmerByteArray)(unsafe.Pointer(&argumentsBytes)), len(argumentsBytes),
		(*cWasmerByteArray)(unsafe.Pointer(&environmentVariablesBytes)), len(environmentVariablesBytes),
		(*cWasmerByteArray)(unsafe.Pointer(&preopenedDirsBytes)), len(preopenedDirsBytes),
		(*cWasmerWasiMapDirEntryT)(unsafe.Pointer(&mappedDirsBytes)), len(mappedDirsBytes),
	)

	return &ImportObject{inner}
}
