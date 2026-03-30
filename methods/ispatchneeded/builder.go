package ispatchneeded

import (
	"bytes"
	_ "embed"
	"reflect"
	"text/template"

	"github.com/samber/lo"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/methods"
	"github.com/outofforest/proton/types"
)

//go:embed template.gotmpl
var tmpl string

type Bool struct {
	Field string
}

type Template struct {
	Field string
}

// Build generates code of Marshal method.
func Build(cfg methods.Config, tm *types.TypeMap) []byte {
	methodName := "isPatchNeeded"
	if len(cfg.IgnoreFields) > 0 {
		methodName += "i"
	}

	data := struct {
		Reflect   string
		FuncName  string
		TypeName  string
		Bools     []Bool
		Templates []Template
	}{
		Reflect:  tm.Import("reflect"),
		FuncName: tm.VarName(cfg.Type, methodName),
		TypeName: tm.TypeName(cfg.Type),
	}

	lo.Must0(helpers.ForEachField(cfg.Type, func(field reflect.StructField) error {
		if cfg.IgnoreFields[field.Name] {
			return nil
		}

		if field.Type.Kind() == reflect.Bool {
			data.Bools = append(data.Bools, Bool{
				Field: field.Name,
			})
			return nil
		}

		data.Templates = append(data.Templates, Template{
			Field: field.Name,
		})
		return nil
	}))

	funcTemplate := lo.Must(template.New("").Parse(tmpl))

	b := &bytes.Buffer{}
	lo.Must0(funcTemplate.Execute(b, data))
	return b.Bytes()
}
