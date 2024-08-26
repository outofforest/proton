package proton

// Marshaler is the interface implemented by marshalers.
type Marshaler interface {
	Marshal(msg Marshalable, buf []byte) (uint64, uint64, error)
	Unmarshal(id uint64, buf []byte) (any, uint64, error)
}

// Marshalable is the interface for marshalable messages.
type Marshalable interface {
	Marshal(buf []byte) uint64
}

// Message is the interface implemented by messages.
type Message interface {
	Marshalable

	Size() uint64
	Unmarshal(b []byte) uint64
}
