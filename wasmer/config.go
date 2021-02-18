package wasmer

// #include <wasmer_wasm.h>
import "C"

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

// UseJITEngine sets the engine to JIT in the configuration.
//
//   config := NewConfig()
//   config.UseJITEngine()
//
func (self *Config) UseJITEngine() *Config {
	C.wasm_config_set_engine(self.inner(), C.JIT)

	return self
}

// UseNativeEngine sets the engine to Native in the configuration.
//
//   config := NewConfig()
//   config.UseNativeEngine()
//
func (self *Config) UseNativeEngine() *Config {
	C.wasm_config_set_engine(self.inner(), C.NATIVE)

	return self
}

// UseCraneliftCompiler sets the compiler to Cranelift in the configuration.
//
//   config := NewConfig()
//   config.UseCraneliftCompiler()
//
func (self *Config) UseCraneliftCompiler() *Config {
	C.wasm_config_set_engine(self.inner(), C.CRANELIFT)

	return self
}

// UseLLVMCompiler sets the compiler to LLVM in the configuration.
//
//   config := NewConfig()
//   config.UseLLVMCompiler()
//
func (self *Config) UseLLVMCompiler() *Config {
	C.wasm_config_set_engine(self.inner(), C.LLVM)

	return self
}

// UseSinglepassCompiler sets the compiler to Singlepass in the
// configuration.
//
//   config := NewConfig()
//   config.UseSinglepassCompiler()
//
func (self *Config) UseSinglepassCompiler() *Config {
	C.wasm_config_set_engine(self.inner(), C.SINGLEPASS)

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
