package size

import (
	"bytes"
	_ "embed"
	"reflect"
	"text/template"
	"text/template/parse"

	"github.com/samber/lo"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/methods"
	"github.com/outofforest/proton/types"
	"github.com/outofforest/proton/types/factory"
)

//go:embed template.gotmpl
var tmpl string

type Template struct {
	Field    string
	Name     string
	Variable string
	Data     any
}

// Build generates code of Size method.
func Build(cfg methods.Config, tm *types.TypeMap) []byte {
	var n uint64
	if cfg.NumOfBooleanFields > 0 {
		n += methods.BitMapLength(cfg.NumOfBooleanFields)
	}

	methodName := "size"
	if len(cfg.IgnoreFields) > 0 {
		methodName += "i"
	}

	data := struct {
		FuncName  string
		TypeName  string
		N         uint64
		Templates []Template
	}{
		FuncName: tm.VarName(cfg.Type, methodName),
		TypeName: tm.TypeName(cfg.Type),
		N:        n,
	}

	trees := map[string]*parse.Tree{}
	varIndex := new(uint64)
	lo.Must0(helpers.ForEachField(cfg.Type, func(field reflect.StructField) error {
		if cfg.IgnoreFields[field.Name] {
			return nil
		}

		// Booleans are handled separately as a series of bits.
		if field.Type.Kind() == reflect.Bool {
			return nil
		}

		builder, err := factory.Get(field.Type, tm)
		if err != nil {
			return err
		}

		n += builder.ConstantSize()

		fieldTrees, fieldData := builder.SizeCode(varIndex)
		if fieldTrees != nil {
			tmplName := types.Var("tmpl", varIndex)
			trees[tmplName] = fieldTrees["size"]
			for k, v := range fieldTrees {
				trees[k] = v
			}
			data.Templates = append(data.Templates, Template{
				Field:    field.Name,
				Name:     tmplName,
				Variable: "m." + field.Name,
				Data:     fieldData,
			})
		}
		return nil
	}))

	funcTemplate := lo.Must(template.New("").Parse(tmpl))
	for k, v := range trees {
		funcTemplate = lo.Must(funcTemplate.AddParseTree(k, v))
	}

	b := &bytes.Buffer{}
	lo.Must0(funcTemplate.Execute(b, data))
	return b.Bytes()
}
