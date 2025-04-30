package pkg1

import (
	"time"
	"unsafe"

	"github.com/outofforest/proton"
	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1/spkg"
	"github.com/outofforest/proton/test/pkg2"
	spkg7 "github.com/outofforest/proton/test/pkg2/spkg"
	"github.com/pkg/errors"
)

const (
	id38 uint64 = iota + 1
	id58
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
	id39
	id36
	id35
	id34
	id33
	id32
	id31
	id30
	id29
	id28
	id26
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
		MsgTime{},
		MsgMixedCustom{},
	}
}

// ID returns ID of message type.
func (m Marshaller) ID(msg any) (uint64, error) {
	switch msg.(type) {
	case *pkg2.SubMsg:
		return id38, nil
	case *MsgUint64:
		return id58, nil
	case *MsgUint32:
		return id57, nil
	case *MsgUint16:
		return id56, nil
	case *MsgUint8:
		return id55, nil
	case *MsgInt64:
		return id54, nil
	case *MsgInt32:
		return id53, nil
	case *MsgInt16:
		return id52, nil
	case *MsgInt8:
		return id51, nil
	case *MsgBool3:
		return id50, nil
	case *MsgBool10:
		return id49, nil
	case *MsgFloat64:
		return id48, nil
	case *MsgFloat32:
		return id47, nil
	case *MsgString:
		return id46, nil
	case *MsgArray:
		return id45, nil
	case *MsgSlice:
		return id44, nil
	case *MsgMap:
		return id43, nil
	case *MsgMapString:
		return id42, nil
	case *MsgStruct:
		return id41, nil
	case *MsgStructAnonymous:
		return id39, nil
	case *MsgArrayUint8:
		return id36, nil
	case *MsgArrayInt8:
		return id35, nil
	case *MsgArrayFloat32:
		return id34, nil
	case *MsgArrayFloat64:
		return id33, nil
	case *MsgSliceUint8:
		return id32, nil
	case *MsgSliceInt8:
		return id31, nil
	case *MsgSliceFloat32:
		return id30, nil
	case *MsgSliceFloat64:
		return id29, nil
	case *MsgMixed:
		return id28, nil
	case *MsgUint64Custom:
		return id26, nil
	case *MsgUint32Custom:
		return id25, nil
	case *MsgUint16Custom:
		return id24, nil
	case *MsgUint8Custom:
		return id23, nil
	case *MsgInt64Custom:
		return id22, nil
	case *MsgInt32Custom:
		return id21, nil
	case *MsgInt16Custom:
		return id20, nil
	case *MsgInt8Custom:
		return id19, nil
	case *MsgBoolCustom:
		return id18, nil
	case *MsgFloat64Custom:
		return id17, nil
	case *MsgFloat32Custom:
		return id16, nil
	case *MsgStringCustom:
		return id15, nil
	case *MsgArrayCustom:
		return id14, nil
	case *MsgSliceCustom:
		return id13, nil
	case *MsgMapCustom:
		return id12, nil
	case *MsgArrayUint8Custom:
		return id11, nil
	case *MsgArrayUint8Custom2:
		return id10, nil
	case *MsgArrayInt8Custom:
		return id9, nil
	case *MsgArrayFloat32Custom:
		return id8, nil
	case *MsgArrayFloat64Custom:
		return id7, nil
	case *MsgSliceUint8Custom:
		return id6, nil
	case *MsgSliceUint8Custom2:
		return id5, nil
	case *MsgSliceInt8Custom:
		return id4, nil
	case *MsgSliceFloat32Custom:
		return id3, nil
	case *MsgSliceFloat64Custom:
		return id2, nil
	case *MsgTime:
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
		return size38(msg2), nil
	case *MsgUint64:
		return size58(msg2), nil
	case *MsgUint32:
		return size57(msg2), nil
	case *MsgUint16:
		return size56(msg2), nil
	case *MsgUint8:
		return size55(msg2), nil
	case *MsgInt64:
		return size54(msg2), nil
	case *MsgInt32:
		return size53(msg2), nil
	case *MsgInt16:
		return size52(msg2), nil
	case *MsgInt8:
		return size51(msg2), nil
	case *MsgBool3:
		return size50(msg2), nil
	case *MsgBool10:
		return size49(msg2), nil
	case *MsgFloat64:
		return size48(msg2), nil
	case *MsgFloat32:
		return size47(msg2), nil
	case *MsgString:
		return size46(msg2), nil
	case *MsgArray:
		return size45(msg2), nil
	case *MsgSlice:
		return size44(msg2), nil
	case *MsgMap:
		return size43(msg2), nil
	case *MsgMapString:
		return size42(msg2), nil
	case *MsgStruct:
		return size41(msg2), nil
	case *MsgStructAnonymous:
		return size39(msg2), nil
	case *MsgArrayUint8:
		return size36(msg2), nil
	case *MsgArrayInt8:
		return size35(msg2), nil
	case *MsgArrayFloat32:
		return size34(msg2), nil
	case *MsgArrayFloat64:
		return size33(msg2), nil
	case *MsgSliceUint8:
		return size32(msg2), nil
	case *MsgSliceInt8:
		return size31(msg2), nil
	case *MsgSliceFloat32:
		return size30(msg2), nil
	case *MsgSliceFloat64:
		return size29(msg2), nil
	case *MsgMixed:
		return size28(msg2), nil
	case *MsgUint64Custom:
		return size26(msg2), nil
	case *MsgUint32Custom:
		return size25(msg2), nil
	case *MsgUint16Custom:
		return size24(msg2), nil
	case *MsgUint8Custom:
		return size23(msg2), nil
	case *MsgInt64Custom:
		return size22(msg2), nil
	case *MsgInt32Custom:
		return size21(msg2), nil
	case *MsgInt16Custom:
		return size20(msg2), nil
	case *MsgInt8Custom:
		return size19(msg2), nil
	case *MsgBoolCustom:
		return size18(msg2), nil
	case *MsgFloat64Custom:
		return size17(msg2), nil
	case *MsgFloat32Custom:
		return size16(msg2), nil
	case *MsgStringCustom:
		return size15(msg2), nil
	case *MsgArrayCustom:
		return size14(msg2), nil
	case *MsgSliceCustom:
		return size13(msg2), nil
	case *MsgMapCustom:
		return size12(msg2), nil
	case *MsgArrayUint8Custom:
		return size11(msg2), nil
	case *MsgArrayUint8Custom2:
		return size10(msg2), nil
	case *MsgArrayInt8Custom:
		return size9(msg2), nil
	case *MsgArrayFloat32Custom:
		return size8(msg2), nil
	case *MsgArrayFloat64Custom:
		return size7(msg2), nil
	case *MsgSliceUint8Custom:
		return size6(msg2), nil
	case *MsgSliceUint8Custom2:
		return size5(msg2), nil
	case *MsgSliceInt8Custom:
		return size4(msg2), nil
	case *MsgSliceFloat32Custom:
		return size3(msg2), nil
	case *MsgSliceFloat64Custom:
		return size2(msg2), nil
	case *MsgTime:
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
		return id38, marshal38(msg2, buf), nil
	case *MsgUint64:
		return id58, marshal58(msg2, buf), nil
	case *MsgUint32:
		return id57, marshal57(msg2, buf), nil
	case *MsgUint16:
		return id56, marshal56(msg2, buf), nil
	case *MsgUint8:
		return id55, marshal55(msg2, buf), nil
	case *MsgInt64:
		return id54, marshal54(msg2, buf), nil
	case *MsgInt32:
		return id53, marshal53(msg2, buf), nil
	case *MsgInt16:
		return id52, marshal52(msg2, buf), nil
	case *MsgInt8:
		return id51, marshal51(msg2, buf), nil
	case *MsgBool3:
		return id50, marshal50(msg2, buf), nil
	case *MsgBool10:
		return id49, marshal49(msg2, buf), nil
	case *MsgFloat64:
		return id48, marshal48(msg2, buf), nil
	case *MsgFloat32:
		return id47, marshal47(msg2, buf), nil
	case *MsgString:
		return id46, marshal46(msg2, buf), nil
	case *MsgArray:
		return id45, marshal45(msg2, buf), nil
	case *MsgSlice:
		return id44, marshal44(msg2, buf), nil
	case *MsgMap:
		return id43, marshal43(msg2, buf), nil
	case *MsgMapString:
		return id42, marshal42(msg2, buf), nil
	case *MsgStruct:
		return id41, marshal41(msg2, buf), nil
	case *MsgStructAnonymous:
		return id39, marshal39(msg2, buf), nil
	case *MsgArrayUint8:
		return id36, marshal36(msg2, buf), nil
	case *MsgArrayInt8:
		return id35, marshal35(msg2, buf), nil
	case *MsgArrayFloat32:
		return id34, marshal34(msg2, buf), nil
	case *MsgArrayFloat64:
		return id33, marshal33(msg2, buf), nil
	case *MsgSliceUint8:
		return id32, marshal32(msg2, buf), nil
	case *MsgSliceInt8:
		return id31, marshal31(msg2, buf), nil
	case *MsgSliceFloat32:
		return id30, marshal30(msg2, buf), nil
	case *MsgSliceFloat64:
		return id29, marshal29(msg2, buf), nil
	case *MsgMixed:
		return id28, marshal28(msg2, buf), nil
	case *MsgUint64Custom:
		return id26, marshal26(msg2, buf), nil
	case *MsgUint32Custom:
		return id25, marshal25(msg2, buf), nil
	case *MsgUint16Custom:
		return id24, marshal24(msg2, buf), nil
	case *MsgUint8Custom:
		return id23, marshal23(msg2, buf), nil
	case *MsgInt64Custom:
		return id22, marshal22(msg2, buf), nil
	case *MsgInt32Custom:
		return id21, marshal21(msg2, buf), nil
	case *MsgInt16Custom:
		return id20, marshal20(msg2, buf), nil
	case *MsgInt8Custom:
		return id19, marshal19(msg2, buf), nil
	case *MsgBoolCustom:
		return id18, marshal18(msg2, buf), nil
	case *MsgFloat64Custom:
		return id17, marshal17(msg2, buf), nil
	case *MsgFloat32Custom:
		return id16, marshal16(msg2, buf), nil
	case *MsgStringCustom:
		return id15, marshal15(msg2, buf), nil
	case *MsgArrayCustom:
		return id14, marshal14(msg2, buf), nil
	case *MsgSliceCustom:
		return id13, marshal13(msg2, buf), nil
	case *MsgMapCustom:
		return id12, marshal12(msg2, buf), nil
	case *MsgArrayUint8Custom:
		return id11, marshal11(msg2, buf), nil
	case *MsgArrayUint8Custom2:
		return id10, marshal10(msg2, buf), nil
	case *MsgArrayInt8Custom:
		return id9, marshal9(msg2, buf), nil
	case *MsgArrayFloat32Custom:
		return id8, marshal8(msg2, buf), nil
	case *MsgArrayFloat64Custom:
		return id7, marshal7(msg2, buf), nil
	case *MsgSliceUint8Custom:
		return id6, marshal6(msg2, buf), nil
	case *MsgSliceUint8Custom2:
		return id5, marshal5(msg2, buf), nil
	case *MsgSliceInt8Custom:
		return id4, marshal4(msg2, buf), nil
	case *MsgSliceFloat32Custom:
		return id3, marshal3(msg2, buf), nil
	case *MsgSliceFloat64Custom:
		return id2, marshal2(msg2, buf), nil
	case *MsgTime:
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
	case id38:
		msg := &pkg2.SubMsg{}
		return msg, unmarshal38(msg, buf), nil
	case id58:
		msg := &MsgUint64{}
		return msg, unmarshal58(msg, buf), nil
	case id57:
		msg := &MsgUint32{}
		return msg, unmarshal57(msg, buf), nil
	case id56:
		msg := &MsgUint16{}
		return msg, unmarshal56(msg, buf), nil
	case id55:
		msg := &MsgUint8{}
		return msg, unmarshal55(msg, buf), nil
	case id54:
		msg := &MsgInt64{}
		return msg, unmarshal54(msg, buf), nil
	case id53:
		msg := &MsgInt32{}
		return msg, unmarshal53(msg, buf), nil
	case id52:
		msg := &MsgInt16{}
		return msg, unmarshal52(msg, buf), nil
	case id51:
		msg := &MsgInt8{}
		return msg, unmarshal51(msg, buf), nil
	case id50:
		msg := &MsgBool3{}
		return msg, unmarshal50(msg, buf), nil
	case id49:
		msg := &MsgBool10{}
		return msg, unmarshal49(msg, buf), nil
	case id48:
		msg := &MsgFloat64{}
		return msg, unmarshal48(msg, buf), nil
	case id47:
		msg := &MsgFloat32{}
		return msg, unmarshal47(msg, buf), nil
	case id46:
		msg := &MsgString{}
		return msg, unmarshal46(msg, buf), nil
	case id45:
		msg := &MsgArray{}
		return msg, unmarshal45(msg, buf), nil
	case id44:
		msg := &MsgSlice{}
		return msg, unmarshal44(msg, buf), nil
	case id43:
		msg := &MsgMap{}
		return msg, unmarshal43(msg, buf), nil
	case id42:
		msg := &MsgMapString{}
		return msg, unmarshal42(msg, buf), nil
	case id41:
		msg := &MsgStruct{}
		return msg, unmarshal41(msg, buf), nil
	case id39:
		msg := &MsgStructAnonymous{}
		return msg, unmarshal39(msg, buf), nil
	case id36:
		msg := &MsgArrayUint8{}
		return msg, unmarshal36(msg, buf), nil
	case id35:
		msg := &MsgArrayInt8{}
		return msg, unmarshal35(msg, buf), nil
	case id34:
		msg := &MsgArrayFloat32{}
		return msg, unmarshal34(msg, buf), nil
	case id33:
		msg := &MsgArrayFloat64{}
		return msg, unmarshal33(msg, buf), nil
	case id32:
		msg := &MsgSliceUint8{}
		return msg, unmarshal32(msg, buf), nil
	case id31:
		msg := &MsgSliceInt8{}
		return msg, unmarshal31(msg, buf), nil
	case id30:
		msg := &MsgSliceFloat32{}
		return msg, unmarshal30(msg, buf), nil
	case id29:
		msg := &MsgSliceFloat64{}
		return msg, unmarshal29(msg, buf), nil
	case id28:
		msg := &MsgMixed{}
		return msg, unmarshal28(msg, buf), nil
	case id26:
		msg := &MsgUint64Custom{}
		return msg, unmarshal26(msg, buf), nil
	case id25:
		msg := &MsgUint32Custom{}
		return msg, unmarshal25(msg, buf), nil
	case id24:
		msg := &MsgUint16Custom{}
		return msg, unmarshal24(msg, buf), nil
	case id23:
		msg := &MsgUint8Custom{}
		return msg, unmarshal23(msg, buf), nil
	case id22:
		msg := &MsgInt64Custom{}
		return msg, unmarshal22(msg, buf), nil
	case id21:
		msg := &MsgInt32Custom{}
		return msg, unmarshal21(msg, buf), nil
	case id20:
		msg := &MsgInt16Custom{}
		return msg, unmarshal20(msg, buf), nil
	case id19:
		msg := &MsgInt8Custom{}
		return msg, unmarshal19(msg, buf), nil
	case id18:
		msg := &MsgBoolCustom{}
		return msg, unmarshal18(msg, buf), nil
	case id17:
		msg := &MsgFloat64Custom{}
		return msg, unmarshal17(msg, buf), nil
	case id16:
		msg := &MsgFloat32Custom{}
		return msg, unmarshal16(msg, buf), nil
	case id15:
		msg := &MsgStringCustom{}
		return msg, unmarshal15(msg, buf), nil
	case id14:
		msg := &MsgArrayCustom{}
		return msg, unmarshal14(msg, buf), nil
	case id13:
		msg := &MsgSliceCustom{}
		return msg, unmarshal13(msg, buf), nil
	case id12:
		msg := &MsgMapCustom{}
		return msg, unmarshal12(msg, buf), nil
	case id11:
		msg := &MsgArrayUint8Custom{}
		return msg, unmarshal11(msg, buf), nil
	case id10:
		msg := &MsgArrayUint8Custom2{}
		return msg, unmarshal10(msg, buf), nil
	case id9:
		msg := &MsgArrayInt8Custom{}
		return msg, unmarshal9(msg, buf), nil
	case id8:
		msg := &MsgArrayFloat32Custom{}
		return msg, unmarshal8(msg, buf), nil
	case id7:
		msg := &MsgArrayFloat64Custom{}
		return msg, unmarshal7(msg, buf), nil
	case id6:
		msg := &MsgSliceUint8Custom{}
		return msg, unmarshal6(msg, buf), nil
	case id5:
		msg := &MsgSliceUint8Custom2{}
		return msg, unmarshal5(msg, buf), nil
	case id4:
		msg := &MsgSliceInt8Custom{}
		return msg, unmarshal4(msg, buf), nil
	case id3:
		msg := &MsgSliceFloat32Custom{}
		return msg, unmarshal3(msg, buf), nil
	case id2:
		msg := &MsgSliceFloat64Custom{}
		return msg, unmarshal2(msg, buf), nil
	case id1:
		msg := &MsgTime{}
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
		helpers.UInt64Size(l, &n)
		n += l * 3
		for _, sv6 := range m.Value {
			for _, av5 := range sv6 {
				l := uint64(len(av5))
				helpers.UInt64Size(l, &n)
				n += l
				for _, sv4 := range av5 {
					l := uint64(len(sv4))
						helpers.UInt64Size(l, &n)
					n += l * 3
					for mk2, mv3 := range sv4 {
						helpers.Int16Size(mk2, &n)
						for _, av1 := range mv3 {
							helpers.UInt64Size(av1, &n)
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

		helpers.UInt64Marshal(uint64(len(m.Value)), b, &o)
		for _, sv6 := range m.Value {
			for _, av5 := range sv6 {
				helpers.UInt64Marshal(uint64(len(av5)), b, &o)
				for _, sv4 := range av5 {
					helpers.UInt64Marshal(uint64(len(sv4)), b, &o)
					for mk2, mv3 := range sv4 {
						helpers.Int16Marshal(mk2, b, &o)
						for _, av1 := range mv3 {
							helpers.UInt64Marshal(av1, b, &o)
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
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([][3][]map[int16]custom.Array, l)
			for i6 := range l {
				for i5 := range 3 {
					var l uint64
					helpers.UInt64Unmarshal(&l, b, &o)
					if l > 0 {
						m.Value[i6][i5] = make([]map[int16]custom.Array, l)
						for i4 := range l {
							var l uint64
							helpers.UInt64Unmarshal(&l, b, &o)
							if l > 0 {
								m.Value[i6][i5][i4] = make(map[int16]custom.Array, l)
							
								var mk2 int16
								var mv3 custom.Array
							
								for range l {
									helpers.Int16Unmarshal(&mk2, b, &o)
									for i1 := range 2 {
										helpers.UInt64Unmarshal(&mv3[i1], b, &o)
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

func size1(m *MsgTime) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int64Size(m.Value.Unix(), &n)
	}
	return n
}

func marshal1(m *MsgTime, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Marshal(m.Value.Unix(), b, &o)
	}

	return o
}

func unmarshal1(m *MsgTime, b []byte) uint64 {
	var o uint64
	{
		// Value

		var vi int64
		helpers.Int64Unmarshal(&vi, b, &o)
		m.Value = time.Unix(vi, 0)
	}

	return o
}

func size2(m *MsgSliceFloat64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l * 8
	}
	return n
}

func marshal2(m *MsgSliceFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l*8], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8))
			o += l * 8
		}
	}

	return o
}

func unmarshal2(m *MsgSliceFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]custom.Float64, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
			o += l * 8
		}
	}

	return o
}

func size3(m *MsgSliceFloat32Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l * 4
	}
	return n
}

func marshal3(m *MsgSliceFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l*4], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4))
			o += l * 4
		}
	}

	return o
}

func unmarshal3(m *MsgSliceFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]custom.Float32, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
			o += l * 4
		}
	}

	return o
}

