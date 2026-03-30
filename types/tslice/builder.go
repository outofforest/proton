package tslice

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
func New(fieldType reflect.Type, elementBuilder types.BuilderFactory, tm *types.TypeMap) Builder {
	return Builder{
		fieldType:      fieldType,
		elementBuilder: elementBuilder,
		tm:             tm,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType      reflect.Type
	elementBuilder types.BuilderFactory
	tm             *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return b.elementBuilder.Dependencies()
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1 // covers the first byte of length
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(varIndex *uint64) (*parse.Tree, any) {
	return t["size"]
	code := `l := uint64(len({{ . }}))
helpers.UInt64Size(l, &n)
`

	constSize := b.elementBuilder.ConstantSize()
	switch {
	case constSize > 1:
		code += fmt.Sprintf("n += l * %d", constSize)
	case constSize == 1:
		code += "n += l"
	}

	elementTpl, elementOK := b.elementBuilder.SizeCode(varIndex)

	if !elementOK {
		return code, true
	}

	if constSize > 0 {
		code += "\n"
	}

	sv := types.Var("sv", varIndex)
	code += fmt.Sprintf("for _, %s := range {{ . }} {\n", sv)

	buf := &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, sv)

	code += types.AddIndent(buf.String(), 1) + "\n}"

	return code, true
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(varIndex *uint64) (*parse.Tree, any) {
	return t["marshal"]
	elementTpl := b.elementBuilder.MarshalCode(varIndex)

	code := `helpers.UInt64Marshal(uint64(len({{ . }})), b, &o)
`

	sv := types.Var("sv", varIndex)
	code += fmt.Sprintf("for _, %s := range {{ . }} {\n", sv)

	buf := &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, sv)
	code += types.AddIndent(buf.String(), 1) + "\n"

	code += "}"

	return code
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(varIndex *uint64) (*parse.Tree, any) {
	return t["unmarshal"]
	elementTpl := b.elementBuilder.UnmarshalCode(varIndex)

	code := `var l uint64
helpers.UInt64Unmarshal(&l, b, &o)
if l > 0 {
`
	code += fmt.Sprintf(`	{{ . }} = make([]%[1]s, l)
`, b.tm.TypeName(b.fieldType.Elem()))

	i := types.Var("i", varIndex)
	code += fmt.Sprintf("	for %[1]s := range l {\n", i)

	buf := &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, fmt.Sprintf("{{ . }}[%s]", i))
	code += types.AddIndent(buf.String(), 2) + "\n"

	code += `	}
}`

	return code
}
