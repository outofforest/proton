package pkg2

import (
	"unsafe"

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
	var n uint64 = 3
	{
		// Value1

		l := uint64(len(m.Value1))
		{
			vi := l
			switch {
			case vi <= 0x7F:
			case vi <= 0x3FFF:
				n++
			case vi <= 0x1FFFFF:
				n += 2
			case vi <= 0xFFFFFFF:
				n += 3
			case vi <= 0x7FFFFFFFF:
				n += 4
			case vi <= 0x3FFFFFFFFFF:
				n += 5
			case vi <= 0x1FFFFFFFFFFFF:
				n += 6
			case vi <= 0xFFFFFFFFFFFFFF:
				n += 7
			default:
				n += 8
			}
		}
		n += l * 2
		for _, mv2 := range m.Value1 {
			{
				l := uint64(len(mv2))
				n += l
				{
					vi := l
					switch {
					case vi <= 0x7F:
					case vi <= 0x3FFF:
						n++
					case vi <= 0x1FFFFF:
						n += 2
					case vi <= 0xFFFFFFF:
						n += 3
					case vi <= 0x7FFFFFFFF:
						n += 4
					case vi <= 0x3FFFFFFFFFF:
						n += 5
					case vi <= 0x1FFFFFFFFFFFF:
						n += 6
					case vi <= 0xFFFFFFFFFFFFFF:
						n += 7
					default:
						n += 8
					}
				}
			}
		}
	}
	{
		// Value2

		{
			l := uint64(len(m.Value2))
			n += l
			{
				vi := l
				switch {
				case vi <= 0x7F:
				case vi <= 0x3FFF:
					n++
				case vi <= 0x1FFFFF:
					n += 2
				case vi <= 0xFFFFFFF:
					n += 3
				case vi <= 0x7FFFFFFFF:
					n += 4
				case vi <= 0x3FFFFFFFFFF:
					n += 5
				case vi <= 0x1FFFFFFFFFFFF:
					n += 6
				case vi <= 0xFFFFFFFFFFFFFF:
					n += 7
				default:
					n += 8
				}
			}
		}
	}
	{
		// Value3

		{
			vi := m.Value3
			switch {
			case vi <= 0x7F:
			case vi <= 0x3FFF:
				n++
			default:
				n += 2
			}
		}
	}
	return n
}

