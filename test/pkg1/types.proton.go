package pkg1

import (
	"unsafe"

	"github.com/outofforest/proton"
	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1/spkg"
	"github.com/outofforest/proton/test/pkg2"
	spkg6 "github.com/outofforest/proton/test/pkg2/spkg"
	"github.com/pkg/errors"
)

const (
	id37 uint64 = iota + 1
	id57
	id56
	id55
	id54
	id53
	id52
	id51
	id50
	id49
	id48
	id47
	id46
	id45
	id44
	id43
	id42
	id41
	id40
	id38
	id35
	id34
	id33
	id32
	id31
	id30
	id29
	id28
	id27
	id25
	id24
	id23
	id22
	id21
	id20
	id19
	id18
	id17
	id16
	id15
	id14
	id13
	id12
	id11
	id10
	id9
	id8
	id7
	id6
	id5
	id4
	id3
	id2
	id1
	id0
)

var _ proton.Marshaller = Marshaller{}

// NewMarshaller creates marshaller.
func NewMarshaller() Marshaller {
	return Marshaller{}
}

// Marshaller marshals and unmarshals messages.
type Marshaller struct {
}

// Messages returns list of the message types supported by marshaller.
func (m Marshaller) Messages() []any {
	return []any {
		pkg2.SubMsg{},
		MsgUint64{},
		MsgUint32{},
		MsgUint16{},
		MsgUint8{},
		MsgInt64{},
		MsgInt32{},
		MsgInt16{},
		MsgInt8{},
		MsgBool3{},
		MsgBool10{},
		MsgFloat64{},
		MsgFloat32{},
		MsgString{},
		MsgArray{},
		MsgSlice{},
		MsgMap{},
		MsgMapString{},
		MsgStruct{},
		MsgStructAnonymous{},
		MsgArrayUint8{},
		MsgArrayInt8{},
		MsgArrayFloat32{},
		MsgArrayFloat64{},
		MsgSliceUint8{},
		MsgSliceInt8{},
		MsgSliceFloat32{},
		MsgSliceFloat64{},
		MsgMixed{},
		MsgUint64Custom{},
		MsgUint32Custom{},
		MsgUint16Custom{},
		MsgUint8Custom{},
		MsgInt64Custom{},
		MsgInt32Custom{},
		MsgInt16Custom{},
		MsgInt8Custom{},
		MsgBoolCustom{},
		MsgFloat64Custom{},
		MsgFloat32Custom{},
		MsgStringCustom{},
		MsgArrayCustom{},
		MsgSliceCustom{},
		MsgMapCustom{},
		MsgArrayUint8Custom{},
		MsgArrayUint8Custom2{},
		MsgArrayInt8Custom{},
		MsgArrayFloat32Custom{},
		MsgArrayFloat64Custom{},
		MsgSliceUint8Custom{},
		MsgSliceUint8Custom2{},
		MsgSliceInt8Custom{},
		MsgSliceFloat32Custom{},
		MsgSliceFloat64Custom{},
		MsgMixedCustom{},
	}
}

// ID returns ID of message type.
func (m Marshaller) ID(msg any) (uint64, error) {
	switch msg.(type) {
	case *pkg2.SubMsg:
		return id37, nil
	case *MsgUint64:
		return id57, nil
	case *MsgUint32:
		return id56, nil
	case *MsgUint16:
		return id55, nil
	case *MsgUint8:
		return id54, nil
	case *MsgInt64:
		return id53, nil
	case *MsgInt32:
		return id52, nil
	case *MsgInt16:
		return id51, nil
	case *MsgInt8:
		return id50, nil
	case *MsgBool3:
		return id49, nil
	case *MsgBool10:
		return id48, nil
	case *MsgFloat64:
		return id47, nil
	case *MsgFloat32:
		return id46, nil
	case *MsgString:
		return id45, nil
	case *MsgArray:
		return id44, nil
	case *MsgSlice:
		return id43, nil
	case *MsgMap:
		return id42, nil
	case *MsgMapString:
		return id41, nil
	case *MsgStruct:
		return id40, nil
	case *MsgStructAnonymous:
		return id38, nil
	case *MsgArrayUint8:
		return id35, nil
	case *MsgArrayInt8:
		return id34, nil
	case *MsgArrayFloat32:
		return id33, nil
	case *MsgArrayFloat64:
		return id32, nil
	case *MsgSliceUint8:
		return id31, nil
	case *MsgSliceInt8:
		return id30, nil
	case *MsgSliceFloat32:
		return id29, nil
	case *MsgSliceFloat64:
		return id28, nil
	case *MsgMixed:
		return id27, nil
	case *MsgUint64Custom:
		return id25, nil
	case *MsgUint32Custom:
		return id24, nil
	case *MsgUint16Custom:
		return id23, nil
	case *MsgUint8Custom:
		return id22, nil
	case *MsgInt64Custom:
		return id21, nil
	case *MsgInt32Custom:
		return id20, nil
	case *MsgInt16Custom:
		return id19, nil
	case *MsgInt8Custom:
		return id18, nil
	case *MsgBoolCustom:
		return id17, nil
	case *MsgFloat64Custom:
		return id16, nil
	case *MsgFloat32Custom:
		return id15, nil
	case *MsgStringCustom:
		return id14, nil
	case *MsgArrayCustom:
		return id13, nil
	case *MsgSliceCustom:
		return id12, nil
	case *MsgMapCustom:
		return id11, nil
	case *MsgArrayUint8Custom:
		return id10, nil
	case *MsgArrayUint8Custom2:
		return id9, nil
	case *MsgArrayInt8Custom:
		return id8, nil
	case *MsgArrayFloat32Custom:
		return id7, nil
	case *MsgArrayFloat64Custom:
		return id6, nil
	case *MsgSliceUint8Custom:
		return id5, nil
	case *MsgSliceUint8Custom2:
		return id4, nil
	case *MsgSliceInt8Custom:
		return id3, nil
	case *MsgSliceFloat32Custom:
		return id2, nil
	case *MsgSliceFloat64Custom:
		return id1, nil
	case *MsgMixedCustom:
		return id0, nil
	default:
		return 0, errors.Errorf("unknown message type %T", msg)
	}
}

