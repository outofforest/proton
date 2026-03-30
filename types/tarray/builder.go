package tarray

import (
	"bytes"
	_ "embed"
	"fmt"
	"reflect"
	"text/template/parse"

	"github.com/samber/lo"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/types"
)

var (
	//go:embed templates.gotmpl
	tmpl string
	t    = lo.Must(parse.Parse("", tmpl, "{{", "}}"))
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

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(varIndex *uint64) *parse.Tree {
	return t["size"]
	elementTpl, elementOK := b.elementBuilder.SizeCode(varIndex)
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

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(varIndex *uint64) *parse.Tree {
	return t["marshal"]
	elementTpl := b.elementBuilder.MarshalCode(varIndex)

	av := types.Var("av", varIndex)
	code := fmt.Sprintf("for _, %s := range {{ . }} {\n", av)

	buf := &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, av)
	code += types.AddIndent(buf.String(), 1)

	code += "\n}"

	return code
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(varIndex *uint64) *parse.Tree {
	return t["unmarshal"]
	elementTpl := b.elementBuilder.UnmarshalCode(varIndex)

	i := types.Var("i", varIndex)
	code := fmt.Sprintf("for %[1]s := range %[2]d {\n", i, b.fieldType.Len())

	buf := &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, fmt.Sprintf("{{ . }}[%s]", i))
	code += types.AddIndent(buf.String(), 1)

	code += "\n}"

	return code
}
