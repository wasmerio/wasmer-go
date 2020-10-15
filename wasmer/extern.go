package wasmer

type Extern interface {
	Type() ExternType
}
