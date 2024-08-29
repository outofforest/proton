package tmap

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/types"
)

// New returns new code builder.
func New(msgType, fieldType reflect.Type, keyBuilder, elementBuilder types.BuilderFactory, tm types.TypeMap) Builder {
	return Builder{
		msgType:        msgType,
		fieldType:      fieldType,
		keyBuilder:     keyBuilder,
		elementBuilder: elementBuilder,
		tm:             tm,
	}
}

// Builder generates the code.
type Builder struct {
	msgType        reflect.Type
	fieldType      reflect.Type
	keyBuilder     types.BuilderFactory
	elementBuilder types.BuilderFactory
	tm             types.TypeMap
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

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(varIndex *uint64) (string, bool) {
	code := "l := uint64(len({{ . }}))\n"

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64SizeCode(), "l")
	code += buf.String() + "\n"

	constSize := b.keyBuilder.ConstantSize() + b.elementBuilder.ConstantSize()
	if constSize > 0 {
		code += fmt.Sprintf("n += l * %d", constSize)
	}

	keyTpl, keyOK := b.keyBuilder.SizeCodeTemplate(varIndex)
	elementTpl, elementOK := b.elementBuilder.SizeCodeTemplate(varIndex)

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

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(varIndex *uint64) string {
	keyTpl := b.keyBuilder.MarshalCodeTemplate(varIndex)
	elementTpl := b.elementBuilder.MarshalCodeTemplate(varIndex)

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64Marshal(), "uint64(len({{ . }}))")
	code := buf.String() + "\n"

	mk := types.Var("mk", varIndex)
	mv := types.Var("mv", varIndex)
	code += fmt.Sprintf("for %[1]s, %[2]s := range {{ . }} {\n", mk, mv)

	buf = &bytes.Buffer{}
	helpers.Execute(buf, keyTpl, mk)
	code += types.AddIndent(buf.String(), 1) + "\n"

	buf = &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, mv)
	code += types.AddIndent(buf.String(), 1) + "\n"

	code += "}"

	return code
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(varIndex *uint64) string {
	keyTpl := b.keyBuilder.UnmarshalCodeTemplate(varIndex)
	elementTpl := b.elementBuilder.UnmarshalCodeTemplate(varIndex)

	code := "var l uint64\n"

	buf := &bytes.Buffer{}
	helpers.Execute(buf, types.UInt64Unmarshal("uint64"), "l")
	code += buf.String() + "\n"

	mk := types.Var("mk", varIndex)
	mv := types.Var("mv", varIndex)

	code += fmt.Sprintf(`{{ . }} = make(%[1]s, l)
var %[2]s %[4]s
var %[3]s %[5]s
`, b.tm.TypeName(b.msgType, b.fieldType), mk, mv, b.tm.TypeName(b.msgType, b.fieldType.Key()), b.tm.TypeName(b.msgType,
		b.fieldType.Elem()))

	code += "for range l {\n"

	buf = &bytes.Buffer{}
	helpers.Execute(buf, keyTpl, mk)
	code += types.AddIndent(buf.String(), 1) + "\n"

	buf = &bytes.Buffer{}
	helpers.Execute(buf, elementTpl, mv)
	code += types.AddIndent(buf.String(), 1) + "\n"

	code += fmt.Sprintf("	{{ . }}[%[1]s] = %[2]s", mk, mv)

	code += "\n}"

	return code
}
