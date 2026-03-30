package tuint8

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
	return 1
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) string {
	t := b.tm.TypeName(b.fieldType)
	if t == "uint8" {
		return `b[o] = {{ . }}
o++`
	}
	return `b[o] = byte({{ . }})
o++`
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) string {
	t := b.tm.TypeName(b.fieldType)
	if t == "uint8" {
		return `{{ . }} = b[o]
o++`
	}
	return fmt.Sprintf(`{{ . }} = %[1]s(b[o])
o++`, t)
}
