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

// Allocators returns the list of types for which massive allocators are needed.
func (b Builder) Allocators() []reflect.Type {
	return nil
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	return `if {{ . }} {
	b[o] = 0x01
} else {
	b[o] = 0x00
}
o++`
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	return `{{ . }} = b[o] != 0x00
o++`
}
