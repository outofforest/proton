package custom

// Uint64 is used in tests.
type Uint64 uint64

// Uint32 is used in tests.
type Uint32 uint32

// Uint16 is used in tests.
type Uint16 uint16

// Uint8 is used in tests.
type Uint8 uint8

// Int64 is used in tests.
type Int64 int64

// Int32 is used in tests.
type Int32 int32

// Int16 is used in tests.
type Int16 int16

// Int8 is used in tests.
type Int8 int8

// Bool is used in tests.
type Bool bool

// Float64 is used in tests.
type Float64 float64

// Float32 is used in tests.
type Float32 float32

// String is used in tests.
type String string

// Array is used in tests.
type Array [2]Uint64

// Slice is used in tests.
type Slice []Uint64

// Map is used in tests.
type Map map[String]Int32

// ArrayUint8 is used in tests.
type ArrayUint8 [5]uint8

// ArrayCustomUint8 is used in tests.
type ArrayCustomUint8 [5]Uint8

// ArrayInt8 is used in tests.
type ArrayInt8 [5]Int8

// ArrayFloat32 is used in tests.
type ArrayFloat32 [5]Float32

// ArrayFloat64 is used in tests.
type ArrayFloat64 [5]Float64

// SliceUint8 is used in tests.
type SliceUint8 []uint8

// SliceCustomUint8 is used in tests.
type SliceCustomUint8 []Uint8

// SliceInt8 is used in tests.
type SliceInt8 []Int8

// SliceFloat32 is used in tests.
type SliceFloat32 []Float32

// SliceFloat64 is used in tests.
type SliceFloat64 []Float64

// Mixed is used in tests.
type Mixed [][3][]map[int16]Array
