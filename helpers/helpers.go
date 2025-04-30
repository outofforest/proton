package helpers

import (
	"io"
	"reflect"
	"text/template"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// ForEachField iterates over fields in the type and calls a function for each public one.
func ForEachField(msgType reflect.Type, fn func(field reflect.StructField) error) error {
	numOfFields := msgType.NumField()
	for i := range numOfFields {
		field := msgType.Field(i)

		if !field.IsExported() {
			continue
		}

		if err := fn(field); err != nil {
			return err
		}
	}
	return nil
}

// Execute executes a template using provided data.
func Execute(b io.Writer, code string, data any) {
	lo.Must0(template.Must(template.New("").Parse(code)).Execute(b, data))
}

// RecoverMarshal recovers from panic inside marshaller.
func RecoverMarshal(err *error) {
	if res := recover(); res != nil {
		*err = errors.Errorf("marshaling message failed: %s", res)
	}
}

// RecoverUnmarshal recovers from panic inside unmarshaller.
func RecoverUnmarshal(err *error) {
	if res := recover(); res != nil {
		*err = errors.Errorf("unmarshaling message failed: %s", res)
	}
}

// Cases when varint takes one byte are ignored because they should be handled in more optimal way by
// `ConstantSize` method of the builders.

// UInt16Size computes size of varint for uint16.
func UInt16Size[T ~uint16](v T, size *uint64) {
	switch {
	case v <= 0x7F:
	case v <= 0x3FFF:
		*size++
	default:
		*size += 2
	}
}

// UInt16Marshal marshals varint for uint16.
func UInt16Marshal[T ~uint16](v T, b []byte, o *uint64) {
	switch {
	case v <= 0x7F:
		b[*o] = byte(v)
		*o++
	case v <= 0x3FFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	default:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	}
}

// UInt16Unmarshal unmarshals varint for uint16.
func UInt16Unmarshal[T ~uint16](v *T, b []byte, o *uint64) {
	*v = T(b[*o] & 0x7F)
	if b[*o]&0x80 == 0 {
		*o++
	} else {
		*v |= T(b[*o+1]&0x7F) << 7
		if b[*o+1]&0x80 == 0 {
			*o += 2
		} else {
			*v |= T(b[*o+2]) << 14
			*o += 3
		}
	}
}

// UInt32Size computes size of varint for uint32.
func UInt32Size[T ~uint32](v T, size *uint64) {
	switch {
	case v <= 0x7F:
	case v <= 0x3FFF:
		*size++
	case v <= 0x1FFFFF:
		*size += 2
	case v <= 0xFFFFFFF:
		*size += 3
	default:
		*size += 4
	}
}

// UInt32Marshal marshals varint for uint32.
func UInt32Marshal[T ~uint32](v T, b []byte, o *uint64) {
	//nolint:dupl
	switch {
	case v <= 0x7F:
		b[*o] = byte(v)
		*o++
	case v <= 0x3FFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	case v <= 0x1FFFFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	case v <= 0xFFFFFFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	default:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	}
}

// UInt32Unmarshal unmarshals varint for uint32.
func UInt32Unmarshal[T ~uint32](v *T, b []byte, o *uint64) {
	*v = T(b[*o] & 0x7F)
	//nolint:nestif
	if b[*o]&0x80 == 0 {
		*o++
	} else {
		*v |= T(b[*o+1]&0x7F) << 7
		if b[*o+1]&0x80 == 0 {
			*o += 2
		} else {
			*v |= T(b[*o+2]&0x7F) << 14
			if b[*o+2]&0x80 == 0 {
				*o += 3
			} else {
				*v |= T(b[*o+3]&0x7F) << 21
				if b[*o+3]&0x80 == 0 {
					*o += 4
				} else {
					*v |= T(b[*o+4]) << 28
					*o += 5
				}
			}
		}
	}
}

// UInt64Size computes size of varint for uint64.
func UInt64Size[T ~uint64](v T, size *uint64) {
	// For last byte last bit is used for data, not for continuation flag.
	// That's why we may fit 64-bit number in 9 bytes, not 10.

	switch {
	case v <= 0x7F:
	case v <= 0x3FFF:
		*size++
	case v <= 0x1FFFFF:
		*size += 2
	case v <= 0xFFFFFFF:
		*size += 3
	case v <= 0x7FFFFFFFF:
		*size += 4
	case v <= 0x3FFFFFFFFFF:
		*size += 5
	case v <= 0x1FFFFFFFFFFFF:
		*size += 6
	case v <= 0xFFFFFFFFFFFFFF:
		*size += 7
	default:
		*size += 8
	}
}

// UInt64Marshal marshals varint for uint64.
func UInt64Marshal[T ~uint64](v T, b []byte, o *uint64) {
	switch {
	case v <= 0x7F:
		b[*o] = byte(v)
		*o++
	case v <= 0x3FFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	case v <= 0x1FFFFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	case v <= 0xFFFFFFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	case v <= 0x7FFFFFFFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	case v <= 0x3FFFFFFFFFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	case v <= 0x1FFFFFFFFFFFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	case v <= 0xFFFFFFFFFFFFFF:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	default:
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v) | 0x80
		*o++
		v >>= 7
		b[*o] = byte(v)
		*o++
	}
}

