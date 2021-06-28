package wasmer

// #include <wasmer.h>
import "C"

// CompilerKind represents the possible compiler types.
type CompilerKind C.wasmer_compiler_t

const (
	// Represents the Cranelift compiler.
	CRANELIFT = CompilerKind(C.CRANELIFT)

	// Represents the LLVM compiler.
	LLVM = CompilerKind(C.LLVM)

	// Represents the Singlepass compiler.
	SINGLEPASS = CompilerKind(C.SINGLEPASS)
)

// Strings returns the CompilerKind as a string.
//
//   CRANELIFT.String() // "cranelift"
//   LLVM.String() // "llvm"
func (self CompilerKind) String() string {
	switch self {
	case CRANELIFT:
		return "cranelift"

	case LLVM:
		return "llvm"

	case SINGLEPASS:
		return "singlepass"
	}
	panic("Unknown compiler")
}

// IsCompilerAvailable checks that the given compiler is available
// in this current version of `wasmer-go`.
//
//   IsCompilerAvailable(CRANELIFT)
func IsCompilerAvailable(compiler CompilerKind) bool {
	return bool(C.wasmer_is_compiler_available(uint32(C.wasmer_compiler_t(compiler))))
}

// EngineKind represents the possible engine types.
type EngineKind C.wasmer_engine_t

const (
	// Represents the Universal engine.
	UNIVERSAL = EngineKind(C.UNIVERSAL)

	// Represents the Dylib engine.
	DYLIB = EngineKind(C.DYLIB)

	// Deprecated constant. Please use UNIVERSAL instead.
	JIT = UNIVERSAL

	// Deprecated constant. Please use DYLIB instead.
	NATIVE = DYLIB
)

// Strings returns the EngineKind as a string.
//
//   JIT.String() // "jit"
//   NATIVE.String() // "native"
func (self EngineKind) String() string {
	switch self {
	case UNIVERSAL:
		return "universal"

	case DYLIB:
		return "dylib"
	}
	panic("Unknown engine")
}

// IsEngineAvailable checks that the given engine is available in this
// current version of `wasmer-go`.
//
//   IsEngineAvailable(UNIVERSAL)
func IsEngineAvailable(engine EngineKind) bool {
	return bool(C.wasmer_is_engine_available(uint32(C.wasmer_engine_t(engine))))
}

// Config holds the compiler and the Engine used by the Store.
type Config struct {
	_inner *C.wasm_config_t
}

// NewConfig instantiates and returns a new Config.
//
//   config := NewConfig()
func NewConfig() *Config {
	config := C.wasm_config_new()

	return &Config{
		_inner: config,
	}
}

func (self *Config) inner() *C.wasm_config_t {
	return self._inner
}

// UseNativeEngine sets the engine to Universal in the configuration.
//
//   config := NewConfig()
//   config.UseUniversalEngine()
//
// This method might fail if the Universal engine isn't
// available. Check `IsEngineAvailable` to learn more.
func (self *Config) UseUniversalEngine() *Config {
	if !IsEngineAvailable(UNIVERSAL) {
		panic("This `wasmer-go` version doesn't include the Universal engine; use `IsEngineAvailable(UNIVERSAL)` to avoid this panic")
	}

	C.wasm_config_set_engine(self.inner(), uint32(C.wasmer_engine_t(UNIVERSAL)))

	return self
}

// UseDylibEngine sets the engine to Dylib in the configuration.
//
//   config := NewConfig()
//   config.UseDylibEngine()
//
// This method might fail if the Dylib engine isn't available. Check
// `IsEngineAvailable` to learn more.
func (self *Config) UseDylibEngine() *Config {
	if !IsEngineAvailable(DYLIB) {
		panic("This `wasmer-go` version doesn't include the DYLIB engine; use `IsEngineAvailable(DYLIB)` to avoid this panic")
	}

	C.wasm_config_set_engine(self.inner(), uint32(C.wasmer_engine_t(DYLIB)))

	return self
}

// UseJITEngine is a deprecated method. Please use UseUniversalEngine
// instead.
func (self *Config) UseJITEngine() *Config {
	return self.UseUniversalEngine()
}

// UseNativeEngine is a deprecated method. Please use
// UseDylibEngine instead.
func (self *Config) UseNativeEngine() *Config {
	return self.UseDylibEngine()
}

// UseCraneliftCompiler sets the compiler to Cranelift in the configuration.
//
//   config := NewConfig()
//   config.UseCraneliftCompiler()
//
// This method might fail if the Cranelift compiler isn't
// available. Check `IsCompilerAvailable` to learn more.
func (self *Config) UseCraneliftCompiler() *Config {
	if !IsCompilerAvailable(CRANELIFT) {
		panic("This `wasmer-go` version doesn't include the Cranelift compiler; use `IsCompilerAvailable(CRANELIFT)` to avoid this panic")
	}

	C.wasm_config_set_compiler(self.inner(), uint32(C.wasmer_compiler_t(CRANELIFT)))

	return self
}

// UseLLVMCompiler sets the compiler to LLVM in the configuration.
//
//   config := NewConfig()
//   config.UseLLVMCompiler()
//
// This method might fail if the LLVM compiler isn't available. Check
// `IsCompilerAvailable` to learn more.
func (self *Config) UseLLVMCompiler() *Config {
	if !IsCompilerAvailable(LLVM) {
		panic("This `wasmer-go` version doesn't include the LLVM compiler; use `IsCompilerAvailable(LLVM)` to avoid this panic")
	}

	C.wasm_config_set_compiler(self.inner(), uint32(C.wasmer_compiler_t(LLVM)))

	return self
}

// UseSinglepassCompiler sets the compiler to Singlepass in the
// configuration.
//
//   config := NewConfig()
//   config.UseSinglepassCompiler()
//
// This method might fail if the Singlepass compiler isn't
// available. Check `IsCompilerAvailable` to learn more.
func (self *Config) UseSinglepassCompiler() *Config {
	if !IsCompilerAvailable(SINGLEPASS) {
		panic("This `wasmer-go` version doesn't include the Singlepass compiler; use `IsCompilerAvailable(SINGLEPASS)` to avoid this panic")
	}

	C.wasm_config_set_compiler(self.inner(), uint32(C.wasmer_compiler_t(SINGLEPASS)))

	return self
}

// Use a specific target for doing cross-compilation.
//
//   triple, _ := NewTriple("aarch64-unknown-linux-gnu")
//   cpuFeatures := NewCpuFeatures()
//   target := NewTarget(triple, cpuFeatures)
//
//   config := NewConfig()
//   config.UseTarget(target)
func (self *Config) UseTarget(target *Target) *Config {
	C.wasm_config_set_target(self.inner(), target.inner())

	return self
}
