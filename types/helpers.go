package types

import (
	"fmt"
	"strings"
)

// Cases when varint takes one byte are ignored because they should be handled in more optimal way by
// `ConstantSize` method of the builders.

// UInt16SizeCode generates code to get size of varint for uint16.
func UInt16SizeCode() string {
	return `{
	vi := {{ . }}
	switch {
	case vi <= 0x7F:
	case vi <= 0x3FFF:
		n++
	default:
		n += 2
	}
}`
}

// UInt16Marshal generates code to marshal varint for uint16.
func UInt16Marshal() string {
	return `{
	vi := {{ . }}
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
}`
}

// UInt16Unmarshal generates code to unmarshal varint for uint16.
func UInt16Unmarshal(typeName string) string {
	return fmt.Sprintf(`{
	vi := %[1]s(b[o] & 0x7F)
	if b[o]&0x80 == 0 {
		o++
	} else {
		vi |= %[1]s(b[o+1]&0x7F) << 7
		if b[o+1]&0x80 == 0 {
			o += 2
		} else {
			vi |= %[1]s(b[o+2]) << 14
			o += 3
		}
	}
	{{ . }} = vi
}`, typeName)
}

// UInt32SizeCode generates code to get size of varint for uint32.
func UInt32SizeCode() string {
	return `{
	vi := {{ . }}
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
}`
}

// UInt32Marshal generates code to marshal varint for uint32.
func UInt32Marshal() string {
	return `{
	vi := {{ . }}
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
}`
}

// UInt32Unmarshal generates code to unmarshal varint for uint32.
func UInt32Unmarshal(typeName string) string {
	return fmt.Sprintf(`{
	vi := %[1]s(b[o] & 0x7F)
	if b[o]&0x80 == 0 {
		o++
	} else {
		vi |= %[1]s(b[o+1]&0x7F) << 7
		if b[o+1]&0x80 == 0 {
			o += 2
		} else {
			vi |= %[1]s(b[o+2]&0x7F) << 14
			if b[o+2]&0x80 == 0 {
				o += 3
			} else {
				vi |= %[1]s(b[o+3]&0x7F) << 21
				if b[o+3]&0x80 == 0 {
					o += 4
				} else {
					vi |= %[1]s(b[o+4]) << 28
					o += 5
				}
			}
		}
	}
	{{ . }} = vi
}`, typeName)
}

// UInt64SizeCode generates code to get size of varint for uint64.
func UInt64SizeCode() string {
	// For last byte last bit is used for data, not for continuation flag.
	// That's why we may fit 64-bit number in 9 bytes, not 10.

	return `{
	vi := {{ . }}
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
}`
}

// UInt64Marshal generates code to marshal varint for uint64.
func UInt64Marshal() string {
	return `{
	vi := {{ . }}
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
}`
}

// UInt64Unmarshal generates code to unmarshal varint for uint64.
func UInt64Unmarshal(typeName string) string {
	return fmt.Sprintf(`{
	vi := %[1]s(b[o] & 0x7F)
	if b[o]&0x80 == 0 {
		o++
	} else {
		vi |= %[1]s(b[o+1]&0x7F) << 7
		if b[o+1]&0x80 == 0 {
			o += 2
		} else {
			vi |= %[1]s(b[o+2]&0x7F) << 14
			if b[o+2]&0x80 == 0 {
				o += 3
			} else {
				vi |= %[1]s(b[o+3]&0x7F) << 21
				if b[o+3]&0x80 == 0 {
					o += 4
				} else {
					vi |= %[1]s(b[o+4]&0x7F) << 28
					if b[o+4]&0x80 == 0 {
						o += 5
					} else {
						vi |= %[1]s(b[o+5]&0x7F) << 35
						if b[o+5]&0x80 == 0 {
							o += 6
						} else {
							vi |= %[1]s(b[o+6]&0x7F) << 42
							if b[o+6]&0x80 == 0 {
								o += 7
							} else {
								vi |= %[1]s(b[o+7]&0x7F) << 49
								if b[o+7]&0x80 == 0 {
									o += 8
								} else {
									vi |= %[1]s(b[o+8]) << 56
									o += 9
								}
							}
						}
					}
				}
			}
		}
	}
	{{ . }} = vi
}`, typeName)
}

// Int16SizeCode generates code to get size of varint for int16.
func Int16SizeCode() string {
	return `{
	vi := uint16({{ . }}) << 1
	if {{ . }} < 0 {
		vi ^= 0xFFFF
	}
	switch {
	case vi <= 0x7F:
	case vi <= 0x3FFF:
		n++
	default:
		n += 2
	}
}`
}

// Int16Marshal generates code to marshal varint for int16.
func Int16Marshal() string {
	return `{
	vi := uint16({{ . }}) << 1
	if {{ . }} < 0 {
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
}`
}

// Int16Unmarshal generates code to unmarshal varint for int16.
func Int16Unmarshal(typeName string) string {
	return fmt.Sprintf(`{
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
	{{ . }} = %[1]s(vi)
}`, typeName)
}

// Int32SizeCode generates code to get size of varint for int32.
func Int32SizeCode() string {
	return `{
	vi := uint32({{ . }}) << 1
	if {{ . }} < 0 {
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
}`
}

// Int32Marshal generates code to marshal varint for int32.
func Int32Marshal() string {
	return `{
	vi := uint32({{ . }}) << 1
	if {{ . }} < 0 {
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
}`
}

// Int32Unmarshal generates code to unmarshal varint for int32.
func Int32Unmarshal(typeName string) string {
	return fmt.Sprintf(`{
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
	{{ . }} = %[1]s(vi)
}`, typeName)
}

// Int64SizeCode generates code to get size of varint for int64.
func Int64SizeCode() string {
	return `{
	vi := uint64({{ . }}) << 1
	if {{ . }} < 0 {
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
}`
}

// Int64Marshal generates code to marshal varint for int64.
func Int64Marshal() string {
	return `{
	vi := uint64({{ . }}) << 1
	if {{ . }} < 0 {
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
}`
}

// Int64Unmarshal generates code to unmarshal varint for int64.
func Int64Unmarshal(typeName string) string {
	return fmt.Sprintf(`{
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
	{{ . }} = %[1]s(vi)
}`, typeName)
}

// AddIndent adds indentation to the code.
func AddIndent(code string, numOfIndentations int) string {
	var indent string
	for range numOfIndentations {
		indent += "\t"
	}

	return indent + strings.ReplaceAll(code, "\n", "\n"+indent)
}

// Align returns spaces needed to align strings.
func Align(v string, l int) string {
	if len(v) >= l {
		return ""
	}
	return strings.Repeat(" ", l-len(v))
}

// Var generates random variable name with provided prefix.
func Var(prefix string, varIndex *uint64) string {
	*varIndex++
	return fmt.Sprintf("%s%d", prefix, *varIndex)
}
