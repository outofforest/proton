package tstruct

import (
	"reflect"
)

// New returns new code builder.
func New(fieldType reflect.Type) Builder {
	return Builder{
		fieldType: fieldType,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType reflect.Type
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return []reflect.Type{b.fieldType}
}

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(_ *uint64) string {
	return "n += {{ . }}.Size()"
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	return "o += {{ . }}.Marshal(b[o:])"
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	return "o += {{ . }}.Unmarshal(b[o:])"
}
