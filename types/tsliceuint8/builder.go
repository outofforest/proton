package tsliceuint8

import (
	_ "embed"
	"fmt"
	"reflect"
	"text/template/parse"

	"github.com/samber/lo"

	"github.com/outofforest/proton/types"
)

var (
	//go:embed templates.gotmpl
	tmpl string
	t    = lo.Must(parse.Parse("", tmpl, "{{", "}}"))
)

const uint8Name = "uint8"

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
	return 1 // covers the first byte of length
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(_ *uint64) (*parse.Tree, any) {
	return t["size"]
	return `l := uint64(len({{ . }}))
helpers.UInt64Size(l, &n)
n += l`, true
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) (*parse.Tree, any) {
	return t["marshal"]
	code := `l := uint64(len({{ . }}))
helpers.UInt64Marshal(l, b, &o)
`

	unsafe := b.tm.Import("unsafe")
	code += "if l > 0 {\n	"

	if b.fieldType.Elem().Name() == uint8Name {
		code += fmt.Sprintf(`copy(b[o:o+l], %[1]s.Slice(&{{ . }}[0], l))
	o += l`, unsafe)
	} else {
		code += fmt.Sprintf(`copy(b[o:o+l], %[1]s.Slice((*byte)(&{{ . }}[0]), l))
	o += l`, unsafe)
	}

	code += "\n}"

	return code
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) (*parse.Tree, any) {
	return t["unmarshal"]
	code := `var l uint64
helpers.UInt64Unmarshal(&l, b, &o)
`

	code += fmt.Sprintf(`if l > 0 {
	{{ . }} = make([]%[1]s, l)
`, b.tm.TypeName(b.fieldType.Elem()))
	if b.fieldType.Elem().Name() == uint8Name {
		code += `	copy({{ . }}, b[o:o+l])`
	} else {
		unsafe := b.tm.Import("unsafe")
		code += fmt.Sprintf(`	copy(%[1]s.Slice((*byte)(&{{ . }}[0]), l), b[o:o+l])`, unsafe)
	}

	code += `
	o += l
}`

	return code
}
