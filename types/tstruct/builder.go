package tstruct

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
func New(fieldType reflect.Type, fieldBuilders []types.BuilderFactory, tm *types.TypeMap) Builder {
	return Builder{
		fieldType:     fieldType,
		fieldBuilders: fieldBuilders,
		tm:            tm,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType     reflect.Type
	fieldBuilders []types.BuilderFactory
	tm            *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return []reflect.Type{b.fieldType}
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(_ *uint64) *parse.Tree {
	return t["size"]
	return fmt.Sprintf("n += %[1]s(&{{ . }})", b.tm.VarName(b.fieldType, "size"))
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(_ *uint64) *parse.Tree {
	return t["marshal"]
	return fmt.Sprintf("o += %[1]s(&{{ . }}, b[o:])", b.tm.VarName(b.fieldType, "marshal"))
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(_ *uint64) *parse.Tree {
	return t["unmarshal"]
	return fmt.Sprintf(`o += %[1]s(&{{ . }}, b[o:])`, b.tm.VarName(b.fieldType, "unmarshal"))
}
