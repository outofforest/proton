package makepatch

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

type Bool struct {
	Field string
	Index uint64
	And   uint64
	Or    uint64
}

type Template struct {
	Field    string
	Index    uint64
	And      uint64
	Or       uint64
	Name     string
	Variable string
	Data     any
}

// Build generates code of Marshal method.
func Build(cfg methods.Config, tm *types.TypeMap) []byte {
	methodName := "makePatch"
	if len(cfg.IgnoreFields) > 0 {
		methodName += "i"
	}

	boolOffset := methods.BitMapLength(cfg.NumOfFields)
	dataOffset := boolOffset + methods.BitMapLength(cfg.NumOfBooleanFields)

	data := struct {
		Reflect    string
		FuncName   string
		TypeName   string
		DataOffset uint64
		Bools      []Bool
		Templates  []Template
	}{
		Reflect:    tm.Import("reflect"),
		FuncName:   tm.VarName(cfg.Type, methodName),
		TypeName:   tm.TypeName(cfg.Type),
		DataOffset: dataOffset,
	}

	var presenceIndex, boolIndex uint64
	trees := map[string]*parse.Tree{}
	varIndex := new(uint64)
	lo.Must0(helpers.ForEachField(cfg.Type, func(field reflect.StructField) error {
		if cfg.IgnoreFields[field.Name] {
			return nil
		}

		if field.Type.Kind() == reflect.Bool {
			byteIndex, bitIndex := methods.BitMapPosition(boolIndex)
			boolIndex++

			data.Bools = append(data.Bools, Bool{
				Field: field.Name,
				Index: boolOffset + byteIndex,
				And:   0xFF ^ (0x01 << bitIndex),
				Or:    0x01 << bitIndex,
			})
			return nil
		}

		byteIndex, bitIndex := methods.BitMapPosition(presenceIndex)
		presenceIndex++

		builder, err := factory.Get(field.Type, tm)
		if err != nil {
			return err
		}

		fieldTrees, fieldData := builder.MarshalCode(varIndex)
		tmplName := types.Var("tmpl", varIndex)
		trees[tmplName] = fieldTrees["marshal"]
		for k, v := range fieldTrees {
			trees[k] = v
		}
		data.Templates = append(data.Templates, Template{
			Field:    field.Name,
			Index:    byteIndex,
			And:      0xFF ^ (0x01 << bitIndex),
			Or:       0x01 << bitIndex,
			Name:     tmplName,
			Variable: "m." + field.Name,
			Data:     fieldData,
		})
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