func size4(m *MsgSliceInt8Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal4(m *MsgSliceInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l))
			o += l
		}
	}

	return o
}

func unmarshal4(m *MsgSliceInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]custom.Int8, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
			o += l
		}
	}

	return o
}

func size5(m *MsgSliceUint8Custom2) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal5(m *MsgSliceUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice((*byte)(&m.Value[0]), l))
			o += l
		}
	}

	return o
}

func unmarshal5(m *MsgSliceUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]custom.Uint8, l)
			copy(unsafe.Slice((*byte)(&m.Value[0]), l), b[o:o+l])
			o += l
		}
	}

	return o
}

func size6(m *MsgSliceUint8Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal6(m *MsgSliceUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice(&m.Value[0], l))
			o += l
		}
	}

	return o
}

func unmarshal6(m *MsgSliceUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]uint8, l)
			copy(m.Value, b[o:o+l])
			o += l
		}
	}

	return o
}

func size7(m *MsgArrayFloat64Custom) uint64 {
	var n uint64 = 40
	return n
}

func marshal7(m *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

func unmarshal7(m *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
		o += 40
	}

	return o
}

func size8(m *MsgArrayFloat32Custom) uint64 {
	var n uint64 = 20
	return n
}

func marshal8(m *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

func unmarshal8(m *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
		o += 20
	}

	return o
}

