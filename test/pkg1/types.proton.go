package pkg1

import (
	"unsafe"

	"github.com/outofforest/mass"
	"github.com/outofforest/proton"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1/spkg"
	spkg1 "github.com/outofforest/proton/test/pkg2/spkg"
	"github.com/pkg/errors"
)

const (
	// IDMsgUint64 is the ID of MsgUint64 message.
	IDMsgUint64 uint64 = iota + 1
	// IDMsgUint32 is the ID of MsgUint32 message.
	IDMsgUint32
	// IDMsgUint16 is the ID of MsgUint16 message.
	IDMsgUint16
	// IDMsgUint8 is the ID of MsgUint8 message.
	IDMsgUint8
	// IDMsgInt64 is the ID of MsgInt64 message.
	IDMsgInt64
	// IDMsgInt32 is the ID of MsgInt32 message.
	IDMsgInt32
	// IDMsgInt16 is the ID of MsgInt16 message.
	IDMsgInt16
	// IDMsgInt8 is the ID of MsgInt8 message.
	IDMsgInt8
	// IDMsgBool3 is the ID of MsgBool3 message.
	IDMsgBool3
	// IDMsgBool10 is the ID of MsgBool10 message.
	IDMsgBool10
	// IDMsgFloat64 is the ID of MsgFloat64 message.
	IDMsgFloat64
	// IDMsgFloat32 is the ID of MsgFloat32 message.
	IDMsgFloat32
	// IDMsgString is the ID of MsgString message.
	IDMsgString
	// IDMsgArray is the ID of MsgArray message.
	IDMsgArray
	// IDMsgSlice is the ID of MsgSlice message.
	IDMsgSlice
	// IDMsgMap is the ID of MsgMap message.
	IDMsgMap
	// IDMsgMapString is the ID of MsgMapString message.
	IDMsgMapString
	// IDMsgStruct is the ID of MsgStruct message.
	IDMsgStruct
	// IDMsgStructAnonymous is the ID of MsgStructAnonymous message.
	IDMsgStructAnonymous
	// IDMsgArrayUint8 is the ID of MsgArrayUint8 message.
	IDMsgArrayUint8
	// IDMsgArrayInt8 is the ID of MsgArrayInt8 message.
	IDMsgArrayInt8
	// IDMsgArrayFloat32 is the ID of MsgArrayFloat32 message.
	IDMsgArrayFloat32
	// IDMsgArrayFloat64 is the ID of MsgArrayFloat64 message.
	IDMsgArrayFloat64
	// IDMsgSliceUint8 is the ID of MsgSliceUint8 message.
	IDMsgSliceUint8
	// IDMsgSliceInt8 is the ID of MsgSliceInt8 message.
	IDMsgSliceInt8
	// IDMsgSliceFloat32 is the ID of MsgSliceFloat32 message.
	IDMsgSliceFloat32
	// IDMsgSliceFloat64 is the ID of MsgSliceFloat64 message.
	IDMsgSliceFloat64
	// IDMsgMixed is the ID of MsgMixed message.
	IDMsgMixed
	// IDMsgUint64Custom is the ID of MsgUint64Custom message.
	IDMsgUint64Custom
	// IDMsgUint32Custom is the ID of MsgUint32Custom message.
	IDMsgUint32Custom
	// IDMsgUint16Custom is the ID of MsgUint16Custom message.
	IDMsgUint16Custom
	// IDMsgUint8Custom is the ID of MsgUint8Custom message.
	IDMsgUint8Custom
	// IDMsgInt64Custom is the ID of MsgInt64Custom message.
	IDMsgInt64Custom
	// IDMsgInt32Custom is the ID of MsgInt32Custom message.
	IDMsgInt32Custom
	// IDMsgInt16Custom is the ID of MsgInt16Custom message.
	IDMsgInt16Custom
	// IDMsgInt8Custom is the ID of MsgInt8Custom message.
	IDMsgInt8Custom
	// IDMsgBoolCustom is the ID of MsgBoolCustom message.
	IDMsgBoolCustom
	// IDMsgFloat64Custom is the ID of MsgFloat64Custom message.
	IDMsgFloat64Custom
	// IDMsgFloat32Custom is the ID of MsgFloat32Custom message.
	IDMsgFloat32Custom
	// IDMsgStringCustom is the ID of MsgStringCustom message.
	IDMsgStringCustom
	// IDMsgArrayCustom is the ID of MsgArrayCustom message.
	IDMsgArrayCustom
	// IDMsgSliceCustom is the ID of MsgSliceCustom message.
	IDMsgSliceCustom
	// IDMsgMapCustom is the ID of MsgMapCustom message.
	IDMsgMapCustom
	// IDMsgArrayUint8Custom is the ID of MsgArrayUint8Custom message.
	IDMsgArrayUint8Custom
	// IDMsgArrayUint8Custom2 is the ID of MsgArrayUint8Custom2 message.
	IDMsgArrayUint8Custom2
	// IDMsgArrayInt8Custom is the ID of MsgArrayInt8Custom message.
	IDMsgArrayInt8Custom
	// IDMsgArrayFloat32Custom is the ID of MsgArrayFloat32Custom message.
	IDMsgArrayFloat32Custom
	// IDMsgArrayFloat64Custom is the ID of MsgArrayFloat64Custom message.
	IDMsgArrayFloat64Custom
	// IDMsgSliceUint8Custom is the ID of MsgSliceUint8Custom message.
	IDMsgSliceUint8Custom
	// IDMsgSliceUint8Custom2 is the ID of MsgSliceUint8Custom2 message.
	IDMsgSliceUint8Custom2
	// IDMsgSliceInt8Custom is the ID of MsgSliceInt8Custom message.
	IDMsgSliceInt8Custom
	// IDMsgSliceFloat32Custom is the ID of MsgSliceFloat32Custom message.
	IDMsgSliceFloat32Custom
	// IDMsgSliceFloat64Custom is the ID of MsgSliceFloat64Custom message.
	IDMsgSliceFloat64Custom
	// IDMsgMixedCustom is the ID of MsgMixedCustom message.
	IDMsgMixedCustom
)

