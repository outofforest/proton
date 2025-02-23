package pkg1

import (
	"unsafe"

	"github.com/pkg/errors"

	"github.com/outofforest/mass"
	"github.com/outofforest/proton"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1/spkg"
	"github.com/outofforest/proton/test/pkg2"
	spkg6 "github.com/outofforest/proton/test/pkg2/spkg"
)

const (
	id51 uint64 = iota + 1
	id72
	id71
	id70
	id69
	id68
	id67
	id66
	id65
	id64
	id63
	id62
	id61
	id60
	id59
	id57
	id56
	id55
	id54
	id52
	id49
	id48
	id47
	id46
	id45
	id43
	id41
	id39
	id34
	id32
	id31
	id30
	id29
	id28
	id27
	id26
	id25
	id24
	id23
	id22
	id21
	id20
	id18
	id17
	id16
	id15
	id14
	id13
	id12
	id11
	id9
	id7
	id5
	id3
	id0
)

var _ proton.Marshaller = Marshaller{}

// NewMarshaller creates marshaller.
func NewMarshaller(capacity uint64) Marshaller {
	return Marshaller{
		mass0:  mass.New[MsgMixedCustom](capacity),
		mass3:  mass.New[MsgSliceFloat64Custom](capacity),
		mass5:  mass.New[MsgSliceFloat32Custom](capacity),
		mass7:  mass.New[MsgSliceInt8Custom](capacity),
		mass9:  mass.New[MsgSliceUint8Custom2](capacity),
		mass11: mass.New[MsgSliceUint8Custom](capacity),
		mass12: mass.New[MsgArrayFloat64Custom](capacity),
		mass13: mass.New[MsgArrayFloat32Custom](capacity),
		mass14: mass.New[MsgArrayInt8Custom](capacity),
		mass15: mass.New[MsgArrayUint8Custom2](capacity),
		mass16: mass.New[MsgArrayUint8Custom](capacity),
		mass17: mass.New[MsgMapCustom](capacity),
		mass18: mass.New[MsgSliceCustom](capacity),
		mass20: mass.New[MsgArrayCustom](capacity),
		mass21: mass.New[MsgStringCustom](capacity),
		mass22: mass.New[MsgFloat32Custom](capacity),
		mass23: mass.New[MsgFloat64Custom](capacity),
		mass24: mass.New[MsgBoolCustom](capacity),
		mass25: mass.New[MsgInt8Custom](capacity),
		mass26: mass.New[MsgInt16Custom](capacity),
		mass27: mass.New[MsgInt32Custom](capacity),
		mass28: mass.New[MsgInt64Custom](capacity),
		mass29: mass.New[MsgUint8Custom](capacity),
		mass30: mass.New[MsgUint16Custom](capacity),
		mass31: mass.New[MsgUint32Custom](capacity),
		mass32: mass.New[MsgUint64Custom](capacity),
		mass34: mass.New[MsgMixed](capacity),
		mass39: mass.New[MsgSliceFloat64](capacity),
		mass41: mass.New[MsgSliceFloat32](capacity),
		mass43: mass.New[MsgSliceInt8](capacity),
		mass45: mass.New[MsgSliceUint8](capacity),
		mass46: mass.New[MsgArrayFloat64](capacity),
		mass47: mass.New[MsgArrayFloat32](capacity),
		mass48: mass.New[MsgArrayInt8](capacity),
		mass49: mass.New[MsgArrayUint8](capacity),
		mass52: mass.New[MsgStructAnonymous](capacity),
		mass54: mass.New[MsgStruct](capacity),
		mass55: mass.New[MsgMapString](capacity),
		mass56: mass.New[MsgMap](capacity),
		mass57: mass.New[MsgSlice](capacity),
		mass59: mass.New[MsgArray](capacity),
		mass60: mass.New[MsgString](capacity),
		mass61: mass.New[MsgFloat32](capacity),
		mass62: mass.New[MsgFloat64](capacity),
		mass63: mass.New[MsgBool10](capacity),
		mass64: mass.New[MsgBool3](capacity),
		mass65: mass.New[MsgInt8](capacity),
		mass66: mass.New[MsgInt16](capacity),
		mass67: mass.New[MsgInt32](capacity),
		mass68: mass.New[MsgInt64](capacity),
		mass69: mass.New[MsgUint8](capacity),
		mass70: mass.New[MsgUint16](capacity),
		mass71: mass.New[MsgUint32](capacity),
		mass72: mass.New[MsgUint64](capacity),
		mass51: mass.New[pkg2.SubMsg](capacity),
		mass2:  mass.New[[3][]map[int16]custom.Array](capacity),
		mass1:  mass.New[map[int16]custom.Array](capacity),
		mass4:  mass.New[custom.Float64](capacity),
		mass6:  mass.New[custom.Float32](capacity),
		mass8:  mass.New[custom.Int8](capacity),
		mass10: mass.New[custom.Uint8](capacity),
		mass19: mass.New[custom.Uint64](capacity),
		mass35: mass.New[string](capacity),
		mass36: mass.New[[32]uint16](capacity),
		mass38: mass.New[[3][]map[int16][2]int64](capacity),
		mass37: mass.New[map[int16][2]int64](capacity),
		mass40: mass.New[float64](capacity),
		mass42: mass.New[float32](capacity),
		mass44: mass.New[int8](capacity),
		mass33: mass.New[spkg.SubMsg](capacity),
		mass53: mass.New[spkg6.SubMsg](capacity),
		mass58: mass.New[bool](capacity),
	}
}