func size9(m *MsgArrayInt8Custom) uint64 {
	var n uint64 = 5
	return n
}

func marshal9(m *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

func unmarshal9(m *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size10(m *MsgArrayUint8Custom2) uint64 {
	var n uint64 = 5
	return n
}

func marshal10(m *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(&m.Value[0]), 5))
		o += 5
	}

	return o
}

func unmarshal10(m *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(&m.Value[0]), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size11(m *MsgArrayUint8Custom) uint64 {
	var n uint64 = 5
	return n
}

func marshal11(m *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

func unmarshal11(m *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
		o += 5
	}

	return o
}

func size12(m *MsgMapCustom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
			helpers.UInt64Size(l, &n)
		n += l * 2
		for mk1, mv2 := range m.Value {
			{
				l := uint64(len(mk1))
				helpers.UInt64Size(l, &n)
				n += l
			}
			helpers.Int32Size(mv2, &n)
		}
	}
	return n
}

func marshal12(m *MsgMapCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(uint64(len(m.Value)), b, &o)
		for mk1, mv2 := range m.Value {
			{
				l := uint64(len(mk1))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], mk1)
				o += l
			}
			helpers.Int32Marshal(mv2, b, &o)
		}
	}

	return o
}

func unmarshal12(m *MsgMapCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make(custom.Map, l)
		
			var mk1 custom.String
			var mv2 custom.Int32
		
			for range l {
				{
					var l uint64
					helpers.UInt64Unmarshal(&l, b, &o)
					if l > 0 {
						mk1 = custom.String(b[o:o+l])
						o += l
					}
				}
				helpers.Int32Unmarshal(&mv2, b, &o)
				m.Value[mk1] = mv2
			}
		}
	}

	return o
}

