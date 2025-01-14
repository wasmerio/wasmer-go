package wasmer

// #include <wasmer.h>
import "C"

// Target represents a triple + CPU features pairs.
type Target struct {
	CPtrBase[*C.wasmer_target_t]
}

func newTarget(target *C.wasmer_target_t) *Target {
	self := &Target{CPtrBase: mkPtr(target)}
	self.SetFinalizer(func(v *C.wasmer_target_t) {
		C.wasmer_target_delete(v)
	})
	return self
}

// NewTarget creates a new target.
//
//	triple, err := NewTriple("aarch64-unknown-linux-gnu")
//	cpuFeatures := NewCpuFeatures()
//	target := NewTarget(triple, cpuFeatures)
func NewTarget(triple *Triple, cpuFeatures *CpuFeatures) *Target {
	return newTarget(C.wasmer_target_new(triple.release(), cpuFeatures.release()))
}

func (self *Target) inner() *C.wasmer_target_t {
	return self.ptr()
}

// Triple; historically such things had three fields, though they have
// added additional fields over time.
type Triple struct {
	CPtrBase[*C.wasmer_triple_t]
}

func newTriple(triple *C.wasmer_triple_t) *Triple {
	self := &Triple{CPtrBase: mkPtr(triple)}
	self.SetFinalizer(func(v *C.wasmer_triple_t) {
		C.wasmer_triple_delete(v)
	})
	return self
}

// NewTriple creates a new triple, otherwise it returns an error
// specifying why the provided triple isn't valid.
//
//	triple, err := NewTriple("aarch64-unknown-linux-gnu")
func NewTriple(triple string) (*Triple, error) {
	cTripleName := newName(triple)
	defer C.wasm_name_delete(&cTripleName)

	var cTriple *C.wasmer_triple_t

	err := maybeNewErrorFromWasmer(func() bool {
		cTriple := C.wasmer_triple_new(&cTripleName)

		return cTriple == nil
	})

	if err != nil {
		return nil, err
	}

	return newTriple(cTriple), nil
}

// NewTripleFromHost creates a new triple from the current host.
func NewTripleFromHost() *Triple {
	return newTriple(C.wasmer_triple_new_from_host())
}

func (self *Triple) inner() *C.wasmer_triple_t {
	return self.ptr()
}

// CpuFeatures holds a set of CPU features. They are identified by
// their stringified names. The reference is the GCC options:
//
// • https://gcc.gnu.org/onlinedocs/gcc/x86-Options.html,
//
// • https://gcc.gnu.org/onlinedocs/gcc/ARM-Options.html,
//
// • https://gcc.gnu.org/onlinedocs/gcc/AArch64-Options.html.
//
// At the time of writing this documentation (it might be outdated in
// the future), the supported featurse are the following:
//
// • sse2,
//
// • sse3,
//
// • ssse3,
//
// • sse4.1,
//
// • sse4.2,
//
// • popcnt,
//
// • avx,
//
// • bmi,
//
// • bmi2,
//
// • avx2,
//
// • avx512dq,
//
// • avx512vl,
//
// • lzcnt.
type CpuFeatures struct {
	CPtrBase[*C.wasmer_cpu_features_t]
}

func newCpuFeatures(cpu_features *C.wasmer_cpu_features_t) *CpuFeatures {
	self := &CpuFeatures{
		CPtrBase: mkPtr(cpu_features),
	}
	self.SetFinalizer(func(v *C.wasmer_cpu_features_t) {
		C.wasmer_cpu_features_delete(v)
	})
	return self
}

// NewCpuFeatures creates a new CpuFeatures, which is a set of CPU
// features.
func NewCpuFeatures() *CpuFeatures {
	return newCpuFeatures(C.wasmer_cpu_features_new())
}

// Add adds a new CPU feature to the existing set.
func (self *CpuFeatures) Add(feature string) error {
	cFeature := newName(feature)
	defer C.wasm_name_delete(&cFeature)

	err := maybeNewErrorFromWasmer(func() bool {
		return false == C.wasmer_cpu_features_add(self.inner(), &cFeature)
	})

	if err != nil {
		return err
	}

	return nil
}

func (self *CpuFeatures) inner() *C.wasmer_cpu_features_t {
	return self.ptr()
}
