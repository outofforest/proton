package size

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/samber/lo"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/methods"
	"github.com/outofforest/proton/types"
	"github.com/outofforest/proton/types/factory"
)

const (
	header = `func {{ .FuncName }}(m *{{ .TypeName }}) uint64 {
`
)

// Build generates code of Size method.
func Build(cfg methods.Config, tm *types.TypeMap) []byte {
	var n uint64
	if cfg.NumOfBooleanFields > 0 {
		n += methods.BitMapLength(cfg.NumOfBooleanFields)
	}

	code := &bytes.Buffer{}
	lo.Must0(helpers.ForEachField(cfg.Type, func(field reflect.StructField) error {
		// Booleans are handled separately as a series of bits
		if field.Type.Kind() == reflect.Bool {
			return nil
		}

		builder, err := factory.Get(field.Type, tm)
		if err != nil {
			return err
		}

		n += builder.ConstantSize()

		sizeCode, ok := builder.SizeCodeTemplate(new(uint64))
		if ok {
			code.WriteString("	{\n		// " + field.Name + "\n\n")
			helpers.Execute(code, types.AddIndent(sizeCode, 2), "m."+field.Name)
			code.WriteString("\n	}")
			code.WriteString("\n")
		}
		return nil
	}))

	b := &bytes.Buffer{}
	helpers.Execute(b, header, struct {
		FuncName string
		TypeName string
	}{
		FuncName: tm.VarName(cfg.Type, "size"),
		TypeName: tm.TypeName(cfg.Type),
	})

	if n > 0 {
		b.WriteString(fmt.Sprintf("	var n uint64 = %d\n", n))
	} else {
		b.WriteString("	var n uint64\n")
	}

	if code.Len() > 0 {
		lo.Must(code.WriteTo(b))
	}

	b.WriteString("	return n\n}")
	return b.Bytes()
}
