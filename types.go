package proton

// Marshaller is the interface implemented by marshallers.
type Marshaller interface {
	Marshal(msg Marshallable, buf []byte) (uint64, uint64, error)
	Unmarshal(id uint64, buf []byte) (any, uint64, error)
}

// Marshallable is the interface for marshallable messages.
type Marshallable interface {
	Marshal(buf []byte) uint64
}

// Message is the interface implemented by messages.
type Message interface {
	Marshallable

	Size() uint64
}