// Marshaller marshals and unmarshals messages.
type Marshaller struct {
	mass0  *mass.Mass[MsgMixedCustom]
	mass3  *mass.Mass[MsgSliceFloat64Custom]
	mass5  *mass.Mass[MsgSliceFloat32Custom]
	mass7  *mass.Mass[MsgSliceInt8Custom]
	mass9  *mass.Mass[MsgSliceUint8Custom2]
	mass11 *mass.Mass[MsgSliceUint8Custom]
	mass12 *mass.Mass[MsgArrayFloat64Custom]
	mass13 *mass.Mass[MsgArrayFloat32Custom]
	mass14 *mass.Mass[MsgArrayInt8Custom]
	mass15 *mass.Mass[MsgArrayUint8Custom2]
	mass16 *mass.Mass[MsgArrayUint8Custom]
	mass17 *mass.Mass[MsgMapCustom]
	mass18 *mass.Mass[MsgSliceCustom]
	mass20 *mass.Mass[MsgArrayCustom]
	mass21 *mass.Mass[MsgStringCustom]
	mass22 *mass.Mass[MsgFloat32Custom]
	mass23 *mass.Mass[MsgFloat64Custom]
	mass24 *mass.Mass[MsgBoolCustom]
	mass25 *mass.Mass[MsgInt8Custom]
	mass26 *mass.Mass[MsgInt16Custom]
	mass27 *mass.Mass[MsgInt32Custom]
	mass28 *mass.Mass[MsgInt64Custom]
	mass29 *mass.Mass[MsgUint8Custom]
	mass30 *mass.Mass[MsgUint16Custom]
	mass31 *mass.Mass[MsgUint32Custom]
	mass32 *mass.Mass[MsgUint64Custom]
	mass34 *mass.Mass[MsgMixed]
	mass39 *mass.Mass[MsgSliceFloat64]
	mass41 *mass.Mass[MsgSliceFloat32]
	mass43 *mass.Mass[MsgSliceInt8]
	mass45 *mass.Mass[MsgSliceUint8]
	mass46 *mass.Mass[MsgArrayFloat64]
	mass47 *mass.Mass[MsgArrayFloat32]
	mass48 *mass.Mass[MsgArrayInt8]
	mass49 *mass.Mass[MsgArrayUint8]
	mass52 *mass.Mass[MsgStructAnonymous]
	mass54 *mass.Mass[MsgStruct]
	mass55 *mass.Mass[MsgMapString]
	mass56 *mass.Mass[MsgMap]
	mass57 *mass.Mass[MsgSlice]
	mass59 *mass.Mass[MsgArray]
	mass60 *mass.Mass[MsgString]
	mass61 *mass.Mass[MsgFloat32]
	mass62 *mass.Mass[MsgFloat64]
	mass63 *mass.Mass[MsgBool10]
	mass64 *mass.Mass[MsgBool3]
	mass65 *mass.Mass[MsgInt8]
	mass66 *mass.Mass[MsgInt16]
	mass67 *mass.Mass[MsgInt32]
	mass68 *mass.Mass[MsgInt64]
	mass69 *mass.Mass[MsgUint8]
	mass70 *mass.Mass[MsgUint16]
	mass71 *mass.Mass[MsgUint32]
	mass72 *mass.Mass[MsgUint64]
	mass51 *mass.Mass[pkg2.SubMsg]
	mass2  *mass.Mass[[3][]map[int16]custom.Array]
	mass1  *mass.Mass[map[int16]custom.Array]
	mass4  *mass.Mass[custom.Float64]
	mass6  *mass.Mass[custom.Float32]
	mass8  *mass.Mass[custom.Int8]
	mass10 *mass.Mass[custom.Uint8]
	mass19 *mass.Mass[custom.Uint64]
	mass35 *mass.Mass[string]
	mass36 *mass.Mass[[32]uint16]
	mass38 *mass.Mass[[3][]map[int16][2]int64]
	mass37 *mass.Mass[map[int16][2]int64]
	mass40 *mass.Mass[float64]
	mass42 *mass.Mass[float32]
	mass44 *mass.Mass[int8]
	mass33 *mass.Mass[spkg.SubMsg]
	mass53 *mass.Mass[spkg6.SubMsg]
	mass58 *mass.Mass[bool]
}