// UInt64Unmarshal unmarshals varint for uint64.
func UInt64Unmarshal[T ~uint64](v *T, b []byte, o *uint64) {
	*v = T(b[*o] & 0x7F)
	//nolint:nestif
	if b[*o]&0x80 == 0 {
		*o++
	} else {
		*v |= T(b[*o+1]&0x7F) << 7
		if b[*o+1]&0x80 == 0 {
			*o += 2
		} else {
			*v |= T(b[*o+2]&0x7F) << 14
			if b[*o+2]&0x80 == 0 {
				*o += 3
			} else {
				*v |= T(b[*o+3]&0x7F) << 21
				if b[*o+3]&0x80 == 0 {
					*o += 4
				} else {
					*v |= T(b[*o+4]&0x7F) << 28
					if b[*o+4]&0x80 == 0 {
						*o += 5
					} else {
						*v |= T(b[*o+5]&0x7F) << 35
						if b[*o+5]&0x80 == 0 {
							*o += 6
						} else {
							*v |= T(b[*o+6]&0x7F) << 42
							if b[*o+6]&0x80 == 0 {
								*o += 7
							} else {
								*v |= T(b[*o+7]&0x7F) << 49
								if b[*o+7]&0x80 == 0 {
									*o += 8
								} else {
									*v |= T(b[*o+8]) << 56
									*o += 9
								}
							}
						}
					}
				}
			}
		}
	}
}

// Int16Size computes size of varint for int16.
func Int16Size[T ~int16](v T, size *uint64) {
	vi := uint16(v) << 1
	if v < 0 {
		vi ^= 0xFFFF
	}
	switch {
	case vi <= 0x7F:
	case vi <= 0x3FFF:
		*size++
	default:
		*size += 2
	}
}

// Int16Marshal marshals varint for int16.
func Int16Marshal[T ~int16](v T, b []byte, o *uint64) {
	vi := uint16(v) << 1
	if v < 0 {
		vi ^= 0xFFFF
	}
	switch {
	case vi <= 0x7F:
		b[*o] = byte(vi)
		*o++
	case vi <= 0x3FFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	default:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	}
}

// Int16Unmarshal unmarshals varint for int16.
func Int16Unmarshal[T ~int16](v *T, b []byte, o *uint64) {
	vi := uint16(b[*o] & 0x7F)
	if b[*o]&0x80 == 0 {
		*o++
	} else {
		vi |= uint16(b[*o+1]&0x7F) << 7
		if b[*o+1]&0x80 == 0 {
			*o += 2
		} else {
			vi |= uint16(b[*o+2]) << 14
			*o += 3
		}
	}
	if vi&0x01 == 0 {
		vi >>= 1
	} else {
		vi >>= 1
		vi ^= 0xFFFF
	}
	*v = T(vi)
}

// Int32Size computes size of varint for int32.
func Int32Size[T ~int32](v T, size *uint64) {
	vi := uint32(v) << 1
	if v < 0 {
		vi ^= 0xFFFFFFFF
	}
	switch {
	case vi <= 0x7F:
	case vi <= 0x3FFF:
		*size++
	case vi <= 0x1FFFFF:
		*size += 2
	case vi <= 0xFFFFFFF:
		*size += 3
	default:
		*size += 4
	}
}

