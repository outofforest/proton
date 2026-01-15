package pkg1

import (
	"reflect"
	"time"
	"unsafe"

	"github.com/outofforest/proton"
	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1/spkg"
	"github.com/outofforest/proton/test/pkg2"
	spkg8 "github.com/outofforest/proton/test/pkg2/spkg"
	"github.com/pkg/errors"
)

const (
	id41 uint64 = iota + 1
	id61
	id60
	id59
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
	id42
	id39
	id38
	id37
	id36
	id35
	id34
	id33
	id32
	id31
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
	id1
	id2
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
		MsgIgnore{},
		MsgNotIgnore{},
		MsgStrings{},
	}
}

// ID returns ID of message type.
func (m Marshaller) ID(msg any) (uint64, error) {
	switch msg.(type) {
	case *pkg2.SubMsg:
		return id41, nil
	case *MsgUint64:
		return id61, nil
	case *MsgUint32:
		return id60, nil
	case *MsgUint16:
		return id59, nil
	case *MsgUint8:
		return id58, nil
	case *MsgInt64:
		return id57, nil
	case *MsgInt32:
		return id56, nil
	case *MsgInt16:
		return id55, nil
	case *MsgInt8:
		return id54, nil
	case *MsgBool3:
		return id53, nil
	case *MsgBool10:
		return id52, nil
	case *MsgFloat64:
		return id51, nil
	case *MsgFloat32:
		return id50, nil
	case *MsgString:
		return id49, nil
	case *MsgArray:
		return id48, nil
	case *MsgSlice:
		return id47, nil
	case *MsgMap:
		return id46, nil
	case *MsgMapString:
		return id45, nil
	case *MsgStruct:
		return id44, nil
	case *MsgStructAnonymous:
		return id42, nil
	case *MsgArrayUint8:
		return id39, nil
	case *MsgArrayInt8:
		return id38, nil
	case *MsgArrayFloat32:
		return id37, nil
	case *MsgArrayFloat64:
		return id36, nil
	case *MsgSliceUint8:
		return id35, nil
	case *MsgSliceInt8:
		return id34, nil
	case *MsgSliceFloat32:
		return id33, nil
	case *MsgSliceFloat64:
		return id32, nil
	case *MsgMixed:
		return id31, nil
	case *MsgUint64Custom:
		return id29, nil
	case *MsgUint32Custom:
		return id28, nil
	case *MsgUint16Custom:
		return id27, nil
	case *MsgUint8Custom:
		return id26, nil
	case *MsgInt64Custom:
		return id25, nil
	case *MsgInt32Custom:
		return id24, nil
	case *MsgInt16Custom:
		return id23, nil
	case *MsgInt8Custom:
		return id22, nil
	case *MsgBoolCustom:
		return id21, nil
	case *MsgFloat64Custom:
		return id20, nil
	case *MsgFloat32Custom:
		return id19, nil
	case *MsgStringCustom:
		return id18, nil
	case *MsgArrayCustom:
		return id17, nil
	case *MsgSliceCustom:
		return id16, nil
	case *MsgMapCustom:
		return id15, nil
	case *MsgArrayUint8Custom:
		return id14, nil
	case *MsgArrayUint8Custom2:
		return id13, nil
	case *MsgArrayInt8Custom:
		return id12, nil
	case *MsgArrayFloat32Custom:
		return id11, nil
	case *MsgArrayFloat64Custom:
		return id10, nil
	case *MsgSliceUint8Custom:
		return id9, nil
	case *MsgSliceUint8Custom2:
		return id8, nil
	case *MsgSliceInt8Custom:
		return id7, nil
	case *MsgSliceFloat32Custom:
		return id6, nil
	case *MsgSliceFloat64Custom:
		return id5, nil
	case *MsgTime:
		return id4, nil
	case *MsgMixedCustom:
		return id3, nil
	case *MsgIgnore:
		return id1, nil
	case *MsgNotIgnore:
		return id2, nil
	case *MsgStrings:
		return id0, nil
	default:
		return 0, errors.Errorf("unknown message type %T", msg)
	}
}

// Size computes the size of marshalled message.
func (m Marshaller) Size(msg any) (uint64, error) {
	switch msg2 := msg.(type) {
	case *pkg2.SubMsg:
		return size41(msg2), nil
	case *MsgUint64:
		return size61(msg2), nil
	case *MsgUint32:
		return size60(msg2), nil
	case *MsgUint16:
		return size59(msg2), nil
	case *MsgUint8:
		return size58(msg2), nil
	case *MsgInt64:
		return size57(msg2), nil
	case *MsgInt32:
		return size56(msg2), nil
	case *MsgInt16:
		return size55(msg2), nil
	case *MsgInt8:
		return size54(msg2), nil
	case *MsgBool3:
		return size53(msg2), nil
	case *MsgBool10:
		return size52(msg2), nil
	case *MsgFloat64:
		return size51(msg2), nil
	case *MsgFloat32:
		return size50(msg2), nil
	case *MsgString:
		return size49(msg2), nil
	case *MsgArray:
		return size48(msg2), nil
	case *MsgSlice:
		return size47(msg2), nil
	case *MsgMap:
		return size46(msg2), nil
	case *MsgMapString:
		return size45(msg2), nil
	case *MsgStruct:
		return size44(msg2), nil
	case *MsgStructAnonymous:
		return size42(msg2), nil
	case *MsgArrayUint8:
		return size39(msg2), nil
	case *MsgArrayInt8:
		return size38(msg2), nil
	case *MsgArrayFloat32:
		return size37(msg2), nil
	case *MsgArrayFloat64:
		return size36(msg2), nil
	case *MsgSliceUint8:
		return size35(msg2), nil
	case *MsgSliceInt8:
		return size34(msg2), nil
	case *MsgSliceFloat32:
		return size33(msg2), nil
	case *MsgSliceFloat64:
		return size32(msg2), nil
	case *MsgMixed:
		return size31(msg2), nil
	case *MsgUint64Custom:
		return size29(msg2), nil
	case *MsgUint32Custom:
		return size28(msg2), nil
	case *MsgUint16Custom:
		return size27(msg2), nil
	case *MsgUint8Custom:
		return size26(msg2), nil
	case *MsgInt64Custom:
		return size25(msg2), nil
	case *MsgInt32Custom:
		return size24(msg2), nil
	case *MsgInt16Custom:
		return size23(msg2), nil
	case *MsgInt8Custom:
		return size22(msg2), nil
	case *MsgBoolCustom:
		return size21(msg2), nil
	case *MsgFloat64Custom:
		return size20(msg2), nil
	case *MsgFloat32Custom:
		return size19(msg2), nil
	case *MsgStringCustom:
		return size18(msg2), nil
	case *MsgArrayCustom:
		return size17(msg2), nil
	case *MsgSliceCustom:
		return size16(msg2), nil
	case *MsgMapCustom:
		return size15(msg2), nil
	case *MsgArrayUint8Custom:
		return size14(msg2), nil
	case *MsgArrayUint8Custom2:
		return size13(msg2), nil
	case *MsgArrayInt8Custom:
		return size12(msg2), nil
	case *MsgArrayFloat32Custom:
		return size11(msg2), nil
	case *MsgArrayFloat64Custom:
		return size10(msg2), nil
	case *MsgSliceUint8Custom:
		return size9(msg2), nil
	case *MsgSliceUint8Custom2:
		return size8(msg2), nil
	case *MsgSliceInt8Custom:
		return size7(msg2), nil
	case *MsgSliceFloat32Custom:
		return size6(msg2), nil
	case *MsgSliceFloat64Custom:
		return size5(msg2), nil
	case *MsgTime:
		return size4(msg2), nil
	case *MsgMixedCustom:
		return size3(msg2), nil
	case *MsgIgnore:
		return sizei1(msg2), nil
	case *MsgNotIgnore:
		return size2(msg2), nil
	case *MsgStrings:
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
		return id41, marshal41(msg2, buf), nil
	case *MsgUint64:
		return id61, marshal61(msg2, buf), nil
	case *MsgUint32:
		return id60, marshal60(msg2, buf), nil
	case *MsgUint16:
		return id59, marshal59(msg2, buf), nil
	case *MsgUint8:
		return id58, marshal58(msg2, buf), nil
	case *MsgInt64:
		return id57, marshal57(msg2, buf), nil
	case *MsgInt32:
		return id56, marshal56(msg2, buf), nil
	case *MsgInt16:
		return id55, marshal55(msg2, buf), nil
	case *MsgInt8:
		return id54, marshal54(msg2, buf), nil
	case *MsgBool3:
		return id53, marshal53(msg2, buf), nil
	case *MsgBool10:
		return id52, marshal52(msg2, buf), nil
	case *MsgFloat64:
		return id51, marshal51(msg2, buf), nil
	case *MsgFloat32:
		return id50, marshal50(msg2, buf), nil
	case *MsgString:
		return id49, marshal49(msg2, buf), nil
	case *MsgArray:
		return id48, marshal48(msg2, buf), nil
	case *MsgSlice:
		return id47, marshal47(msg2, buf), nil
	case *MsgMap:
		return id46, marshal46(msg2, buf), nil
	case *MsgMapString:
		return id45, marshal45(msg2, buf), nil
	case *MsgStruct:
		return id44, marshal44(msg2, buf), nil
	case *MsgStructAnonymous:
		return id42, marshal42(msg2, buf), nil
	case *MsgArrayUint8:
		return id39, marshal39(msg2, buf), nil
	case *MsgArrayInt8:
		return id38, marshal38(msg2, buf), nil
	case *MsgArrayFloat32:
		return id37, marshal37(msg2, buf), nil
	case *MsgArrayFloat64:
		return id36, marshal36(msg2, buf), nil
	case *MsgSliceUint8:
		return id35, marshal35(msg2, buf), nil
	case *MsgSliceInt8:
		return id34, marshal34(msg2, buf), nil
	case *MsgSliceFloat32:
		return id33, marshal33(msg2, buf), nil
	case *MsgSliceFloat64:
		return id32, marshal32(msg2, buf), nil
	case *MsgMixed:
		return id31, marshal31(msg2, buf), nil
	case *MsgUint64Custom:
		return id29, marshal29(msg2, buf), nil
	case *MsgUint32Custom:
		return id28, marshal28(msg2, buf), nil
	case *MsgUint16Custom:
		return id27, marshal27(msg2, buf), nil
	case *MsgUint8Custom:
		return id26, marshal26(msg2, buf), nil
	case *MsgInt64Custom:
		return id25, marshal25(msg2, buf), nil
	case *MsgInt32Custom:
		return id24, marshal24(msg2, buf), nil
	case *MsgInt16Custom:
		return id23, marshal23(msg2, buf), nil
	case *MsgInt8Custom:
		return id22, marshal22(msg2, buf), nil
	case *MsgBoolCustom:
		return id21, marshal21(msg2, buf), nil
	case *MsgFloat64Custom:
		return id20, marshal20(msg2, buf), nil
	case *MsgFloat32Custom:
		return id19, marshal19(msg2, buf), nil
	case *MsgStringCustom:
		return id18, marshal18(msg2, buf), nil
	case *MsgArrayCustom:
		return id17, marshal17(msg2, buf), nil
	case *MsgSliceCustom:
		return id16, marshal16(msg2, buf), nil
	case *MsgMapCustom:
		return id15, marshal15(msg2, buf), nil
	case *MsgArrayUint8Custom:
		return id14, marshal14(msg2, buf), nil
	case *MsgArrayUint8Custom2:
		return id13, marshal13(msg2, buf), nil
	case *MsgArrayInt8Custom:
		return id12, marshal12(msg2, buf), nil
	case *MsgArrayFloat32Custom:
		return id11, marshal11(msg2, buf), nil
	case *MsgArrayFloat64Custom:
		return id10, marshal10(msg2, buf), nil
	case *MsgSliceUint8Custom:
		return id9, marshal9(msg2, buf), nil
	case *MsgSliceUint8Custom2:
		return id8, marshal8(msg2, buf), nil
	case *MsgSliceInt8Custom:
		return id7, marshal7(msg2, buf), nil
	case *MsgSliceFloat32Custom:
		return id6, marshal6(msg2, buf), nil
	case *MsgSliceFloat64Custom:
		return id5, marshal5(msg2, buf), nil
	case *MsgTime:
		return id4, marshal4(msg2, buf), nil
	case *MsgMixedCustom:
		return id3, marshal3(msg2, buf), nil
	case *MsgIgnore:
		return id1, marshali1(msg2, buf), nil
	case *MsgNotIgnore:
		return id2, marshal2(msg2, buf), nil
	case *MsgStrings:
		return id0, marshal0(msg2, buf), nil
	default:
		return 0, 0, errors.Errorf("unknown message type %T", msg)
	}
}

