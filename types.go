package proton

// Marshaller is the interface implemented by marshallers.
type Marshaller interface {
	Messages() []any
	ID(msg any) (uint64, error)
	Size(msg any) (uint64, error)
	Marshal(msg any, buf []byte) (uint64, uint64, error)
	Unmarshal(id uint64, buf []byte) (any, uint64, error)
	MakePatch(msgDst, msgSrc any, buf []byte) (uint64, uint64, error)
	ApplyPatch(msg any, buf []byte) (uint64, error)
}
