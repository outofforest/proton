package marshal

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
	header = `func {{ .FuncName}}(m *{{ .TypeName }}, b []byte) uint64 {
`
)

// Build generates code of Marshal method.
func Build(cfg methods.Config, tm *types.TypeMap) []byte {
	code := &bytes.Buffer{}

	offset := methods.BitMapLength(cfg.NumOfBooleanFields)
	if offset == 0 {
		_, _ = fmt.Fprint(code, "	var o uint64\n")
	} else {
		_, _ = fmt.Fprintf(code, "	var o uint64 = %d\n", offset)
	}

	var boolIndex uint64
	lo.Must0(helpers.ForEachField(cfg.Type, func(field reflect.StructField) error {
		if cfg.IgnoreFields[field.Name] {
			return nil
		}

		if field.Type.Kind() == reflect.Bool {
			byteIndex, bitIndex := methods.BitMapPosition(boolIndex)
			boolIndex++

			_, _ = fmt.Fprintf(code, `	{
		// %[1]s

		if m.%[1]s {
			b[%[2]d] |= 0x%02[3]X
		} else {
			b[%[2]d] &= 0x%02[4]X
		}
	}
`, field.Name, byteIndex, 0x01<<bitIndex, 0xFF^(0x01<<bitIndex))
			return nil
		}

		builder, err := factory.Get(field.Type, tm)
		if err != nil {
			return err
		}

		marshalCode := builder.MarshalCodeTemplate(new(uint64))

		_, _ = fmt.Fprint(code, "	{\n		// "+field.Name+"\n\n")
		helpers.Execute(code, types.AddIndent(marshalCode, 2), "m."+field.Name)
		_, _ = fmt.Fprint(code, "\n	}")
		_, _ = fmt.Fprint(code, "\n")

		return nil
	}))

	methodName := "marshal"
	if len(cfg.IgnoreFields) > 0 {
		methodName += "i"
	}

	b := &bytes.Buffer{}
	helpers.Execute(b, header, struct {
		FuncName string
		TypeName string
	}{
		FuncName: tm.VarName(cfg.Type, methodName),
		TypeName: tm.TypeName(cfg.Type),
	})

	if code.Len() > 0 {
		lo.Must(code.WriteTo(b))
	}

	_, _ = fmt.Fprint(b, "\n	return o\n}")
	return b.Bytes()
}