// Size computes the size of marshalled message.
func (m Marshaller) Size(msg any) (uint64, error) {
	switch msg2 := msg.(type) {
	case *MsgMixedCustom:
		return size0(msg2), nil
	case *MsgSliceFloat64Custom:
		return size3(msg2), nil
	case *MsgSliceFloat32Custom:
		return size5(msg2), nil
	case *MsgSliceInt8Custom:
		return size7(msg2), nil
	case *MsgSliceUint8Custom2:
		return size9(msg2), nil
	case *MsgSliceUint8Custom:
		return size11(msg2), nil
	case *MsgArrayFloat64Custom:
		return size12(msg2), nil
	case *MsgArrayFloat32Custom:
		return size13(msg2), nil
	case *MsgArrayInt8Custom:
		return size14(msg2), nil
	case *MsgArrayUint8Custom2:
		return size15(msg2), nil
	case *MsgArrayUint8Custom:
		return size16(msg2), nil
	case *MsgMapCustom:
		return size17(msg2), nil
	case *MsgSliceCustom:
		return size18(msg2), nil
	case *MsgArrayCustom:
		return size20(msg2), nil
	case *MsgStringCustom:
		return size21(msg2), nil
	case *MsgFloat32Custom:
		return size22(msg2), nil
	case *MsgFloat64Custom:
		return size23(msg2), nil
	case *MsgBoolCustom:
		return size24(msg2), nil
	case *MsgInt8Custom:
		return size25(msg2), nil
	case *MsgInt16Custom:
		return size26(msg2), nil
	case *MsgInt32Custom:
		return size27(msg2), nil
	case *MsgInt64Custom:
		return size28(msg2), nil
	case *MsgUint8Custom:
		return size29(msg2), nil
	case *MsgUint16Custom:
		return size30(msg2), nil
	case *MsgUint32Custom:
		return size31(msg2), nil
	case *MsgUint64Custom:
		return size32(msg2), nil
	case *MsgMixed:
		return size34(msg2), nil
	case *MsgSliceFloat64:
		return size39(msg2), nil
	case *MsgSliceFloat32:
		return size41(msg2), nil
	case *MsgSliceInt8:
		return size43(msg2), nil
	case *MsgSliceUint8:
		return size45(msg2), nil
	case *MsgArrayFloat64:
		return size46(msg2), nil
	case *MsgArrayFloat32:
		return size47(msg2), nil
	case *MsgArrayInt8:
		return size48(msg2), nil
	case *MsgArrayUint8:
		return size49(msg2), nil
	case *MsgStructAnonymous:
		return size52(msg2), nil
	case *MsgStruct:
		return size54(msg2), nil
	case *MsgMapString:
		return size55(msg2), nil
	case *MsgMap:
		return size56(msg2), nil
	case *MsgSlice:
		return size57(msg2), nil
	case *MsgArray:
		return size59(msg2), nil
	case *MsgString:
		return size60(msg2), nil
	case *MsgFloat32:
		return size61(msg2), nil
	case *MsgFloat64:
		return size62(msg2), nil
	case *MsgBool10:
		return size63(msg2), nil
	case *MsgBool3:
		return size64(msg2), nil
	case *MsgInt8:
		return size65(msg2), nil
	case *MsgInt16:
		return size66(msg2), nil
	case *MsgInt32:
		return size67(msg2), nil
	case *MsgInt64:
		return size68(msg2), nil
	case *MsgUint8:
		return size69(msg2), nil
	case *MsgUint16:
		return size70(msg2), nil
	case *MsgUint32:
		return size71(msg2), nil
	case *MsgUint64:
		return size72(msg2), nil
	case *pkg2.SubMsg:
		return size51(msg2), nil
	default:
		return 0, errors.Errorf("unknown message type %T", msg)
	}
}