// Size computes the size of marshalled message.
func (m Marshaller) Size(msg any) (uint64, error) {
	switch msg2 := msg.(type) {
	case *pkg2.SubMsg:
		return size37(msg2), nil
	case *MsgUint64:
		return size57(msg2), nil
	case *MsgUint32:
		return size56(msg2), nil
	case *MsgUint16:
		return size55(msg2), nil
	case *MsgUint8:
		return size54(msg2), nil
	case *MsgInt64:
		return size53(msg2), nil
	case *MsgInt32:
		return size52(msg2), nil
	case *MsgInt16:
		return size51(msg2), nil
	case *MsgInt8:
		return size50(msg2), nil
	case *MsgBool3:
		return size49(msg2), nil
	case *MsgBool10:
		return size48(msg2), nil
	case *MsgFloat64:
		return size47(msg2), nil
	case *MsgFloat32:
		return size46(msg2), nil
	case *MsgString:
		return size45(msg2), nil
	case *MsgArray:
		return size44(msg2), nil
	case *MsgSlice:
		return size43(msg2), nil
	case *MsgMap:
		return size42(msg2), nil
	case *MsgMapString:
		return size41(msg2), nil
	case *MsgStruct:
		return size40(msg2), nil
	case *MsgStructAnonymous:
		return size38(msg2), nil
	case *MsgArrayUint8:
		return size35(msg2), nil
	case *MsgArrayInt8:
		return size34(msg2), nil
	case *MsgArrayFloat32:
		return size33(msg2), nil
	case *MsgArrayFloat64:
		return size32(msg2), nil
	case *MsgSliceUint8:
		return size31(msg2), nil
	case *MsgSliceInt8:
		return size30(msg2), nil
	case *MsgSliceFloat32:
		return size29(msg2), nil
	case *MsgSliceFloat64:
		return size28(msg2), nil
	case *MsgMixed:
		return size27(msg2), nil
	case *MsgUint64Custom:
		return size25(msg2), nil
	case *MsgUint32Custom:
		return size24(msg2), nil
	case *MsgUint16Custom:
		return size23(msg2), nil
	case *MsgUint8Custom:
		return size22(msg2), nil
	case *MsgInt64Custom:
		return size21(msg2), nil
	case *MsgInt32Custom:
		return size20(msg2), nil
	case *MsgInt16Custom:
		return size19(msg2), nil
	case *MsgInt8Custom:
		return size18(msg2), nil
	case *MsgBoolCustom:
		return size17(msg2), nil
	case *MsgFloat64Custom:
		return size16(msg2), nil
	case *MsgFloat32Custom:
		return size15(msg2), nil
	case *MsgStringCustom:
		return size14(msg2), nil
	case *MsgArrayCustom:
		return size13(msg2), nil
	case *MsgSliceCustom:
		return size12(msg2), nil
	case *MsgMapCustom:
		return size11(msg2), nil
	case *MsgArrayUint8Custom:
		return size10(msg2), nil
	case *MsgArrayUint8Custom2:
		return size9(msg2), nil
	case *MsgArrayInt8Custom:
		return size8(msg2), nil
	case *MsgArrayFloat32Custom:
		return size7(msg2), nil
	case *MsgArrayFloat64Custom:
		return size6(msg2), nil
	case *MsgSliceUint8Custom:
		return size5(msg2), nil
	case *MsgSliceUint8Custom2:
		return size4(msg2), nil
	case *MsgSliceInt8Custom:
		return size3(msg2), nil
	case *MsgSliceFloat32Custom:
		return size2(msg2), nil
	case *MsgSliceFloat64Custom:
		return size1(msg2), nil
	case *MsgMixedCustom:
		return size0(msg2), nil
	default:
		return 0, errors.Errorf("unknown message type %T", msg)
	}
}

// Marshal marshals message.
func (m Marshaller) Marshal(msg any, buf []byte) (retID, retSize uint64, retErr error) {
	defer helpers.RecoverMarshal(&retErr)

	switch msg2 := msg.(type) {
	case *pkg2.SubMsg:
		return id37, marshal37(msg2, buf), nil
	case *MsgUint64:
		return id57, marshal57(msg2, buf), nil
	case *MsgUint32:
		return id56, marshal56(msg2, buf), nil
	case *MsgUint16:
		return id55, marshal55(msg2, buf), nil
	case *MsgUint8:
		return id54, marshal54(msg2, buf), nil
	case *MsgInt64:
		return id53, marshal53(msg2, buf), nil
	case *MsgInt32:
		return id52, marshal52(msg2, buf), nil
	case *MsgInt16:
		return id51, marshal51(msg2, buf), nil
	case *MsgInt8:
		return id50, marshal50(msg2, buf), nil
	case *MsgBool3:
		return id49, marshal49(msg2, buf), nil
	case *MsgBool10:
		return id48, marshal48(msg2, buf), nil
	case *MsgFloat64:
		return id47, marshal47(msg2, buf), nil
	case *MsgFloat32:
		return id46, marshal46(msg2, buf), nil
	case *MsgString:
		return id45, marshal45(msg2, buf), nil
	case *MsgArray:
		return id44, marshal44(msg2, buf), nil
	case *MsgSlice:
		return id43, marshal43(msg2, buf), nil
	case *MsgMap:
		return id42, marshal42(msg2, buf), nil
	case *MsgMapString:
		return id41, marshal41(msg2, buf), nil
	case *MsgStruct:
		return id40, marshal40(msg2, buf), nil
	case *MsgStructAnonymous:
		return id38, marshal38(msg2, buf), nil
	case *MsgArrayUint8:
		return id35, marshal35(msg2, buf), nil
	case *MsgArrayInt8:
		return id34, marshal34(msg2, buf), nil
	case *MsgArrayFloat32:
		return id33, marshal33(msg2, buf), nil
	case *MsgArrayFloat64:
		return id32, marshal32(msg2, buf), nil
	case *MsgSliceUint8:
		return id31, marshal31(msg2, buf), nil
	case *MsgSliceInt8:
		return id30, marshal30(msg2, buf), nil
	case *MsgSliceFloat32:
		return id29, marshal29(msg2, buf), nil
	case *MsgSliceFloat64:
		return id28, marshal28(msg2, buf), nil
	case *MsgMixed:
		return id27, marshal27(msg2, buf), nil
	case *MsgUint64Custom:
		return id25, marshal25(msg2, buf), nil
	case *MsgUint32Custom:
		return id24, marshal24(msg2, buf), nil
	case *MsgUint16Custom:
		return id23, marshal23(msg2, buf), nil
	case *MsgUint8Custom:
		return id22, marshal22(msg2, buf), nil
	case *MsgInt64Custom:
		return id21, marshal21(msg2, buf), nil
	case *MsgInt32Custom:
		return id20, marshal20(msg2, buf), nil
	case *MsgInt16Custom:
		return id19, marshal19(msg2, buf), nil
	case *MsgInt8Custom:
		return id18, marshal18(msg2, buf), nil
	case *MsgBoolCustom:
		return id17, marshal17(msg2, buf), nil
	case *MsgFloat64Custom:
		return id16, marshal16(msg2, buf), nil
	case *MsgFloat32Custom:
		return id15, marshal15(msg2, buf), nil
	case *MsgStringCustom:
		return id14, marshal14(msg2, buf), nil
	case *MsgArrayCustom:
		return id13, marshal13(msg2, buf), nil
	case *MsgSliceCustom:
		return id12, marshal12(msg2, buf), nil
	case *MsgMapCustom:
		return id11, marshal11(msg2, buf), nil
	case *MsgArrayUint8Custom:
		return id10, marshal10(msg2, buf), nil
	case *MsgArrayUint8Custom2:
		return id9, marshal9(msg2, buf), nil
	case *MsgArrayInt8Custom:
		return id8, marshal8(msg2, buf), nil
	case *MsgArrayFloat32Custom:
		return id7, marshal7(msg2, buf), nil
	case *MsgArrayFloat64Custom:
		return id6, marshal6(msg2, buf), nil
	case *MsgSliceUint8Custom:
		return id5, marshal5(msg2, buf), nil
	case *MsgSliceUint8Custom2:
		return id4, marshal4(msg2, buf), nil
	case *MsgSliceInt8Custom:
		return id3, marshal3(msg2, buf), nil
	case *MsgSliceFloat32Custom:
		return id2, marshal2(msg2, buf), nil
	case *MsgSliceFloat64Custom:
		return id1, marshal1(msg2, buf), nil
	case *MsgMixedCustom:
		return id0, marshal0(msg2, buf), nil
	default:
		return 0, 0, errors.Errorf("unknown message type %T", msg)
	}
}

