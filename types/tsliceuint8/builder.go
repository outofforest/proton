package tsliceuint8

import (
	"fmt"
	"reflect"

	"github.com/outofforest/proton/types"
)

const uint8Name = "uint8"

// New returns new code builder.
func New(fieldType reflect.Type, tm *types.TypeMap) Builder {
	return Builder{
		fieldType: fieldType,
		tm:        tm,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType reflect.Type
	tm        *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return nil
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1 // covers the first byte of length
}

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(_ *uint64) (string, bool) {
	return `l := uint64(len({{ . }}))
helpers.UInt64Size(l, &n)
n += l`, true
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	code := `l := uint64(len({{ . }}))
helpers.UInt64Marshal(l, b, &o)
`

	unsafe := b.tm.Import("unsafe")
	code += "if l > 0 {\n	"

	if b.fieldType.Elem().Name() == uint8Name {
		code += fmt.Sprintf(`copy(b[o:o+l], %[1]s.Slice(&{{ . }}[0], l))
	o += l`, unsafe)
	} else {
		code += fmt.Sprintf(`copy(b[o:o+l], %[1]s.Slice((*byte)(&{{ . }}[0]), l))
	o += l`, unsafe)
	}

	code += "\n}"

	return code
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	code := `var l uint64
helpers.UInt64Unmarshal(&l, b, &o)
`

	code += fmt.Sprintf(`if l > 0 {
	{{ . }} = make([]%[1]s, l)
`, b.tm.TypeName(b.fieldType.Elem()))
	if b.fieldType.Elem().Name() == uint8Name {
		code += `	copy({{ . }}, b[o:o+l])`
	} else {
		unsafe := b.tm.Import("unsafe")
		code += fmt.Sprintf(`	copy(%[1]s.Slice((*byte)(&{{ . }}[0]), l), b[o:o+l])`, unsafe)
	}

	code += `
	o += l
}`

	return code
}