var _ proton.Marshaller = Marshaller{}

// NewMarshaller creates marshaller.
func NewMarshaller(capacity uint64) Marshaller {
	return Marshaller{
		massMsgMixedCustom:             mass.New[MsgMixedCustom](capacity),
		massMsgSliceFloat64Custom:      mass.New[MsgSliceFloat64Custom](capacity),
		massMsgSliceFloat32Custom:      mass.New[MsgSliceFloat32Custom](capacity),
		massMsgSliceInt8Custom:         mass.New[MsgSliceInt8Custom](capacity),
		massMsgSliceUint8Custom2:       mass.New[MsgSliceUint8Custom2](capacity),
		massMsgSliceUint8Custom:        mass.New[MsgSliceUint8Custom](capacity),
		massMsgArrayFloat64Custom:      mass.New[MsgArrayFloat64Custom](capacity),
		massMsgArrayFloat32Custom:      mass.New[MsgArrayFloat32Custom](capacity),
		massMsgArrayInt8Custom:         mass.New[MsgArrayInt8Custom](capacity),
		massMsgArrayUint8Custom2:       mass.New[MsgArrayUint8Custom2](capacity),
		massMsgArrayUint8Custom:        mass.New[MsgArrayUint8Custom](capacity),
		massMsgMapCustom:               mass.New[MsgMapCustom](capacity),
		massMsgSliceCustom:             mass.New[MsgSliceCustom](capacity),
		massMsgArrayCustom:             mass.New[MsgArrayCustom](capacity),
		massMsgStringCustom:            mass.New[MsgStringCustom](capacity),
		massMsgFloat32Custom:           mass.New[MsgFloat32Custom](capacity),
		massMsgFloat64Custom:           mass.New[MsgFloat64Custom](capacity),
		massMsgBoolCustom:              mass.New[MsgBoolCustom](capacity),
		massMsgInt8Custom:              mass.New[MsgInt8Custom](capacity),
		massMsgInt16Custom:             mass.New[MsgInt16Custom](capacity),
		massMsgInt32Custom:             mass.New[MsgInt32Custom](capacity),
		massMsgInt64Custom:             mass.New[MsgInt64Custom](capacity),
		massMsgUint8Custom:             mass.New[MsgUint8Custom](capacity),
		massMsgUint16Custom:            mass.New[MsgUint16Custom](capacity),
		massMsgUint32Custom:            mass.New[MsgUint32Custom](capacity),
		massMsgUint64Custom:            mass.New[MsgUint64Custom](capacity),
		massMsgMixed:                   mass.New[MsgMixed](capacity),
		massMsgSliceFloat64:            mass.New[MsgSliceFloat64](capacity),
		massMsgSliceFloat32:            mass.New[MsgSliceFloat32](capacity),
		massMsgSliceInt8:               mass.New[MsgSliceInt8](capacity),
		massMsgSliceUint8:              mass.New[MsgSliceUint8](capacity),
		massMsgArrayFloat64:            mass.New[MsgArrayFloat64](capacity),
		massMsgArrayFloat32:            mass.New[MsgArrayFloat32](capacity),
		massMsgArrayInt8:               mass.New[MsgArrayInt8](capacity),
		massMsgArrayUint8:              mass.New[MsgArrayUint8](capacity),
		massMsgStructAnonymous:         mass.New[MsgStructAnonymous](capacity),
		massMsgStruct:                  mass.New[MsgStruct](capacity),
		massMsgMapString:               mass.New[MsgMapString](capacity),
		massMsgMap:                     mass.New[MsgMap](capacity),
		massMsgSlice:                   mass.New[MsgSlice](capacity),
		massMsgArray:                   mass.New[MsgArray](capacity),
		massMsgString:                  mass.New[MsgString](capacity),
		massMsgFloat32:                 mass.New[MsgFloat32](capacity),
		massMsgFloat64:                 mass.New[MsgFloat64](capacity),
		massMsgBool10:                  mass.New[MsgBool10](capacity),
		massMsgBool3:                   mass.New[MsgBool3](capacity),
		massMsgInt8:                    mass.New[MsgInt8](capacity),
		massMsgInt16:                   mass.New[MsgInt16](capacity),
		massMsgInt32:                   mass.New[MsgInt32](capacity),
		massMsgInt64:                   mass.New[MsgInt64](capacity),
		massMsgUint8:                   mass.New[MsgUint8](capacity),
		massMsgUint16:                  mass.New[MsgUint16](capacity),
		massMsgUint32:                  mass.New[MsgUint32](capacity),
		massMsgUint64:                  mass.New[MsgUint64](capacity),
		massA3BABMapAInt16BCustomArray: mass.New[[3][]map[int16]custom.Array](capacity),
		massMapAInt16BCustomArray:      mass.New[map[int16]custom.Array](capacity),
		massCustomFloat64:              mass.New[custom.Float64](capacity),
		massCustomFloat32:              mass.New[custom.Float32](capacity),
		massCustomInt8:                 mass.New[custom.Int8](capacity),
		massCustomUint8:                mass.New[custom.Uint8](capacity),
		massCustomUint64:               mass.New[custom.Uint64](capacity),
		massString:                     mass.New[string](capacity),
		massA32BUint16:                 mass.New[[32]uint16](capacity),
		massA3BABMapAInt16BA2BInt64:    mass.New[[3][]map[int16][2]int64](capacity),
		massMapAInt16BA2BInt64:         mass.New[map[int16][2]int64](capacity),
		massFloat64:                    mass.New[float64](capacity),
		massFloat32:                    mass.New[float32](capacity),
		massInt8:                       mass.New[int8](capacity),
		massSpkgSubMsg:                 mass.New[spkg.SubMsg](capacity),
		massSpkg1SubMsg:                mass.New[spkg1.SubMsg](capacity),
		massBool:                       mass.New[bool](capacity),
	}
}