// Unmarshal unmarshals message.
func (m Marshaller) Unmarshal(id uint64, buf []byte) (retMsg any, retSize uint64, retErr error) {
	defer helpers.RecoverUnmarshal(&retErr)

	switch id {
	case id37:
		msg := &pkg2.SubMsg{}
		return msg, unmarshal37(msg, buf), nil
	case id57:
		msg := &MsgUint64{}
		return msg, unmarshal57(msg, buf), nil
	case id56:
		msg := &MsgUint32{}
		return msg, unmarshal56(msg, buf), nil
	case id55:
		msg := &MsgUint16{}
		return msg, unmarshal55(msg, buf), nil
	case id54:
		msg := &MsgUint8{}
		return msg, unmarshal54(msg, buf), nil
	case id53:
		msg := &MsgInt64{}
		return msg, unmarshal53(msg, buf), nil
	case id52:
		msg := &MsgInt32{}
		return msg, unmarshal52(msg, buf), nil
	case id51:
		msg := &MsgInt16{}
		return msg, unmarshal51(msg, buf), nil
	case id50:
		msg := &MsgInt8{}
		return msg, unmarshal50(msg, buf), nil
	case id49:
		msg := &MsgBool3{}
		return msg, unmarshal49(msg, buf), nil
	case id48:
		msg := &MsgBool10{}
		return msg, unmarshal48(msg, buf), nil
	case id47:
		msg := &MsgFloat64{}
		return msg, unmarshal47(msg, buf), nil
	case id46:
		msg := &MsgFloat32{}
		return msg, unmarshal46(msg, buf), nil
	case id45:
		msg := &MsgString{}
		return msg, unmarshal45(msg, buf), nil
	case id44:
		msg := &MsgArray{}
		return msg, unmarshal44(msg, buf), nil
	case id43:
		msg := &MsgSlice{}
		return msg, unmarshal43(msg, buf), nil
	case id42:
		msg := &MsgMap{}
		return msg, unmarshal42(msg, buf), nil
	case id41:
		msg := &MsgMapString{}
		return msg, unmarshal41(msg, buf), nil
	case id40:
		msg := &MsgStruct{}
		return msg, unmarshal40(msg, buf), nil
	case id38:
		msg := &MsgStructAnonymous{}
		return msg, unmarshal38(msg, buf), nil
	case id35:
		msg := &MsgArrayUint8{}
		return msg, unmarshal35(msg, buf), nil
	case id34:
		msg := &MsgArrayInt8{}
		return msg, unmarshal34(msg, buf), nil
	case id33:
		msg := &MsgArrayFloat32{}
		return msg, unmarshal33(msg, buf), nil
	case id32:
		msg := &MsgArrayFloat64{}
		return msg, unmarshal32(msg, buf), nil
	case id31:
		msg := &MsgSliceUint8{}
		return msg, unmarshal31(msg, buf), nil
	case id30:
		msg := &MsgSliceInt8{}
		return msg, unmarshal30(msg, buf), nil
	case id29:
		msg := &MsgSliceFloat32{}
		return msg, unmarshal29(msg, buf), nil
	case id28:
		msg := &MsgSliceFloat64{}
		return msg, unmarshal28(msg, buf), nil
	case id27:
		msg := &MsgMixed{}
		return msg, unmarshal27(msg, buf), nil
	case id25:
		msg := &MsgUint64Custom{}
		return msg, unmarshal25(msg, buf), nil
	case id24:
		msg := &MsgUint32Custom{}
		return msg, unmarshal24(msg, buf), nil
	case id23:
		msg := &MsgUint16Custom{}
		return msg, unmarshal23(msg, buf), nil
	case id22:
		msg := &MsgUint8Custom{}
		return msg, unmarshal22(msg, buf), nil
	case id21:
		msg := &MsgInt64Custom{}
		return msg, unmarshal21(msg, buf), nil
	case id20:
		msg := &MsgInt32Custom{}
		return msg, unmarshal20(msg, buf), nil
	case id19:
		msg := &MsgInt16Custom{}
		return msg, unmarshal19(msg, buf), nil
	case id18:
		msg := &MsgInt8Custom{}
		return msg, unmarshal18(msg, buf), nil
	case id17:
		msg := &MsgBoolCustom{}
		return msg, unmarshal17(msg, buf), nil
	case id16:
		msg := &MsgFloat64Custom{}
		return msg, unmarshal16(msg, buf), nil
	case id15:
		msg := &MsgFloat32Custom{}
		return msg, unmarshal15(msg, buf), nil
	case id14:
		msg := &MsgStringCustom{}
		return msg, unmarshal14(msg, buf), nil
	case id13:
		msg := &MsgArrayCustom{}
		return msg, unmarshal13(msg, buf), nil
	case id12:
		msg := &MsgSliceCustom{}
		return msg, unmarshal12(msg, buf), nil
	case id11:
		msg := &MsgMapCustom{}
		return msg, unmarshal11(msg, buf), nil
	case id10:
		msg := &MsgArrayUint8Custom{}
		return msg, unmarshal10(msg, buf), nil
	case id9:
		msg := &MsgArrayUint8Custom2{}
		return msg, unmarshal9(msg, buf), nil
	case id8:
		msg := &MsgArrayInt8Custom{}
		return msg, unmarshal8(msg, buf), nil
	case id7:
		msg := &MsgArrayFloat32Custom{}
		return msg, unmarshal7(msg, buf), nil
	case id6:
		msg := &MsgArrayFloat64Custom{}
		return msg, unmarshal6(msg, buf), nil
	case id5:
		msg := &MsgSliceUint8Custom{}
		return msg, unmarshal5(msg, buf), nil
	case id4:
		msg := &MsgSliceUint8Custom2{}
		return msg, unmarshal4(msg, buf), nil
	case id3:
		msg := &MsgSliceInt8Custom{}
		return msg, unmarshal3(msg, buf), nil
	case id2:
		msg := &MsgSliceFloat32Custom{}
		return msg, unmarshal2(msg, buf), nil
	case id1:
		msg := &MsgSliceFloat64Custom{}
		return msg, unmarshal1(msg, buf), nil
	case id0:
		msg := &MsgMixedCustom{}
		return msg, unmarshal0(msg, buf), nil
	default:
		return nil, 0, errors.Errorf("unknown ID %d", id)
	}
}