// Unmarshal unmarshals message.
func (m Marshaller) Unmarshal(id uint64, buf []byte) (retMsg any, retSize uint64, retErr error) {
	defer helpers.RecoverUnmarshal(&retErr)

	switch id {
	case id41:
		msg := &pkg2.SubMsg{}
		return msg, unmarshal41(msg, buf), nil
	case id61:
		msg := &MsgUint64{}
		return msg, unmarshal61(msg, buf), nil
	case id60:
		msg := &MsgUint32{}
		return msg, unmarshal60(msg, buf), nil
	case id59:
		msg := &MsgUint16{}
		return msg, unmarshal59(msg, buf), nil
	case id58:
		msg := &MsgUint8{}
		return msg, unmarshal58(msg, buf), nil
	case id57:
		msg := &MsgInt64{}
		return msg, unmarshal57(msg, buf), nil
	case id56:
		msg := &MsgInt32{}
		return msg, unmarshal56(msg, buf), nil
	case id55:
		msg := &MsgInt16{}
		return msg, unmarshal55(msg, buf), nil
	case id54:
		msg := &MsgInt8{}
		return msg, unmarshal54(msg, buf), nil
	case id53:
		msg := &MsgBool3{}
		return msg, unmarshal53(msg, buf), nil
	case id52:
		msg := &MsgBool10{}
		return msg, unmarshal52(msg, buf), nil
	case id51:
		msg := &MsgFloat64{}
		return msg, unmarshal51(msg, buf), nil
	case id50:
		msg := &MsgFloat32{}
		return msg, unmarshal50(msg, buf), nil
	case id49:
		msg := &MsgString{}
		return msg, unmarshal49(msg, buf), nil
	case id48:
		msg := &MsgArray{}
		return msg, unmarshal48(msg, buf), nil
	case id47:
		msg := &MsgSlice{}
		return msg, unmarshal47(msg, buf), nil
	case id46:
		msg := &MsgMap{}
		return msg, unmarshal46(msg, buf), nil
	case id45:
		msg := &MsgMapString{}
		return msg, unmarshal45(msg, buf), nil
	case id44:
		msg := &MsgStruct{}
		return msg, unmarshal44(msg, buf), nil
	case id42:
		msg := &MsgStructAnonymous{}
		return msg, unmarshal42(msg, buf), nil
	case id39:
		msg := &MsgArrayUint8{}
		return msg, unmarshal39(msg, buf), nil
	case id38:
		msg := &MsgArrayInt8{}
		return msg, unmarshal38(msg, buf), nil
	case id37:
		msg := &MsgArrayFloat32{}
		return msg, unmarshal37(msg, buf), nil
	case id36:
		msg := &MsgArrayFloat64{}
		return msg, unmarshal36(msg, buf), nil
	case id35:
		msg := &MsgSliceUint8{}
		return msg, unmarshal35(msg, buf), nil
	case id34:
		msg := &MsgSliceInt8{}
		return msg, unmarshal34(msg, buf), nil
	case id33:
		msg := &MsgSliceFloat32{}
		return msg, unmarshal33(msg, buf), nil
	case id32:
		msg := &MsgSliceFloat64{}
		return msg, unmarshal32(msg, buf), nil
	case id31:
		msg := &MsgMixed{}
		return msg, unmarshal31(msg, buf), nil
	case id29:
		msg := &MsgUint64Custom{}
		return msg, unmarshal29(msg, buf), nil
	case id28:
		msg := &MsgUint32Custom{}
		return msg, unmarshal28(msg, buf), nil
	case id27:
		msg := &MsgUint16Custom{}
		return msg, unmarshal27(msg, buf), nil
	case id26:
		msg := &MsgUint8Custom{}
		return msg, unmarshal26(msg, buf), nil
	case id25:
		msg := &MsgInt64Custom{}
		return msg, unmarshal25(msg, buf), nil
	case id24:
		msg := &MsgInt32Custom{}
		return msg, unmarshal24(msg, buf), nil
	case id23:
		msg := &MsgInt16Custom{}
		return msg, unmarshal23(msg, buf), nil
	case id22:
		msg := &MsgInt8Custom{}
		return msg, unmarshal22(msg, buf), nil
	case id21:
		msg := &MsgBoolCustom{}
		return msg, unmarshal21(msg, buf), nil
	case id20:
		msg := &MsgFloat64Custom{}
		return msg, unmarshal20(msg, buf), nil
	case id19:
		msg := &MsgFloat32Custom{}
		return msg, unmarshal19(msg, buf), nil
	case id18:
		msg := &MsgStringCustom{}
		return msg, unmarshal18(msg, buf), nil
	case id17:
		msg := &MsgArrayCustom{}
		return msg, unmarshal17(msg, buf), nil
	case id16:
		msg := &MsgSliceCustom{}
		return msg, unmarshal16(msg, buf), nil
	case id15:
		msg := &MsgMapCustom{}
		return msg, unmarshal15(msg, buf), nil
	case id14:
		msg := &MsgArrayUint8Custom{}
		return msg, unmarshal14(msg, buf), nil
	case id13:
		msg := &MsgArrayUint8Custom2{}
		return msg, unmarshal13(msg, buf), nil
	case id12:
		msg := &MsgArrayInt8Custom{}
		return msg, unmarshal12(msg, buf), nil
	case id11:
		msg := &MsgArrayFloat32Custom{}
		return msg, unmarshal11(msg, buf), nil
	case id10:
		msg := &MsgArrayFloat64Custom{}
		return msg, unmarshal10(msg, buf), nil
	case id9:
		msg := &MsgSliceUint8Custom{}
		return msg, unmarshal9(msg, buf), nil
	case id8:
		msg := &MsgSliceUint8Custom2{}
		return msg, unmarshal8(msg, buf), nil
	case id7:
		msg := &MsgSliceInt8Custom{}
		return msg, unmarshal7(msg, buf), nil
	case id6:
		msg := &MsgSliceFloat32Custom{}
		return msg, unmarshal6(msg, buf), nil
	case id5:
		msg := &MsgSliceFloat64Custom{}
		return msg, unmarshal5(msg, buf), nil
	case id4:
		msg := &MsgTime{}
		return msg, unmarshal4(msg, buf), nil
	case id3:
		msg := &MsgMixedCustom{}
		return msg, unmarshal3(msg, buf), nil
	case id1:
		msg := &MsgIgnore{}
		return msg, unmarshali1(msg, buf), nil
	case id2:
		msg := &MsgNotIgnore{}
		return msg, unmarshal2(msg, buf), nil
	case id0:
		msg := &MsgStrings{}
		return msg, unmarshal0(msg, buf), nil
	default:
		return nil, 0, errors.Errorf("unknown ID %d", id)
	}
}

