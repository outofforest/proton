package proton

// Marshalable is the interface for marshalable messages.
type Marshalable interface {
	Marshal(buf []byte) uint64
}