func size0(m *MsgMixedCustom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l * 3
		for _, sv6 := range m.Value {
			for _, av5 := range sv6 {
				l := uint64(len(av5))
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
				n += l
				for _, sv4 := range av5 {
					l := uint64(len(sv4))
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
					n += l * 3
					for mk2, mv3 := range sv4 {
						{
							vi := uint16(mk2) << 1
							if mk2 < 0 {
								vi ^= 0xFFFF
							}
							switch {
							case vi <= 0x7F:
							case vi <= 0x3FFF:
								n++
							default:
								n += 2
							}
						}
						for _, av1 := range mv3 {
							{
								vi := av1
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
			}
		}
	}
	return n
}

func marshal0(m *MsgMixedCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint64(len(m.Value))
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
		for _, sv6 := range m.Value {
			for _, av5 := range sv6 {
				{
					vi := uint64(len(av5))
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
				for _, sv4 := range av5 {
					{
						vi := uint64(len(sv4))
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
					for mk2, mv3 := range sv4 {
						{
							vi := uint16(mk2) << 1
							if mk2 < 0 {
								vi ^= 0xFFFF
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
						for _, av1 := range mv3 {
							{
								vi := av1
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
						}
					}
				}
			}
		}
	}

	return o
}

func unmarshal0(m *MsgMixedCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([][3][]map[int16]custom.Array, l)
			for i6 := range l {
				for i5 := range 3 {
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
						m.Value[i6][i5] = make([]map[int16]custom.Array, l)
						for i4 := range l {
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
								m.Value[i6][i5][i4] = make(map[int16]custom.Array, l)
							
								var mk2 int16
								var mv3 custom.Array
							
								for range l {
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
										if vi&0x01 == 0 {
											vi >>= 1
										} else {
											vi >>= 1
											vi ^= 0xFFFF
										}
										mk2 = int16(vi)
									}
									for i1 := range 2 {
										{
											vi := custom.Uint64(b[o] & 0x7F)
											if b[o]&0x80 == 0 {
												o++
											} else {
												vi |= custom.Uint64(b[o+1]&0x7F) << 7
												if b[o+1]&0x80 == 0 {
													o += 2
												} else {
													vi |= custom.Uint64(b[o+2]&0x7F) << 14
													if b[o+2]&0x80 == 0 {
														o += 3
													} else {
														vi |= custom.Uint64(b[o+3]&0x7F) << 21
														if b[o+3]&0x80 == 0 {
															o += 4
														} else {
															vi |= custom.Uint64(b[o+4]&0x7F) << 28
															if b[o+4]&0x80 == 0 {
																o += 5
															} else {
																vi |= custom.Uint64(b[o+5]&0x7F) << 35
																if b[o+5]&0x80 == 0 {
																	o += 6
																} else {
																	vi |= custom.Uint64(b[o+6]&0x7F) << 42
																	if b[o+6]&0x80 == 0 {
																		o += 7
																	} else {
																		vi |= custom.Uint64(b[o+7]&0x7F) << 49
																		if b[o+7]&0x80 == 0 {
																			o += 8
																		} else {
																			vi |= custom.Uint64(b[o+8]) << 56
																			o += 9
																		}
																	}
																}
															}
														}
													}
												}
											}
											mv3[i1] = vi
										}
									}
									m.Value[i6][i5][i4][mk2] = mv3
								}
							}
						}
					}
				}
			}
		}
	}

	return o
}

func size1(m *MsgSliceFloat64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l * 8
	}
	return n
}

func marshal1(m *MsgSliceFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l*8], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8))
			o += l * 8
		}
	}

	return o
}

func unmarshal1(m *MsgSliceFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]custom.Float64, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
			o += l * 8
		}
	}

	return o
}

func size2(m *MsgSliceFloat32Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l * 4
	}
	return n
}

func marshal2(m *MsgSliceFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l*4], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4))
			o += l * 4
		}
	}

	return o
}

func unmarshal2(m *MsgSliceFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]custom.Float32, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
			o += l * 4
		}
	}

	return o
}

func size3(m *MsgSliceInt8Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l
	}
	return n
}

func marshal3(m *MsgSliceInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l))
			o += l
		}
	}

	return o
}

func unmarshal3(m *MsgSliceInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]custom.Int8, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
			o += l
		}
	}

	return o
}

func size4(m *MsgSliceUint8Custom2) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l
	}
	return n
}

func marshal4(m *MsgSliceUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice((*byte)(&m.Value[0]), l))
			o += l
		}
	}

	return o
}

func unmarshal4(m *MsgSliceUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]custom.Uint8, l)
			copy(unsafe.Slice((*byte)(&m.Value[0]), l), b[o:o+l])
			o += l
		}
	}

	return o
}

func size5(m *MsgSliceUint8Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l
	}
	return n
}

func marshal5(m *MsgSliceUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice(&m.Value[0], l))
			o += l
		}
	}

	return o
}

func unmarshal5(m *MsgSliceUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]uint8, l)
			copy(m.Value, b[o:o+l])
			o += l
		}
	}

	return o
}

func size6(m *MsgArrayFloat64Custom) uint64 {
	var n uint64 = 40
	return n
}

func marshal6(m *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

func unmarshal6(m *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
		o += 40
	}

	return o
}

func size7(m *MsgArrayFloat32Custom) uint64 {
	var n uint64 = 20
	return n
}

func marshal7(m *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

func unmarshal7(m *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
		o += 20
	}

	return o
}

func size8(m *MsgArrayInt8Custom) uint64 {
	var n uint64 = 5
	return n
}

func marshal8(m *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

func unmarshal8(m *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size9(m *MsgArrayUint8Custom2) uint64 {
	var n uint64 = 5
	return n
}

func marshal9(m *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(&m.Value[0]), 5))
		o += 5
	}

	return o
}

func unmarshal9(m *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(&m.Value[0]), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size10(m *MsgArrayUint8Custom) uint64 {
	var n uint64 = 5
	return n
}

func marshal10(m *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

func unmarshal10(m *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
		o += 5
	}

	return o
}

func size11(m *MsgMapCustom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		for mk1, mv2 := range m.Value {
			{
				l := uint64(len(mk1))
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
			{
				vi := uint32(mv2) << 1
				if mv2 < 0 {
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
	}
	return n
}

func marshal11(m *MsgMapCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint64(len(m.Value))
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
		for mk1, mv2 := range m.Value {
			{
				l := uint64(len(mk1))
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
				copy(b[o:o+l], mk1)
				o += l
			}
			{
				vi := uint32(mv2) << 1
				if mv2 < 0 {
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
	}

	return o
}

func unmarshal11(m *MsgMapCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make(custom.Map, l)
		
			var mk1 custom.String
			var mv2 custom.Int32
		
			for range l {
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
						mk1 = custom.String(b[o:o+l])
						o += l
					}
				}
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
					mv2 = custom.Int32(vi)
				}
				m.Value[mk1] = mv2
			}
		}
	}

	return o
}

func size12(m *MsgSliceCustom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l
		for _, sv1 := range m.Value {
			{
				vi := sv1
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
	return n
}

func marshal12(m *MsgSliceCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint64(len(m.Value))
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
		for _, sv1 := range m.Value {
			{
				vi := sv1
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
		}
	}

	return o
}

func unmarshal12(m *MsgSliceCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]custom.Uint64, l)
			for i1 := range l {
				{
					vi := custom.Uint64(b[o] & 0x7F)
					if b[o]&0x80 == 0 {
						o++
					} else {
						vi |= custom.Uint64(b[o+1]&0x7F) << 7
						if b[o+1]&0x80 == 0 {
							o += 2
						} else {
							vi |= custom.Uint64(b[o+2]&0x7F) << 14
							if b[o+2]&0x80 == 0 {
								o += 3
							} else {
								vi |= custom.Uint64(b[o+3]&0x7F) << 21
								if b[o+3]&0x80 == 0 {
									o += 4
								} else {
									vi |= custom.Uint64(b[o+4]&0x7F) << 28
									if b[o+4]&0x80 == 0 {
										o += 5
									} else {
										vi |= custom.Uint64(b[o+5]&0x7F) << 35
										if b[o+5]&0x80 == 0 {
											o += 6
										} else {
											vi |= custom.Uint64(b[o+6]&0x7F) << 42
											if b[o+6]&0x80 == 0 {
												o += 7
											} else {
												vi |= custom.Uint64(b[o+7]&0x7F) << 49
												if b[o+7]&0x80 == 0 {
													o += 8
												} else {
													vi |= custom.Uint64(b[o+8]) << 56
													o += 9
												}
											}
										}
									}
								}
							}
						}
					}
					m.Value[i1] = vi
				}
			}
		}
	}

	return o
}

