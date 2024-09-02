package tslicefloat64

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/types"
)

// New returns new code builder.
func New(msgType, fieldType reflect.Type, tm types.TypeMap) Builder {
	return Builder{
		msgType:   msgType,
		fieldType: fieldType,
		tm:        tm,
	}
}

// Builder generates the code.
type Builder struct {
	msgType   reflect.Type
	fieldType reflect.Type
	tm        types.TypeMap
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
	code := "l := uint64(len({{ . }}))\n"

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64SizeCode(), "l")
	code += buf.String() + "\n"

	code += "n += l * 8"

	return code, true
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	code := "l := uint64(len({{ . }}))\n"

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64Marshal(), "l")
	code += buf.String() + "\n"

	unsafe := b.tm.Import("unsafe")
	code += fmt.Sprintf(`if l > 0 {
	copy(b[o:o+l*8], %[1]s.Slice((*byte)(%[1]s.Pointer(&{{ . }}[0])), l*8))
	o += l * 8
}`, unsafe)

	return code
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	code := "var l uint64\n"

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64Unmarshal("uint64"), "l")
	code += buf.String() + "\n"

	unsafe := b.tm.Import("unsafe")
	code += fmt.Sprintf(`if l > 0 {
	{{ . }} = make(%[2]s, l)
	copy(%[1]s.Slice((*byte)(%[1]s.Pointer(&{{ . }}[0])), l*8), b[o:o+l*8])
	o += l * 8
} else {
	{{ . }} = nil
}`, unsafe, b.tm.TypeName(b.msgType, b.fieldType))

	return code
}