// Int32Marshal marshals varint for int32.
func Int32Marshal[T ~int32](v T, b []byte, o *uint64) {
	vi := uint32(v) << 1
	if v < 0 {
		vi ^= 0xFFFFFFFF
	}
	//nolint:dupl
	switch {
	case vi <= 0x7F:
		b[*o] = byte(vi)
		*o++
	case vi <= 0x3FFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0x1FFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0xFFFFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	default:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	}
}

// Int32Unmarshal unmarshals varint for int32.
func Int32Unmarshal[T ~int32](v *T, b []byte, o *uint64) {
	vi := uint32(b[*o] & 0x7F)
	//nolint:nestif
	if b[*o]&0x80 == 0 {
		*o++
	} else {
		vi |= uint32(b[*o+1]&0x7F) << 7
		if b[*o+1]&0x80 == 0 {
			*o += 2
		} else {
			vi |= uint32(b[*o+2]&0x7F) << 14
			if b[*o+2]&0x80 == 0 {
				*o += 3
			} else {
				vi |= uint32(b[*o+3]&0x7F) << 21
				if b[*o+3]&0x80 == 0 {
					*o += 4
				} else {
					vi |= uint32(b[*o+4]) << 28
					*o += 5
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
	*v = T(vi)
}

// Int64Size computes size of varint for int64.
func Int64Size[T ~int64](v T, size *uint64) {
	vi := uint64(v) << 1
	if v < 0 {
		vi ^= 0xFFFFFFFFFFFFFFFF
	}
	switch {
	case vi <= 0x7F:
	case vi <= 0x3FFF:
		*size++
	case vi <= 0x1FFFFF:
		*size += 2
	case vi <= 0xFFFFFFF:
		*size += 3
	case vi <= 0x7FFFFFFFF:
		*size += 4
	case vi <= 0x3FFFFFFFFFF:
		*size += 5
	case vi <= 0x1FFFFFFFFFFFF:
		*size += 6
	case vi <= 0xFFFFFFFFFFFFFF:
		*size += 7
	case vi <= 0x7FFFFFFFFFFFFFFF:
		*size += 8
	default:
		*size += 9
	}
}

// Int64Marshal marshals varint for int64.
func Int64Marshal[T ~int64](v T, b []byte, o *uint64) {
	vi := uint64(v) << 1
	if v < 0 {
		vi ^= 0xFFFFFFFFFFFFFFFF
	}
	switch {
	case vi <= 0x7F:
		b[*o] = byte(vi)
		*o++
	case vi <= 0x3FFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0x1FFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0xFFFFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0x7FFFFFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0x3FFFFFFFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0x1FFFFFFFFFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0xFFFFFFFFFFFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	case vi <= 0x7FFFFFFFFFFFFFFF:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	default:
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi) | 0x80
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
		vi >>= 7
		b[*o] = byte(vi)
		*o++
	}
}

// Int64Unmarshal unmarshals varint for int64.
func Int64Unmarshal[T ~int64](v *T, b []byte, o *uint64) {
	vi := uint64(b[*o] & 0x7F)
	//nolint:nestif
	if b[*o]&0x80 == 0 {
		*o++
	} else {
		vi |= uint64(b[*o+1]&0x7F) << 7
		if b[*o+1]&0x80 == 0 {
			*o += 2
		} else {
			vi |= uint64(b[*o+2]&0x7F) << 14
			if b[*o+2]&0x80 == 0 {
				*o += 3
			} else {
				vi |= uint64(b[*o+3]&0x7F) << 21
				if b[*o+3]&0x80 == 0 {
					*o += 4
				} else {
					vi |= uint64(b[*o+4]&0x7F) << 28
					if b[*o+4]&0x80 == 0 {
						*o += 5
					} else {
						vi |= uint64(b[*o+5]&0x7F) << 35
						if b[*o+5]&0x80 == 0 {
							*o += 6
						} else {
							vi |= uint64(b[*o+6]&0x7F) << 42
							if b[*o+6]&0x80 == 0 {
								*o += 7
							} else {
								vi |= uint64(b[*o+7]&0x7F) << 49
								if b[*o+7]&0x80 == 0 {
									*o += 8
								} else {
									vi |= uint64(b[*o+8]&0x7F) << 56
									if b[*o+8]&0x80 == 0 {
										*o += 9
									} else {
										vi |= uint64(b[*o+9]) << 63
										*o += 10
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
	*v = T(vi)
}
