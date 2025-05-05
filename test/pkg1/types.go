package pkg1

import (
	"time"

	"github.com/outofforest/proton/test/custom"
	spkg1 "github.com/outofforest/proton/test/pkg1/spkg"
	"github.com/outofforest/proton/test/pkg2"
	spkg2 "github.com/outofforest/proton/test/pkg2/spkg"
)

// MsgUint64 is used in tests.
type MsgUint64 struct {
	Value uint64
}

// MsgUint64Custom is used in tests.
type MsgUint64Custom struct {
	Value custom.Uint64
}

// MsgUint32 is used in tests.
type MsgUint32 struct {
	Value uint32
}

// MsgUint32Custom is used in tests.
type MsgUint32Custom struct {
	Value custom.Uint32
}

// MsgUint16 is used in tests.
type MsgUint16 struct {
	Value uint16
}

// MsgUint16Custom is used in tests.
type MsgUint16Custom struct {
	Value custom.Uint16
}

// MsgUint8 is used in tests.
type MsgUint8 struct {
	Value uint8
}

// MsgUint8Custom is used in tests.
type MsgUint8Custom struct {
	Value custom.Uint8
}

// MsgInt64 is used in tests.
type MsgInt64 struct {
	Value int64
}

// MsgInt64Custom is used in tests.
type MsgInt64Custom struct {
	Value custom.Int64
}

// MsgInt32 is used in tests.
type MsgInt32 struct {
	Value int32
}

// MsgInt32Custom is used in tests.
type MsgInt32Custom struct {
	Value custom.Int32
}

// MsgInt16 is used in tests.
type MsgInt16 struct {
	Value int16
}

// MsgInt16Custom is used in tests.
type MsgInt16Custom struct {
	Value custom.Int16
}

// MsgInt8 is used in tests.
type MsgInt8 struct {
	Value int8
}

// MsgInt8Custom is used in tests.
type MsgInt8Custom struct {
	Value custom.Int8
}

// MsgBool3 is used in tests.
type MsgBool3 struct {
	Value1 bool
	Value2 bool
	Value3 bool
}

// MsgBool10 is used in tests.
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

// MsgBoolCustom is used in tests.
type MsgBoolCustom struct {
	Value custom.Bool
}

// MsgFloat64 is used in tests.
type MsgFloat64 struct {
	Value float64
}

// MsgFloat64Custom is used in tests.
type MsgFloat64Custom struct {
	Value custom.Float64
}

// MsgFloat32 is used in tests.
type MsgFloat32 struct {
	Value float32
}

// MsgFloat32Custom is used in tests.
type MsgFloat32Custom struct {
	Value custom.Float32
}

// MsgString is used in tests.
type MsgString struct {
	Value string
}

// MsgStringCustom is used in tests.
type MsgStringCustom struct {
	Value custom.String
}

// MsgArray is used in tests.
type MsgArray struct {
	Value [3]bool
}

// MsgArrayCustom is used in tests.
type MsgArrayCustom struct {
	Value custom.Array
}

// MsgSlice is used in tests.
type MsgSlice struct {
	Value []bool
}

// MsgSliceCustom is used in tests.
type MsgSliceCustom struct {
	Value custom.Slice
}

// MsgMap is used in tests.
type MsgMap struct {
	Value map[uint64]uint64
}

// MsgMapString is used in tests.
type MsgMapString struct {
	Value map[string]string
}

// MsgMapCustom is used in tests.
type MsgMapCustom struct {
	Value custom.Map
}

// MsgStruct is used in tests.
type MsgStruct struct {
	Value1 SubMsg
	Value2 pkg2.SubMsg
}

// MsgStructAnonymous is used in tests.
type MsgStructAnonymous struct {
	SubMsg
	Value2 pkg2.SubMsg
}

// MsgArrayUint8 is used in tests.
type MsgArrayUint8 struct {
	Value [5]uint8
}

// MsgArrayUint8Custom is used in tests.
type MsgArrayUint8Custom struct {
	Value custom.ArrayUint8
}

// MsgArrayUint8Custom2 is used in tests.
type MsgArrayUint8Custom2 struct {
	Value custom.ArrayCustomUint8
}

// MsgArrayInt8 is used in tests.
type MsgArrayInt8 struct {
	Value [5]int8
}

// MsgArrayInt8Custom is used in tests.
type MsgArrayInt8Custom struct {
	Value custom.ArrayInt8
}

// MsgArrayFloat32 is used in tests.
type MsgArrayFloat32 struct {
	Value [5]float32
}

// MsgArrayFloat32Custom is used in tests.
type MsgArrayFloat32Custom struct {
	Value custom.ArrayFloat32
}

// MsgArrayFloat64 is used in tests.
type MsgArrayFloat64 struct {
	Value [5]float64
}

// MsgArrayFloat64Custom is used in tests.
type MsgArrayFloat64Custom struct {
	Value custom.ArrayFloat64
}

// MsgSliceUint8 is used in tests.
type MsgSliceUint8 struct {
	Value []uint8
}

// MsgSliceUint8Custom is used in tests.
type MsgSliceUint8Custom struct {
	Value custom.SliceUint8
}

// MsgSliceUint8Custom2 is used in tests.
type MsgSliceUint8Custom2 struct {
	Value custom.SliceCustomUint8
}

// MsgSliceInt8 is used in tests.
type MsgSliceInt8 struct {
	Value []int8
}

// MsgSliceInt8Custom is used in tests.
type MsgSliceInt8Custom struct {
	Value custom.SliceInt8
}

// MsgSliceFloat32 is used in tests.
type MsgSliceFloat32 struct {
	Value []float32
}

// MsgSliceFloat32Custom is used in tests.
type MsgSliceFloat32Custom struct {
	Value custom.SliceFloat32
}

// MsgSliceFloat64 is used in tests.
type MsgSliceFloat64 struct {
	Value []float64
}

// MsgSliceFloat64Custom is used in tests.
type MsgSliceFloat64Custom struct {
	Value custom.SliceFloat64
}

// MsgTime is used in tests.
type MsgTime struct {
	Value time.Time
}

// MsgMixed is used in tests.
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

// MsgMixedCustom is used in tests.
type MsgMixedCustom struct {
	Value custom.Mixed
}

// SubMsg is used in tests.
type SubMsg struct {
	Value1 []float32
	Value2 []spkg1.SubMsg
	Value3 []spkg2.SubMsg
}

// MsgIgnore is used in tests.
type MsgIgnore struct {
	Value1        bool
	Value2Ignored bool
	Value3        string
	Value4Ignored string
}

// MsgNotIgnore is used in tests.
type MsgNotIgnore struct {
	SubMsg MsgIgnore
}

// MsgStrings is used in tests.
type MsgStrings struct {
	Value1  string
	Value2  string
	Value3  string
	Value4  string
	Value5  string
	Value6  string
	Value7  string
	Value8  string
	Value9  string
	Value10 string
}
