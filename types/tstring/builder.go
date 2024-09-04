package tstring

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

// Allocators returns the list of types for which massive allocators are needed.
func (b Builder) Allocators() []reflect.Type {
	return nil
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1 // size always takes at least one byte
}

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(_ *uint64) (string, bool) {
	code := `{
	l := uint64(len({{ . }}))
	n += l
`

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64SizeCode(), "l")
	code += types.AddIndent(buf.String(), 1)

	code += `
}`
	return code, true
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	code := `{
	l := uint64(len({{ . }}))
`

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64Marshal(), "l")
	code += types.AddIndent(buf.String(), 1) + "\n"

	code += `	copy(b[o:o+l], {{ . }})
	o += l
}`

	return code
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	unsafe := b.tm.Import("unsafe")
	code := `{
	var l uint64
`

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64Unmarshal("uint64"), "l")
	code += types.AddIndent(buf.String(), 1) + "\n"

	code += `	if l > 0 {
`
	if b.fieldType.Name() == "string" {
		code += fmt.Sprintf("		{{ . }} = %[1]s.String((*byte)(%[1]s.Pointer(&b[o])), l)\n", unsafe)
	} else {
		code += fmt.Sprintf("		{{ . }} = %[2]s(%[1]s.String((*byte)(%[1]s.Pointer(&b[o])), l))\n", unsafe,
			b.tm.TypeName(b.msgType, b.fieldType))
	}
	code += `		o += l
	} else {
		{{ . }} = "" 
	}
}`

	return code
}