// Marshal marshals message.
func (m Marshaller) Marshal(msg any, buf []byte) (retID, retSize uint64, retErr error) {
	defer func() {
		if res := recover(); res != nil {
			retErr = errors.Errorf("marshaling message failed: %s", res)
		}
	}()

	switch msg2 := msg.(type) {
	case *MsgMixedCustom:
		return id0, marshal0(msg2, buf), nil
	case *MsgSliceFloat64Custom:
		return id3, marshal3(msg2, buf), nil
	case *MsgSliceFloat32Custom:
		return id5, marshal5(msg2, buf), nil
	case *MsgSliceInt8Custom:
		return id7, marshal7(msg2, buf), nil
	case *MsgSliceUint8Custom2:
		return id9, marshal9(msg2, buf), nil
	case *MsgSliceUint8Custom:
		return id11, marshal11(msg2, buf), nil
	case *MsgArrayFloat64Custom:
		return id12, marshal12(msg2, buf), nil
	case *MsgArrayFloat32Custom:
		return id13, marshal13(msg2, buf), nil
	case *MsgArrayInt8Custom:
		return id14, marshal14(msg2, buf), nil
	case *MsgArrayUint8Custom2:
		return id15, marshal15(msg2, buf), nil
	case *MsgArrayUint8Custom:
		return id16, marshal16(msg2, buf), nil
	case *MsgMapCustom:
		return id17, marshal17(msg2, buf), nil
	case *MsgSliceCustom:
		return id18, marshal18(msg2, buf), nil
	case *MsgArrayCustom:
		return id20, marshal20(msg2, buf), nil
	case *MsgStringCustom:
		return id21, marshal21(msg2, buf), nil
	case *MsgFloat32Custom:
		return id22, marshal22(msg2, buf), nil
	case *MsgFloat64Custom:
		return id23, marshal23(msg2, buf), nil
	case *MsgBoolCustom:
		return id24, marshal24(msg2, buf), nil
	case *MsgInt8Custom:
		return id25, marshal25(msg2, buf), nil
	case *MsgInt16Custom:
		return id26, marshal26(msg2, buf), nil
	case *MsgInt32Custom:
		return id27, marshal27(msg2, buf), nil
	case *MsgInt64Custom:
		return id28, marshal28(msg2, buf), nil
	case *MsgUint8Custom:
		return id29, marshal29(msg2, buf), nil
	case *MsgUint16Custom:
		return id30, marshal30(msg2, buf), nil
	case *MsgUint32Custom:
		return id31, marshal31(msg2, buf), nil
	case *MsgUint64Custom:
		return id32, marshal32(msg2, buf), nil
	case *MsgMixed:
		return id34, marshal34(msg2, buf), nil
	case *MsgSliceFloat64:
		return id39, marshal39(msg2, buf), nil
	case *MsgSliceFloat32:
		return id41, marshal41(msg2, buf), nil
	case *MsgSliceInt8:
		return id43, marshal43(msg2, buf), nil
	case *MsgSliceUint8:
		return id45, marshal45(msg2, buf), nil
	case *MsgArrayFloat64:
		return id46, marshal46(msg2, buf), nil
	case *MsgArrayFloat32:
		return id47, marshal47(msg2, buf), nil
	case *MsgArrayInt8:
		return id48, marshal48(msg2, buf), nil
	case *MsgArrayUint8:
		return id49, marshal49(msg2, buf), nil
	case *MsgStructAnonymous:
		return id52, marshal52(msg2, buf), nil
	case *MsgStruct:
		return id54, marshal54(msg2, buf), nil
	case *MsgMapString:
		return id55, marshal55(msg2, buf), nil
	case *MsgMap:
		return id56, marshal56(msg2, buf), nil
	case *MsgSlice:
		return id57, marshal57(msg2, buf), nil
	case *MsgArray:
		return id59, marshal59(msg2, buf), nil
	case *MsgString:
		return id60, marshal60(msg2, buf), nil
	case *MsgFloat32:
		return id61, marshal61(msg2, buf), nil
	case *MsgFloat64:
		return id62, marshal62(msg2, buf), nil
	case *MsgBool10:
		return id63, marshal63(msg2, buf), nil
	case *MsgBool3:
		return id64, marshal64(msg2, buf), nil
	case *MsgInt8:
		return id65, marshal65(msg2, buf), nil
	case *MsgInt16:
		return id66, marshal66(msg2, buf), nil
	case *MsgInt32:
		return id67, marshal67(msg2, buf), nil
	case *MsgInt64:
		return id68, marshal68(msg2, buf), nil
	case *MsgUint8:
		return id69, marshal69(msg2, buf), nil
	case *MsgUint16:
		return id70, marshal70(msg2, buf), nil
	case *MsgUint32:
		return id71, marshal71(msg2, buf), nil
	case *MsgUint64:
		return id72, marshal72(msg2, buf), nil
	case *pkg2.SubMsg:
		return id51, marshal51(msg2, buf), nil
	default:
		return 0, 0, errors.Errorf("unknown message type %T", msg)
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
	case id0:
		msg := m.mass0.New()
		return msg, unmarshal0(
			msg,
			buf,
			m.mass2,
			m.mass1,
		), nil
	case id3:
		msg := m.mass3.New()
		return msg, unmarshal3(
			msg,
			buf,
			m.mass4,
		), nil
	case id5:
		msg := m.mass5.New()
		return msg, unmarshal5(
			msg,
			buf,
			m.mass6,
		), nil
	case id7:
		msg := m.mass7.New()
		return msg, unmarshal7(
			msg,
			buf,
			m.mass8,
		), nil
	case id9:
		msg := m.mass9.New()
		return msg, unmarshal9(
			msg,
			buf,
			m.mass10,
		), nil
	case id11:
		msg := m.mass11.New()
		return msg, unmarshal11(
			msg,
			buf,
		), nil
	case id12:
		msg := m.mass12.New()
		return msg, unmarshal12(
			msg,
			buf,
		), nil
	case id13:
		msg := m.mass13.New()
		return msg, unmarshal13(
			msg,
			buf,
		), nil
	case id14:
		msg := m.mass14.New()
		return msg, unmarshal14(
			msg,
			buf,
		), nil
	case id15:
		msg := m.mass15.New()
		return msg, unmarshal15(
			msg,
			buf,
		), nil
	case id16:
		msg := m.mass16.New()
		return msg, unmarshal16(
			msg,
			buf,
		), nil
	case id17:
		msg := m.mass17.New()
		return msg, unmarshal17(
			msg,
			buf,
		), nil
	case id18:
		msg := m.mass18.New()
		return msg, unmarshal18(
			msg,
			buf,
			m.mass19,
		), nil
	case id20:
		msg := m.mass20.New()
		return msg, unmarshal20(
			msg,
			buf,
		), nil
	case id21:
		msg := m.mass21.New()
		return msg, unmarshal21(
			msg,
			buf,
		), nil
	case id22:
		msg := m.mass22.New()
		return msg, unmarshal22(
			msg,
			buf,
		), nil
	case id23:
		msg := m.mass23.New()
		return msg, unmarshal23(
			msg,
			buf,
		), nil
	case id24:
		msg := m.mass24.New()
		return msg, unmarshal24(
			msg,
			buf,
		), nil
	case id25:
		msg := m.mass25.New()
		return msg, unmarshal25(
			msg,
			buf,
		), nil
	case id26:
		msg := m.mass26.New()
		return msg, unmarshal26(
			msg,
			buf,
		), nil
	case id27:
		msg := m.mass27.New()
		return msg, unmarshal27(
			msg,
			buf,
		), nil
	case id28:
		msg := m.mass28.New()
		return msg, unmarshal28(
			msg,
			buf,
		), nil
	case id29:
		msg := m.mass29.New()
		return msg, unmarshal29(
			msg,
			buf,
		), nil
	case id30:
		msg := m.mass30.New()
		return msg, unmarshal30(
			msg,
			buf,
		), nil
	case id31:
		msg := m.mass31.New()
		return msg, unmarshal31(
			msg,
			buf,
		), nil
	case id32:
		msg := m.mass32.New()
		return msg, unmarshal32(
			msg,
			buf,
		), nil
	case id34:
		msg := m.mass34.New()
		return msg, unmarshal34(
			msg,
			buf,
			m.mass35,
			m.mass36,
			m.mass38,
			m.mass37,
		), nil
	case id39:
		msg := m.mass39.New()
		return msg, unmarshal39(
			msg,
			buf,
			m.mass40,
		), nil
	case id41:
		msg := m.mass41.New()
		return msg, unmarshal41(
			msg,
			buf,
			m.mass42,
		), nil
	case id43:
		msg := m.mass43.New()
		return msg, unmarshal43(
			msg,
			buf,
			m.mass44,
		), nil
	case id45:
		msg := m.mass45.New()
		return msg, unmarshal45(
			msg,
			buf,
		), nil
	case id46:
		msg := m.mass46.New()
		return msg, unmarshal46(
			msg,
			buf,
		), nil
	case id47:
		msg := m.mass47.New()
		return msg, unmarshal47(
			msg,
			buf,
		), nil
	case id48:
		msg := m.mass48.New()
		return msg, unmarshal48(
			msg,
			buf,
		), nil
	case id49:
		msg := m.mass49.New()
		return msg, unmarshal49(
			msg,
			buf,
		), nil
	case id52:
		msg := m.mass52.New()
		return msg, unmarshal52(
			msg,
			buf,
			m.mass42,
			m.mass33,
			m.mass53,
		), nil
	case id54:
		msg := m.mass54.New()
		return msg, unmarshal54(
			msg,
			buf,
			m.mass42,
			m.mass33,
			m.mass53,
		), nil
	case id55:
		msg := m.mass55.New()
		return msg, unmarshal55(
			msg,
			buf,
		), nil
	case id56:
		msg := m.mass56.New()
		return msg, unmarshal56(
			msg,
			buf,
		), nil
	case id57:
		msg := m.mass57.New()
		return msg, unmarshal57(
			msg,
			buf,
			m.mass58,
		), nil
	case id59:
		msg := m.mass59.New()
		return msg, unmarshal59(
			msg,
			buf,
		), nil
	case id60:
		msg := m.mass60.New()
		return msg, unmarshal60(
			msg,
			buf,
		), nil
	case id61:
		msg := m.mass61.New()
		return msg, unmarshal61(
			msg,
			buf,
		), nil
	case id62:
		msg := m.mass62.New()
		return msg, unmarshal62(
			msg,
			buf,
		), nil
	case id63:
		msg := m.mass63.New()
		return msg, unmarshal63(
			msg,
			buf,
		), nil
	case id64:
		msg := m.mass64.New()
		return msg, unmarshal64(
			msg,
			buf,
		), nil
	case id65:
		msg := m.mass65.New()
		return msg, unmarshal65(
			msg,
			buf,
		), nil
	case id66:
		msg := m.mass66.New()
		return msg, unmarshal66(
			msg,
			buf,
		), nil
	case id67:
		msg := m.mass67.New()
		return msg, unmarshal67(
			msg,
			buf,
		), nil
	case id68:
		msg := m.mass68.New()
		return msg, unmarshal68(
			msg,
			buf,
		), nil
	case id69:
		msg := m.mass69.New()
		return msg, unmarshal69(
			msg,
			buf,
		), nil
	case id70:
		msg := m.mass70.New()
		return msg, unmarshal70(
			msg,
			buf,
		), nil
	case id71:
		msg := m.mass71.New()
		return msg, unmarshal71(
			msg,
			buf,
		), nil
	case id72:
		msg := m.mass72.New()
		return msg, unmarshal72(
			msg,
			buf,
		), nil
	case id51:
		msg := m.mass51.New()
		return msg, unmarshal51(
			msg,
			buf,
		), nil
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

func unmarshal0(
	m *MsgMixedCustom,
	b []byte,
	mass2 *mass.Mass[[3][]map[int16]custom.Array],
	mass1 *mass.Mass[map[int16]custom.Array],
) uint64 {
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
			m.Value = mass2.NewSlice(l)
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
						m.Value[i6][i5] = mass1.NewSlice(l)
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
							} else {
								m.Value[i6][i5][i4] = nil
							}
						}
					} else {
						m.Value[i6][i5] = nil
					}
				}
			}
		} else {
			m.Value = nil
		}
	}

	return o
}

