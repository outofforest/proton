package tmap

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
func New(fieldType reflect.Type, keyBuilder, elementBuilder types.BuilderFactory, tm *types.TypeMap) Builder {
	return Builder{
		fieldType:      fieldType,
		keyBuilder:     keyBuilder,
		elementBuilder: elementBuilder,
		tm:             tm,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType      reflect.Type
	keyBuilder     types.BuilderFactory
	elementBuilder types.BuilderFactory
	tm             *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	dependencies := map[reflect.Type]bool{}
	for _, d := range b.keyBuilder.Dependencies() {
		dependencies[d] = true
	}
	for _, d := range b.elementBuilder.Dependencies() {
		dependencies[d] = true
	}

	res := make([]reflect.Type, 0, len(dependencies))
	for d := range dependencies {
		res = append(res, d)
	}
	return res
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1 // covers the first byte of length
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(varIndex *uint64) (string, bool) {
	code := `l := uint64(len({{ . }}))
	helpers.UInt64Size(l, &n)
`

	constSize := b.keyBuilder.ConstantSize() + b.elementBuilder.ConstantSize()
	if constSize > 0 {
		code += fmt.Sprintf("n += l * %d", constSize)
	}

	keyTpl, keyOK := b.keyBuilder.SizeCode(varIndex)
	elementTpl, elementOK := b.elementBuilder.SizeCode(varIndex)

	if keyOK || elementOK {
		if constSize > 0 {
			code += "\n"
		}

		mk := types.Var("mk", varIndex)
		mv := types.Var("mv", varIndex)
		switch {
		case keyOK && elementOK:
			code += fmt.Sprintf("for %[1]s, %[2]s := range {{ . }} {\n", mk, mv)
		case keyOK:
			code += fmt.Sprintf("for %s := range {{ . }} {\n", mk)
		case elementOK:
			code += fmt.Sprintf("for _, %s := range {{ . }} {\n", mv)
		}

		if keyOK {
			b := &bytes.Buffer{}
			helpers.Execute(b, keyTpl, mk)
			code += types.AddIndent(b.String(), 1)
		}
		if keyOK && elementOK {
			code += "\n"
		}
		if elementOK {
			b := &bytes.Buffer{}
			helpers.Execute(b, elementTpl, mv)
			code += types.AddIndent(b.String(), 1)
		}
		code += "\n}"
	}

	return code, true
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(varIndex *uint64) string {
	keyTpl := b.keyBuilder.MarshalCode(varIndex)
	elementTpl := b.elementBuilder.MarshalCode(varIndex)

	code := `helpers.UInt64Marshal(uint64(len({{ . }})), b, &o)
`

	mk := types.Var("mk", varIndex)
	mv := types.Var("mv", varIndex)
	code += fmt.Sprintf("for %[1]s, %[2]s := range {{ . }} {\n", mk, mv)

	buf := &bytes.Buffer{}
	helpers.Execute(buf, keyTpl, mk)
	code += types.AddIndent(buf.String(), 1) + "\n"

	buf = &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, mv)
	code += types.AddIndent(buf.String(), 1) + "\n"

	code += "}"

	return code
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(varIndex *uint64) string {
	keyTpl := b.keyBuilder.UnmarshalCode(varIndex)
	elementTpl := b.elementBuilder.UnmarshalCode(varIndex)

	code := `var l uint64
helpers.UInt64Unmarshal(&l, b, &o)
`

	mk := types.Var("mk", varIndex)
	mv := types.Var("mv", varIndex)

	code += fmt.Sprintf(`if l > 0 {
	{{ . }} = make(%[1]s, l)

	var %[2]s %[4]s
	var %[3]s %[5]s

	for range l {
`, b.tm.TypeName(b.fieldType), mk, mv, b.tm.TypeName(b.fieldType.Key()), b.tm.TypeName(b.fieldType.Elem()))

	buf := &bytes.Buffer{}
	helpers.Execute(buf, keyTpl, mk)
	code += types.AddIndent(buf.String(), 2) + "\n"

	buf = &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, mv)
	code += types.AddIndent(buf.String(), 2) + "\n"

	code += fmt.Sprintf("		{{ . }}[%[1]s] = %[2]s", mk, mv)

	code += `
	}
}`

	return code
}
