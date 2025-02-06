package spkg

import (
	"github.com/outofforest/mass"
	"github.com/outofforest/proton"
	"github.com/pkg/errors"
)

const (
	// IDSubMsg is the ID of SubMsg message.
	IDSubMsg uint64 = iota + 1
)

var _ proton.Marshaller = Marshaller{}

// NewMarshaller creates marshaller.
func NewMarshaller(capacity uint64) Marshaller {
	return Marshaller{
		massSubMsg: mass.New[SubMsg](capacity),
	}
}

// Marshaller marshals and unmarshals messages.
type Marshaller struct {
	massSubMsg *mass.Mass[SubMsg]
}

// Marshal marshals message.
func (m Marshaller) Marshal(msg proton.Marshallable, buf []byte) (retID, retSize uint64, retErr error) {
	defer func() {
		if res := recover(); res != nil {
			retErr = errors.Errorf("marshaling message failed: %s", res)
		}
	}()

	switch msg2 := msg.(type) {
	case *SubMsg:
		return IDSubMsg, msg2.Marshal(buf), nil
	default:
		return 0, 0, errors.Errorf("unknown message type %T", m)
	}
}

// Unmarshal unmarshals message.
func (m Marshaller) Unmarshal(id uint64, buf []byte) (retMsg any, retSize uint64, retErr error) {
	defer func() {
		if res := recover(); res != nil {
			retErr = errors.Errorf("unmarshaling message failed: %s", res)
		}
	}()

	switch id {
	case IDSubMsg:
		msg := m.massSubMsg.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	default:
		return nil, 0, errors.Errorf("unknown ID %d", id)
	}
}

var _ proton.Message = &SubMsg{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *SubMsg) Size() uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := uint32(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFFFFFF
			}
			switch {
			case vi <= 0x7F:
			case vi <= 0x3FFF:
				n++
			case vi <= 0x1FFFFF:
				n += 2
			case vi <= 0xFFFFFFF:
				n += 3
			default:
				n += 4
			}
		}
	}
	return n
}

// Marshal marshals the structure.
func (m *SubMsg) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint32(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFFFFFF
			}
			switch {
			case vi <= 0x7F:
				b[o] = byte(vi)
				o++
			case vi <= 0x3FFF:
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi)
				o++
			case vi <= 0x1FFFFF:
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi)
				o++
			case vi <= 0xFFFFFFF:
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi)
				o++
			default:
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi)
				o++
			}
		}
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *SubMsg) Unmarshal(
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint32(b[o] & 0x7F)
			if b[o]&0x80 == 0 {
				o++
			} else {
				vi |= uint32(b[o+1]&0x7F) << 7
				if b[o+1]&0x80 == 0 {
					o += 2
				} else {
					vi |= uint32(b[o+2]&0x7F) << 14
					if b[o+2]&0x80 == 0 {
						o += 3
					} else {
						vi |= uint32(b[o+3]&0x7F) << 21
						if b[o+3]&0x80 == 0 {
							o += 4
						} else {
							vi |= uint32(b[o+4]) << 28
							o += 5
						}
					}
				}
			}
			if vi&0x01 == 0 {
				vi >>= 1
			} else {
				vi >>= 1
				vi ^= 0xFFFFFFFF
			}
			m.Value = int32(vi)
		}
	}

	return o
}
