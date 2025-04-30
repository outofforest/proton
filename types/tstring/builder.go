package tstring

import (
	"fmt"
	"reflect"

	"github.com/outofforest/proton/types"
)

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
	return 1 // size always takes at least one byte
}

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(_ *uint64) (string, bool) {
	return `{
	l := uint64(len({{ . }}))
	helpers.UInt64Size(l, &n)
	n += l
}`, true
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	return `{
	l := uint64(len({{ . }}))
	helpers.UInt64Marshal(l, b, &o)
	copy(b[o:o+l], {{ . }})
	o += l
}`
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	code := `{
	var l uint64
	helpers.UInt64Unmarshal(&l, b, &o)
`

	code += fmt.Sprintf(`	if l > 0 {
		{{ . }} = %[1]s(b[o:o+l])
		o += l
	}
}`, b.tm.TypeName(b.fieldType))

	return code
}