// MakePatch creates a patch.
func (m Marshaller) MakePatch(msgDst, msgSrc any, buf []byte) (retID, retSize uint64, retErr error) {
	defer helpers.RecoverMakePatch(&retErr)

	switch msg2 := msgDst.(type) {
	case *pkg2.SubMsg:
		return id41, makePatch41(msg2, msgSrc.(*pkg2.SubMsg), buf), nil
	case *MsgUint64:
		return id61, makePatch61(msg2, msgSrc.(*MsgUint64), buf), nil
	case *MsgUint32:
		return id60, makePatch60(msg2, msgSrc.(*MsgUint32), buf), nil
	case *MsgUint16:
		return id59, makePatch59(msg2, msgSrc.(*MsgUint16), buf), nil
	case *MsgUint8:
		return id58, makePatch58(msg2, msgSrc.(*MsgUint8), buf), nil
	case *MsgInt64:
		return id57, makePatch57(msg2, msgSrc.(*MsgInt64), buf), nil
	case *MsgInt32:
		return id56, makePatch56(msg2, msgSrc.(*MsgInt32), buf), nil
	case *MsgInt16:
		return id55, makePatch55(msg2, msgSrc.(*MsgInt16), buf), nil
	case *MsgInt8:
		return id54, makePatch54(msg2, msgSrc.(*MsgInt8), buf), nil
	case *MsgBool3:
		return id53, makePatch53(msg2, msgSrc.(*MsgBool3), buf), nil
	case *MsgBool10:
		return id52, makePatch52(msg2, msgSrc.(*MsgBool10), buf), nil
	case *MsgFloat64:
		return id51, makePatch51(msg2, msgSrc.(*MsgFloat64), buf), nil
	case *MsgFloat32:
		return id50, makePatch50(msg2, msgSrc.(*MsgFloat32), buf), nil
	case *MsgString:
		return id49, makePatch49(msg2, msgSrc.(*MsgString), buf), nil
	case *MsgArray:
		return id48, makePatch48(msg2, msgSrc.(*MsgArray), buf), nil
	case *MsgSlice:
		return id47, makePatch47(msg2, msgSrc.(*MsgSlice), buf), nil
	case *MsgMap:
		return id46, makePatch46(msg2, msgSrc.(*MsgMap), buf), nil
	case *MsgMapString:
		return id45, makePatch45(msg2, msgSrc.(*MsgMapString), buf), nil
	case *MsgStruct:
		return id44, makePatch44(msg2, msgSrc.(*MsgStruct), buf), nil
	case *MsgStructAnonymous:
		return id42, makePatch42(msg2, msgSrc.(*MsgStructAnonymous), buf), nil
	case *MsgArrayUint8:
		return id39, makePatch39(msg2, msgSrc.(*MsgArrayUint8), buf), nil
	case *MsgArrayInt8:
		return id38, makePatch38(msg2, msgSrc.(*MsgArrayInt8), buf), nil
	case *MsgArrayFloat32:
		return id37, makePatch37(msg2, msgSrc.(*MsgArrayFloat32), buf), nil
	case *MsgArrayFloat64:
		return id36, makePatch36(msg2, msgSrc.(*MsgArrayFloat64), buf), nil
	case *MsgSliceUint8:
		return id35, makePatch35(msg2, msgSrc.(*MsgSliceUint8), buf), nil
	case *MsgSliceInt8:
		return id34, makePatch34(msg2, msgSrc.(*MsgSliceInt8), buf), nil
	case *MsgSliceFloat32:
		return id33, makePatch33(msg2, msgSrc.(*MsgSliceFloat32), buf), nil
	case *MsgSliceFloat64:
		return id32, makePatch32(msg2, msgSrc.(*MsgSliceFloat64), buf), nil
	case *MsgMixed:
		return id31, makePatch31(msg2, msgSrc.(*MsgMixed), buf), nil
	case *MsgUint64Custom:
		return id29, makePatch29(msg2, msgSrc.(*MsgUint64Custom), buf), nil
	case *MsgUint32Custom:
		return id28, makePatch28(msg2, msgSrc.(*MsgUint32Custom), buf), nil
	case *MsgUint16Custom:
		return id27, makePatch27(msg2, msgSrc.(*MsgUint16Custom), buf), nil
	case *MsgUint8Custom:
		return id26, makePatch26(msg2, msgSrc.(*MsgUint8Custom), buf), nil
	case *MsgInt64Custom:
		return id25, makePatch25(msg2, msgSrc.(*MsgInt64Custom), buf), nil
	case *MsgInt32Custom:
		return id24, makePatch24(msg2, msgSrc.(*MsgInt32Custom), buf), nil
	case *MsgInt16Custom:
		return id23, makePatch23(msg2, msgSrc.(*MsgInt16Custom), buf), nil
	case *MsgInt8Custom:
		return id22, makePatch22(msg2, msgSrc.(*MsgInt8Custom), buf), nil
	case *MsgBoolCustom:
		return id21, makePatch21(msg2, msgSrc.(*MsgBoolCustom), buf), nil
	case *MsgFloat64Custom:
		return id20, makePatch20(msg2, msgSrc.(*MsgFloat64Custom), buf), nil
	case *MsgFloat32Custom:
		return id19, makePatch19(msg2, msgSrc.(*MsgFloat32Custom), buf), nil
	case *MsgStringCustom:
		return id18, makePatch18(msg2, msgSrc.(*MsgStringCustom), buf), nil
	case *MsgArrayCustom:
		return id17, makePatch17(msg2, msgSrc.(*MsgArrayCustom), buf), nil
	case *MsgSliceCustom:
		return id16, makePatch16(msg2, msgSrc.(*MsgSliceCustom), buf), nil
	case *MsgMapCustom:
		return id15, makePatch15(msg2, msgSrc.(*MsgMapCustom), buf), nil
	case *MsgArrayUint8Custom:
		return id14, makePatch14(msg2, msgSrc.(*MsgArrayUint8Custom), buf), nil
	case *MsgArrayUint8Custom2:
		return id13, makePatch13(msg2, msgSrc.(*MsgArrayUint8Custom2), buf), nil
	case *MsgArrayInt8Custom:
		return id12, makePatch12(msg2, msgSrc.(*MsgArrayInt8Custom), buf), nil
	case *MsgArrayFloat32Custom:
		return id11, makePatch11(msg2, msgSrc.(*MsgArrayFloat32Custom), buf), nil
	case *MsgArrayFloat64Custom:
		return id10, makePatch10(msg2, msgSrc.(*MsgArrayFloat64Custom), buf), nil
	case *MsgSliceUint8Custom:
		return id9, makePatch9(msg2, msgSrc.(*MsgSliceUint8Custom), buf), nil
	case *MsgSliceUint8Custom2:
		return id8, makePatch8(msg2, msgSrc.(*MsgSliceUint8Custom2), buf), nil
	case *MsgSliceInt8Custom:
		return id7, makePatch7(msg2, msgSrc.(*MsgSliceInt8Custom), buf), nil
	case *MsgSliceFloat32Custom:
		return id6, makePatch6(msg2, msgSrc.(*MsgSliceFloat32Custom), buf), nil
	case *MsgSliceFloat64Custom:
		return id5, makePatch5(msg2, msgSrc.(*MsgSliceFloat64Custom), buf), nil
	case *MsgTime:
		return id4, makePatch4(msg2, msgSrc.(*MsgTime), buf), nil
	case *MsgMixedCustom:
		return id3, makePatch3(msg2, msgSrc.(*MsgMixedCustom), buf), nil
	case *MsgIgnore:
		return id1, makePatchi1(msg2, msgSrc.(*MsgIgnore), buf), nil
	case *MsgNotIgnore:
		return id2, makePatch2(msg2, msgSrc.(*MsgNotIgnore), buf), nil
	case *MsgStrings:
		return id0, makePatch0(msg2, msgSrc.(*MsgStrings), buf), nil
	default:
		return 0, 0, errors.Errorf("unknown message type %T", msgDst)
	}
}

// ApplyPatch applies patch.
func (m Marshaller) ApplyPatch(msg any, buf []byte) (retSize uint64, retErr error) {
	defer helpers.RecoverApplyPatch(&retErr)

	switch msg2 := msg.(type) {
	case *pkg2.SubMsg:
		return applyPatch41(msg2, buf), nil
	case *MsgUint64:
		return applyPatch61(msg2, buf), nil
	case *MsgUint32:
		return applyPatch60(msg2, buf), nil
	case *MsgUint16:
		return applyPatch59(msg2, buf), nil
	case *MsgUint8:
		return applyPatch58(msg2, buf), nil
	case *MsgInt64:
		return applyPatch57(msg2, buf), nil
	case *MsgInt32:
		return applyPatch56(msg2, buf), nil
	case *MsgInt16:
		return applyPatch55(msg2, buf), nil
	case *MsgInt8:
		return applyPatch54(msg2, buf), nil
	case *MsgBool3:
		return applyPatch53(msg2, buf), nil
	case *MsgBool10:
		return applyPatch52(msg2, buf), nil
	case *MsgFloat64:
		return applyPatch51(msg2, buf), nil
	case *MsgFloat32:
		return applyPatch50(msg2, buf), nil
	case *MsgString:
		return applyPatch49(msg2, buf), nil
	case *MsgArray:
		return applyPatch48(msg2, buf), nil
	case *MsgSlice:
		return applyPatch47(msg2, buf), nil
	case *MsgMap:
		return applyPatch46(msg2, buf), nil
	case *MsgMapString:
		return applyPatch45(msg2, buf), nil
	case *MsgStruct:
		return applyPatch44(msg2, buf), nil
	case *MsgStructAnonymous:
		return applyPatch42(msg2, buf), nil
	case *MsgArrayUint8:
		return applyPatch39(msg2, buf), nil
	case *MsgArrayInt8:
		return applyPatch38(msg2, buf), nil
	case *MsgArrayFloat32:
		return applyPatch37(msg2, buf), nil
	case *MsgArrayFloat64:
		return applyPatch36(msg2, buf), nil
	case *MsgSliceUint8:
		return applyPatch35(msg2, buf), nil
	case *MsgSliceInt8:
		return applyPatch34(msg2, buf), nil
	case *MsgSliceFloat32:
		return applyPatch33(msg2, buf), nil
	case *MsgSliceFloat64:
		return applyPatch32(msg2, buf), nil
	case *MsgMixed:
		return applyPatch31(msg2, buf), nil
	case *MsgUint64Custom:
		return applyPatch29(msg2, buf), nil
	case *MsgUint32Custom:
		return applyPatch28(msg2, buf), nil
	case *MsgUint16Custom:
		return applyPatch27(msg2, buf), nil
	case *MsgUint8Custom:
		return applyPatch26(msg2, buf), nil
	case *MsgInt64Custom:
		return applyPatch25(msg2, buf), nil
	case *MsgInt32Custom:
		return applyPatch24(msg2, buf), nil
	case *MsgInt16Custom:
		return applyPatch23(msg2, buf), nil
	case *MsgInt8Custom:
		return applyPatch22(msg2, buf), nil
	case *MsgBoolCustom:
		return applyPatch21(msg2, buf), nil
	case *MsgFloat64Custom:
		return applyPatch20(msg2, buf), nil
	case *MsgFloat32Custom:
		return applyPatch19(msg2, buf), nil
	case *MsgStringCustom:
		return applyPatch18(msg2, buf), nil
	case *MsgArrayCustom:
		return applyPatch17(msg2, buf), nil
	case *MsgSliceCustom:
		return applyPatch16(msg2, buf), nil
	case *MsgMapCustom:
		return applyPatch15(msg2, buf), nil
	case *MsgArrayUint8Custom:
		return applyPatch14(msg2, buf), nil
	case *MsgArrayUint8Custom2:
		return applyPatch13(msg2, buf), nil
	case *MsgArrayInt8Custom:
		return applyPatch12(msg2, buf), nil
	case *MsgArrayFloat32Custom:
		return applyPatch11(msg2, buf), nil
	case *MsgArrayFloat64Custom:
		return applyPatch10(msg2, buf), nil
	case *MsgSliceUint8Custom:
		return applyPatch9(msg2, buf), nil
	case *MsgSliceUint8Custom2:
		return applyPatch8(msg2, buf), nil
	case *MsgSliceInt8Custom:
		return applyPatch7(msg2, buf), nil
	case *MsgSliceFloat32Custom:
		return applyPatch6(msg2, buf), nil
	case *MsgSliceFloat64Custom:
		return applyPatch5(msg2, buf), nil
	case *MsgTime:
		return applyPatch4(msg2, buf), nil
	case *MsgMixedCustom:
		return applyPatch3(msg2, buf), nil
	case *MsgIgnore:
		return applyPatchi1(msg2, buf), nil
	case *MsgNotIgnore:
		return applyPatch2(msg2, buf), nil
	case *MsgStrings:
		return applyPatch0(msg2, buf), nil
	default:
		return 0, errors.Errorf("unknown message type %T", msg)
	}
}