// Marshaller marshals and unmarshals messages.
type Marshaller struct {
	massMsgMixedCustom             *mass.Mass[MsgMixedCustom]
	massMsgSliceFloat64Custom      *mass.Mass[MsgSliceFloat64Custom]
	massMsgSliceFloat32Custom      *mass.Mass[MsgSliceFloat32Custom]
	massMsgSliceInt8Custom         *mass.Mass[MsgSliceInt8Custom]
	massMsgSliceUint8Custom2       *mass.Mass[MsgSliceUint8Custom2]
	massMsgSliceUint8Custom        *mass.Mass[MsgSliceUint8Custom]
	massMsgArrayFloat64Custom      *mass.Mass[MsgArrayFloat64Custom]
	massMsgArrayFloat32Custom      *mass.Mass[MsgArrayFloat32Custom]
	massMsgArrayInt8Custom         *mass.Mass[MsgArrayInt8Custom]
	massMsgArrayUint8Custom2       *mass.Mass[MsgArrayUint8Custom2]
	massMsgArrayUint8Custom        *mass.Mass[MsgArrayUint8Custom]
	massMsgMapCustom               *mass.Mass[MsgMapCustom]
	massMsgSliceCustom             *mass.Mass[MsgSliceCustom]
	massMsgArrayCustom             *mass.Mass[MsgArrayCustom]
	massMsgStringCustom            *mass.Mass[MsgStringCustom]
	massMsgFloat32Custom           *mass.Mass[MsgFloat32Custom]
	massMsgFloat64Custom           *mass.Mass[MsgFloat64Custom]
	massMsgBoolCustom              *mass.Mass[MsgBoolCustom]
	massMsgInt8Custom              *mass.Mass[MsgInt8Custom]
	massMsgInt16Custom             *mass.Mass[MsgInt16Custom]
	massMsgInt32Custom             *mass.Mass[MsgInt32Custom]
	massMsgInt64Custom             *mass.Mass[MsgInt64Custom]
	massMsgUint8Custom             *mass.Mass[MsgUint8Custom]
	massMsgUint16Custom            *mass.Mass[MsgUint16Custom]
	massMsgUint32Custom            *mass.Mass[MsgUint32Custom]
	massMsgUint64Custom            *mass.Mass[MsgUint64Custom]
	massMsgMixed                   *mass.Mass[MsgMixed]
	massMsgSliceFloat64            *mass.Mass[MsgSliceFloat64]
	massMsgSliceFloat32            *mass.Mass[MsgSliceFloat32]
	massMsgSliceInt8               *mass.Mass[MsgSliceInt8]
	massMsgSliceUint8              *mass.Mass[MsgSliceUint8]
	massMsgArrayFloat64            *mass.Mass[MsgArrayFloat64]
	massMsgArrayFloat32            *mass.Mass[MsgArrayFloat32]
	massMsgArrayInt8               *mass.Mass[MsgArrayInt8]
	massMsgArrayUint8              *mass.Mass[MsgArrayUint8]
	massMsgStructAnonymous         *mass.Mass[MsgStructAnonymous]
	massMsgStruct                  *mass.Mass[MsgStruct]
	massMsgMapString               *mass.Mass[MsgMapString]
	massMsgMap                     *mass.Mass[MsgMap]
	massMsgSlice                   *mass.Mass[MsgSlice]
	massMsgArray                   *mass.Mass[MsgArray]
	massMsgString                  *mass.Mass[MsgString]
	massMsgFloat32                 *mass.Mass[MsgFloat32]
	massMsgFloat64                 *mass.Mass[MsgFloat64]
	massMsgBool10                  *mass.Mass[MsgBool10]
	massMsgBool3                   *mass.Mass[MsgBool3]
	massMsgInt8                    *mass.Mass[MsgInt8]
	massMsgInt16                   *mass.Mass[MsgInt16]
	massMsgInt32                   *mass.Mass[MsgInt32]
	massMsgInt64                   *mass.Mass[MsgInt64]
	massMsgUint8                   *mass.Mass[MsgUint8]
	massMsgUint16                  *mass.Mass[MsgUint16]
	massMsgUint32                  *mass.Mass[MsgUint32]
	massMsgUint64                  *mass.Mass[MsgUint64]
	massA3BABMapAInt16BCustomArray *mass.Mass[[3][]map[int16]custom.Array]
	massMapAInt16BCustomArray      *mass.Mass[map[int16]custom.Array]
	massCustomFloat64              *mass.Mass[custom.Float64]
	massCustomFloat32              *mass.Mass[custom.Float32]
	massCustomInt8                 *mass.Mass[custom.Int8]
	massCustomUint8                *mass.Mass[custom.Uint8]
	massCustomUint64               *mass.Mass[custom.Uint64]
	massString                     *mass.Mass[string]
	massA32BUint16                 *mass.Mass[[32]uint16]
	massA3BABMapAInt16BA2BInt64    *mass.Mass[[3][]map[int16][2]int64]
	massMapAInt16BA2BInt64         *mass.Mass[map[int16][2]int64]
	massFloat64                    *mass.Mass[float64]
	massFloat32                    *mass.Mass[float32]
	massInt8                       *mass.Mass[int8]
	massSpkgSubMsg                 *mass.Mass[spkg.SubMsg]
	massSpkg1SubMsg                *mass.Mass[spkg1.SubMsg]
	massBool                       *mass.Mass[bool]
}