func size13(m *MsgSliceCustom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
		for _, sv1 := range m.Value {
			helpers.UInt64Size(sv1, &n)
		}
	}
	return n
}

func marshal13(m *MsgSliceCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(uint64(len(m.Value)), b, &o)
		for _, sv1 := range m.Value {
			helpers.UInt64Marshal(sv1, b, &o)
		}
	}

	return o
}

func unmarshal13(m *MsgSliceCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]custom.Uint64, l)
			for i1 := range l {
				helpers.UInt64Unmarshal(&m.Value[i1], b, &o)
			}
		}
	}

	return o
}

func size14(m *MsgArrayCustom) uint64 {
	var n uint64 = 2
	{
		// Value

		for _, av1 := range m.Value {
			helpers.UInt64Size(av1, &n)
		}
	}
	return n
}

func marshal14(m *MsgArrayCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		for _, av1 := range m.Value {
			helpers.UInt64Marshal(av1, b, &o)
		}
	}

	return o
}

func unmarshal14(m *MsgArrayCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		for i1 := range 2 {
			helpers.UInt64Unmarshal(&m.Value[i1], b, &o)
		}
	}

	return o
}

func size15(m *MsgStringCustom) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			l := uint64(len(m.Value))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	return n
}

func marshal15(m *MsgStringCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value)
			o += l
		}
	}

	return o
}