func size0(m *MsgStrings) uint64 {
	var n uint64 = 10
	{
		// Value1

		{
			l := uint64(len(m.Value1))
			helpers.UInt64Size(l, &n)
			n += l
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

		{
			l := uint64(len(m.Value3))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	{
		// Value4

		{
			l := uint64(len(m.Value4))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	{
		// Value5

		{
			l := uint64(len(m.Value5))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	{
		// Value6

		{
			l := uint64(len(m.Value6))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	{
		// Value7

		{
			l := uint64(len(m.Value7))
			helpers.UInt64Size(l, &n)
			n += l
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
	{
		// Value9

		{
			l := uint64(len(m.Value9))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	{
		// Value10

		{
			l := uint64(len(m.Value10))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	return n
}

func marshal0(m *MsgStrings, b []byte) uint64 {
	var o uint64
	{
		// Value1

		{
			l := uint64(len(m.Value1))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value1)
			o += l
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

		{
			l := uint64(len(m.Value3))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value3)
			o += l
		}
	}
	{
		// Value4

		{
			l := uint64(len(m.Value4))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value4)
			o += l
		}
	}
	{
		// Value5

		{
			l := uint64(len(m.Value5))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value5)
			o += l
		}
	}
	{
		// Value6

		{
			l := uint64(len(m.Value6))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value6)
			o += l
		}
	}
	{
		// Value7

		{
			l := uint64(len(m.Value7))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value7)
			o += l
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
	{
		// Value9

		{
			l := uint64(len(m.Value9))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value9)
			o += l
		}
	}
	{
		// Value10

		{
			l := uint64(len(m.Value10))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value10)
			o += l
		}
	}

	return o
}

func unmarshal0(m *MsgStrings, b []byte) uint64 {
	var o uint64
	{
		// Value1

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value1 = string(b[o:o+l])
				o += l
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

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value3 = string(b[o:o+l])
				o += l
			}
		}
	}
	{
		// Value4

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value4 = string(b[o:o+l])
				o += l
			}
		}
	}
	{
		// Value5

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value5 = string(b[o:o+l])
				o += l
			}
		}
	}
	{
		// Value6

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value6 = string(b[o:o+l])
				o += l
			}
		}
	}
	{
		// Value7

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value7 = string(b[o:o+l])
				o += l
			}
		}
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
	{
		// Value9

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value9 = string(b[o:o+l])
				o += l
			}
		}
	}
	{
		// Value10

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value10 = string(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func makePatch0(m, mSrc *MsgStrings, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if reflect.DeepEqual(m.Value1, mSrc.Value1) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			{
				l := uint64(len(m.Value1))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value1)
				o += l
			}
		}
	}
	{
		// Value2

		if reflect.DeepEqual(m.Value2, mSrc.Value2) {
			b[0] &= 0xFD
		} else {
			b[0] |= 0x02
			{
				l := uint64(len(m.Value2))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value2)
				o += l
			}
		}
	}
	{
		// Value3

		if reflect.DeepEqual(m.Value3, mSrc.Value3) {
			b[0] &= 0xFB
		} else {
			b[0] |= 0x04
			{
				l := uint64(len(m.Value3))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value3)
				o += l
			}
		}
	}
	{
		// Value4

		if reflect.DeepEqual(m.Value4, mSrc.Value4) {
			b[0] &= 0xF7
		} else {
			b[0] |= 0x08
			{
				l := uint64(len(m.Value4))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value4)
				o += l
			}
		}
	}
	{
		// Value5

		if reflect.DeepEqual(m.Value5, mSrc.Value5) {
			b[0] &= 0xEF
		} else {
			b[0] |= 0x10
			{
				l := uint64(len(m.Value5))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value5)
				o += l
			}
		}
	}
	{
		// Value6

		if reflect.DeepEqual(m.Value6, mSrc.Value6) {
			b[0] &= 0xDF
		} else {
			b[0] |= 0x20
			{
				l := uint64(len(m.Value6))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value6)
				o += l
			}
		}
	}
	{
		// Value7

		if reflect.DeepEqual(m.Value7, mSrc.Value7) {
			b[0] &= 0xBF
		} else {
			b[0] |= 0x40
			{
				l := uint64(len(m.Value7))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value7)
				o += l
			}
		}
	}
	{
		// Value8

		if reflect.DeepEqual(m.Value8, mSrc.Value8) {
			b[0] &= 0x7F
		} else {
			b[0] |= 0x80
			{
				l := uint64(len(m.Value8))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value8)
				o += l
			}
		}
	}
	{
		// Value9

		if reflect.DeepEqual(m.Value9, mSrc.Value9) {
			b[1] &= 0xFE
		} else {
			b[1] |= 0x01
			{
				l := uint64(len(m.Value9))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value9)
				o += l
			}
		}
	}
	{
		// Value10

		if reflect.DeepEqual(m.Value10, mSrc.Value10) {
			b[1] &= 0xFD
		} else {
			b[1] |= 0x02
			{
				l := uint64(len(m.Value10))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value10)
				o += l
			}
		}
	}

	return o
}

func applyPatch0(m *MsgStrings, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if b[0]&0x01 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value1 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value2

		if b[0]&0x02 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value2 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value3

		if b[0]&0x04 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value3 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value4

		if b[0]&0x08 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value4 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value5

		if b[0]&0x10 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value5 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value6

		if b[0]&0x20 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value6 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value7

		if b[0]&0x40 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value7 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value8

		if b[0]&0x80 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value8 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value9

		if b[1]&0x01 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value9 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value10

		if b[1]&0x02 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value10 = string(b[o:o+l])
					o += l
				}
			}
		}
	}

	return o
}

func size2(m *MsgNotIgnore) uint64 {
	var n uint64
	{
		// SubMsg

		n += size1(&m.SubMsg)
	}
	return n
}

func marshal2(m *MsgNotIgnore, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += marshal1(&m.SubMsg, b[o:])
	}

	return o
}

func unmarshal2(m *MsgNotIgnore, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += unmarshal1(&m.SubMsg, b[o:])
	}

	return o
}

func makePatch2(m, mSrc *MsgNotIgnore, b []byte) uint64 {
	var o uint64 = 1
	{
		// SubMsg

		if reflect.DeepEqual(m.SubMsg, mSrc.SubMsg) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			o += marshal1(&m.SubMsg, b[o:])
		}
	}

	return o
}

func applyPatch2(m *MsgNotIgnore, b []byte) uint64 {
	var o uint64 = 1
	{
		// SubMsg

		if b[0]&0x01 != 0 {
			o += unmarshal1(&m.SubMsg, b[o:])
		}
	}

	return o
}

func size1(m *MsgIgnore) uint64 {
	var n uint64 = 3
	{
		// Value3

		{
			l := uint64(len(m.Value3))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	{
		// Value4Ignored

		{
			l := uint64(len(m.Value4Ignored))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	return n
}

func marshal1(m *MsgIgnore, b []byte) uint64 {
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
		// Value2Ignored

		if m.Value2Ignored {
			b[0] |= 0x02
		} else {
			b[0] &= 0xFD
		}
	}
	{
		// Value3

		{
			l := uint64(len(m.Value3))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value3)
			o += l
		}
	}
	{
		// Value4Ignored

		{
			l := uint64(len(m.Value4Ignored))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value4Ignored)
			o += l
		}
	}

	return o
}

func unmarshal1(m *MsgIgnore, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		m.Value1 = b[0]&0x01 != 0
	}
	{
		// Value2Ignored

		m.Value2Ignored = b[0]&0x02 != 0
	}
	{
		// Value3

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value3 = string(b[o:o+l])
				o += l
			}
		}
	}
	{
		// Value4Ignored

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value4Ignored = string(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func makePatch1(m, mSrc *MsgIgnore, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if m.Value1 == mSrc.Value1 {
			b[1] &= 0xFE
		} else {
			b[1] |= 0x01
		}
	}
	{
		// Value2Ignored

		if m.Value2Ignored == mSrc.Value2Ignored {
			b[1] &= 0xFD
		} else {
			b[1] |= 0x02
		}
	}
	{
		// Value3

		if reflect.DeepEqual(m.Value3, mSrc.Value3) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			{
				l := uint64(len(m.Value3))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value3)
				o += l
			}
		}
	}
	{
		// Value4Ignored

		if reflect.DeepEqual(m.Value4Ignored, mSrc.Value4Ignored) {
			b[0] &= 0xFD
		} else {
			b[0] |= 0x02
			{
				l := uint64(len(m.Value4Ignored))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value4Ignored)
				o += l
			}
		}
	}

	return o
}