// Marshal marshals message.
func (m Marshaller) Marshal(msg proton.Marshallable, buf []byte) (retID, retSize uint64, retErr error) {
	defer func() {
		if res := recover(); res != nil {
			retErr = errors.Errorf("marshaling message failed: %s", res)
		}
	}()

	switch msg2 := msg.(type) {
	case *MsgMixedCustom:
		return IDMsgMixedCustom, msg2.Marshal(buf), nil
	case *MsgSliceFloat64Custom:
		return IDMsgSliceFloat64Custom, msg2.Marshal(buf), nil
	case *MsgSliceFloat32Custom:
		return IDMsgSliceFloat32Custom, msg2.Marshal(buf), nil
	case *MsgSliceInt8Custom:
		return IDMsgSliceInt8Custom, msg2.Marshal(buf), nil
	case *MsgSliceUint8Custom2:
		return IDMsgSliceUint8Custom2, msg2.Marshal(buf), nil
	case *MsgSliceUint8Custom:
		return IDMsgSliceUint8Custom, msg2.Marshal(buf), nil
	case *MsgArrayFloat64Custom:
		return IDMsgArrayFloat64Custom, msg2.Marshal(buf), nil
	case *MsgArrayFloat32Custom:
		return IDMsgArrayFloat32Custom, msg2.Marshal(buf), nil
	case *MsgArrayInt8Custom:
		return IDMsgArrayInt8Custom, msg2.Marshal(buf), nil
	case *MsgArrayUint8Custom2:
		return IDMsgArrayUint8Custom2, msg2.Marshal(buf), nil
	case *MsgArrayUint8Custom:
		return IDMsgArrayUint8Custom, msg2.Marshal(buf), nil
	case *MsgMapCustom:
		return IDMsgMapCustom, msg2.Marshal(buf), nil
	case *MsgSliceCustom:
		return IDMsgSliceCustom, msg2.Marshal(buf), nil
	case *MsgArrayCustom:
		return IDMsgArrayCustom, msg2.Marshal(buf), nil
	case *MsgStringCustom:
		return IDMsgStringCustom, msg2.Marshal(buf), nil
	case *MsgFloat32Custom:
		return IDMsgFloat32Custom, msg2.Marshal(buf), nil
	case *MsgFloat64Custom:
		return IDMsgFloat64Custom, msg2.Marshal(buf), nil
	case *MsgBoolCustom:
		return IDMsgBoolCustom, msg2.Marshal(buf), nil
	case *MsgInt8Custom:
		return IDMsgInt8Custom, msg2.Marshal(buf), nil
	case *MsgInt16Custom:
		return IDMsgInt16Custom, msg2.Marshal(buf), nil
	case *MsgInt32Custom:
		return IDMsgInt32Custom, msg2.Marshal(buf), nil
	case *MsgInt64Custom:
		return IDMsgInt64Custom, msg2.Marshal(buf), nil
	case *MsgUint8Custom:
		return IDMsgUint8Custom, msg2.Marshal(buf), nil
	case *MsgUint16Custom:
		return IDMsgUint16Custom, msg2.Marshal(buf), nil
	case *MsgUint32Custom:
		return IDMsgUint32Custom, msg2.Marshal(buf), nil
	case *MsgUint64Custom:
		return IDMsgUint64Custom, msg2.Marshal(buf), nil
	case *MsgMixed:
		return IDMsgMixed, msg2.Marshal(buf), nil
	case *MsgSliceFloat64:
		return IDMsgSliceFloat64, msg2.Marshal(buf), nil
	case *MsgSliceFloat32:
		return IDMsgSliceFloat32, msg2.Marshal(buf), nil
	case *MsgSliceInt8:
		return IDMsgSliceInt8, msg2.Marshal(buf), nil
	case *MsgSliceUint8:
		return IDMsgSliceUint8, msg2.Marshal(buf), nil
	case *MsgArrayFloat64:
		return IDMsgArrayFloat64, msg2.Marshal(buf), nil
	case *MsgArrayFloat32:
		return IDMsgArrayFloat32, msg2.Marshal(buf), nil
	case *MsgArrayInt8:
		return IDMsgArrayInt8, msg2.Marshal(buf), nil
	case *MsgArrayUint8:
		return IDMsgArrayUint8, msg2.Marshal(buf), nil
	case *MsgStructAnonymous:
		return IDMsgStructAnonymous, msg2.Marshal(buf), nil
	case *MsgStruct:
		return IDMsgStruct, msg2.Marshal(buf), nil
	case *MsgMapString:
		return IDMsgMapString, msg2.Marshal(buf), nil
	case *MsgMap:
		return IDMsgMap, msg2.Marshal(buf), nil
	case *MsgSlice:
		return IDMsgSlice, msg2.Marshal(buf), nil
	case *MsgArray:
		return IDMsgArray, msg2.Marshal(buf), nil
	case *MsgString:
		return IDMsgString, msg2.Marshal(buf), nil
	case *MsgFloat32:
		return IDMsgFloat32, msg2.Marshal(buf), nil
	case *MsgFloat64:
		return IDMsgFloat64, msg2.Marshal(buf), nil
	case *MsgBool10:
		return IDMsgBool10, msg2.Marshal(buf), nil
	case *MsgBool3:
		return IDMsgBool3, msg2.Marshal(buf), nil
	case *MsgInt8:
		return IDMsgInt8, msg2.Marshal(buf), nil
	case *MsgInt16:
		return IDMsgInt16, msg2.Marshal(buf), nil
	case *MsgInt32:
		return IDMsgInt32, msg2.Marshal(buf), nil
	case *MsgInt64:
		return IDMsgInt64, msg2.Marshal(buf), nil
	case *MsgUint8:
		return IDMsgUint8, msg2.Marshal(buf), nil
	case *MsgUint16:
		return IDMsgUint16, msg2.Marshal(buf), nil
	case *MsgUint32:
		return IDMsgUint32, msg2.Marshal(buf), nil
	case *MsgUint64:
		return IDMsgUint64, msg2.Marshal(buf), nil
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
	case IDMsgMixedCustom:
		msg := m.massMsgMixedCustom.New()
		return msg, msg.Unmarshal(
			buf,
			m.massA3BABMapAInt16BCustomArray,
			m.massMapAInt16BCustomArray,
		), nil
	case IDMsgSliceFloat64Custom:
		msg := m.massMsgSliceFloat64Custom.New()
		return msg, msg.Unmarshal(
			buf,
			m.massCustomFloat64,
		), nil
	case IDMsgSliceFloat32Custom:
		msg := m.massMsgSliceFloat32Custom.New()
		return msg, msg.Unmarshal(
			buf,
			m.massCustomFloat32,
		), nil
	case IDMsgSliceInt8Custom:
		msg := m.massMsgSliceInt8Custom.New()
		return msg, msg.Unmarshal(
			buf,
			m.massCustomInt8,
		), nil
	case IDMsgSliceUint8Custom2:
		msg := m.massMsgSliceUint8Custom2.New()
		return msg, msg.Unmarshal(
			buf,
			m.massCustomUint8,
		), nil
	case IDMsgSliceUint8Custom:
		msg := m.massMsgSliceUint8Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayFloat64Custom:
		msg := m.massMsgArrayFloat64Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayFloat32Custom:
		msg := m.massMsgArrayFloat32Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayInt8Custom:
		msg := m.massMsgArrayInt8Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayUint8Custom2:
		msg := m.massMsgArrayUint8Custom2.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayUint8Custom:
		msg := m.massMsgArrayUint8Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgMapCustom:
		msg := m.massMsgMapCustom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgSliceCustom:
		msg := m.massMsgSliceCustom.New()
		return msg, msg.Unmarshal(
			buf,
			m.massCustomUint64,
		), nil
	case IDMsgArrayCustom:
		msg := m.massMsgArrayCustom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgStringCustom:
		msg := m.massMsgStringCustom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgFloat32Custom:
		msg := m.massMsgFloat32Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgFloat64Custom:
		msg := m.massMsgFloat64Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgBoolCustom:
		msg := m.massMsgBoolCustom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgInt8Custom:
		msg := m.massMsgInt8Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgInt16Custom:
		msg := m.massMsgInt16Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgInt32Custom:
		msg := m.massMsgInt32Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgInt64Custom:
		msg := m.massMsgInt64Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgUint8Custom:
		msg := m.massMsgUint8Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgUint16Custom:
		msg := m.massMsgUint16Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgUint32Custom:
		msg := m.massMsgUint32Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgUint64Custom:
		msg := m.massMsgUint64Custom.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgMixed:
		msg := m.massMsgMixed.New()
		return msg, msg.Unmarshal(
			buf,
			m.massString,
			m.massA32BUint16,
			m.massA3BABMapAInt16BA2BInt64,
			m.massMapAInt16BA2BInt64,
		), nil
	case IDMsgSliceFloat64:
		msg := m.massMsgSliceFloat64.New()
		return msg, msg.Unmarshal(
			buf,
			m.massFloat64,
		), nil
	case IDMsgSliceFloat32:
		msg := m.massMsgSliceFloat32.New()
		return msg, msg.Unmarshal(
			buf,
			m.massFloat32,
		), nil
	case IDMsgSliceInt8:
		msg := m.massMsgSliceInt8.New()
		return msg, msg.Unmarshal(
			buf,
			m.massInt8,
		), nil
	case IDMsgSliceUint8:
		msg := m.massMsgSliceUint8.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayFloat64:
		msg := m.massMsgArrayFloat64.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayFloat32:
		msg := m.massMsgArrayFloat32.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayInt8:
		msg := m.massMsgArrayInt8.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgArrayUint8:
		msg := m.massMsgArrayUint8.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgStructAnonymous:
		msg := m.massMsgStructAnonymous.New()
		return msg, msg.Unmarshal(
			buf,
			m.massFloat32,
			m.massSpkgSubMsg,
			m.massSpkg1SubMsg,
		), nil
	case IDMsgStruct:
		msg := m.massMsgStruct.New()
		return msg, msg.Unmarshal(
			buf,
			m.massFloat32,
			m.massSpkgSubMsg,
			m.massSpkg1SubMsg,
		), nil
	case IDMsgMapString:
		msg := m.massMsgMapString.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgMap:
		msg := m.massMsgMap.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgSlice:
		msg := m.massMsgSlice.New()
		return msg, msg.Unmarshal(
			buf,
			m.massBool,
		), nil
	case IDMsgArray:
		msg := m.massMsgArray.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgString:
		msg := m.massMsgString.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgFloat32:
		msg := m.massMsgFloat32.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgFloat64:
		msg := m.massMsgFloat64.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgBool10:
		msg := m.massMsgBool10.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgBool3:
		msg := m.massMsgBool3.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgInt8:
		msg := m.massMsgInt8.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgInt16:
		msg := m.massMsgInt16.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgInt32:
		msg := m.massMsgInt32.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgInt64:
		msg := m.massMsgInt64.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgUint8:
		msg := m.massMsgUint8.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgUint16:
		msg := m.massMsgUint16.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgUint32:
		msg := m.massMsgUint32.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	case IDMsgUint64:
		msg := m.massMsgUint64.New()
		return msg, msg.Unmarshal(
			buf,
		), nil
	default:
		return nil, 0, errors.Errorf("unknown ID %d", id)
	}
}

