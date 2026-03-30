package tslicefloat64

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
	return 1 // covers the first byte of length
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(_ *uint64) (map[string]*parse.Tree, any) {
	return t
	return `l := uint64(len({{ . }}))
helpers.UInt64Size(l, &n)
n += l * 8`, true
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) (map[string]*parse.Tree, any) {
	return t
	code := `l := uint64(len({{ . }}))
helpers.UInt64Marshal(l, b, &o)
`

	unsafe := b.tm.Import("unsafe")
	code += fmt.Sprintf(`if l > 0 {
	copy(b[o:o+l*8], %[1]s.Slice((*byte)(%[1]s.Pointer(&{{ . }}[0])), l*8))
	o += l * 8
}`, unsafe)

	return code
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) (map[string]*parse.Tree, any) {
	return t
	code := `var l uint64
helpers.UInt64Unmarshal(&l, b, &o)
`

	unsafe := b.tm.Import("unsafe")
	code += fmt.Sprintf(`if l > 0 {
	{{ . }} = make([]%[2]s, l)
	copy(%[1]s.Slice((*byte)(%[1]s.Pointer(&{{ . }}[0])), l*8), b[o:o+l*8])
	o += l * 8
}`, unsafe, b.tm.TypeName(b.fieldType.Elem()))

	return code
}