func applyPatch1(m *MsgIgnore, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if b[1]&0x01 != 0 {
			m.Value1 = !m.Value1
		}
	}
	{
		// Value2Ignored

		if b[1]&0x02 != 0 {
			m.Value2Ignored = !m.Value2Ignored
		}
	}
	{
		// Value3

		if b[0]&0x01 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value3 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value4Ignored

		if b[0]&0x02 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value4Ignored = string(b[o:o+l])
					o += l
				}
			}
		}
	}

	return o
}

func sizei1(m *MsgIgnore) uint64 {
	var n uint64 = 2
	{
		// Value3

		{
			l := uint64(len(m.Value3))
			helpers.UInt64Size(l, &n)
			n += l
		}
	}
	return n
}

func marshali1(m *MsgIgnore, b []byte) uint64 {
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
		// Value3

		{
			l := uint64(len(m.Value3))
			helpers.UInt64Marshal(l, b, &o)
			copy(b[o:o+l], m.Value3)
			o += l
		}
	}

	return o
}

func unmarshali1(m *MsgIgnore, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		m.Value1 = b[0]&0x01 != 0
	}
	{
		// Value3

		{
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value3 = string(b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func makePatchi1(m, mSrc *MsgIgnore, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if m.Value1 == mSrc.Value1 {
			b[1] &= 0xFE
		} else {
			b[1] |= 0x01
		}
	}
	{
		// Value3

		if reflect.DeepEqual(m.Value3, mSrc.Value3) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			{
				l := uint64(len(m.Value3))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value3)
				o += l
			}
		}
	}

	return o
}

func applyPatchi1(m *MsgIgnore, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if b[1]&0x01 != 0 {
			m.Value1 = !m.Value1
		}
	}
	{
		// Value3

		if b[0]&0x01 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value3 = string(b[o:o+l])
					o += l
				}
			}
		}
	}

	return o
}

func size3(m *MsgMixedCustom) uint64 {
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

func marshal3(m *MsgMixedCustom, b []byte) uint64 {
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

func unmarshal3(m *MsgMixedCustom, b []byte) uint64 {
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

func makePatch3(m, mSrc *MsgMixedCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
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
	}

	return o
}

func applyPatch3(m *MsgMixedCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
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
	}

	return o
}

func size4(m *MsgTime) uint64 {
	var n uint64 = 2
	{
		// Value

		helpers.Int64Size(m.Value.Unix() - -62135596800, &n)
		helpers.UInt32Size(uint32(m.Value.Nanosecond()), &n)
	}
	return n
}

func marshal4(m *MsgTime, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Marshal(m.Value.Unix() - -62135596800, b, &o)
		helpers.UInt32Marshal(uint32(m.Value.Nanosecond()), b, &o)
	}

	return o
}

func unmarshal4(m *MsgTime, b []byte) uint64 {
	var o uint64
	{
		// Value

		var seconds int64
		var nanoseconds uint32
		helpers.Int64Unmarshal(&seconds, b, &o)
		helpers.UInt32Unmarshal(&nanoseconds, b, &o)
		m.Value = time.Unix(seconds + -62135596800, int64(nanoseconds))
	}

	return o
}

func makePatch4(m, mSrc *MsgTime, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.Int64Marshal(m.Value.Unix() - -62135596800, b, &o)
			helpers.UInt32Marshal(uint32(m.Value.Nanosecond()), b, &o)
		}
	}

	return o
}

func applyPatch4(m *MsgTime, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var seconds int64
			var nanoseconds uint32
			helpers.Int64Unmarshal(&seconds, b, &o)
			helpers.UInt32Unmarshal(&nanoseconds, b, &o)
			m.Value = time.Unix(seconds + -62135596800, int64(nanoseconds))
		}
	}

	return o
}

func size5(m *MsgSliceFloat64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l * 8
	}
	return n
}

func marshal5(m *MsgSliceFloat64Custom, b []byte) uint64 {
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

func unmarshal5(m *MsgSliceFloat64Custom, b []byte) uint64 {
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

func makePatch5(m, mSrc *MsgSliceFloat64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l*8], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8))
				o += l * 8
			}
		}
	}

	return o
}

func applyPatch5(m *MsgSliceFloat64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]custom.Float64, l)
				copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
				o += l * 8
			}
		}
	}

	return o
}

func size6(m *MsgSliceFloat32Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l * 4
	}
	return n
}

func marshal6(m *MsgSliceFloat32Custom, b []byte) uint64 {
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

func unmarshal6(m *MsgSliceFloat32Custom, b []byte) uint64 {
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

func makePatch6(m, mSrc *MsgSliceFloat32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l*4], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4))
				o += l * 4
			}
		}
	}

	return o
}

func applyPatch6(m *MsgSliceFloat32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]custom.Float32, l)
				copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
				o += l * 4
			}
		}
	}

	return o
}

func size7(m *MsgSliceInt8Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal7(m *MsgSliceInt8Custom, b []byte) uint64 {
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

func unmarshal7(m *MsgSliceInt8Custom, b []byte) uint64 {
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

func makePatch7(m, mSrc *MsgSliceInt8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l))
				o += l
			}
		}
	}

	return o
}

func applyPatch7(m *MsgSliceInt8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]custom.Int8, l)
				copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size8(m *MsgSliceUint8Custom2) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal8(m *MsgSliceUint8Custom2, b []byte) uint64 {
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

func unmarshal8(m *MsgSliceUint8Custom2, b []byte) uint64 {
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

func makePatch8(m, mSrc *MsgSliceUint8Custom2, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l], unsafe.Slice((*byte)(&m.Value[0]), l))
				o += l
			}
		}
	}

	return o
}

func applyPatch8(m *MsgSliceUint8Custom2, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]custom.Uint8, l)
				copy(unsafe.Slice((*byte)(&m.Value[0]), l), b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size9(m *MsgSliceUint8Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal9(m *MsgSliceUint8Custom, b []byte) uint64 {
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

func unmarshal9(m *MsgSliceUint8Custom, b []byte) uint64 {
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

func makePatch9(m, mSrc *MsgSliceUint8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l], unsafe.Slice(&m.Value[0], l))
				o += l
			}
		}
	}

	return o
}

func applyPatch9(m *MsgSliceUint8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]uint8, l)
				copy(m.Value, b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size10(m *MsgArrayFloat64Custom) uint64 {
	var n uint64 = 40
	return n
}

func marshal10(m *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

func unmarshal10(m *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
		o += 40
	}

	return o
}

func makePatch10(m, mSrc *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
			o += 40
		}
	}

	return o
}

func applyPatch10(m *MsgArrayFloat64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
			o += 40
		}
	}

	return o
}

func size11(m *MsgArrayFloat32Custom) uint64 {
	var n uint64 = 20
	return n
}

func marshal11(m *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

func unmarshal11(m *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
		o += 20
	}

	return o
}

func makePatch11(m, mSrc *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
			o += 20
		}
	}

	return o
}

func applyPatch11(m *MsgArrayFloat32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
			o += 20
		}
	}

	return o
}

func size12(m *MsgArrayInt8Custom) uint64 {
	var n uint64 = 5
	return n
}

func marshal12(m *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

func unmarshal12(m *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
		o += 5
	}

	return o
}

func makePatch12(m, mSrc *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
			o += 5
		}
	}

	return o
}

func applyPatch12(m *MsgArrayInt8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
			o += 5
		}
	}

	return o
}

func size13(m *MsgArrayUint8Custom2) uint64 {
	var n uint64 = 5
	return n
}

func marshal13(m *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(&m.Value[0]), 5))
		o += 5
	}

	return o
}

func unmarshal13(m *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(&m.Value[0]), 5), b[o:o+5])
		o += 5
	}

	return o
}

func makePatch13(m, mSrc *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+5], unsafe.Slice((*byte)(&m.Value[0]), 5))
			o += 5
		}
	}

	return o
}

func applyPatch13(m *MsgArrayUint8Custom2, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice((*byte)(&m.Value[0]), 5), b[o:o+5])
			o += 5
		}
	}

	return o
}

func size14(m *MsgArrayUint8Custom) uint64 {
	var n uint64 = 5
	return n
}

func marshal14(m *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

func unmarshal14(m *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
		o += 5
	}

	return o
}

func makePatch14(m, mSrc *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
			o += 5
		}
	}

	return o
}

func applyPatch14(m *MsgArrayUint8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
			o += 5
		}
	}

	return o
}

func size15(m *MsgMapCustom) uint64 {
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

func marshal15(m *MsgMapCustom, b []byte) uint64 {
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

func unmarshal15(m *MsgMapCustom, b []byte) uint64 {
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

func makePatch15(m, mSrc *MsgMapCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
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
	}

	return o
}

func applyPatch15(m *MsgMapCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
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
	}

	return o
}

func size16(m *MsgSliceCustom) uint64 {
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

func marshal16(m *MsgSliceCustom, b []byte) uint64 {
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

func unmarshal16(m *MsgSliceCustom, b []byte) uint64 {
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

func makePatch16(m, mSrc *MsgSliceCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt64Marshal(uint64(len(m.Value)), b, &o)
			for _, sv1 := range m.Value {
				helpers.UInt64Marshal(sv1, b, &o)
			}
		}
	}

	return o
}

func applyPatch16(m *MsgSliceCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]custom.Uint64, l)
				for i1 := range l {
					helpers.UInt64Unmarshal(&m.Value[i1], b, &o)
				}
			}
		}
	}

	return o
}

