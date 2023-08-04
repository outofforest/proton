package test

// CustomUint64 is used in tests
type CustomUint64 uint64

// CustomUint32 is used in tests
type CustomUint32 uint32

// CustomUint16 is used in tests
type CustomUint16 uint16

// CustomUint8 is used in tests
type CustomUint8 uint8

// CustomInt64 is used in tests
type CustomInt64 int64

// CustomInt32 is used in tests
type CustomInt32 int32

// CustomInt16 is used in tests
type CustomInt16 int16

// CustomInt8 is used in tests
type CustomInt8 int8

// CustomBool is used in tests
type CustomBool bool

// CustomFloat64 is used in tests
type CustomFloat64 float64

// CustomFloat32 is used in tests
type CustomFloat32 float32

// CustomString is used in tests
type CustomString string

// CustomArray is used in tests
type CustomArray [2]CustomUint64

// CustomSlice is used in tests
type CustomSlice []CustomUint64

// CustomMap is used in tests
type CustomMap map[CustomString]CustomInt32

// CustomArrayUint8 is used in tests
type CustomArrayUint8 [5]uint8

// CustomArrayCustomUint8 is used in tests
type CustomArrayCustomUint8 [5]CustomUint8

// CustomArrayInt8 is used in tests
type CustomArrayInt8 [5]CustomInt8

// CustomArrayFloat32 is used in tests
type CustomArrayFloat32 [5]CustomFloat32

// CustomArrayFloat64 is used in tests
type CustomArrayFloat64 [5]CustomFloat64

// CustomSliceUint8 is used in tests
type CustomSliceUint8 []uint8

// CustomSliceCustomUint8 is used in tests
type CustomSliceCustomUint8 []CustomUint8

// CustomSliceInt8 is used in tests
type CustomSliceInt8 []CustomInt8

// CustomSliceFloat32 is used in tests
type CustomSliceFloat32 []CustomFloat32

// CustomSliceFloat64 is used in tests
type CustomSliceFloat64 []CustomFloat64

// CustomMixed is used in tests
type CustomMixed [][3][]map[int16]CustomArray