func size13(m *MsgArrayCustom) uint64 {
	var n uint64 = 2
	{
		// Value

		for _, av1 := range m.Value {
			{
				vi := av1
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
	return n
}

func marshal13(m *MsgArrayCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		for _, av1 := range m.Value {
			{
				vi := av1
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
		}
	}

	return o
}

func unmarshal13(m *MsgArrayCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		for i1 := range 2 {
			{
				vi := custom.Uint64(b[o] & 0x7F)
				if b[o]&0x80 == 0 {
					o++
				} else {
					vi |= custom.Uint64(b[o+1]&0x7F) << 7
					if b[o+1]&0x80 == 0 {
						o += 2
					} else {
						vi |= custom.Uint64(b[o+2]&0x7F) << 14
						if b[o+2]&0x80 == 0 {
							o += 3
						} else {
							vi |= custom.Uint64(b[o+3]&0x7F) << 21
							if b[o+3]&0x80 == 0 {
								o += 4
							} else {
								vi |= custom.Uint64(b[o+4]&0x7F) << 28
								if b[o+4]&0x80 == 0 {
									o += 5
								} else {
									vi |= custom.Uint64(b[o+5]&0x7F) << 35
									if b[o+5]&0x80 == 0 {
										o += 6
									} else {
										vi |= custom.Uint64(b[o+6]&0x7F) << 42
										if b[o+6]&0x80 == 0 {
											o += 7
										} else {
											vi |= custom.Uint64(b[o+7]&0x7F) << 49
											if b[o+7]&0x80 == 0 {
												o += 8
											} else {
												vi |= custom.Uint64(b[o+8]) << 56
												o += 9
											}
										}
									}
								}
							}
						}
					}
				}
				m.Value[i1] = vi
			}
		}
	}

	return o
}

func size14(m *MsgStringCustom) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			l := uint64(len(m.Value))
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
	return n
}

func marshal14(m *MsgStringCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			l := uint64(len(m.Value))
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
			copy(b[o:o+l], m.Value)
			o += l
		}
	}

	return o
}

func unmarshal14(m *MsgStringCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
				m.Value = custom.String(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size15(m *MsgFloat32Custom) uint64 {
	var n uint64 = 4
	return n
}

func marshal15(m *MsgFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

func unmarshal15(m *MsgFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*custom.Float32)(unsafe.Pointer(&b[o]))
		o += 4
	}

	return o
}

func size16(m *MsgFloat64Custom) uint64 {
	var n uint64 = 8
	return n
}

func marshal16(m *MsgFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

func unmarshal16(m *MsgFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*custom.Float64)(unsafe.Pointer(&b[o]))
		o += 8
	}

	return o
}

func size17(m *MsgBoolCustom) uint64 {
	var n uint64 = 1
	return n
}

func marshal17(m *MsgBoolCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if m.Value {
			b[0] |= 0x01
		} else {
			b[0] &= 0xFE
		}
	}

	return o
}

func unmarshal17(m *MsgBoolCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		m.Value = b[0]&0x01 != 0
	}

	return o
}

func size18(m *MsgInt8Custom) uint64 {
	var n uint64 = 1
	return n
}

func marshal18(m *MsgInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal18(m *MsgInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = custom.Int8(b[o])
		o++
	}

	return o
}

func size19(m *MsgInt16Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := uint16(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFF
			}
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

func marshal19(m *MsgInt16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint16(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFF
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

func unmarshal19(m *MsgInt16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			if vi&0x01 == 0 {
				vi >>= 1
			} else {
				vi >>= 1
				vi ^= 0xFFFF
			}
			m.Value = custom.Int16(vi)
		}
	}

	return o
}

func size20(m *MsgInt32Custom) uint64 {
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

func marshal20(m *MsgInt32Custom, b []byte) uint64 {
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

func unmarshal20(m *MsgInt32Custom, b []byte) uint64 {
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
			m.Value = custom.Int32(vi)
		}
	}

	return o
}

func size21(m *MsgInt64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := uint64(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFFFFFFFFFFFFFF
			}
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
			case vi <= 0x7FFFFFFFFFFFFFFF:
				n += 8
			default:
				n += 9
			}
		}
	}
	return n
}

func marshal21(m *MsgInt64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint64(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFFFFFFFFFFFFFF
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
			case vi <= 0x7FFFFFFFFFFFFFFF:
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
				vi >>= 7
				b[o] = byte(vi)
				o++
			}
		}
	}

	return o
}

func unmarshal21(m *MsgInt64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

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
											vi |= uint64(b[o+8]&0x7F) << 56
											if b[o+8]&0x80 == 0 {
												o += 9
											} else {
												vi |= uint64(b[o+9]) << 63
												o += 10
											}
										}
									}
								}
							}
						}
					}
				}
			}
			if vi&0x01 == 0 {
				vi >>= 1
			} else {
				vi >>= 1
				vi ^= 0xFFFFFFFFFFFFFFFF
			}
			m.Value = custom.Int64(vi)
		}
	}

	return o
}

func size22(m *MsgUint8Custom) uint64 {
	var n uint64 = 1
	return n
}

func marshal22(m *MsgUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal22(m *MsgUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = custom.Uint8(b[o])
		o++
	}

	return o
}

func size23(m *MsgUint16Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := m.Value
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

func marshal23(m *MsgUint16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := m.Value
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

func unmarshal23(m *MsgUint16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := custom.Uint16(b[o] & 0x7F)
			if b[o]&0x80 == 0 {
				o++
			} else {
				vi |= custom.Uint16(b[o+1]&0x7F) << 7
				if b[o+1]&0x80 == 0 {
					o += 2
				} else {
					vi |= custom.Uint16(b[o+2]) << 14
					o += 3
				}
			}
			m.Value = vi
		}
	}

	return o
}

func size24(m *MsgUint32Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := m.Value
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

func marshal24(m *MsgUint32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := m.Value
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

func unmarshal24(m *MsgUint32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := custom.Uint32(b[o] & 0x7F)
			if b[o]&0x80 == 0 {
				o++
			} else {
				vi |= custom.Uint32(b[o+1]&0x7F) << 7
				if b[o+1]&0x80 == 0 {
					o += 2
				} else {
					vi |= custom.Uint32(b[o+2]&0x7F) << 14
					if b[o+2]&0x80 == 0 {
						o += 3
					} else {
						vi |= custom.Uint32(b[o+3]&0x7F) << 21
						if b[o+3]&0x80 == 0 {
							o += 4
						} else {
							vi |= custom.Uint32(b[o+4]) << 28
							o += 5
						}
					}
				}
			}
			m.Value = vi
		}
	}

	return o
}

func size25(m *MsgUint64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := m.Value
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
	return n
}

func marshal25(m *MsgUint64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := m.Value
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
	}

	return o
}

func unmarshal25(m *MsgUint64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := custom.Uint64(b[o] & 0x7F)
			if b[o]&0x80 == 0 {
				o++
			} else {
				vi |= custom.Uint64(b[o+1]&0x7F) << 7
				if b[o+1]&0x80 == 0 {
					o += 2
				} else {
					vi |= custom.Uint64(b[o+2]&0x7F) << 14
					if b[o+2]&0x80 == 0 {
						o += 3
					} else {
						vi |= custom.Uint64(b[o+3]&0x7F) << 21
						if b[o+3]&0x80 == 0 {
							o += 4
						} else {
							vi |= custom.Uint64(b[o+4]&0x7F) << 28
							if b[o+4]&0x80 == 0 {
								o += 5
							} else {
								vi |= custom.Uint64(b[o+5]&0x7F) << 35
								if b[o+5]&0x80 == 0 {
									o += 6
								} else {
									vi |= custom.Uint64(b[o+6]&0x7F) << 42
									if b[o+6]&0x80 == 0 {
										o += 7
									} else {
										vi |= custom.Uint64(b[o+7]&0x7F) << 49
										if b[o+7]&0x80 == 0 {
											o += 8
										} else {
											vi |= custom.Uint64(b[o+8]) << 56
											o += 9
										}
									}
								}
							}
						}
					}
				}
			}
			m.Value = vi
		}
	}

	return o
}

