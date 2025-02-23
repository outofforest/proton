package proton

// Marshaller is the interface implemented by marshallers.
type Marshaller interface {
	Marshal(msg any, buf []byte) (uint64, uint64, error)
	Unmarshal(id uint64, buf []byte) (any, uint64, error)
}