func unmarshal15(m *MsgStringCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = custom.String(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size16(m *MsgFloat32Custom) uint64 {
	var n uint64 = 4
	return n
}

func marshal16(m *MsgFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

func unmarshal16(m *MsgFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*custom.Float32)(unsafe.Pointer(&b[o]))
		o += 4
	}

	return o
}

func size17(m *MsgFloat64Custom) uint64 {
	var n uint64 = 8
	return n
}

func marshal17(m *MsgFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

func unmarshal17(m *MsgFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*custom.Float64)(unsafe.Pointer(&b[o]))
		o += 8
	}

	return o
}

func size18(m *MsgBoolCustom) uint64 {
	var n uint64 = 1
	return n
}

func marshal18(m *MsgBoolCustom, b []byte) uint64 {
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

func unmarshal18(m *MsgBoolCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		m.Value = b[0]&0x01 != 0
	}

	return o
}

func size19(m *MsgInt8Custom) uint64 {
	var n uint64 = 1
	return n
}

func marshal19(m *MsgInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal19(m *MsgInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = custom.Int8(b[o])
		o++
	}

	return o
}

func size20(m *MsgInt16Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int16Size(m.Value, &n)
	}
	return n
}

func marshal20(m *MsgInt16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal20(m *MsgInt16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size21(m *MsgInt32Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int32Size(m.Value, &n)
	}
	return n
}

func marshal21(m *MsgInt32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal21(m *MsgInt32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size22(m *MsgInt64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int64Size(m.Value, &n)
	}
	return n
}

func marshal22(m *MsgInt64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal22(m *MsgInt64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size23(m *MsgUint8Custom) uint64 {
	var n uint64 = 1
	return n
}

func marshal23(m *MsgUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal23(m *MsgUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = custom.Uint8(b[o])
		o++
	}

	return o
}

func size24(m *MsgUint16Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt16Size(m.Value, &n)
	}
	return n
}

func marshal24(m *MsgUint16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal24(m *MsgUint16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size25(m *MsgUint32Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt32Size(m.Value, &n)
	}
	return n
}

func marshal25(m *MsgUint32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal25(m *MsgUint32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size26(m *MsgUint64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt64Size(m.Value, &n)
	}
	return n
}

func marshal26(m *MsgUint64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal26(m *MsgUint64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size28(m *MsgMixed) uint64 {
	var n uint64 = 18
	{
		// Value1

		l := uint64(len(m.Value1))
			helpers.UInt64Size(l, &n)
		n += l * 1
		for mk1, mv2 := range m.Value1 {
			{
				l := uint64(len(mk1))
				helpers.UInt64Size(l, &n)
				n += l
			}
			n += size27(&mv2)
		}
	}
	{
		// Value2

		l := uint64(len(m.Value2))
			helpers.UInt64Size(l, &n)
		n += l * 2
		for _, mv3 := range m.Value2 {
			l := uint64(len(mv3))
			helpers.UInt64Size(l, &n)
			n += l
			for _, sv1 := range mv3 {
				{
					l := uint64(len(sv1))
					helpers.UInt64Size(l, &n)
					n += l
				}
			}
		}
	}
	{
		// Value3

		l := uint64(len(m.Value3))
		helpers.UInt64Size(l, &n)
		n += l * 32
		for _, sv2 := range m.Value3 {
			for _, av1 := range sv2 {
				helpers.UInt16Size(av1, &n)
			}
		}
	}
	{
		// Value4

		for _, av1 := range m.Value4 {
			l := uint64(len(av1))
				helpers.UInt64Size(l, &n)
			n += l * 5
		}
	}
	{
		// Value5

		l := uint64(len(m.Value5))
		helpers.UInt64Size(l, &n)
		n += l * 3
		for _, sv6 := range m.Value5 {
			for _, av5 := range sv6 {
				l := uint64(len(av5))
				helpers.UInt64Size(l, &n)
				n += l
				for _, sv4 := range av5 {
					l := uint64(len(sv4))
						helpers.UInt64Size(l, &n)
					n += l * 3
					for mk2, mv3 := range sv4 {
						helpers.Int16Size(mk2, &n)
						for _, av1 := range mv3 {
							helpers.Int64Size(av1, &n)
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
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	return n
}

func marshal28(m *MsgMixed, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		helpers.UInt64Marshal(uint64(len(m.Value1)), b, &o)
		for mk1, mv2 := range m.Value1 {
			{
				l := uint64(len(mk1))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], mk1)
				o += l
			}
			o += marshal27(&mv2, b[o:])
		}
	}
	{
		// Value2

		helpers.UInt64Marshal(uint64(len(m.Value2)), b, &o)
		for mk2, mv3 := range m.Value2 {
			b[o] = mk2
			o++
			helpers.UInt64Marshal(uint64(len(mv3)), b, &o)
			for _, sv1 := range mv3 {
				{
					l := uint64(len(sv1))
					helpers.UInt64Marshal(l, b, &o)
					copy(b[o:o+l], sv1)
					o += l
				}
			}
		}
	}
	{
		// Value3

		helpers.UInt64Marshal(uint64(len(m.Value3)), b, &o)
		for _, sv2 := range m.Value3 {
			for _, av1 := range sv2 {
				helpers.UInt16Marshal(av1, b, &o)
			}
		}
	}
	{
		// Value4

		for _, av3 := range m.Value4 {
			helpers.UInt64Marshal(uint64(len(av3)), b, &o)
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

		helpers.UInt64Marshal(uint64(len(m.Value5)), b, &o)
		for _, sv6 := range m.Value5 {
			for _, av5 := range sv6 {
				helpers.UInt64Marshal(uint64(len(av5)), b, &o)
				for _, sv4 := range av5 {
					helpers.UInt64Marshal(uint64(len(sv4)), b, &o)
					for mk2, mv3 := range sv4 {
						helpers.Int16Marshal(mk2, b, &o)
						for _, av1 := range mv3 {
							helpers.Int64Marshal(av1, b, &o)
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
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value8)
			o += l
		}
	}

	return o
}

func unmarshal28(m *MsgMixed, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value1 = make(map[string]spkg.SubMsg, l)
		
			var mk1 string
			var mv2 spkg.SubMsg
		
			for range l {
				{
					var l uint64
					helpers.UInt64Unmarshal(&l, b, &o)
					if l > 0 {
						mk1 = string(b[o:o+l])
						o += l
					}
				}
				o += unmarshal27(&mv2, b[o:])
				m.Value1[mk1] = mv2
			}
		}
	}
	{
		// Value2

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value2 = make(map[uint8][]string, l)
		
			var mk2 uint8
			var mv3 []string
		
			for range l {
				mk2 = b[o]
				o++
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					mv3 = make([]string, l)
					for i1 := range l {
						{
							var l uint64
							helpers.UInt64Unmarshal(&l, b, &o)
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
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value3 = make([][32]uint16, l)
			for i2 := range l {
				for i1 := range 32 {
					helpers.UInt16Unmarshal(&m.Value3[i2][i1], b, &o)
				}
			}
		}
	}
	{
		// Value4

		for i3 := range 12 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
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
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value5 = make([][3][]map[int16][2]int64, l)
			for i6 := range l {
				for i5 := range 3 {
					var l uint64
					helpers.UInt64Unmarshal(&l, b, &o)
					if l > 0 {
						m.Value5[i6][i5] = make([]map[int16][2]int64, l)
						for i4 := range l {
							var l uint64
							helpers.UInt64Unmarshal(&l, b, &o)
							if l > 0 {
								m.Value5[i6][i5][i4] = make(map[int16][2]int64, l)
							
								var mk2 int16
								var mv3 [2]int64
							
								for range l {
									helpers.Int16Unmarshal(&mk2, b, &o)
									for i1 := range 2 {
										helpers.Int64Unmarshal(&mv3[i1], b, &o)
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
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value8 = string(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size27(m *spkg.SubMsg) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int32Size(m.Value, &n)
	}
	return n
}

func marshal27(m *spkg.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal27(m *spkg.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size29(m *MsgSliceFloat64) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l * 8
	}
	return n
}

func marshal29(m *MsgSliceFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l*8], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8))
			o += l * 8
		}
	}

	return o
}

func unmarshal29(m *MsgSliceFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]float64, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
			o += l * 8
		}
	}

	return o
}

func size30(m *MsgSliceFloat32) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l * 4
	}
	return n
}

func marshal30(m *MsgSliceFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l*4], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4))
			o += l * 4
		}
	}

	return o
}

func unmarshal30(m *MsgSliceFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]float32, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
			o += l * 4
		}
	}

	return o
}

func size31(m *MsgSliceInt8) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal31(m *MsgSliceInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l))
			o += l
		}
	}

	return o
}