var _ proton.Message = &MsgMixedCustom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgMixedCustom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgMixedCustom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgMixedCustom) Unmarshal(
	b []byte,
	massA3BABMapAInt16BCustomArray *mass.Mass[[3][]map[int16]custom.Array],
	massMapAInt16BCustomArray *mass.Mass[map[int16]custom.Array],
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
			m.Value = massA3BABMapAInt16BCustomArray.NewSlice(l)
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
						m.Value[i6][i5] = massMapAInt16BCustomArray.NewSlice(l)
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

var _ proton.Message = &MsgSliceFloat64Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceFloat64Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceFloat64Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceFloat64Custom) Unmarshal(
	b []byte,
	massCustomFloat64 *mass.Mass[custom.Float64],
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
			m.Value = massCustomFloat64.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
			o += l * 8
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgSliceFloat32Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceFloat32Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceFloat32Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceFloat32Custom) Unmarshal(
	b []byte,
	massCustomFloat32 *mass.Mass[custom.Float32],
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
			m.Value = massCustomFloat32.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
			o += l * 4
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgSliceInt8Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceInt8Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceInt8Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceInt8Custom) Unmarshal(
	b []byte,
	massCustomInt8 *mass.Mass[custom.Int8],
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
			m.Value = massCustomInt8.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgSliceUint8Custom2{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceUint8Custom2) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceUint8Custom2) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceUint8Custom2) Unmarshal(
	b []byte,
	massCustomUint8 *mass.Mass[custom.Uint8],
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
			m.Value = massCustomUint8.NewSlice(l)
			copy(unsafe.Slice((*byte)(&m.Value[0]), l), b[o:o+l])
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgSliceUint8Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceUint8Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceUint8Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceUint8Custom) Unmarshal(
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
			m.Value = b[o:o+l]
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgArrayFloat64Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayFloat64Custom) Size() uint64 {
	var n uint64 = 40
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayFloat64Custom) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayFloat64Custom) Unmarshal(
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

var _ proton.Message = &MsgArrayFloat32Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayFloat32Custom) Size() uint64 {
	var n uint64 = 20
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayFloat32Custom) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayFloat32Custom) Unmarshal(
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

var _ proton.Message = &MsgArrayInt8Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayInt8Custom) Size() uint64 {
	var n uint64 = 5
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayInt8Custom) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayInt8Custom) Unmarshal(
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

var _ proton.Message = &MsgArrayUint8Custom2{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayUint8Custom2) Size() uint64 {
	var n uint64 = 5
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayUint8Custom2) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(&m.Value[0]), 5))
		o += 5
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayUint8Custom2) Unmarshal(
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

var _ proton.Message = &MsgArrayUint8Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayUint8Custom) Size() uint64 {
	var n uint64 = 5
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayUint8Custom) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayUint8Custom) Unmarshal(
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

var _ proton.Message = &MsgMapCustom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgMapCustom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgMapCustom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgMapCustom) Unmarshal(
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

var _ proton.Message = &MsgSliceCustom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceCustom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceCustom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceCustom) Unmarshal(
	b []byte,
	massCustomUint64 *mass.Mass[custom.Uint64],
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
			m.Value = massCustomUint64.NewSlice(l)
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

var _ proton.Message = &MsgArrayCustom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayCustom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgArrayCustom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgArrayCustom) Unmarshal(
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

var _ proton.Message = &MsgStringCustom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgStringCustom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgStringCustom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgStringCustom) Unmarshal(
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

var _ proton.Message = &MsgFloat32Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgFloat32Custom) Size() uint64 {
	var n uint64 = 4
	return n
}

