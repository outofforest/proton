package tarray

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
func New(fieldType reflect.Type, elementBuilder types.BuilderFactory) Builder {
	return Builder{
		fieldType:      fieldType,
		elementBuilder: elementBuilder,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType      reflect.Type
	elementBuilder types.BuilderFactory
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return b.elementBuilder.Dependencies()
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return uint64(b.fieldType.Len()) * b.elementBuilder.ConstantSize()
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(varIndex *uint64) (map[string]*parse.Tree, any) {
	elementTrees, elementData := b.elementBuilder.SizeCode(varIndex)
	if elementTrees == nil {
		return nil, nil
	}
	elementTreeName := types.Var("tmpl", varIndex)
	trees := maps.Clone(elementTrees)
	trees[elementTreeName] = elementTrees["size"]
	trees["size"] = t["size"]
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
	variable := types.Var("av", varIndex)
	i := types.Var("i", varIndex)
	return trees, struct {
		ElementVariable string
		Len             int
		I               string
		Variable        string
		Data            any
		Template        string
	}{
		ElementVariable: variable,
		Len:             b.fieldType.Len(),
		I:               i,
		Variable:        variable + "[" + i + "]",
		Data:            elementData,
		Template:        elementTreeName,
	}
}
