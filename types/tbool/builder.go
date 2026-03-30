package tbool

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
	return 1
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) *parse.Tree {
	return t["marshal"]
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) *parse.Tree {
	return t["unmarshal"]
}
