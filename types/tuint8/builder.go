package tuint8

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
	return 1
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	t := b.tm.TypeName(b.fieldType)
	if t == "uint8" {
		return `b[o] = {{ . }}
o++`
	}
	return `b[o] = byte({{ . }})
o++`
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	t := b.tm.TypeName(b.fieldType)
	if t == "uint8" {
		return `{{ . }} = b[o]
o++`
	}
	return fmt.Sprintf(`{{ . }} = %[1]s(b[o])
o++`, t)
}
