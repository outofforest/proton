package tmap

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
func New(fieldType reflect.Type, keyBuilder, elementBuilder types.BuilderFactory, tm *types.TypeMap) Builder {
	return Builder{
		fieldType:      fieldType,
		keyBuilder:     keyBuilder,
		elementBuilder: elementBuilder,
		tm:             tm,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType      reflect.Type
	keyBuilder     types.BuilderFactory
	elementBuilder types.BuilderFactory
	tm             *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	dependencies := map[reflect.Type]bool{}
	for _, d := range b.keyBuilder.Dependencies() {
		dependencies[d] = true
	}
	for _, d := range b.elementBuilder.Dependencies() {
		dependencies[d] = true
	}

	res := make([]reflect.Type, 0, len(dependencies))
	for d := range dependencies {
		res = append(res, d)
	}
	return res
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1 // covers the first byte of length
}

// SizeCode returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCode(varIndex *uint64) (map[string]*parse.Tree, any) {
	keyTrees, keyData := b.keyBuilder.SizeCode(varIndex)
	elementTrees, elementData := b.elementBuilder.SizeCode(varIndex)
	trees := maps.Clone(t)
	var keyTreeName, elementTreeName string
	if keyTrees != nil {
		keyTreeName = types.Var("tmpl", varIndex)
		trees[keyTreeName] = keyTrees["size"]
		for k, v := range keyTrees {
			if k != "size" {
				trees[k] = v
			}
		}
	}
	if elementTrees != nil {
		elementTreeName = types.Var("tmpl", varIndex)
		trees[elementTreeName] = elementTrees["size"]
		for k, v := range elementTrees {
			if k != "size" {
				trees[k] = v
			}
		}
	}
	return t, struct {
		ConstSize uint64
		Key       struct {
			Variable string
			Data     any
		}
		Element struct {
			Variable string
			Data     any
		}
		KeyTemplate     string
		ElementTemplate string
	}{
		ConstSize: b.keyBuilder.ConstantSize() + b.elementBuilder.ConstantSize(),
		Key: struct {
			Variable string
			Data     any
		}{
			Variable: types.Var("mk", varIndex),
			Data:     keyData,
		},
		Element: struct {
			Variable string
			Data     any
		}{
			Variable: types.Var("mv", varIndex),
			Data:     elementData,
		},
		KeyTemplate:     keyTreeName,
		ElementTemplate: elementTreeName,
	}
}

// MarshalCode returns code template marshaling the data.
func (b Builder) MarshalCode(varIndex *uint64) (map[string]*parse.Tree, any) {
	keyTrees, keyData := b.keyBuilder.MarshalCode(varIndex)
	elementTrees, elementData := b.elementBuilder.MarshalCode(varIndex)
	trees := maps.Clone(t)
	var keyTreeName, elementTreeName string
	if keyTrees != nil {
		keyTreeName = types.Var("tmpl", varIndex)
		trees[keyTreeName] = keyTrees["marshal"]
		for k, v := range keyTrees {
			if k != "marshal" {
				trees[k] = v
			}
		}
	}
	if elementTrees != nil {
		elementTreeName = types.Var("tmpl", varIndex)
		trees[elementTreeName] = elementTrees["marshal"]
		for k, v := range elementTrees {
			if k != "marshal" {
				trees[k] = v
			}
		}
	}
	return t, struct {
		Key struct {
			Variable string
			Data     any
		}
		Element struct {
			Variable string
			Data     any
		}
		KeyTemplate     string
		ElementTemplate string
	}{
		Key: struct {
			Variable string
			Data     any
		}{
			Variable: types.Var("mk", varIndex),
			Data:     keyData,
		},
		Element: struct {
			Variable string
			Data     any
		}{
			Variable: types.Var("mv", varIndex),
			Data:     elementData,
		},
		KeyTemplate:     keyTreeName,
		ElementTemplate: elementTreeName,
	}
}

// UnmarshalCode returns code template unmarshaling the data.
func (b Builder) UnmarshalCode(varIndex *uint64) (map[string]*parse.Tree, any) {
	keyTrees, keyData := b.keyBuilder.UnmarshalCode(varIndex)
	elementTrees, elementData := b.elementBuilder.UnmarshalCode(varIndex)
	trees := maps.Clone(t)
	var keyTreeName, elementTreeName string
	if keyTrees != nil {
		keyTreeName = types.Var("tmpl", varIndex)
		trees[keyTreeName] = keyTrees["marshal"]
		for k, v := range keyTrees {
			if k != "marshal" {
				trees[k] = v
			}
		}
	}
	if elementTrees != nil {
		elementTreeName = types.Var("tmpl", varIndex)
		trees[elementTreeName] = elementTrees["marshal"]
		for k, v := range elementTrees {
			if k != "marshal" {
				trees[k] = v
			}
		}
	}
	return t, struct {
		Type string
		Key  struct {
			Variable string
			Data     any
		}
		Element struct {
			Variable string
			Data     any
		}
		KeyType         string
		ElementType     string
		KeyTemplate     string
		ElementTemplate string
	}{
		Type: b.tm.TypeName(b.fieldType),
		Key: struct {
			Variable string
			Data     any
		}{
			Variable: types.Var("mk", varIndex),
			Data:     keyData,
		},
		Element: struct {
			Variable string
			Data     any
		}{
			Variable: types.Var("mv", varIndex),
			Data:     elementData,
		},
		KeyType:         b.tm.TypeName(b.fieldType.Key()),
		ElementType:     b.tm.TypeName(b.fieldType.Elem()),
		KeyTemplate:     keyTreeName,
		ElementTemplate: elementTreeName,
	}
}
