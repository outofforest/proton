package tstruct

import (
	"fmt"
	"reflect"

	"github.com/outofforest/proton/types"
)

// New returns new code builder.
func New(fieldType reflect.Type, fieldBuilders []types.BuilderFactory, tm *types.TypeMap) Builder {
	return Builder{
		fieldType:     fieldType,
		fieldBuilders: fieldBuilders,
		tm:            tm,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType     reflect.Type
	fieldBuilders []types.BuilderFactory
	tm            *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return []reflect.Type{b.fieldType}
}

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(_ *uint64) string {
	return fmt.Sprintf("n += %[1]s(&{{ . }})", b.tm.VarName(b.fieldType, "size"))
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	return fmt.Sprintf("o += %[1]s(&{{ . }}, b[o:])", b.tm.VarName(b.fieldType, "marshal"))
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	return fmt.Sprintf(`o += %[1]s(&{{ . }}, b[o:])
`, b.tm.VarName(b.fieldType, "unmarshal"))
}