func size27(m *MsgMixed) uint64 {
	var n uint64 = 18
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
		n += l * 1
		for mk1, mv2 := range m.Value1 {
			{
				l := uint64(len(mk1))
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
			n += size26(&mv2)
		}
	}
	{
		// Value2

		l := uint64(len(m.Value2))
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
		for _, mv3 := range m.Value2 {
			l := uint64(len(mv3))
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
			n += l
			for _, sv1 := range mv3 {
				{
					l := uint64(len(sv1))
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
	}
	{
		// Value3

		l := uint64(len(m.Value3))
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
		n += l * 32
		for _, sv2 := range m.Value3 {
			for _, av1 := range sv2 {
				{
					vi := av1
					switch {
					case vi <= 0x7F:
					case vi <= 0x3FFF:
						n++
					default:
						n += 2
					}
				}
			}
		}
	}
	{
		// Value4

		for _, av1 := range m.Value4 {
			l := uint64(len(av1))
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
			n += l * 5
		}
	}
	{
		// Value5

		l := uint64(len(m.Value5))
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
		n += l * 3
		for _, sv6 := range m.Value5 {
			for _, av5 := range sv6 {
				l := uint64(len(av5))
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
				n += l
				for _, sv4 := range av5 {
					l := uint64(len(sv4))
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
					n += l * 3
					for mk2, mv3 := range sv4 {
						{
							vi := uint16(mk2) << 1
							if mk2 < 0 {
								vi ^= 0xFFFF
							}
							switch {
							case vi <= 0x7F:
							case vi <= 0x3FFF:
								n++
							default:
								n += 2
							}
						}
						for _, av1 := range mv3 {
							{
								vi := uint64(av1) << 1
								if av1 < 0 {
									vi ^= 0xFFFFFFFFFFFFFFFF
								}
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
								case vi <= 0x7FFFFFFFFFFFFFFF:
									n += 8
								default:
									n += 9
								}
							}
						}
					}
				}
			}
		}
	}
	{
		// Value8

		{
			l := uint64(len(m.Value8))
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
	return n
}

func marshal27(m *MsgMixed, b []byte) uint64 {
	var o uint64 = 1
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
			{
				l := uint64(len(mk1))
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
				copy(b[o:o+l], mk1)
				o += l
			}
			o += marshal26(&mv2, b[o:])
		}
	}
	{
		// Value2

		{
			vi := uint64(len(m.Value2))
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
		for mk2, mv3 := range m.Value2 {
			b[o] = mk2
			o++
			{
				vi := uint64(len(mv3))
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
			for _, sv1 := range mv3 {
				{
					l := uint64(len(sv1))
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
					copy(b[o:o+l], sv1)
					o += l
				}
			}
		}
	}
	{
		// Value3

		{
			vi := uint64(len(m.Value3))
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
		for _, sv2 := range m.Value3 {
			for _, av1 := range sv2 {
				{
					vi := av1
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
		}
	}
	{
		// Value4

		for _, av3 := range m.Value4 {
			{
				vi := uint64(len(av3))
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
			for mk1, mv2 := range av3 {
				b[o] = byte(mk1)
				o++
				*(*float32)(unsafe.Pointer(&b[o])) = mv2
				o += 4
			}
		}
	}
	{
		// Value5

		{
			vi := uint64(len(m.Value5))
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
		for _, sv6 := range m.Value5 {
			for _, av5 := range sv6 {
				{
					vi := uint64(len(av5))
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
				for _, sv4 := range av5 {
					{
						vi := uint64(len(sv4))
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
					for mk2, mv3 := range sv4 {
						{
							vi := uint16(mk2) << 1
							if mk2 < 0 {
								vi ^= 0xFFFF
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
						for _, av1 := range mv3 {
							{
								vi := uint64(av1) << 1
								if av1 < 0 {
									vi ^= 0xFFFFFFFFFFFFFFFF
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
								case vi <= 0x7FFFFFFFFFFFFFFF:
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
									vi >>= 7
									b[o] = byte(vi)
									o++
								}
							}
						}
					}
				}
			}
		}
	}
	{
		// Value6

		if m.Value6 {
			b[0] |= 0x01
		} else {
			b[0] &= 0xFE
		}
	}
	{
		// Value7

		if m.Value7 {
			b[0] |= 0x02
		} else {
			b[0] &= 0xFD
		}
	}
	{
		// Value8

		{
			l := uint64(len(m.Value8))
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
			copy(b[o:o+l], m.Value8)
			o += l
		}
	}

	return o
}

func unmarshal27(m *MsgMixed, b []byte) uint64 {
	var o uint64 = 1
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
			m.Value1 = make(map[string]spkg.SubMsg, l)
		
			var mk1 string
			var mv2 spkg.SubMsg
		
			for range l {
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
						mk1 = string(b[o:o+l])
						o += l
					}
				}
				o += unmarshal26(&mv2, b[o:])
				
				m.Value1[mk1] = mv2
			}
		}
	}
	{
		// Value2

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
			m.Value2 = make(map[uint8][]string, l)
		
			var mk2 uint8
			var mv3 []string
		
			for range l {
				mk2 = b[o]
				o++
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
					mv3 = make([]string, l)
					for i1 := range l {
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
								mv3[i1] = string(b[o:o+l])
								o += l
							}
						}
					}
				}
				m.Value2[mk2] = mv3
			}
		}
	}
	{
		// Value3

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
			m.Value3 = make([][32]uint16, l)
			for i2 := range l {
				for i1 := range 32 {
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
						m.Value3[i2][i1] = vi
					}
				}
			}
		}
	}
	{
		// Value4

		for i3 := range 12 {
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
				m.Value4[i3] = make(map[int8]float32, l)
			
				var mk1 int8
				var mv2 float32
			
				for range l {
					mk1 = int8(b[o])
					o++
					mv2 = *(*float32)(unsafe.Pointer(&b[o]))
					o += 4
					m.Value4[i3][mk1] = mv2
				}
			}
		}
	}
	{
		// Value5

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
			m.Value5 = make([][3][]map[int16][2]int64, l)
			for i6 := range l {
				for i5 := range 3 {
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
						m.Value5[i6][i5] = make([]map[int16][2]int64, l)
						for i4 := range l {
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
								m.Value5[i6][i5][i4] = make(map[int16][2]int64, l)
							
								var mk2 int16
								var mv3 [2]int64
							
								for range l {
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
										if vi&0x01 == 0 {
											vi >>= 1
										} else {
											vi >>= 1
											vi ^= 0xFFFF
										}
										mk2 = int16(vi)
									}
									for i1 := range 2 {
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
																			vi |= uint64(b[o+8]&0x7F) << 56
																			if b[o+8]&0x80 == 0 {
																				o += 9
																			} else {
																				vi |= uint64(b[o+9]) << 63
																				o += 10
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
											if vi&0x01 == 0 {
												vi >>= 1
											} else {
												vi >>= 1
												vi ^= 0xFFFFFFFFFFFFFFFF
											}
											mv3[i1] = int64(vi)
										}
									}
									m.Value5[i6][i5][i4][mk2] = mv3
								}
							}
						}
					}
				}
			}
		}
	}
	{
		// Value6

		m.Value6 = b[0]&0x01 != 0
	}
	{
		// Value7

		m.Value7 = b[0]&0x02 != 0
	}
	{
		// Value8

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
				m.Value8 = string(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size26(m *spkg.SubMsg) uint64 {
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

func marshal26(m *spkg.SubMsg, b []byte) uint64 {
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

func unmarshal26(m *spkg.SubMsg, b []byte) uint64 {
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

func size28(m *MsgSliceFloat64) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l * 8
	}
	return n
}

func marshal28(m *MsgSliceFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l*8], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8))
			o += l * 8
		}
	}

	return o
}