// Marshal marshals the structure.
func (m *MsgFloat32Custom) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgFloat32Custom) Unmarshal(
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

var _ proton.Message = &MsgFloat64Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgFloat64Custom) Size() uint64 {
	var n uint64 = 8
	return n
}

// Marshal marshals the structure.
func (m *MsgFloat64Custom) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*custom.Float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgFloat64Custom) Unmarshal(
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

var _ proton.Message = &MsgBoolCustom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgBoolCustom) Size() uint64 {
	var n uint64 = 1
	return n
}

// Marshal marshals the structure.
func (m *MsgBoolCustom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgBoolCustom) Unmarshal(
	b []byte,
) uint64 {
	var o uint64 = 1
	{
		// Value

		m.Value = b[0]&0x01 != 0
	}

	return o
}

var _ proton.Message = &MsgInt8Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgInt8Custom) Size() uint64 {
	var n uint64 = 1
	return n
}

// Marshal marshals the structure.
func (m *MsgInt8Custom) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgInt8Custom) Unmarshal(
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

var _ proton.Message = &MsgInt16Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgInt16Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgInt16Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgInt16Custom) Unmarshal(
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

var _ proton.Message = &MsgInt32Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgInt32Custom) Size() uint64 {
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
func (m *MsgInt32Custom) Marshal(b []byte) uint64 {
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
func (m *MsgInt32Custom) Unmarshal(
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

var _ proton.Message = &MsgInt64Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgInt64Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgInt64Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgInt64Custom) Unmarshal(
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

var _ proton.Message = &MsgUint8Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgUint8Custom) Size() uint64 {
	var n uint64 = 1
	return n
}

// Marshal marshals the structure.
func (m *MsgUint8Custom) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgUint8Custom) Unmarshal(
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

var _ proton.Message = &MsgUint16Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgUint16Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgUint16Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgUint16Custom) Unmarshal(
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

var _ proton.Message = &MsgUint32Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgUint32Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgUint32Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgUint32Custom) Unmarshal(
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

var _ proton.Message = &MsgUint64Custom{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgUint64Custom) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgUint64Custom) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgUint64Custom) Unmarshal(
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

var _ proton.Message = &MsgMixed{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgMixed) Size() uint64 {
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
			n += mv2.Size()
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

// Marshal marshals the structure.
func (m *MsgMixed) Marshal(b []byte) uint64 {
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
			o += mv2.Marshal(b[o:])
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

// Unmarshal unmarshals the structure.
func (m *MsgMixed) Unmarshal(
	b []byte,
	massString *mass.Mass[string],
	massA32BUint16 *mass.Mass[[32]uint16],
	massA3BABMapAInt16BA2BInt64 *mass.Mass[[3][]map[int16][2]int64],
	massMapAInt16BA2BInt64 *mass.Mass[map[int16][2]int64],
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
				o += mv2.Unmarshal(
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
					mv3 = massString.NewSlice(l)
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
			m.Value3 = massA32BUint16.NewSlice(l)
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
			m.Value5 = massA3BABMapAInt16BA2BInt64.NewSlice(l)
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
						m.Value5[i6][i5] = massMapAInt16BA2BInt64.NewSlice(l)
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

var _ proton.Message = &MsgSliceFloat64{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceFloat64) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceFloat64) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceFloat64) Unmarshal(
	b []byte,
	massFloat64 *mass.Mass[float64],
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
			m.Value = massFloat64.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*8), b[o:o+l*8])
			o += l * 8
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgSliceFloat32{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceFloat32) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceFloat32) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceFloat32) Unmarshal(
	b []byte,
	massFloat32 *mass.Mass[float32],
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
			m.Value = massFloat32.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l*4), b[o:o+l*4])
			o += l * 4
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgSliceInt8{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceInt8) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceInt8) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceInt8) Unmarshal(
	b []byte,
	massInt8 *mass.Mass[int8],
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
			m.Value = massInt8.NewSlice(l)
			copy(unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), l), b[o:o+l])
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgSliceUint8{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSliceUint8) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSliceUint8) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSliceUint8) Unmarshal(
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
			m.Value = b[o:o+l]
			o += l
		} else {
			m.Value = nil
		}
	}

	return o
}