func size17(m *MsgArrayCustom) uint64 {
	var n uint64 = 2
	{
		// Value

		for _, av1 := range m.Value {
			helpers.UInt64Size(av1, &n)
		}
	}
	return n
}

func marshal17(m *MsgArrayCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		for _, av1 := range m.Value {
			helpers.UInt64Marshal(av1, b, &o)
		}
	}

	return o
}

func unmarshal17(m *MsgArrayCustom, b []byte) uint64 {
	var o uint64
	{
		// Value

		for i1 := range 2 {
			helpers.UInt64Unmarshal(&m.Value[i1], b, &o)
		}
	}

	return o
}

func makePatch17(m, mSrc *MsgArrayCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			for _, av1 := range m.Value {
				helpers.UInt64Marshal(av1, b, &o)
			}
		}
	}

	return o
}

func applyPatch17(m *MsgArrayCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			for i1 := range 2 {
				helpers.UInt64Unmarshal(&m.Value[i1], b, &o)
			}
		}
	}

	return o
}

func size18(m *MsgStringCustom) uint64 {
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

func marshal18(m *MsgStringCustom, b []byte) uint64 {
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

func unmarshal18(m *MsgStringCustom, b []byte) uint64 {
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

func makePatch18(m, mSrc *MsgStringCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			{
				l := uint64(len(m.Value))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value)
				o += l
			}
		}
	}

	return o
}

func applyPatch18(m *MsgStringCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value = custom.String(b[o:o+l])
					o += l
				}
			}
		}
	}

	return o
}

func size19(m *MsgFloat32Custom) uint64 {
	var n uint64 = 4
	return n
}

func marshal19(m *MsgFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

func unmarshal19(m *MsgFloat32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*custom.Float32)(unsafe.Pointer(&b[o]))
		o += 4
	}

	return o
}

func makePatch19(m, mSrc *MsgFloat32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			*(*custom.Float32)(unsafe.Pointer(&b[o])) = m.Value
			o += 4
		}
	}

	return o
}

func applyPatch19(m *MsgFloat32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = *(*custom.Float32)(unsafe.Pointer(&b[o]))
			o += 4
		}
	}

	return o
}

func size20(m *MsgFloat64Custom) uint64 {
	var n uint64 = 8
	return n
}

func marshal20(m *MsgFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

func unmarshal20(m *MsgFloat64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*custom.Float64)(unsafe.Pointer(&b[o]))
		o += 8
	}

	return o
}

func makePatch20(m, mSrc *MsgFloat64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			*(*custom.Float64)(unsafe.Pointer(&b[o])) = m.Value
			o += 8
		}
	}

	return o
}

func applyPatch20(m *MsgFloat64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = *(*custom.Float64)(unsafe.Pointer(&b[o]))
			o += 8
		}
	}

	return o
}

func size21(m *MsgBoolCustom) uint64 {
	var n uint64 = 1
	return n
}

func marshal21(m *MsgBoolCustom, b []byte) uint64 {
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

func unmarshal21(m *MsgBoolCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		m.Value = b[0]&0x01 != 0
	}

	return o
}

func makePatch21(m, mSrc *MsgBoolCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if m.Value == mSrc.Value {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
		}
	}

	return o
}

func applyPatch21(m *MsgBoolCustom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = !m.Value
		}
	}

	return o
}

func size22(m *MsgInt8Custom) uint64 {
	var n uint64 = 1
	return n
}

func marshal22(m *MsgInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal22(m *MsgInt8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = custom.Int8(b[o])
		o++
	}

	return o
}

func makePatch22(m, mSrc *MsgInt8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			b[o] = byte(m.Value)
			o++
		}
	}

	return o
}

func applyPatch22(m *MsgInt8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = custom.Int8(b[o])
			o++
		}
	}

	return o
}

func size23(m *MsgInt16Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int16Size(m.Value, &n)
	}
	return n
}

func marshal23(m *MsgInt16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal23(m *MsgInt16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch23(m, mSrc *MsgInt16Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.Int16Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch23(m *MsgInt16Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.Int16Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size24(m *MsgInt32Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int32Size(m.Value, &n)
	}
	return n
}

func marshal24(m *MsgInt32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal24(m *MsgInt32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch24(m, mSrc *MsgInt32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.Int32Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch24(m *MsgInt32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.Int32Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size25(m *MsgInt64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int64Size(m.Value, &n)
	}
	return n
}

func marshal25(m *MsgInt64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal25(m *MsgInt64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch25(m, mSrc *MsgInt64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.Int64Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch25(m *MsgInt64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.Int64Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size26(m *MsgUint8Custom) uint64 {
	var n uint64 = 1
	return n
}

func marshal26(m *MsgUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal26(m *MsgUint8Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = custom.Uint8(b[o])
		o++
	}

	return o
}

func makePatch26(m, mSrc *MsgUint8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			b[o] = byte(m.Value)
			o++
		}
	}

	return o
}

func applyPatch26(m *MsgUint8Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = custom.Uint8(b[o])
			o++
		}
	}

	return o
}

func size27(m *MsgUint16Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt16Size(m.Value, &n)
	}
	return n
}

func marshal27(m *MsgUint16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal27(m *MsgUint16Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch27(m, mSrc *MsgUint16Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt16Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch27(m *MsgUint16Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.UInt16Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size28(m *MsgUint32Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt32Size(m.Value, &n)
	}
	return n
}

func marshal28(m *MsgUint32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal28(m *MsgUint32Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch28(m, mSrc *MsgUint32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt32Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch28(m *MsgUint32Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.UInt32Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size29(m *MsgUint64Custom) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt64Size(m.Value, &n)
	}
	return n
}

func marshal29(m *MsgUint64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal29(m *MsgUint64Custom, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch29(m, mSrc *MsgUint64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt64Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch29(m *MsgUint64Custom, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.UInt64Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size31(m *MsgMixed) uint64 {
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
			n += size30(&mv2)
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

func marshal31(m *MsgMixed, b []byte) uint64 {
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
			o += marshal30(&mv2, b[o:])
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

func unmarshal31(m *MsgMixed, b []byte) uint64 {
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
				o += unmarshal30(&mv2, b[o:])
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

func makePatch31(m, mSrc *MsgMixed, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if reflect.DeepEqual(m.Value1, mSrc.Value1) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt64Marshal(uint64(len(m.Value1)), b, &o)
			for mk1, mv2 := range m.Value1 {
				{
					l := uint64(len(mk1))
					helpers.UInt64Marshal(l, b, &o)
					copy(b[o:o+l], mk1)
					o += l
				}
				o += marshal30(&mv2, b[o:])
			}
		}
	}
	{
		// Value2

		if reflect.DeepEqual(m.Value2, mSrc.Value2) {
			b[0] &= 0xFD
		} else {
			b[0] |= 0x02
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
	}
	{
		// Value3

		if reflect.DeepEqual(m.Value3, mSrc.Value3) {
			b[0] &= 0xFB
		} else {
			b[0] |= 0x04
			helpers.UInt64Marshal(uint64(len(m.Value3)), b, &o)
			for _, sv2 := range m.Value3 {
				for _, av1 := range sv2 {
					helpers.UInt16Marshal(av1, b, &o)
				}
			}
		}
	}
	{
		// Value4

		if reflect.DeepEqual(m.Value4, mSrc.Value4) {
			b[0] &= 0xF7
		} else {
			b[0] |= 0x08
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
	}
	{
		// Value5

		if reflect.DeepEqual(m.Value5, mSrc.Value5) {
			b[0] &= 0xEF
		} else {
			b[0] |= 0x10
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
	}
	{
		// Value6

		if m.Value6 == mSrc.Value6 {
			b[1] &= 0xFE
		} else {
			b[1] |= 0x01
		}
	}
	{
		// Value7

		if m.Value7 == mSrc.Value7 {
			b[1] &= 0xFD
		} else {
			b[1] |= 0x02
		}
	}
	{
		// Value8

		if reflect.DeepEqual(m.Value8, mSrc.Value8) {
			b[0] &= 0xDF
		} else {
			b[0] |= 0x20
			{
				l := uint64(len(m.Value8))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value8)
				o += l
			}
		}
	}

	return o
}

func applyPatch31(m *MsgMixed, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if b[0]&0x01 != 0 {
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
					o += unmarshal30(&mv2, b[o:])
					m.Value1[mk1] = mv2
				}
			}
		}
	}
	{
		// Value2

		if b[0]&0x02 != 0 {
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
	}
	{
		// Value3

		if b[0]&0x04 != 0 {
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
	}
	{
		// Value4

		if b[0]&0x08 != 0 {
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
	}
	{
		// Value5

		if b[0]&0x10 != 0 {
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
	}
	{
		// Value6

		if b[1]&0x01 != 0 {
			m.Value6 = !m.Value6
		}
	}
	{
		// Value7

		if b[1]&0x02 != 0 {
			m.Value7 = !m.Value7
		}
	}
	{
		// Value8

		if b[0]&0x20 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value8 = string(b[o:o+l])
					o += l
				}
			}
		}
	}

	return o
}

func size30(m *spkg.SubMsg) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int32Size(m.Value, &n)
	}
	return n
}

func marshal30(m *spkg.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal30(m *spkg.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size32(m *MsgSliceFloat64) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l * 8
	}
	return n
}

func marshal32(m *MsgSliceFloat64, b []byte) uint64 {
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

func unmarshal32(m *MsgSliceFloat64, b []byte) uint64 {
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

func makePatch32(m, mSrc *MsgSliceFloat64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l*8], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8))
				o += l * 8
			}
		}
	}

	return o
}

func applyPatch32(m *MsgSliceFloat64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]float64, l)
				copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
				o += l * 8
			}
		}
	}

	return o
}

func size33(m *MsgSliceFloat32) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l * 4
	}
	return n
}

func marshal33(m *MsgSliceFloat32, b []byte) uint64 {
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

func unmarshal33(m *MsgSliceFloat32, b []byte) uint64 {
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

func makePatch33(m, mSrc *MsgSliceFloat32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l*4], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4))
				o += l * 4
			}
		}
	}

	return o
}

func applyPatch33(m *MsgSliceFloat32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]float32, l)
				copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
				o += l * 4
			}
		}
	}

	return o
}

