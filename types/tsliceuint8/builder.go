package tsliceuint8

import (
	_ "embed"
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
func (b Builder) SizeCode(_ *uint64) (map[string]*parse.Tree, any) {
	return t, nil
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) (map[string]*parse.Tree, any) {
	return t, struct {
		Unsafe string
		Type   string
	}{
		Unsafe: b.tm.Import("unsafe"),
		Type:   b.tm.TypeName(b.fieldType.Elem()),
	}
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) (map[string]*parse.Tree, any) {
	return t, struct {
		Unsafe string
		Type   string
	}{
		Unsafe: b.tm.Import("unsafe"),
		Type:   b.tm.TypeName(b.fieldType.Elem()),
	}
}