var _ proton.Message = &MsgArrayFloat64{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayFloat64) Size() uint64 {
	var n uint64 = 40
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayFloat64) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+40], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 40))
		o += 40
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayFloat64) Unmarshal(
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

var _ proton.Message = &MsgArrayFloat32{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayFloat32) Size() uint64 {
	var n uint64 = 20
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayFloat32) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+20], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 20))
		o += 20
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayFloat32) Unmarshal(
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

var _ proton.Message = &MsgArrayInt8{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayInt8) Size() uint64 {
	var n uint64 = 5
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayInt8) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice((*byte)(unsafe.Pointer(&m.Value[0])), 5))
		o += 5
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayInt8) Unmarshal(
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

var _ proton.Message = &MsgArrayUint8{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArrayUint8) Size() uint64 {
	var n uint64 = 5
	return n
}

// Marshal marshals the structure.
func (m *MsgArrayUint8) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		copy(b[o:o+5], unsafe.Slice(&m.Value[0], 5))
		o += 5
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgArrayUint8) Unmarshal(
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

var _ proton.Message = &MsgStructAnonymous{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgStructAnonymous) Size() uint64 {
	var n uint64
	{
		// SubMsg

		n += m.SubMsg.Size()
	}
	{
		// Value2

		n += m.Value2.Size()
	}
	return n
}

// Marshal marshals the structure.
func (m *MsgStructAnonymous) Marshal(b []byte) uint64 {
	var o uint64
	{
		// SubMsg

		o += m.SubMsg.Marshal(b[o:])
	}
	{
		// Value2

		o += m.Value2.Marshal(b[o:])
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgStructAnonymous) Unmarshal(
	b []byte,
	massFloat32 *mass.Mass[float32],
	massSpkgSubMsg *mass.Mass[spkg.SubMsg],
	massSpkg1SubMsg *mass.Mass[spkg1.SubMsg],
) uint64 {
	var o uint64
	{
		// SubMsg

		o += m.SubMsg.Unmarshal(
			b[o:],
			massFloat32,
			massSpkgSubMsg,
			massSpkg1SubMsg,
		)
	}
	{
		// Value2

		o += m.Value2.Unmarshal(
			b[o:],
		)
	}

	return o
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
			n += sv1.Size()
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
			n += sv1.Size()
		}
	}
	return n
}

