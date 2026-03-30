package tstring

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
	return 1 // size always takes at least one byte
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(_ *uint64) *parse.Tree {
	return t["size"]
	return `{
	l := uint64(len({{ . }}))
	helpers.UInt64Size(l, &n)
	n += l
}`, true
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) *parse.Tree {
	return t["marshal"]
	return `{
	l := uint64(len({{ . }}))
	helpers.UInt64Marshal(l, b, &o)
	copy(b[o:o+l], {{ . }})
	o += l
}`
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) *parse.Tree {
	return t["unmarshal"]
	code := `{
	var l uint64
	helpers.UInt64Unmarshal(&l, b, &o)
`

	code += fmt.Sprintf(`	if l > 0 {
		{{ . }} = %[1]s(b[o:o+l])
		o += l
	}
}`, b.tm.TypeName(b.fieldType))

	return code
}
