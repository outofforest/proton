package tuint32

import (
	_ "embed"
	"reflect"
	"text/template/parse"

	"github.com/samber/lo"
)

var (
	//go:embed templates.gotmpl
	tmpl string
	t    = lo.Must(parse.Parse("", tmpl, "{{", "}}"))
)

// New returns new code builder.
func New() Builder {
	return Builder{}
}

// Builder generates the code.
type Builder struct{}

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
func (b Builder) SizeCode(_ *uint64) (*parse.Tree, any) {
	return t["size"], nil
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) (*parse.Tree, any) {
	return t["marshal"], nil
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) (*parse.Tree, any) {
	return t["unmarshal"], nil
}