func unmarshal31(m *MsgSliceInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]int8, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
			o += l
		}
	}

	return o
}

func size32(m *MsgSliceUint8) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal32(m *MsgSliceUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l], unsafe.Slice(&m.Value[0], l))
			o += l
		}
	}

	return o
}

func unmarshal32(m *MsgSliceUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make([]uint8, l)
			copy(m.Value, b[o:o+l])
			o += l
		}
	}

	return o
}

func size33(m *MsgArrayFloat64) uint64 {
	var n uint64 = 40
	return n
}

func marshal33(m *MsgArrayFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

func unmarshal33(m *MsgArrayFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
		o += 40
	}

	return o
}

func size34(m *MsgArrayFloat32) uint64 {
	var n uint64 = 20
	return n
}

func marshal34(m *MsgArrayFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

func unmarshal34(m *MsgArrayFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
		o += 20
	}

	return o
}

func size35(m *MsgArrayInt8) uint64 {
	var n uint64 = 5
	return n
}

func marshal35(m *MsgArrayInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

func unmarshal35(m *MsgArrayInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size36(m *MsgArrayUint8) uint64 {
	var n uint64 = 5
	return n
}

func marshal36(m *MsgArrayUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

func unmarshal36(m *MsgArrayUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
		o += 5
	}

	return o
}

func size39(m *MsgStructAnonymous) uint64 {
	var n uint64
	{
		// SubMsg

		n += size37(&m.SubMsg)
	}
	{
		// Value2

		n += size38(&m.Value2)
	}
	return n
}

func marshal39(m *MsgStructAnonymous, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += marshal37(&m.SubMsg, b[o:])
	}
	{
		// Value2

		o += marshal38(&m.Value2, b[o:])
	}

	return o
}

func unmarshal39(m *MsgStructAnonymous, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += unmarshal37(&m.SubMsg, b[o:])
	}
	{
		// Value2

		o += unmarshal38(&m.Value2, b[o:])
	}

	return o
}