func unmarshal28(m *MsgSliceFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]float64, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
			o += l * 8
		}
	}

	return o
}

func size29(m *MsgSliceFloat32) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l * 4
	}
	return n
}

func marshal29(m *MsgSliceFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l*4], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4))
			o += l * 4
		}
	}

	return o
}

func unmarshal29(m *MsgSliceFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]float32, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
			o += l * 4
		}
	}

	return o
}

func size30(m *MsgSliceInt8) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l
	}
	return n
}

func marshal30(m *MsgSliceInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l))
			o += l
		}
	}

	return o
}

func unmarshal30(m *MsgSliceInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]int8, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
			o += l
		}
	}

	return o
}

func size31(m *MsgSliceUint8) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l
	}
	return n
}

func marshal31(m *MsgSliceUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
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
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice(&m.Value[0], l))
			o += l
		}
	}

	return o
}

func unmarshal31(m *MsgSliceUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]uint8, l)
			copy(m.Value, b[o:o+l])
			o += l
		}
	}

	return o
}

func size32(m *MsgArrayFloat64) uint64 {
	var n uint64 = 40
	return n
}

func marshal32(m *MsgArrayFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

func unmarshal32(m *MsgArrayFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
		o += 40
	}

	return o
}

func size33(m *MsgArrayFloat32) uint64 {
	var n uint64 = 20
	return n
}

func marshal33(m *MsgArrayFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

func unmarshal33(m *MsgArrayFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
		o += 20
	}

	return o
}

func size34(m *MsgArrayInt8) uint64 {
	var n uint64 = 5
	return n
}

func marshal34(m *MsgArrayInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

func unmarshal34(m *MsgArrayInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size35(m *MsgArrayUint8) uint64 {
	var n uint64 = 5
	return n
}

func marshal35(m *MsgArrayUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

func unmarshal35(m *MsgArrayUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
		o += 5
	}

	return o
}

func size38(m *MsgStructAnonymous) uint64 {
	var n uint64
	{
		// SubMsg

		n += size36(&m.SubMsg)
	}
	{
		// Value2

		n += size37(&m.Value2)
	}
	return n
}

func marshal38(m *MsgStructAnonymous, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += marshal36(&m.SubMsg, b[o:])
	}
	{
		// Value2

		o += marshal37(&m.Value2, b[o:])
	}

	return o
}

func unmarshal38(m *MsgStructAnonymous, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += unmarshal36(&m.SubMsg, b[o:])
		
	}
	{
		// Value2

		o += unmarshal37(&m.Value2, b[o:])
		
	}

	return o
}

func size36(m *SubMsg) uint64 {
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
		n += l * 4
	}
	{
		// Value2

		l := uint64(len(m.Value2))
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
		for _, sv1 := range m.Value2 {
			n += size26(&sv1)
		}
	}
	{
		// Value3

		l := uint64(len(m.Value3))
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
		for _, sv1 := range m.Value3 {
			n += size39(&sv1)
		}
	}
	return n
}

func marshal36(m *SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value1

		l := uint64(len(m.Value1))
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
		if l > 0 {
			copy(b[o:o+l*4], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value1[0])), l*4))
			o += l * 4
		}
	}
	{
		// Value2

		{
			vi := uint64(len(m.Value2))
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
		for _, sv1 := range m.Value2 {
			o += marshal26(&sv1, b[o:])
		}
	}
	{
		// Value3

		{
			vi := uint64(len(m.Value3))
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
		for _, sv1 := range m.Value3 {
			o += marshal39(&sv1, b[o:])
		}
	}

	return o
}

func unmarshal36(m *SubMsg, b []byte) uint64 {
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
			m.Value1 = make([]float32, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value1[0])), l*4), b[o:o+l*4])
			o += l * 4
		}
	}
	{
		// Value2

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
			m.Value2 = make([]spkg.SubMsg, l)
			for i1 := range l {
				o += unmarshal26(&m.Value2[i1], b[o:])
				
			}
		}
	}
	{
		// Value3

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
			m.Value3 = make([]spkg6.SubMsg, l)
			for i1 := range l {
				o += unmarshal39(&m.Value3[i1], b[o:])
				
			}
		}
	}

	return o
}

func size39(m *spkg6.SubMsg) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := uint16(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFF
			}
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

func marshal39(m *spkg6.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint16(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFF
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

func unmarshal39(m *spkg6.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			if vi&0x01 == 0 {
				vi >>= 1
			} else {
				vi >>= 1
				vi ^= 0xFFFF
			}
			m.Value = int16(vi)
		}
	}

	return o
}

func size40(m *MsgStruct) uint64 {
	var n uint64
	{
		// Value1

		n += size36(&m.Value1)
	}
	{
		// Value2

		n += size37(&m.Value2)
	}
	return n
}

func marshal40(m *MsgStruct, b []byte) uint64 {
	var o uint64
	{
		// Value1

		o += marshal36(&m.Value1, b[o:])
	}
	{
		// Value2

		o += marshal37(&m.Value2, b[o:])
	}

	return o
}

func unmarshal40(m *MsgStruct, b []byte) uint64 {
	var o uint64
	{
		// Value1

		o += unmarshal36(&m.Value1, b[o:])
		
	}
	{
		// Value2

		o += unmarshal37(&m.Value2, b[o:])
		
	}

	return o
}

func size41(m *MsgMapString) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		for mk1, mv2 := range m.Value {
			{
				l := uint64(len(mk1))
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
	return n
}

func marshal41(m *MsgMapString, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint64(len(m.Value))
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
		for mk1, mv2 := range m.Value {
			{
				l := uint64(len(mk1))
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
				copy(b[o:o+l], mk1)
				o += l
			}
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

	return o
}

func unmarshal41(m *MsgMapString, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make(map[string]string, l)
		
			var mk1 string
			var mv2 string
		
			for range l {
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
						mk1 = string(b[o:o+l])
						o += l
					}
				}
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
						mv2 = string(b[o:o+l])
						o += l
					}
				}
				m.Value[mk1] = mv2
			}
		}
	}

	return o
}

func size42(m *MsgMap) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		for mk1, mv2 := range m.Value {
			{
				vi := mk1
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
			{
				vi := mv2
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
	return n
}

func marshal42(m *MsgMap, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint64(len(m.Value))
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
		for mk1, mv2 := range m.Value {
			{
				vi := mk1
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
			{
				vi := mv2
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
		}
	}

	return o
}

func unmarshal42(m *MsgMap, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make(map[uint64]uint64, l)
		
			var mk1 uint64
			var mv2 uint64
		
			for range l {
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
					mk1 = vi
				}
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
					mv2 = vi
				}
				m.Value[mk1] = mv2
			}
		}
	}

	return o
}

func size43(m *MsgSlice) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
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
		n += l
	}
	return n
}

func marshal43(m *MsgSlice, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint64(len(m.Value))
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
		for _, sv1 := range m.Value {
			if sv1 {
				b[o] = 0x01
			} else {
				b[o] = 0x00
			}
			o++
		}
	}

	return o
}

func unmarshal43(m *MsgSlice, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = make([]bool, l)
			for i1 := range l {
				m.Value[i1] = b[o] != 0x00
				o++
			}
		}
	}

	return o
}

func size44(m *MsgArray) uint64 {
	var n uint64 = 3
	return n
}

func marshal44(m *MsgArray, b []byte) uint64 {
	var o uint64
	{
		// Value

		for _, av1 := range m.Value {
			if av1 {
				b[o] = 0x01
			} else {
				b[o] = 0x00
			}
			o++
		}
	}

	return o
}

