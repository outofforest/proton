package tbool

import "reflect"

// New returns new code builder.
func New() Builder {
	return Builder{}
}

// Builder generates the code.
type Builder struct{}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return nil
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) string {
	return `if {{ . }} {
	b[o] = 0x01
} else {
	b[o] = 0x00
}
o++`
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) string {
	return `{{ . }} = b[o] != 0x00
o++`
}
