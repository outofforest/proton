package tfloat32

import (
	"fmt"
	"reflect"

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
	return 4
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	t := b.tm.TypeName(b.msgType, b.fieldType)
	unsafe := b.tm.Import("unsafe")
	return fmt.Sprintf(`*(*%[1]s)(%[2]s.Pointer(&b[o])) = {{ . }}
o += 4`, t, unsafe)
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	t := b.tm.TypeName(b.msgType, b.fieldType)
	unsafe := b.tm.Import("unsafe")
	return fmt.Sprintf(`{{ . }} = *(*%[1]s)(%[2]s.Pointer(&b[o]))
o += 4`, t, unsafe)
}