func size34(m *MsgSliceInt8) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal34(m *MsgSliceInt8, b []byte) uint64 {
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

func unmarshal34(m *MsgSliceInt8, b []byte) uint64 {
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

func makePatch34(m, mSrc *MsgSliceInt8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l))
				o += l
			}
		}
	}

	return o
}

func applyPatch34(m *MsgSliceInt8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]int8, l)
				copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size35(m *MsgSliceUint8) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal35(m *MsgSliceUint8, b []byte) uint64 {
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

func unmarshal35(m *MsgSliceUint8, b []byte) uint64 {
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

func makePatch35(m, mSrc *MsgSliceUint8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			l := uint64(len(m.Value))
			helpers.UInt64Marshal(l, b, &o)
			if l > 0 {
				copy(b[o:o+l], unsafe.Slice(&m.Value[0], l))
				o += l
			}
		}
	}

	return o
}

func applyPatch35(m *MsgSliceUint8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			var l uint64
			helpers.UInt64Unmarshal(&l, b, &o)
			if l > 0 {
				m.Value = make([]uint8, l)
				copy(m.Value, b[o:o+l])
				o += l
			}
		}
	}

	return o
}

func size36(m *MsgArrayFloat64) uint64 {
	var n uint64 = 40
	return n
}

func marshal36(m *MsgArrayFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

func unmarshal36(m *MsgArrayFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
		o += 40
	}

	return o
}

func makePatch36(m, mSrc *MsgArrayFloat64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
			o += 40
		}
	}

	return o
}

func applyPatch36(m *MsgArrayFloat64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40), b[o:o+40])
			o += 40
		}
	}

	return o
}

func size37(m *MsgArrayFloat32) uint64 {
	var n uint64 = 20
	return n
}

func marshal37(m *MsgArrayFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

func unmarshal37(m *MsgArrayFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
		o += 20
	}

	return o
}

func makePatch37(m, mSrc *MsgArrayFloat32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
			o += 20
		}
	}

	return o
}

func applyPatch37(m *MsgArrayFloat32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20), b[o:o+20])
			o += 20
		}
	}

	return o
}

func size38(m *MsgArrayInt8) uint64 {
	var n uint64 = 5
	return n
}

func marshal38(m *MsgArrayInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

func unmarshal38(m *MsgArrayInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
		o += 5
	}

	return o
}

func makePatch38(m, mSrc *MsgArrayInt8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
			o += 5
		}
	}

	return o
}

func applyPatch38(m *MsgArrayInt8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5), b[o:o+5])
			o += 5
		}
	}

	return o
}

func size39(m *MsgArrayUint8) uint64 {
	var n uint64 = 5
	return n
}

func marshal39(m *MsgArrayUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

func unmarshal39(m *MsgArrayUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
		o += 5
	}

	return o
}

func makePatch39(m, mSrc *MsgArrayUint8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
			o += 5
		}
	}

	return o
}

func applyPatch39(m *MsgArrayUint8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			copy(unsafe.Slice(&m.Value[0], 5), b[o:o+5])
			o += 5
		}
	}

	return o
}

func size42(m *MsgStructAnonymous) uint64 {
	var n uint64
	{
		// SubMsg

		n += size40(&m.SubMsg)
	}
	{
		// Value2

		n += size41(&m.Value2)
	}
	return n
}

func marshal42(m *MsgStructAnonymous, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += marshal40(&m.SubMsg, b[o:])
	}
	{
		// Value2

		o += marshal41(&m.Value2, b[o:])
	}

	return o
}

func unmarshal42(m *MsgStructAnonymous, b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += unmarshal40(&m.SubMsg, b[o:])
	}
	{
		// Value2

		o += unmarshal41(&m.Value2, b[o:])
	}

	return o
}

func makePatch42(m, mSrc *MsgStructAnonymous, b []byte) uint64 {
	var o uint64 = 1
	{
		// SubMsg

		if reflect.DeepEqual(m.SubMsg, mSrc.SubMsg) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			o += marshal40(&m.SubMsg, b[o:])
		}
	}
	{
		// Value2

		if reflect.DeepEqual(m.Value2, mSrc.Value2) {
			b[0] &= 0xFD
		} else {
			b[0] |= 0x02
			o += marshal41(&m.Value2, b[o:])
		}
	}

	return o
}

func applyPatch42(m *MsgStructAnonymous, b []byte) uint64 {
	var o uint64 = 1
	{
		// SubMsg

		if b[0]&0x01 != 0 {
			o += unmarshal40(&m.SubMsg, b[o:])
		}
	}
	{
		// Value2

		if b[0]&0x02 != 0 {
			o += unmarshal41(&m.Value2, b[o:])
		}
	}

	return o
}

func size40(m *SubMsg) uint64 {
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
			n += size30(&sv1)
		}
	}
	{
		// Value3

		l := uint64(len(m.Value3))
		helpers.UInt64Size(l, &n)
		for _, sv1 := range m.Value3 {
			n += size43(&sv1)
		}
	}
	return n
}

func marshal40(m *SubMsg, b []byte) uint64 {
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
			o += marshal30(&sv1, b[o:])
		}
	}
	{
		// Value3

		helpers.UInt64Marshal(uint64(len(m.Value3)), b, &o)
		for _, sv1 := range m.Value3 {
			o += marshal43(&sv1, b[o:])
		}
	}

	return o
}

func unmarshal40(m *SubMsg, b []byte) uint64 {
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
				o += unmarshal30(&m.Value2[i1], b[o:])
			}
		}
	}
	{
		// Value3

		var l uint64
		helpers.UInt64Unmarshal(&l, b, &o)
		if l > 0 {
			m.Value3 = make([]spkg8.SubMsg, l)
			for i1 := range l {
				o += unmarshal43(&m.Value3[i1], b[o:])
			}
		}
	}

	return o
}

func size43(m *spkg8.SubMsg) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int16Size(m.Value, &n)
	}
	return n
}

func marshal43(m *spkg8.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal43(m *spkg8.SubMsg, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func size44(m *MsgStruct) uint64 {
	var n uint64
	{
		// Value1

		n += size40(&m.Value1)
	}
	{
		// Value2

		n += size41(&m.Value2)
	}
	return n
}

func marshal44(m *MsgStruct, b []byte) uint64 {
	var o uint64
	{
		// Value1

		o += marshal40(&m.Value1, b[o:])
	}
	{
		// Value2

		o += marshal41(&m.Value2, b[o:])
	}

	return o
}

func unmarshal44(m *MsgStruct, b []byte) uint64 {
	var o uint64
	{
		// Value1

		o += unmarshal40(&m.Value1, b[o:])
	}
	{
		// Value2

		o += unmarshal41(&m.Value2, b[o:])
	}

	return o
}

func makePatch44(m, mSrc *MsgStruct, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		if reflect.DeepEqual(m.Value1, mSrc.Value1) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			o += marshal40(&m.Value1, b[o:])
		}
	}
	{
		// Value2

		if reflect.DeepEqual(m.Value2, mSrc.Value2) {
			b[0] &= 0xFD
		} else {
			b[0] |= 0x02
			o += marshal41(&m.Value2, b[o:])
		}
	}

	return o
}

func applyPatch44(m *MsgStruct, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		if b[0]&0x01 != 0 {
			o += unmarshal40(&m.Value1, b[o:])
		}
	}
	{
		// Value2

		if b[0]&0x02 != 0 {
			o += unmarshal41(&m.Value2, b[o:])
		}
	}

	return o
}

func size45(m *MsgMapString) uint64 {
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

func marshal45(m *MsgMapString, b []byte) uint64 {
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

func unmarshal45(m *MsgMapString, b []byte) uint64 {
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

func makePatch45(m, mSrc *MsgMapString, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
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
	}

	return o
}

func applyPatch45(m *MsgMapString, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
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
	}

	return o
}

func size46(m *MsgMap) uint64 {
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

func marshal46(m *MsgMap, b []byte) uint64 {
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

func unmarshal46(m *MsgMap, b []byte) uint64 {
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

func makePatch46(m, mSrc *MsgMap, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt64Marshal(uint64(len(m.Value)), b, &o)
			for mk1, mv2 := range m.Value {
				helpers.UInt64Marshal(mk1, b, &o)
				helpers.UInt64Marshal(mv2, b, &o)
			}
		}
	}

	return o
}

func applyPatch46(m *MsgMap, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
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
	}

	return o
}

func size47(m *MsgSlice) uint64 {
	var n uint64 = 1
	{
		// Value

		l := uint64(len(m.Value))
		helpers.UInt64Size(l, &n)
		n += l
	}
	return n
}

func marshal47(m *MsgSlice, b []byte) uint64 {
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

func unmarshal47(m *MsgSlice, b []byte) uint64 {
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

func makePatch47(m, mSrc *MsgSlice, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
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
	}

	return o
}

func applyPatch47(m *MsgSlice, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
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
	}

	return o
}

func size48(m *MsgArray) uint64 {
	var n uint64 = 3
	return n
}

func marshal48(m *MsgArray, b []byte) uint64 {
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

func unmarshal48(m *MsgArray, b []byte) uint64 {
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

func makePatch48(m, mSrc *MsgArray, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			for _, av1 := range m.Value {
				if av1 {
					b[o] = 0x01
				} else {
					b[o] = 0x00
				}
				o++
			}
		}
	}

	return o
}

func applyPatch48(m *MsgArray, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			for i1 := range 3 {
				m.Value[i1] = b[o] != 0x00
				o++
			}
		}
	}

	return o
}

func size49(m *MsgString) uint64 {
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

func marshal49(m *MsgString, b []byte) uint64 {
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

func unmarshal49(m *MsgString, b []byte) uint64 {
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

func makePatch49(m, mSrc *MsgString, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			{
				l := uint64(len(m.Value))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value)
				o += l
			}
		}
	}

	return o
}

func applyPatch49(m *MsgString, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value = string(b[o:o+l])
					o += l
				}
			}
		}
	}

	return o
}

func size50(m *MsgFloat32) uint64 {
	var n uint64 = 4
	return n
}

