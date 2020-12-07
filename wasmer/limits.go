package wasmer

// #include <wasmer_wasm.h>
//
// uint32_t limit_max_unbound() {
//     return wasm_limits_max_default;
// }
import "C"
import "runtime"

func LimitMaxUnbound() uint32 {
	return uint32(C.limit_max_unbound())
}

type Limits struct {
	_inner C.wasm_limits_t
}

func newLimits(pointer *C.wasm_limits_t, ownedBy interface{}) *Limits {
	limits, err := NewLimits(uint32(pointer.min), uint32(pointer.max))

	if err != nil {
		return nil
	}

	if ownedBy != nil {
		runtime.KeepAlive(ownedBy)
	}

	return limits
}

func NewLimits(minimum uint32, maximum uint32) (*Limits, error) {
	if minimum > maximum {
		return nil, newErrorWith("The minimum limit is greater than the maximum one")
	}

	return &Limits{
		_inner: C.wasm_limits_t{
			min: C.uint32_t(minimum),
			max: C.uint32_t(maximum),
		},
	}, nil
}

func (self *Limits) inner() *C.wasm_limits_t {
	return &self._inner
}

func (self *Limits) Minimum() uint32 {
	return uint32(self.inner().min)
}

func (self *Limits) Maximum() uint32 {
	return uint32(self.inner().max)
}
