package tslice

import (
	_ "embed"
	"maps"
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
func New(fieldType reflect.Type, elementBuilder types.BuilderFactory, tm *types.TypeMap) Builder {
	return Builder{
		fieldType:      fieldType,
		elementBuilder: elementBuilder,
		tm:             tm,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType      reflect.Type
	elementBuilder types.BuilderFactory
	tm             *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return b.elementBuilder.Dependencies()
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1 // covers the first byte of length
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(varIndex *uint64) (map[string]*parse.Tree, any) {
	elementTrees, elementData := b.elementBuilder.SizeCode(varIndex)
	trees := t
	var elementTreeName string
	if elementTrees != nil {
		elementTreeName = types.Var("tmpl", varIndex)
		trees = maps.Clone(elementTrees)
		trees[elementTreeName] = elementTrees["size"]
		trees["size"] = t["size"]
	}
	return t, struct {
		ConstSize uint64
		Variable  string
		Data      any
		Template  string
	}{
		ConstSize: b.elementBuilder.ConstantSize(),
		Variable:  types.Var("sv", varIndex),
		Data:      elementData,
		Template:  elementTreeName,
	}
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(varIndex *uint64) (map[string]*parse.Tree, any) {
	elementTrees, elementData := b.elementBuilder.MarshalCode(varIndex)
	elementTreeName := types.Var("tmpl", varIndex)
	trees := maps.Clone(elementTrees)
	trees[elementTreeName] = elementTrees["marshal"]
	trees["marshal"] = t["marshal"]
	return trees, struct {
		Variable string
		Data     any
		Template string
	}{
		Variable: types.Var("av", varIndex),
		Data:     elementData,
		Template: elementTreeName,
	}
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(varIndex *uint64) (map[string]*parse.Tree, any) {
	elementTrees, elementData := b.elementBuilder.UnmarshalCode(varIndex)
	elementTreeName := types.Var("tmpl", varIndex)
	trees := maps.Clone(elementTrees)
	trees[elementTreeName] = elementTrees["unmarshal"]
	trees["unmarshal"] = t["unmarshal"]
	variable := types.Var("sv", varIndex)
	i := types.Var("i", varIndex)
	return trees, struct {
		Type            string
		ElementVariable string
		Len             int
		I               string
		Variable        string
		Data            any
		Template        string
	}{
		Type:            b.tm.TypeName(b.fieldType.Elem()),
		ElementVariable: variable,
		Len:             b.fieldType.Len(),
		I:               i,
		Variable:        "(*" + variable + ")[" + i + "]",
		Data:            elementData,
		Template:        elementTreeName,
	}
}
