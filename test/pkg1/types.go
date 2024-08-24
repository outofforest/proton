package pkg1

import (
	"github.com/outofforest/proton/test"
	spkg1 "github.com/outofforest/proton/test/pkg1/spkg"
	"github.com/outofforest/proton/test/pkg2"
	spkg2 "github.com/outofforest/proton/test/pkg2/spkg"
)

// MsgUint64 is used in tests
type MsgUint64 struct {
	Value uint64
}

// MsgUint64Custom is used in tests
type MsgUint64Custom struct {
	Value test.CustomUint64
}

// MsgUint32 is used in tests
type MsgUint32 struct {
	Value uint32
}

// MsgUint32Custom is used in tests
type MsgUint32Custom struct {
	Value test.CustomUint32
}

// MsgUint16 is used in tests
type MsgUint16 struct {
	Value uint16
}

// MsgUint16Custom is used in tests
type MsgUint16Custom struct {
	Value test.CustomUint16
}

// MsgUint8 is used in tests
type MsgUint8 struct {
	Value uint8
}

// MsgUint8Custom is used in tests
type MsgUint8Custom struct {
	Value test.CustomUint8
}

// MsgInt64 is used in tests
type MsgInt64 struct {
	Value int64
}

// MsgInt64Custom is used in tests
type MsgInt64Custom struct {
	Value test.CustomInt64
}

// MsgInt32 is used in tests
type MsgInt32 struct {
	Value int32
}

// MsgInt32Custom is used in tests
type MsgInt32Custom struct {
	Value test.CustomInt32
}

// MsgInt16 is used in tests
type MsgInt16 struct {
	Value int16
}

// MsgInt16Custom is used in tests
type MsgInt16Custom struct {
	Value test.CustomInt16
}

// MsgInt8 is used in tests
type MsgInt8 struct {
	Value int8
}

// MsgInt8Custom is used in tests
type MsgInt8Custom struct {
	Value test.CustomInt8
}

// MsgBool3 is used in tests
type MsgBool3 struct {
	Value1 bool
	Value2 bool
	Value3 bool
}

// MsgBool10 is used in tests
type MsgBool10 struct {
	Value1  bool
	Value2  bool
	Value3  bool
	Value4  bool
	Value5  bool
	Value6  bool
	Value7  bool
	Value8  bool
	Value9  bool
	Value10 bool
}

// MsgBoolCustom is used in tests
type MsgBoolCustom struct {
	Value test.CustomBool
}

// MsgFloat64 is used in tests
type MsgFloat64 struct {
	Value float64
}

// MsgFloat64Custom is used in tests
type MsgFloat64Custom struct {
	Value test.CustomFloat64
}

// MsgFloat32 is used in tests
type MsgFloat32 struct {
	Value float32
}

// MsgFloat32Custom is used in tests
type MsgFloat32Custom struct {
	Value test.CustomFloat32
}

// MsgString is used in tests
type MsgString struct {
	Value string
}

// MsgStringCustom is used in tests
type MsgStringCustom struct {
	Value test.CustomString
}

// MsgArray is used in tests
type MsgArray struct {
	Value [3]bool
}

// MsgArrayCustom is used in tests
type MsgArrayCustom struct {
	Value test.CustomArray
}

// MsgSlice is used in tests
type MsgSlice struct {
	Value []bool
}

// MsgSliceCustom is used in tests
type MsgSliceCustom struct {
	Value test.CustomSlice
}

// MsgMap is used in tests
type MsgMap struct {
	Value map[uint64]uint64
}

// MsgMapString is used in tests
type MsgMapString struct {
	Value map[string]string
}

// MsgMapCustom is used in tests
type MsgMapCustom struct {
	Value test.CustomMap
}

// MsgStruct is used in tests
type MsgStruct struct {
	Value1 SubMsg
	Value2 pkg2.SubMsg
}

// MsgStructAnonymous is used in tests
type MsgStructAnonymous struct {
	SubMsg
	Value2 pkg2.SubMsg
}

// MsgArrayUint8 is used in tests
type MsgArrayUint8 struct {
	Value [5]uint8
}

// MsgArrayUint8Custom is used in tests
type MsgArrayUint8Custom struct {
	Value test.CustomArrayUint8
}

// MsgArrayUint8Custom2 is used in tests
type MsgArrayUint8Custom2 struct {
	Value test.CustomArrayCustomUint8
}

// MsgArrayInt8 is used in tests
type MsgArrayInt8 struct {
	Value [5]int8
}

// MsgArrayInt8Custom is used in tests
type MsgArrayInt8Custom struct {
	Value test.CustomArrayInt8
}

// MsgArrayFloat32 is used in tests
type MsgArrayFloat32 struct {
	Value [5]float32
}

// MsgArrayFloat32Custom is used in tests
type MsgArrayFloat32Custom struct {
	Value test.CustomArrayFloat32
}

// MsgArrayFloat64 is used in tests
type MsgArrayFloat64 struct {
	Value [5]float64
}

// MsgArrayFloat64Custom is used in tests
type MsgArrayFloat64Custom struct {
	Value test.CustomArrayFloat64
}

// MsgSliceUint8 is used in tests
type MsgSliceUint8 struct {
	Value []uint8
}

// MsgSliceUint8Custom is used in tests
type MsgSliceUint8Custom struct {
	Value test.CustomSliceUint8
}

// MsgSliceUint8Custom2 is used in tests
type MsgSliceUint8Custom2 struct {
	Value test.CustomSliceCustomUint8
}

// MsgSliceInt8 is used in tests
type MsgSliceInt8 struct {
	Value []int8
}

// MsgSliceInt8Custom is used in tests
type MsgSliceInt8Custom struct {
	Value test.CustomSliceInt8
}

// MsgSliceFloat32 is used in tests
type MsgSliceFloat32 struct {
	Value []float32
}

// MsgSliceFloat32Custom is used in tests
type MsgSliceFloat32Custom struct {
	Value test.CustomSliceFloat32
}

// MsgSliceFloat64 is used in tests
type MsgSliceFloat64 struct {
	Value []float64
}

// MsgSliceFloat64Custom is used in tests
type MsgSliceFloat64Custom struct {
	Value test.CustomSliceFloat64
}

// MsgMixed is used in tests
type MsgMixed struct {
	Value1 map[string]spkg1.SubMsg
	Value2 map[uint8][]string
	Value3 [][32]uint16
	Value4 [12]map[int8]float32
	Value5 [][3][]map[int16][2]int64
	Value6 bool
	Value7 bool
	Value8 string
}

// MsgMixedCustom is used in tests
type MsgMixedCustom struct {
	Value test.CustomMixed
}

// SubMsg is used in tests
type SubMsg struct {
	Value1 []float32
	Value2 []spkg1.SubMsg
	Value3 []spkg2.SubMsg
}