func unmarshal44(m *MsgArray, b []byte) uint64 {
	var o uint64
	{
		// Value

		for i1 := range 3 {
			m.Value[i1] = b[o] != 0x00
			o++
		}
	}

	return o
}

func size45(m *MsgString) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			l := uint64(len(m.Value))
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
	return n
}

func marshal45(m *MsgString, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			l := uint64(len(m.Value))
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
			copy(b[o:o+l], m.Value)
			o += l
		}
	}

	return o
}

func unmarshal45(m *MsgString, b []byte) uint64 {
	var o uint64
	{
		// Value

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
				m.Value = string(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size46(m *MsgFloat32) uint64 {
	var n uint64 = 4
	return n
}

func marshal46(m *MsgFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

func unmarshal46(m *MsgFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*float32)(unsafe.Pointer(&b[o]))
		o += 4
	}

	return o
}

func size47(m *MsgFloat64) uint64 {
	var n uint64 = 8
	return n
}

func marshal47(m *MsgFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

func unmarshal47(m *MsgFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*float64)(unsafe.Pointer(&b[o]))
		o += 8
	}

	return o
}

func size48(m *MsgBool10) uint64 {
	var n uint64 = 2
	return n
}

func marshal48(m *MsgBool10, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if m.Value1 {
			b[0] |= 0x01
		} else {
			b[0] &= 0xFE
		}
	}
	{
		// Value2

		if m.Value2 {
			b[0] |= 0x02
		} else {
			b[0] &= 0xFD
		}
	}
	{
		// Value3

		if m.Value3 {
			b[0] |= 0x04
		} else {
			b[0] &= 0xFB
		}
	}
	{
		// Value4

		if m.Value4 {
			b[0] |= 0x08
		} else {
			b[0] &= 0xF7
		}
	}
	{
		// Value5

		if m.Value5 {
			b[0] |= 0x10
		} else {
			b[0] &= 0xEF
		}
	}
	{
		// Value6

		if m.Value6 {
			b[0] |= 0x20
		} else {
			b[0] &= 0xDF
		}
	}
	{
		// Value7

		if m.Value7 {
			b[0] |= 0x40
		} else {
			b[0] &= 0xBF
		}
	}
	{
		// Value8

		if m.Value8 {
			b[0] |= 0x80
		} else {
			b[0] &= 0x7F
		}
	}
	{
		// Value9

		if m.Value9 {
			b[1] |= 0x01
		} else {
			b[1] &= 0xFE
		}
	}
	{
		// Value10

		if m.Value10 {
			b[1] |= 0x02
		} else {
			b[1] &= 0xFD
		}
	}

	return o
}

func unmarshal48(m *MsgBool10, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		m.Value1 = b[0]&0x01 != 0
	}
	{
		// Value2

		m.Value2 = b[0]&0x02 != 0
	}
	{
		// Value3

		m.Value3 = b[0]&0x04 != 0
	}
	{
		// Value4

		m.Value4 = b[0]&0x08 != 0
	}
	{
		// Value5

		m.Value5 = b[0]&0x10 != 0
	}
	{
		// Value6

		m.Value6 = b[0]&0x20 != 0
	}
	{
		// Value7

		m.Value7 = b[0]&0x40 != 0
	}
	{
		// Value8

		m.Value8 = b[0]&0x80 != 0
	}
	{
		// Value9

		m.Value9 = b[1]&0x01 != 0
	}
	{
		// Value10

		m.Value10 = b[1]&0x02 != 0
	}

	return o
}

func size49(m *MsgBool3) uint64 {
	var n uint64 = 1
	return n
}

func marshal49(m *MsgBool3, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		if m.Value1 {
			b[0] |= 0x01
		} else {
			b[0] &= 0xFE
		}
	}
	{
		// Value2

		if m.Value2 {
			b[0] |= 0x02
		} else {
			b[0] &= 0xFD
		}
	}
	{
		// Value3

		if m.Value3 {
			b[0] |= 0x04
		} else {
			b[0] &= 0xFB
		}
	}

	return o
}

func unmarshal49(m *MsgBool3, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		m.Value1 = b[0]&0x01 != 0
	}
	{
		// Value2

		m.Value2 = b[0]&0x02 != 0
	}
	{
		// Value3

		m.Value3 = b[0]&0x04 != 0
	}

	return o
}

func size50(m *MsgInt8) uint64 {
	var n uint64 = 1
	return n
}

func marshal50(m *MsgInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal50(m *MsgInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = int8(b[o])
		o++
	}

	return o
}

func size51(m *MsgInt16) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := uint16(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFF
			}
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

func marshal51(m *MsgInt16, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint16(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFF
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

func unmarshal51(m *MsgInt16, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			if vi&0x01 == 0 {
				vi >>= 1
			} else {
				vi >>= 1
				vi ^= 0xFFFF
			}
			m.Value = int16(vi)
		}
	}

	return o
}

func size52(m *MsgInt32) uint64 {
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

func marshal52(m *MsgInt32, b []byte) uint64 {
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

func unmarshal52(m *MsgInt32, b []byte) uint64 {
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

func size53(m *MsgInt64) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := uint64(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFFFFFFFFFFFFFF
			}
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
			case vi <= 0x7FFFFFFFFFFFFFFF:
				n += 8
			default:
				n += 9
			}
		}
	}
	return n
}

func marshal53(m *MsgInt64, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := uint64(m.Value) << 1
			if m.Value < 0 {
				vi ^= 0xFFFFFFFFFFFFFFFF
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
			case vi <= 0x7FFFFFFFFFFFFFFF:
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
				vi >>= 7
				b[o] = byte(vi)
				o++
			}
		}
	}

	return o
}

func unmarshal53(m *MsgInt64, b []byte) uint64 {
	var o uint64
	{
		// Value

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
											vi |= uint64(b[o+8]&0x7F) << 56
											if b[o+8]&0x80 == 0 {
												o += 9
											} else {
												vi |= uint64(b[o+9]) << 63
												o += 10
											}
										}
									}
								}
							}
						}
					}
				}
			}
			if vi&0x01 == 0 {
				vi >>= 1
			} else {
				vi >>= 1
				vi ^= 0xFFFFFFFFFFFFFFFF
			}
			m.Value = int64(vi)
		}
	}

	return o
}

func size54(m *MsgUint8) uint64 {
	var n uint64 = 1
	return n
}

func marshal54(m *MsgUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = m.Value
		o++
	}

	return o
}

func unmarshal54(m *MsgUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = b[o]
		o++
	}

	return o
}

func size55(m *MsgUint16) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := m.Value
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

func marshal55(m *MsgUint16, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := m.Value
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

func unmarshal55(m *MsgUint16, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = vi
		}
	}

	return o
}

func size56(m *MsgUint32) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := m.Value
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

func marshal56(m *MsgUint32, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := m.Value
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

func unmarshal56(m *MsgUint32, b []byte) uint64 {
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
			m.Value = vi
		}
	}

	return o
}

func size57(m *MsgUint64) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			vi := m.Value
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
	return n
}

func marshal57(m *MsgUint64, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			vi := m.Value
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
	}

	return o
}

func unmarshal57(m *MsgUint64, b []byte) uint64 {
	var o uint64
	{
		// Value

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
			m.Value = vi
		}
	}

	return o
}

func size37(m *pkg2.SubMsg) uint64 {
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

func marshal37(m *pkg2.SubMsg, b []byte) uint64 {
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

func unmarshal37(m *pkg2.SubMsg, b []byte) uint64 {
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
						mv2 = string(b[o:o+l])
						o += l
					}
				}
				m.Value1[mk1] = mv2
			}
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
				m.Value2 = string(b[o:o+l])
				o += l
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