// Marshal marshals the structure.
func (m *SubMsg) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value1

		{
			vi := uint64(len(m.Value1))
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
			case vi <= 0x7FFFFFFFF:
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
			case vi <= 0x3FFFFFFFFFF:
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
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi)
				o++
			case vi <= 0x1FFFFFFFFFFFF:
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
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi) | 0x80
				o++
				vi >>= 7
				b[o] = byte(vi)
				o++
			case vi <= 0xFFFFFFFFFFFFFF:
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
		for mk1, mv2 := range m.Value1 {
			if mk1 {
				b[o] = 0x01
			} else {
				b[o] = 0x00
			}
			o++
			{
				l := uint64(len(mv2))
				{
					vi := l
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
					case vi <= 0x7FFFFFFFF:
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
					case vi <= 0x3FFFFFFFFFF:
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
						b[o] = byte(vi) | 0x80
						o++
						vi >>= 7
						b[o] = byte(vi)
						o++
					case vi <= 0x1FFFFFFFFFFFF:
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
						b[o] = byte(vi) | 0x80
						o++
						vi >>= 7
						b[o] = byte(vi) | 0x80
						o++
						vi >>= 7
						b[o] = byte(vi)
						o++
					case vi <= 0xFFFFFFFFFFFFFF:
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
				copy(b[o:o+l], mv2)
				o += l
			}
		}
	}
	{
		// Value2

		{
			l := uint64(len(m.Value2))
			{
				vi := l
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
				case vi <= 0x7FFFFFFFF:
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
				case vi <= 0x3FFFFFFFFFF:
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
					b[o] = byte(vi) | 0x80
					o++
					vi >>= 7
					b[o] = byte(vi)
					o++
				case vi <= 0x1FFFFFFFFFFFF:
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
					b[o] = byte(vi) | 0x80
					o++
					vi >>= 7
					b[o] = byte(vi) | 0x80
					o++
					vi >>= 7
					b[o] = byte(vi)
					o++
				case vi <= 0xFFFFFFFFFFFFFF:
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
			copy(b[o:o+l], m.Value2)
			o += l
		}
	}
	{
		// Value3

		{
			vi := m.Value3
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
			default:
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
		// Value1

		var l uint64
		{
			vi := uint64(b[o] & 0x7F)
			if b[o]&0x80 == 0 {
				o++
			} else {
				vi |= uint64(b[o+1]&0x7F) << 7
				if b[o+1]&0x80 == 0 {
					o += 2
				} else {
					vi |= uint64(b[o+2]&0x7F) << 14
					if b[o+2]&0x80 == 0 {
						o += 3
					} else {
						vi |= uint64(b[o+3]&0x7F) << 21
						if b[o+3]&0x80 == 0 {
							o += 4
						} else {
							vi |= uint64(b[o+4]&0x7F) << 28
							if b[o+4]&0x80 == 0 {
								o += 5
							} else {
								vi |= uint64(b[o+5]&0x7F) << 35
								if b[o+5]&0x80 == 0 {
									o += 6
								} else {
									vi |= uint64(b[o+6]&0x7F) << 42
									if b[o+6]&0x80 == 0 {
										o += 7
									} else {
										vi |= uint64(b[o+7]&0x7F) << 49
										if b[o+7]&0x80 == 0 {
											o += 8
										} else {
											vi |= uint64(b[o+8]) << 56
											o += 9
										}
									}
								}
							}
						}
					}
				}
			}
			l = vi
		}
		if l > 0 {
			m.Value1 = make(map[bool]string, l)
		
			var mk1 bool
			var mv2 string
		
			for range l {
				mk1 = b[o] != 0x00
				o++
				{
					var l uint64
					{
						vi := uint64(b[o] & 0x7F)
						if b[o]&0x80 == 0 {
							o++
						} else {
							vi |= uint64(b[o+1]&0x7F) << 7
							if b[o+1]&0x80 == 0 {
								o += 2
							} else {
								vi |= uint64(b[o+2]&0x7F) << 14
								if b[o+2]&0x80 == 0 {
									o += 3
								} else {
									vi |= uint64(b[o+3]&0x7F) << 21
									if b[o+3]&0x80 == 0 {
										o += 4
									} else {
										vi |= uint64(b[o+4]&0x7F) << 28
										if b[o+4]&0x80 == 0 {
											o += 5
										} else {
											vi |= uint64(b[o+5]&0x7F) << 35
											if b[o+5]&0x80 == 0 {
												o += 6
											} else {
												vi |= uint64(b[o+6]&0x7F) << 42
												if b[o+6]&0x80 == 0 {
													o += 7
												} else {
													vi |= uint64(b[o+7]&0x7F) << 49
													if b[o+7]&0x80 == 0 {
														o += 8
													} else {
														vi |= uint64(b[o+8]) << 56
														o += 9
													}
												}
											}
										}
									}
								}
							}
						}
						l = vi
					}
					if l > 0 {
						mv2 = unsafe.String((*byte)(unsafe.Pointer(&b[o])), l)
						o += l
					} else {
						mv2 = "" 
					}
				}
				m.Value1[mk1] = mv2
			}
		} else {
			m.Value1 = nil
		}
	}
	{
		// Value2

		{
			var l uint64
			{
				vi := uint64(b[o] & 0x7F)
				if b[o]&0x80 == 0 {
					o++
				} else {
					vi |= uint64(b[o+1]&0x7F) << 7
					if b[o+1]&0x80 == 0 {
						o += 2
					} else {
						vi |= uint64(b[o+2]&0x7F) << 14
						if b[o+2]&0x80 == 0 {
							o += 3
						} else {
							vi |= uint64(b[o+3]&0x7F) << 21
							if b[o+3]&0x80 == 0 {
								o += 4
							} else {
								vi |= uint64(b[o+4]&0x7F) << 28
								if b[o+4]&0x80 == 0 {
									o += 5
								} else {
									vi |= uint64(b[o+5]&0x7F) << 35
									if b[o+5]&0x80 == 0 {
										o += 6
									} else {
										vi |= uint64(b[o+6]&0x7F) << 42
										if b[o+6]&0x80 == 0 {
											o += 7
										} else {
											vi |= uint64(b[o+7]&0x7F) << 49
											if b[o+7]&0x80 == 0 {
												o += 8
											} else {
												vi |= uint64(b[o+8]) << 56
												o += 9
											}
										}
									}
								}
							}
						}
					}
				}
				l = vi
			}
			if l > 0 {
				m.Value2 = unsafe.String((*byte)(unsafe.Pointer(&b[o])), l)
				o += l
			} else {
				m.Value2 = "" 
			}
		}
	}
	{
		// Value3

		{
			vi := uint16(b[o] & 0x7F)
			if b[o]&0x80 == 0 {
				o++
			} else {
				vi |= uint16(b[o+1]&0x7F) << 7
				if b[o+1]&0x80 == 0 {
					o += 2
				} else {
					vi |= uint16(b[o+2]) << 14
					o += 3
				}
			}
			m.Value3 = vi
		}
	}

	return o
}
