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
	header = `// Marshal marshals the structure.
func (m *{{ .TypeName }}) Marshal(b []byte) uint64 {
`
)

// Build generates code of Marshal method.
func Build(cfg methods.Config, tm types.TypeMap) []byte {
	code := &bytes.Buffer{}

	offset := methods.BitMapLength(cfg.NumOfBooleanFields)
	if offset == 0 {
		code.WriteString("	var o uint64\n")
	} else {
		code.WriteString(fmt.Sprintf("	var o uint64 = %d\n", offset))
	}

	var boolIndex uint64
	lo.Must0(helpers.ForEachField(cfg.Type, func(field reflect.StructField) error {
		if field.Type.Kind() == reflect.Bool {
			byteIndex, bitIndex := methods.BitMapPosition(boolIndex)
			boolIndex++

			code.WriteString(fmt.Sprintf(`	{
		// %[1]s

		if m.%[1]s {
			b[%[2]d] |= 0x%02[3]X
		} else {
			b[%[2]d] &= 0x%02[4]X
		}
	}
`, field.Name, byteIndex, 0x01<<bitIndex, 0xFF^(0x01<<bitIndex)))
			return nil
		}

		builder, err := factory.Get(cfg.Type, field.Type, tm)
		if err != nil {
			return err
		}

		marshalCode := builder.MarshalCodeTemplate()

		code.WriteString("	{\n		// " + field.Name + "\n\n")
		helpers.Execute(code, types.AddIndent(marshalCode, 2), "m."+field.Name)
		code.WriteString("\n	}")
		code.WriteString("\n")

		return nil
	}))

	b := &bytes.Buffer{}
	helpers.Execute(b, header, struct {
		TypeName string
	}{
		TypeName: cfg.Type.Name(),
	})

	if code.Len() > 0 {
		lo.Must(code.WriteTo(b))
	}

	b.WriteString("\n	return o\n}")
	return b.Bytes()
}