func size3(m *MsgSliceFloat64Custom) uint64 {
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

func marshal3(m *MsgSliceFloat64Custom, b []byte) uint64 {
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

func unmarshal3(
	m *MsgSliceFloat64Custom,
	b []byte,
	mass4 *mass.Mass[custom.Float64],
) uint64 {
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
			m.Value = mass4.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
			o += l * 8
		} else {
			m.Value = nil
		}
	}

	return o
}

func size5(m *MsgSliceFloat32Custom) uint64 {
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

func marshal5(m *MsgSliceFloat32Custom, b []byte) uint64 {
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

func unmarshal5(
	m *MsgSliceFloat32Custom,
	b []byte,
	mass6 *mass.Mass[custom.Float32],
) uint64 {
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
			m.Value = mass6.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
			o += l * 4
		} else {
			m.Value = nil
		}
	}

	return o
}

func size7(m *MsgSliceInt8Custom) uint64 {
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

func marshal7(m *MsgSliceInt8Custom, b []byte) uint64 {
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

func unmarshal7(
	m *MsgSliceInt8Custom,
	b []byte,
	mass8 *mass.Mass[custom.Int8],
) uint64 {
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
			m.Value = mass8.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

func size9(m *MsgSliceUint8Custom2) uint64 {
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

func marshal9(m *MsgSliceUint8Custom2, b []byte) uint64 {
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

func unmarshal9(
	m *MsgSliceUint8Custom2,
	b []byte,
	mass10 *mass.Mass[custom.Uint8],
) uint64 {
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
			m.Value = mass10.NewSlice(l)
			copy(unsafe.Slice((*byte)(&m.Value[0]), l), b[o:o+l])
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

func size11(m *MsgSliceUint8Custom) uint64 {
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

func marshal11(m *MsgSliceUint8Custom, b []byte) uint64 {
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

func unmarshal11(
	m *MsgSliceUint8Custom,
	b []byte,
) uint64 {
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
			m.Value = b[o : o+l]
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

func size12(m *MsgArrayFloat64Custom) uint64 {
	var n uint64 = 40
	return n
}

func marshal12(m *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

func unmarshal12(
	m *MsgArrayFloat64Custom,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
		o += 40
	}

	return o
}

func size13(m *MsgArrayFloat32Custom) uint64 {
	var n uint64 = 20
	return n
}

func marshal13(m *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

func unmarshal13(
	m *MsgArrayFloat32Custom,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
		o += 20
	}

	return o
}

func size14(m *MsgArrayInt8Custom) uint64 {
	var n uint64 = 5
	return n
}

func marshal14(m *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

func unmarshal14(
	m *MsgArrayInt8Custom,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size15(m *MsgArrayUint8Custom2) uint64 {
	var n uint64 = 5
	return n
}

func marshal15(m *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(&m.Value[0]), 5))
		o += 5
	}

	return o
}

func unmarshal15(
	m *MsgArrayUint8Custom2,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(&m.Value[0]), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size16(m *MsgArrayUint8Custom) uint64 {
	var n uint64 = 5
	return n
}

func marshal16(m *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

func unmarshal16(
	m *MsgArrayUint8Custom,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
		o += 5
	}

	return o
}

func size17(m *MsgMapCustom) uint64 {
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

func marshal17(m *MsgMapCustom, b []byte) uint64 {
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

func unmarshal17(
	m *MsgMapCustom,
	b []byte,
) uint64 {
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
						mk1 = custom.String(unsafe.String((*byte)(unsafe.Pointer(&b[o])), l))
						o += l
					} else {
						mk1 = ""
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
		} else {
			m.Value = nil
		}
	}

	return o
}

func size18(m *MsgSliceCustom) uint64 {
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

func marshal18(m *MsgSliceCustom, b []byte) uint64 {
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

func unmarshal18(
	m *MsgSliceCustom,
	b []byte,
	mass19 *mass.Mass[custom.Uint64],
) uint64 {
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
			m.Value = mass19.NewSlice(l)
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
		} else {
			m.Value = nil
		}
	}

	return o
}

func size20(m *MsgArrayCustom) uint64 {
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

func marshal20(m *MsgArrayCustom, b []byte) uint64 {
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

func unmarshal20(
	m *MsgArrayCustom,
	b []byte,
) uint64 {
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

func size21(m *MsgStringCustom) uint64 {
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

func marshal21(m *MsgStringCustom, b []byte) uint64 {
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

func unmarshal21(
	m *MsgStringCustom,
	b []byte,
) uint64 {
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
				m.Value = custom.String(unsafe.String((*byte)(unsafe.Pointer(&b[o])), l))
				o += l
			} else {
				m.Value = ""
			}
		}
	}

	return o
}

func size22(m *MsgFloat32Custom) uint64 {
	var n uint64 = 4
	return n
}

func marshal22(m *MsgFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

func unmarshal22(
	m *MsgFloat32Custom,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*custom.Float32)(unsafe.Pointer(&b[o]))
		o += 4
	}

	return o
}

func size23(m *MsgFloat64Custom) uint64 {
	var n uint64 = 8
	return n
}

func marshal23(m *MsgFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

func unmarshal23(
	m *MsgFloat64Custom,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*custom.Float64)(unsafe.Pointer(&b[o]))
		o += 8
	}

	return o
}

func size24(m *MsgBoolCustom) uint64 {
	var n uint64 = 1
	return n
}

func marshal24(m *MsgBoolCustom, b []byte) uint64 {
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

func unmarshal24(
	m *MsgBoolCustom,
	b []byte,
) uint64 {
	var o uint64 = 1
	{
		// Value

		m.Value = b[0]&0x01 != 0
	}

	return o
}

func size25(m *MsgInt8Custom) uint64 {
	var n uint64 = 1
	return n
}

func marshal25(m *MsgInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal25(
	m *MsgInt8Custom,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		m.Value = custom.Int8(b[o])
		o++
	}

	return o
}

func size26(m *MsgInt16Custom) uint64 {
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

func marshal26(m *MsgInt16Custom, b []byte) uint64 {
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

func unmarshal26(
	m *MsgInt16Custom,
	b []byte,
) uint64 {
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

func size27(m *MsgInt32Custom) uint64 {
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

func marshal27(m *MsgInt32Custom, b []byte) uint64 {
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

func unmarshal27(
	m *MsgInt32Custom,
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
			m.Value = custom.Int32(vi)
		}
	}

	return o
}

func size28(m *MsgInt64Custom) uint64 {
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

func marshal28(m *MsgInt64Custom, b []byte) uint64 {
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

func unmarshal28(
	m *MsgInt64Custom,
	b []byte,
) uint64 {
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

func size29(m *MsgUint8Custom) uint64 {
	var n uint64 = 1
	return n
}

func marshal29(m *MsgUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal29(
	m *MsgUint8Custom,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		m.Value = custom.Uint8(b[o])
		o++
	}

	return o
}

func size30(m *MsgUint16Custom) uint64 {
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

func marshal30(m *MsgUint16Custom, b []byte) uint64 {
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

func unmarshal30(
	m *MsgUint16Custom,
	b []byte,
) uint64 {
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

func size31(m *MsgUint32Custom) uint64 {
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

func marshal31(m *MsgUint32Custom, b []byte) uint64 {
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

func unmarshal31(
	m *MsgUint32Custom,
	b []byte,
) uint64 {
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

func size32(m *MsgUint64Custom) uint64 {
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

func marshal32(m *MsgUint64Custom, b []byte) uint64 {
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

func unmarshal32(
	m *MsgUint64Custom,
	b []byte,
) uint64 {
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

func size34(m *MsgMixed) uint64 {
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
			n += size33(&mv2)
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

func marshal34(m *MsgMixed, b []byte) uint64 {
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
			o += marshal33(&mv2, b[o:])
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

func unmarshal34(
	m *MsgMixed,
	b []byte,
	mass35 *mass.Mass[string],
	mass36 *mass.Mass[[32]uint16],
	mass38 *mass.Mass[[3][]map[int16][2]int64],
	mass37 *mass.Mass[map[int16][2]int64],
) uint64 {
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
						mk1 = unsafe.String((*byte)(unsafe.Pointer(&b[o])), l)
						o += l
					} else {
						mk1 = ""
					}
				}
				o += unmarshal33(
					&mv2,
					b[o:],
				)
				m.Value1[mk1] = mv2
			}
		} else {
			m.Value1 = nil
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
					mv3 = mass35.NewSlice(l)
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
								mv3[i1] = unsafe.String((*byte)(unsafe.Pointer(&b[o])), l)
								o += l
							} else {
								mv3[i1] = ""
							}
						}
					}
				} else {
					mv3 = nil
				}
				m.Value2[mk2] = mv3
			}
		} else {
			m.Value2 = nil
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
			m.Value3 = mass36.NewSlice(l)
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
		} else {
			m.Value3 = nil
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
			} else {
				m.Value4[i3] = nil
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
			m.Value5 = mass38.NewSlice(l)
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
						m.Value5[i6][i5] = mass37.NewSlice(l)
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
							} else {
								m.Value5[i6][i5][i4] = nil
							}
						}
					} else {
						m.Value5[i6][i5] = nil
					}
				}
			}
		} else {
			m.Value5 = nil
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
				m.Value8 = unsafe.String((*byte)(unsafe.Pointer(&b[o])), l)
				o += l
			} else {
				m.Value8 = ""
			}
		}
	}

	return o
}

func size33(m *spkg.SubMsg) uint64 {
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

func marshal33(m *spkg.SubMsg, b []byte) uint64 {
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

func unmarshal33(
	m *spkg.SubMsg,
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

func size39(m *MsgSliceFloat64) uint64 {
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

func marshal39(m *MsgSliceFloat64, b []byte) uint64 {
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

func unmarshal39(
	m *MsgSliceFloat64,
	b []byte,
	mass40 *mass.Mass[float64],
) uint64 {
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
			m.Value = mass40.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
			o += l * 8
		} else {
			m.Value = nil
		}
	}

	return o
}

func size41(m *MsgSliceFloat32) uint64 {
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

func marshal41(m *MsgSliceFloat32, b []byte) uint64 {
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

func unmarshal41(
	m *MsgSliceFloat32,
	b []byte,
	mass42 *mass.Mass[float32],
) uint64 {
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
			m.Value = mass42.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
			o += l * 4
		} else {
			m.Value = nil
		}
	}

	return o
}

func size43(m *MsgSliceInt8) uint64 {
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

func marshal43(m *MsgSliceInt8, b []byte) uint64 {
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

func unmarshal43(
	m *MsgSliceInt8,
	b []byte,
	mass44 *mass.Mass[int8],
) uint64 {
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
			m.Value = mass44.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

func size45(m *MsgSliceUint8) uint64 {
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

func marshal45(m *MsgSliceUint8, b []byte) uint64 {
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

func unmarshal45(
	m *MsgSliceUint8,
	b []byte,
) uint64 {
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
			m.Value = b[o : o+l]
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

func size46(m *MsgArrayFloat64) uint64 {
	var n uint64 = 40
	return n
}

func marshal46(m *MsgArrayFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

func unmarshal46(
	m *MsgArrayFloat64,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
		o += 40
	}

	return o
}

func size47(m *MsgArrayFloat32) uint64 {
	var n uint64 = 20
	return n
}

func marshal47(m *MsgArrayFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

func unmarshal47(
	m *MsgArrayFloat32,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
		o += 20
	}

	return o
}

func size48(m *MsgArrayInt8) uint64 {
	var n uint64 = 5
	return n
}

func marshal48(m *MsgArrayInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

func unmarshal48(
	m *MsgArrayInt8,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
		o += 5
	}

	return o
}

func size49(m *MsgArrayUint8) uint64 {
	var n uint64 = 5
	return n
}

func marshal49(m *MsgArrayUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

func unmarshal49(
	m *MsgArrayUint8,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
		o += 5
	}

	return o
}

func size52(m *MsgStructAnonymous) uint64 {
	var n uint64
	{
		// SubMsg

		n += size50(&m.SubMsg)
	}
	{
		// Value2

		n += size51(&m.Value2)
	}
	return n
}

func marshal52(m *MsgStructAnonymous, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += marshal50(&m.SubMsg, b[o:])
	}
	{
		// Value2

		o += marshal51(&m.Value2, b[o:])
	}

	return o
}

func unmarshal52(
	m *MsgStructAnonymous,
	b []byte,
	mass42 *mass.Mass[float32],
	mass33 *mass.Mass[spkg.SubMsg],
	mass53 *mass.Mass[spkg6.SubMsg],
) uint64 {
	var o uint64
	{
		// SubMsg

		o += unmarshal50(
			&m.SubMsg,
			b[o:],
			mass42,
			mass33,
			mass53,
		)
	}
	{
		// Value2

		o += unmarshal51(
			&m.Value2,
			b[o:],
		)
	}

	return o
}

func size50(m *SubMsg) uint64 {
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
			n += size33(&sv1)
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
			n += size53(&sv1)
		}
	}
	return n
}

func marshal50(m *SubMsg, b []byte) uint64 {
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
			o += marshal33(&sv1, b[o:])
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
			o += marshal53(&sv1, b[o:])
		}
	}

	return o
}

func unmarshal50(
	m *SubMsg,
	b []byte,
	mass42 *mass.Mass[float32],
	mass33 *mass.Mass[spkg.SubMsg],
	mass53 *mass.Mass[spkg6.SubMsg],
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
			m.Value1 = mass42.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value1[0])), l*4), b[o:o+l*4])
			o += l * 4
		} else {
			m.Value1 = nil
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
			m.Value2 = mass33.NewSlice(l)
			for i1 := range l {
				o += unmarshal33(
					&m.Value2[i1],
					b[o:],
				)
			}
		} else {
			m.Value2 = nil
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
			m.Value3 = mass53.NewSlice(l)
			for i1 := range l {
				o += unmarshal53(
					&m.Value3[i1],
					b[o:],
				)
			}
		} else {
			m.Value3 = nil
		}
	}

	return o
}

func size53(m *spkg6.SubMsg) uint64 {
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

func marshal53(m *spkg6.SubMsg, b []byte) uint64 {
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

func unmarshal53(
	m *spkg6.SubMsg,
	b []byte,
) uint64 {
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

func size54(m *MsgStruct) uint64 {
	var n uint64
	{
		// Value1

		n += size50(&m.Value1)
	}
	{
		// Value2

		n += size51(&m.Value2)
	}
	return n
}

func marshal54(m *MsgStruct, b []byte) uint64 {
	var o uint64
	{
		// Value1

		o += marshal50(&m.Value1, b[o:])
	}
	{
		// Value2

		o += marshal51(&m.Value2, b[o:])
	}

	return o
}

func unmarshal54(
	m *MsgStruct,
	b []byte,
	mass42 *mass.Mass[float32],
	mass33 *mass.Mass[spkg.SubMsg],
	mass53 *mass.Mass[spkg6.SubMsg],
) uint64 {
	var o uint64
	{
		// Value1

		o += unmarshal50(
			&m.Value1,
			b[o:],
			mass42,
			mass33,
			mass53,
		)
	}
	{
		// Value2

		o += unmarshal51(
			&m.Value2,
			b[o:],
		)
	}

	return o
}

func size55(m *MsgMapString) uint64 {
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

func marshal55(m *MsgMapString, b []byte) uint64 {
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

func unmarshal55(
	m *MsgMapString,
	b []byte,
) uint64 {
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
						mk1 = unsafe.String((*byte)(unsafe.Pointer(&b[o])), l)
						o += l
					} else {
						mk1 = ""
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
						mv2 = unsafe.String((*byte)(unsafe.Pointer(&b[o])), l)
						o += l
					} else {
						mv2 = ""
					}
				}
				m.Value[mk1] = mv2
			}
		} else {
			m.Value = nil
		}
	}

	return o
}

func size56(m *MsgMap) uint64 {
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

func marshal56(m *MsgMap, b []byte) uint64 {
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

func unmarshal56(
	m *MsgMap,
	b []byte,
) uint64 {
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
		} else {
			m.Value = nil
		}
	}

	return o
}

func size57(m *MsgSlice) uint64 {
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

func marshal57(m *MsgSlice, b []byte) uint64 {
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

func unmarshal57(
	m *MsgSlice,
	b []byte,
	mass58 *mass.Mass[bool],
) uint64 {
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
			m.Value = mass58.NewSlice(l)
			for i1 := range l {
				m.Value[i1] = b[o] != 0x00
				o++
			}
		} else {
			m.Value = nil
		}
	}

	return o
}

func size59(m *MsgArray) uint64 {
	var n uint64 = 3
	return n
}

func marshal59(m *MsgArray, b []byte) uint64 {
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

func unmarshal59(
	m *MsgArray,
	b []byte,
) uint64 {
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

func size60(m *MsgString) uint64 {
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

func marshal60(m *MsgString, b []byte) uint64 {
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

func unmarshal60(
	m *MsgString,
	b []byte,
) uint64 {
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
				m.Value = unsafe.String((*byte)(unsafe.Pointer(&b[o])), l)
				o += l
			} else {
				m.Value = ""
			}
		}
	}

	return o
}

func size61(m *MsgFloat32) uint64 {
	var n uint64 = 4
	return n
}

func marshal61(m *MsgFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

func unmarshal61(
	m *MsgFloat32,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*float32)(unsafe.Pointer(&b[o]))
		o += 4
	}

	return o
}

func size62(m *MsgFloat64) uint64 {
	var n uint64 = 8
	return n
}

func marshal62(m *MsgFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

func unmarshal62(
	m *MsgFloat64,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*float64)(unsafe.Pointer(&b[o]))
		o += 8
	}

	return o
}

func size63(m *MsgBool10) uint64 {
	var n uint64 = 2
	return n
}

func marshal63(m *MsgBool10, b []byte) uint64 {
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

func unmarshal63(
	m *MsgBool10,
	b []byte,
) uint64 {
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

func size64(m *MsgBool3) uint64 {
	var n uint64 = 1
	return n
}

func marshal64(m *MsgBool3, b []byte) uint64 {
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

func unmarshal64(
	m *MsgBool3,
	b []byte,
) uint64 {
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

func size65(m *MsgInt8) uint64 {
	var n uint64 = 1
	return n
}

func marshal65(m *MsgInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal65(
	m *MsgInt8,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		m.Value = int8(b[o])
		o++
	}

	return o
}

func size66(m *MsgInt16) uint64 {
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

func marshal66(m *MsgInt16, b []byte) uint64 {
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

func unmarshal66(
	m *MsgInt16,
	b []byte,
) uint64 {
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

func size67(m *MsgInt32) uint64 {
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

func marshal67(m *MsgInt32, b []byte) uint64 {
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

func unmarshal67(
	m *MsgInt32,
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

func size68(m *MsgInt64) uint64 {
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

func marshal68(m *MsgInt64, b []byte) uint64 {
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

func unmarshal68(
	m *MsgInt64,
	b []byte,
) uint64 {
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

func size69(m *MsgUint8) uint64 {
	var n uint64 = 1
	return n
}

func marshal69(m *MsgUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = m.Value
		o++
	}

	return o
}

func unmarshal69(
	m *MsgUint8,
	b []byte,
) uint64 {
	var o uint64
	{
		// Value

		m.Value = b[o]
		o++
	}

	return o
}

func size70(m *MsgUint16) uint64 {
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

func marshal70(m *MsgUint16, b []byte) uint64 {
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

func unmarshal70(
	m *MsgUint16,
	b []byte,
) uint64 {
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

func size71(m *MsgUint32) uint64 {
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

func marshal71(m *MsgUint32, b []byte) uint64 {
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

func unmarshal71(
	m *MsgUint32,
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
			m.Value = vi
		}
	}

	return o
}

func size72(m *MsgUint64) uint64 {
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

func marshal72(m *MsgUint64, b []byte) uint64 {
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

func unmarshal72(
	m *MsgUint64,
	b []byte,
) uint64 {
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

func size51(m *pkg2.SubMsg) uint64 {
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

func marshal51(m *pkg2.SubMsg, b []byte) uint64 {
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

func unmarshal51(
	m *pkg2.SubMsg,
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