func size37(m *SubMsg) uint64 {
	var n uint64 = 3
	{
		// Value1

		l := uint64(len(m.Value1))
		helpers.UInt64Size(l, &n)
		n += l * 4
	}
	{
		// Value2

		l := uint64(len(m.Value2))
		helpers.UInt64Size(l, &n)
		for _, sv1 := range m.Value2 {
			n += size27(&sv1)
		}
	}
	{
		// Value3

		l := uint64(len(m.Value3))
		helpers.UInt64Size(l, &n)
		for _, sv1 := range m.Value3 {
			n += size40(&sv1)
		}
	}
	return n
}

func marshal37(m *SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value1

		l := uint64(len(m.Value1))
		helpers.UInt64Marshal(l, b, &o)
		if l > 0 {
			copy(b[o:o+l*4], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value1[0])), l*4))
			o += l * 4
		}
	}
	{
		// Value2

		helpers.UInt64Marshal(uint64(len(m.Value2)), b, &o)
		for _, sv1 := range m.Value2 {
			o += marshal27(&sv1, b[o:])
		}
	}
	{
		// Value3

		helpers.UInt64Marshal(uint64(len(m.Value3)), b, &o)
		for _, sv1 := range m.Value3 {
			o += marshal40(&sv1, b[o:])
		}
	}

	return o
}

func unmarshal37(m *SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value1

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value1 = make([]float32, l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value1[0])), l*4), b[o:o+l*4])
			o += l * 4
		}
	}
	{
		// Value2

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value2 = make([]spkg.SubMsg, l)
			for i1 := range l {
				o += unmarshal27(&m.Value2[i1], b[o:])
			}
		}
	}
	{
		// Value3

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value3 = make([]spkg7.SubMsg, l)
			for i1 := range l {
				o += unmarshal40(&m.Value3[i1], b[o:])
			}
		}
	}

	return o
}

func size40(m *spkg7.SubMsg) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int16Size(m.Value, &n)
	}
	return n
}

func marshal40(m *spkg7.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal40(m *spkg7.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size41(m *MsgStruct) uint64 {
	var n uint64
	{
		// Value1

		n += size37(&m.Value1)
	}
	{
		// Value2

		n += size38(&m.Value2)
	}
	return n
}

func marshal41(m *MsgStruct, b []byte) uint64 {
	var o uint64
	{
		// Value1

		o += marshal37(&m.Value1, b[o:])
	}
	{
		// Value2

		o += marshal38(&m.Value2, b[o:])
	}

	return o
}

func unmarshal41(m *MsgStruct, b []byte) uint64 {
	var o uint64
	{
		// Value1

		o += unmarshal37(&m.Value1, b[o:])
	}
	{
		// Value2

		o += unmarshal38(&m.Value2, b[o:])
	}

	return o
}

func size42(m *MsgMapString) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
			helpers.UInt64Size(l, &n)
		n += l * 2
		for mk1, mv2 := range m.Value {
			{
				l := uint64(len(mk1))
				helpers.UInt64Size(l, &n)
				n += l
			}
			{
				l := uint64(len(mv2))
				helpers.UInt64Size(l, &n)
				n += l
			}
		}
	}
	return n
}

func marshal42(m *MsgMapString, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(uint64(len(m.Value)), b, &o)
		for mk1, mv2 := range m.Value {
			{
				l := uint64(len(mk1))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], mk1)
				o += l
			}
			{
				l := uint64(len(mv2))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], mv2)
				o += l
			}
		}
	}

	return o
}

func unmarshal42(m *MsgMapString, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make(map[string]string, l)
		
			var mk1 string
			var mv2 string
		
			for range l {
				{
					var l uint64
					helpers.UInt64Unmarshal(&l, b, &o)
					if l > 0 {
						mk1 = string(b[o:o+l])
						o += l
					}
				}
				{
					var l uint64
					helpers.UInt64Unmarshal(&l, b, &o)
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

func size43(m *MsgMap) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
			helpers.UInt64Size(l, &n)
		n += l * 2
		for mk1, mv2 := range m.Value {
			helpers.UInt64Size(mk1, &n)
			helpers.UInt64Size(mv2, &n)
		}
	}
	return n
}

func marshal43(m *MsgMap, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(uint64(len(m.Value)), b, &o)
		for mk1, mv2 := range m.Value {
			helpers.UInt64Marshal(mk1, b, &o)
			helpers.UInt64Marshal(mv2, b, &o)
		}
	}

	return o
}

