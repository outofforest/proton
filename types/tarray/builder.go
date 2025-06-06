package tarray

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/types"
)

// New returns new code builder.
func New(fieldType reflect.Type, elementBuilder types.BuilderFactory) Builder {
	return Builder{
		fieldType:      fieldType,
		elementBuilder: elementBuilder,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType      reflect.Type
	elementBuilder types.BuilderFactory
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return b.elementBuilder.Dependencies()
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return uint64(b.fieldType.Len()) * b.elementBuilder.ConstantSize()
}

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(varIndex *uint64) (string, bool) {
	elementTpl, elementOK := b.elementBuilder.SizeCodeTemplate(varIndex)
	if !elementOK {
		return "", false
	}

	av := types.Var("av", varIndex)
	code := fmt.Sprintf("for _, %s := range {{ . }} {\n", av)

	buf := &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, av)

	code += types.AddIndent(buf.String(), 1) + "\n}"

	return code, true
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(varIndex *uint64) string {
	elementTpl := b.elementBuilder.MarshalCodeTemplate(varIndex)

	av := types.Var("av", varIndex)
	code := fmt.Sprintf("for _, %s := range {{ . }} {\n", av)

	buf := &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, av)
	code += types.AddIndent(buf.String(), 1)

	code += "\n}"

	return code
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(varIndex *uint64) string {
	elementTpl := b.elementBuilder.UnmarshalCodeTemplate(varIndex)

	i := types.Var("i", varIndex)
	code := fmt.Sprintf("for %[1]s := range %[2]d {\n", i, b.fieldType.Len())

	buf := &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, fmt.Sprintf("{{ . }}[%s]", i))
	code += types.AddIndent(buf.String(), 1)

	code += "\n}"

	return code
}