// Marshal marshals the structure.
func (m *SubMsg) Marshal(b []byte) uint64 {
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
			o += sv1.Marshal(b[o:])
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
			o += sv1.Marshal(b[o:])
		}
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *SubMsg) Unmarshal(
	b []byte,
	massFloat32 *mass.Mass[float32],
	massSpkgSubMsg *mass.Mass[spkg.SubMsg],
	massSpkg1SubMsg *mass.Mass[spkg1.SubMsg],
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
			m.Value1 = massFloat32.NewSlice(l)
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
			m.Value2 = massSpkgSubMsg.NewSlice(l)
			for i1 := range l {
				o += m.Value2[i1].Unmarshal(
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
			m.Value3 = massSpkg1SubMsg.NewSlice(l)
			for i1 := range l {
				o += m.Value3[i1].Unmarshal(
					b[o:],
				)
			}
		} else {
			m.Value3 = nil
		}
	}

	return o
}

var _ proton.Message = &MsgStruct{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgStruct) Size() uint64 {
	var n uint64
	{
		// Value1

		n += m.Value1.Size()
	}
	{
		// Value2

		n += m.Value2.Size()
	}
	return n
}

// Marshal marshals the structure.
func (m *MsgStruct) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value1

		o += m.Value1.Marshal(b[o:])
	}
	{
		// Value2

		o += m.Value2.Marshal(b[o:])
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgStruct) Unmarshal(
	b []byte,
	massFloat32 *mass.Mass[float32],
	massSpkgSubMsg *mass.Mass[spkg.SubMsg],
	massSpkg1SubMsg *mass.Mass[spkg1.SubMsg],
) uint64 {
	var o uint64
	{
		// Value1

		o += m.Value1.Unmarshal(
			b[o:],
			massFloat32,
			massSpkgSubMsg,
			massSpkg1SubMsg,
		)
	}
	{
		// Value2

		o += m.Value2.Unmarshal(
			b[o:],
		)
	}

	return o
}

var _ proton.Message = &MsgMapString{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgMapString) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgMapString) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgMapString) Unmarshal(
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

var _ proton.Message = &MsgMap{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgMap) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgMap) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgMap) Unmarshal(
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

var _ proton.Message = &MsgSlice{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgSlice) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgSlice) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgSlice) Unmarshal(
	b []byte,
	massBool *mass.Mass[bool],
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
			m.Value = massBool.NewSlice(l)
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

var _ proton.Message = &MsgArray{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgArray) Size() uint64 {
	var n uint64 = 3
	return n
}

// Marshal marshals the structure.
func (m *MsgArray) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgArray) Unmarshal(
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

var _ proton.Message = &MsgString{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgString) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgString) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgString) Unmarshal(
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

var _ proton.Message = &MsgFloat32{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgFloat32) Size() uint64 {
	var n uint64 = 4
	return n
}

// Marshal marshals the structure.
func (m *MsgFloat32) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float32)(unsafe.Pointer(&b[o])) = m.Value
		o += 4
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgFloat32) Unmarshal(
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

var _ proton.Message = &MsgFloat64{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgFloat64) Size() uint64 {
	var n uint64 = 8
	return n
}

// Marshal marshals the structure.
func (m *MsgFloat64) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		*(*float64)(unsafe.Pointer(&b[o])) = m.Value
		o += 8
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgFloat64) Unmarshal(
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

var _ proton.Message = &MsgBool10{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgBool10) Size() uint64 {
	var n uint64 = 2
	return n
}

// Marshal marshals the structure.
func (m *MsgBool10) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgBool10) Unmarshal(
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

var _ proton.Message = &MsgBool3{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgBool3) Size() uint64 {
	var n uint64 = 1
	return n
}

// Marshal marshals the structure.
func (m *MsgBool3) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgBool3) Unmarshal(
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

var _ proton.Message = &MsgInt8{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgInt8) Size() uint64 {
	var n uint64 = 1
	return n
}

// Marshal marshals the structure.
func (m *MsgInt8) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = byte(m.Value)
		o++
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgInt8) Unmarshal(
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

var _ proton.Message = &MsgInt16{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgInt16) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgInt16) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgInt16) Unmarshal(
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

var _ proton.Message = &MsgInt32{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgInt32) Size() uint64 {
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
func (m *MsgInt32) Marshal(b []byte) uint64 {
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
func (m *MsgInt32) Unmarshal(
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

var _ proton.Message = &MsgInt64{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgInt64) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgInt64) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgInt64) Unmarshal(
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

var _ proton.Message = &MsgUint8{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgUint8) Size() uint64 {
	var n uint64 = 1
	return n
}

// Marshal marshals the structure.
func (m *MsgUint8) Marshal(b []byte) uint64 {
	var o uint64
	{
		// Value

		b[o] = m.Value
		o++
	}

	return o
}

// Unmarshal unmarshals the structure.
func (m *MsgUint8) Unmarshal(
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

var _ proton.Message = &MsgUint16{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgUint16) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgUint16) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgUint16) Unmarshal(
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

var _ proton.Message = &MsgUint32{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgUint32) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgUint32) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgUint32) Unmarshal(
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

var _ proton.Message = &MsgUint64{}

// Size computes the required size of the buffer for marshaling the structure.
func (m *MsgUint64) Size() uint64 {
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

// Marshal marshals the structure.
func (m *MsgUint64) Marshal(b []byte) uint64 {
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

// Unmarshal unmarshals the structure.
func (m *MsgUint64) Unmarshal(
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