func unmarshal43(m *MsgMap, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value = make(map[uint64]uint64, l)
		
			var mk1 uint64
			var mv2 uint64
		
			for range l {
				helpers.UInt64Unmarshal(&mk1, b, &o)
				helpers.UInt64Unmarshal(&mv2, b, &o)
				m.Value[mk1] = mv2
			}
		}
	}

	return o
}

func size44(m *MsgSlice) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal44(m *MsgSlice, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(uint64(len(m.Value)), b, &o)
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

func unmarshal44(m *MsgSlice, b []byte) uint64 {
	var o uint64
	{
		// Value

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
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

func size45(m *MsgArray) uint64 {
	var n uint64 = 3
	return n
}

func marshal45(m *MsgArray, b []byte) uint64 {
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

func unmarshal45(m *MsgArray, b []byte) uint64 {
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

func size46(m *MsgString) uint64 {
	var n uint64 = 1
	{
		// Value

		{
			l := uint64(len(m.Value))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	return n
}

func marshal46(m *MsgString, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value)
			o += l
		}
	}

	return o
}

func unmarshal46(m *MsgString, b []byte) uint64 {
	var o uint64
	{
		// Value

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = string(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size47(m *MsgFloat32) uint64 {
	var n uint64 = 4
	return n
}

func marshal47(m *MsgFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

func unmarshal47(m *MsgFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*float32)(unsafe.Pointer(&b[o]))
		o += 4
	}

	return o
}

func size48(m *MsgFloat64) uint64 {
	var n uint64 = 8
	return n
}

func marshal48(m *MsgFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

func unmarshal48(m *MsgFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*float64)(unsafe.Pointer(&b[o]))
		o += 8
	}

	return o
}

func size49(m *MsgBool10) uint64 {
	var n uint64 = 2
	return n
}

func marshal49(m *MsgBool10, b []byte) uint64 {
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

func unmarshal49(m *MsgBool10, b []byte) uint64 {
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

func size50(m *MsgBool3) uint64 {
	var n uint64 = 1
	return n
}

func marshal50(m *MsgBool3, b []byte) uint64 {
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

func unmarshal50(m *MsgBool3, b []byte) uint64 {
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

func size51(m *MsgInt8) uint64 {
	var n uint64 = 1
	return n
}

func marshal51(m *MsgInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal51(m *MsgInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = int8(b[o])
		o++
	}

	return o
}

func size52(m *MsgInt16) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int16Size(m.Value, &n)
	}
	return n
}

func marshal52(m *MsgInt16, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal52(m *MsgInt16, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size53(m *MsgInt32) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int32Size(m.Value, &n)
	}
	return n
}

func marshal53(m *MsgInt32, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal53(m *MsgInt32, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size54(m *MsgInt64) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int64Size(m.Value, &n)
	}
	return n
}

func marshal54(m *MsgInt64, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal54(m *MsgInt64, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size55(m *MsgUint8) uint64 {
	var n uint64 = 1
	return n
}

func marshal55(m *MsgUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = m.Value
		o++
	}

	return o
}

func unmarshal55(m *MsgUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = b[o]
		o++
	}

	return o
}

func size56(m *MsgUint16) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt16Size(m.Value, &n)
	}
	return n
}

func marshal56(m *MsgUint16, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal56(m *MsgUint16, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size57(m *MsgUint32) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt32Size(m.Value, &n)
	}
	return n
}

func marshal57(m *MsgUint32, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal57(m *MsgUint32, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size58(m *MsgUint64) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt64Size(m.Value, &n)
	}
	return n
}

func marshal58(m *MsgUint64, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal58(m *MsgUint64, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size38(m *pkg2.SubMsg) uint64 {
	var n uint64 = 3
	{
		// Value1

		l := uint64(len(m.Value1))
			helpers.UInt64Size(l, &n)
		n += l * 2
		for _, mv2 := range m.Value1 {
			{
				l := uint64(len(mv2))
				helpers.UInt64Size(l, &n)
				n += l
			}
		}
	}
	{
		// Value2

		{
			l := uint64(len(m.Value2))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	{
		// Value3

		helpers.UInt16Size(m.Value3, &n)
	}
	return n
}

func marshal38(m *pkg2.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value1

		helpers.UInt64Marshal(uint64(len(m.Value1)), b, &o)
		for mk1, mv2 := range m.Value1 {
			if mk1 {
				b[o] = 0x01
			} else {
				b[o] = 0x00
			}
			o++
			{
				l := uint64(len(mv2))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], mv2)
				o += l
			}
		}
	}
	{
		// Value2

		{
			l := uint64(len(m.Value2))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value2)
			o += l
		}
	}
	{
		// Value3

		helpers.UInt16Marshal(m.Value3, b, &o)
	}

	return o
}

func unmarshal38(m *pkg2.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value1

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value1 = make(map[bool]string, l)
		
			var mk1 bool
			var mv2 string
		
			for range l {
				mk1 = b[o] != 0x00
				o++
				{
					var l uint64
					helpers.UInt64Unmarshal(&l, b, &o)
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
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value2 = string(b[o:o+l])
				o += l
			}
		}
	}
	{
		// Value3

		helpers.UInt16Unmarshal(&m.Value3, b, &o)
	}

	return o
}