func marshal50(m *MsgFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

func unmarshal50(m *MsgFloat32, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*float32)(unsafe.Pointer(&b[o]))
		o += 4
	}

	return o
}

func makePatch50(m, mSrc *MsgFloat32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			*(*float32)(unsafe.Pointer(&b[o])) = m.Value
			o += 4
		}
	}

	return o
}

func applyPatch50(m *MsgFloat32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = *(*float32)(unsafe.Pointer(&b[o]))
			o += 4
		}
	}

	return o
}

func size51(m *MsgFloat64) uint64 {
	var n uint64 = 8
	return n
}

func marshal51(m *MsgFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

func unmarshal51(m *MsgFloat64, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = *(*float64)(unsafe.Pointer(&b[o]))
		o += 8
	}

	return o
}

func makePatch51(m, mSrc *MsgFloat64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			*(*float64)(unsafe.Pointer(&b[o])) = m.Value
			o += 8
		}
	}

	return o
}

func applyPatch51(m *MsgFloat64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = *(*float64)(unsafe.Pointer(&b[o]))
			o += 8
		}
	}

	return o
}

func size52(m *MsgBool10) uint64 {
	var n uint64 = 2
	return n
}

func marshal52(m *MsgBool10, b []byte) uint64 {
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

func unmarshal52(m *MsgBool10, b []byte) uint64 {
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

func makePatch52(m, mSrc *MsgBool10, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if m.Value1 == mSrc.Value1 {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
		}
	}
	{
		// Value2

		if m.Value2 == mSrc.Value2 {
			b[0] &= 0xFD
		} else {
			b[0] |= 0x02
		}
	}
	{
		// Value3

		if m.Value3 == mSrc.Value3 {
			b[0] &= 0xFB
		} else {
			b[0] |= 0x04
		}
	}
	{
		// Value4

		if m.Value4 == mSrc.Value4 {
			b[0] &= 0xF7
		} else {
			b[0] |= 0x08
		}
	}
	{
		// Value5

		if m.Value5 == mSrc.Value5 {
			b[0] &= 0xEF
		} else {
			b[0] |= 0x10
		}
	}
	{
		// Value6

		if m.Value6 == mSrc.Value6 {
			b[0] &= 0xDF
		} else {
			b[0] |= 0x20
		}
	}
	{
		// Value7

		if m.Value7 == mSrc.Value7 {
			b[0] &= 0xBF
		} else {
			b[0] |= 0x40
		}
	}
	{
		// Value8

		if m.Value8 == mSrc.Value8 {
			b[0] &= 0x7F
		} else {
			b[0] |= 0x80
		}
	}
	{
		// Value9

		if m.Value9 == mSrc.Value9 {
			b[1] &= 0xFE
		} else {
			b[1] |= 0x01
		}
	}
	{
		// Value10

		if m.Value10 == mSrc.Value10 {
			b[1] &= 0xFD
		} else {
			b[1] |= 0x02
		}
	}

	return o
}

func applyPatch52(m *MsgBool10, b []byte) uint64 {
	var o uint64 = 2
	{
		// Value1

		if b[0]&0x01 != 0 {
			m.Value1 = !m.Value1
		}
	}
	{
		// Value2

		if b[0]&0x02 != 0 {
			m.Value2 = !m.Value2
		}
	}
	{
		// Value3

		if b[0]&0x04 != 0 {
			m.Value3 = !m.Value3
		}
	}
	{
		// Value4

		if b[0]&0x08 != 0 {
			m.Value4 = !m.Value4
		}
	}
	{
		// Value5

		if b[0]&0x10 != 0 {
			m.Value5 = !m.Value5
		}
	}
	{
		// Value6

		if b[0]&0x20 != 0 {
			m.Value6 = !m.Value6
		}
	}
	{
		// Value7

		if b[0]&0x40 != 0 {
			m.Value7 = !m.Value7
		}
	}
	{
		// Value8

		if b[0]&0x80 != 0 {
			m.Value8 = !m.Value8
		}
	}
	{
		// Value9

		if b[1]&0x01 != 0 {
			m.Value9 = !m.Value9
		}
	}
	{
		// Value10

		if b[1]&0x02 != 0 {
			m.Value10 = !m.Value10
		}
	}

	return o
}

func size53(m *MsgBool3) uint64 {
	var n uint64 = 1
	return n
}

func marshal53(m *MsgBool3, b []byte) uint64 {
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

func unmarshal53(m *MsgBool3, b []byte) uint64 {
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

func makePatch53(m, mSrc *MsgBool3, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		if m.Value1 == mSrc.Value1 {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
		}
	}
	{
		// Value2

		if m.Value2 == mSrc.Value2 {
			b[0] &= 0xFD
		} else {
			b[0] |= 0x02
		}
	}
	{
		// Value3

		if m.Value3 == mSrc.Value3 {
			b[0] &= 0xFB
		} else {
			b[0] |= 0x04
		}
	}

	return o
}

func applyPatch53(m *MsgBool3, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		if b[0]&0x01 != 0 {
			m.Value1 = !m.Value1
		}
	}
	{
		// Value2

		if b[0]&0x02 != 0 {
			m.Value2 = !m.Value2
		}
	}
	{
		// Value3

		if b[0]&0x04 != 0 {
			m.Value3 = !m.Value3
		}
	}

	return o
}

func size54(m *MsgInt8) uint64 {
	var n uint64 = 1
	return n
}

func marshal54(m *MsgInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

func unmarshal54(m *MsgInt8, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = int8(b[o])
		o++
	}

	return o
}

func makePatch54(m, mSrc *MsgInt8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			b[o] = byte(m.Value)
			o++
		}
	}

	return o
}

func applyPatch54(m *MsgInt8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = int8(b[o])
			o++
		}
	}

	return o
}

func size55(m *MsgInt16) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int16Size(m.Value, &n)
	}
	return n
}

func marshal55(m *MsgInt16, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal55(m *MsgInt16, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch55(m, mSrc *MsgInt16, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.Int16Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch55(m *MsgInt16, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.Int16Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size56(m *MsgInt32) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int32Size(m.Value, &n)
	}
	return n
}

func marshal56(m *MsgInt32, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal56(m *MsgInt32, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch56(m, mSrc *MsgInt32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.Int32Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch56(m *MsgInt32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.Int32Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size57(m *MsgInt64) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.Int64Size(m.Value, &n)
	}
	return n
}

func marshal57(m *MsgInt64, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal57(m *MsgInt64, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.Int64Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch57(m, mSrc *MsgInt64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.Int64Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch57(m *MsgInt64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.Int64Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size58(m *MsgUint8) uint64 {
	var n uint64 = 1
	return n
}

func marshal58(m *MsgUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = m.Value
		o++
	}

	return o
}

func unmarshal58(m *MsgUint8, b []byte) uint64 {
	var o uint64
	{
		// Value

		m.Value = b[o]
		o++
	}

	return o
}

func makePatch58(m, mSrc *MsgUint8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			b[o] = m.Value
			o++
		}
	}

	return o
}

func applyPatch58(m *MsgUint8, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			m.Value = b[o]
			o++
		}
	}

	return o
}

func size59(m *MsgUint16) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt16Size(m.Value, &n)
	}
	return n
}

func marshal59(m *MsgUint16, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt16Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal59(m *MsgUint16, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt16Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch59(m, mSrc *MsgUint16, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt16Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch59(m *MsgUint16, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.UInt16Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size60(m *MsgUint32) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt32Size(m.Value, &n)
	}
	return n
}

func marshal60(m *MsgUint32, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt32Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal60(m *MsgUint32, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt32Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch60(m, mSrc *MsgUint32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt32Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch60(m *MsgUint32, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.UInt32Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size61(m *MsgUint64) uint64 {
	var n uint64 = 1
	{
		// Value

		helpers.UInt64Size(m.Value, &n)
	}
	return n
}

func marshal61(m *MsgUint64, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Marshal(m.Value, b, &o)
	}

	return o
}

func unmarshal61(m *MsgUint64, b []byte) uint64 {
	var o uint64
	{
		// Value

		helpers.UInt64Unmarshal(&m.Value, b, &o)
	}

	return o
}

func makePatch61(m, mSrc *MsgUint64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if reflect.DeepEqual(m.Value, mSrc.Value) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
			helpers.UInt64Marshal(m.Value, b, &o)
		}
	}

	return o
}

func applyPatch61(m *MsgUint64, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value

		if b[0]&0x01 != 0 {
			helpers.UInt64Unmarshal(&m.Value, b, &o)
		}
	}

	return o
}

func size41(m *pkg2.SubMsg) uint64 {
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

func marshal41(m *pkg2.SubMsg, b []byte) uint64 {
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

func unmarshal41(m *pkg2.SubMsg, b []byte) uint64 {
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

func makePatch41(m, mSrc *pkg2.SubMsg, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		if reflect.DeepEqual(m.Value1, mSrc.Value1) {
			b[0] &= 0xFE
		} else {
			b[0] |= 0x01
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
	}
	{
		// Value2

		if reflect.DeepEqual(m.Value2, mSrc.Value2) {
			b[0] &= 0xFD
		} else {
			b[0] |= 0x02
			{
				l := uint64(len(m.Value2))
				helpers.UInt64Marshal(l, b, &o)
				copy(b[o:o+l], m.Value2)
				o += l
			}
		}
	}
	{
		// Value3

		if reflect.DeepEqual(m.Value3, mSrc.Value3) {
			b[0] &= 0xFB
		} else {
			b[0] |= 0x04
			helpers.UInt16Marshal(m.Value3, b, &o)
		}
	}

	return o
}

func applyPatch41(m *pkg2.SubMsg, b []byte) uint64 {
	var o uint64 = 1
	{
		// Value1

		if b[0]&0x01 != 0 {
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
	}
	{
		// Value2

		if b[0]&0x02 != 0 {
			{
				var l uint64
				helpers.UInt64Unmarshal(&l, b, &o)
				if l > 0 {
					m.Value2 = string(b[o:o+l])
					o += l
				}
			}
		}
	}
	{
		// Value3

		if b[0]&0x04 != 0 {
			helpers.UInt16Unmarshal(&m.Value3, b, &o)
		}
	}

	return o
}
