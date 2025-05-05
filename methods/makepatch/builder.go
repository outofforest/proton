package makepatch

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
	header = `func {{ .FuncName}}(m, mSrc *{{ .TypeName }}, b []byte) uint64 {
`
)

// Build generates code of Marshal method.
func Build(cfg methods.Config, tm *types.TypeMap) []byte {
	code := &bytes.Buffer{}

	boolOffset := methods.BitMapLength(cfg.NumOfFields)
	dataOffset := boolOffset + methods.BitMapLength(cfg.NumOfBooleanFields)
	if dataOffset == 0 {
		_, _ = fmt.Fprint(code, "	var o uint64\n")
	} else {
		_, _ = fmt.Fprintf(code, "	var o uint64 = %d\n", dataOffset)
	}

	var presenceIndex, boolIndex uint64
	lo.Must0(helpers.ForEachField(cfg.Type, func(field reflect.StructField) error {
		if cfg.IgnoreFields[field.Name] {
			return nil
		}

		if field.Type.Kind() == reflect.Bool {
			byteIndex, bitIndex := methods.BitMapPosition(boolIndex)
			boolIndex++

			_, _ = fmt.Fprintf(code, `	{
		// %[1]s

		if m.%[1]s == mSrc.%[1]s {
			b[%[2]d] &= 0x%02[4]X
		} else {
			b[%[2]d] |= 0x%02[3]X
		}
	}
`, field.Name, boolOffset+byteIndex, 0x01<<bitIndex, 0xFF^(0x01<<bitIndex))
			return nil
		}

		byteIndex, bitIndex := methods.BitMapPosition(presenceIndex)
		presenceIndex++

		builder, err := factory.Get(field.Type, tm)
		if err != nil {
			return err
		}

		marshalCode := builder.MarshalCodeTemplate(new(uint64))

		reflect := tm.Import("reflect")
		_, _ = fmt.Fprintf(code, `	{
		// %[2]s

		if %[1]s.DeepEqual(m.%[2]s, mSrc.%[2]s) {
			b[%[3]d] &= 0x%02[5]X
		} else {
			b[%[3]d] |= 0x%02[4]X
`, reflect, field.Name, byteIndex, 0x01<<bitIndex, 0xFF^(0x01<<bitIndex))
		helpers.Execute(code, types.AddIndent(marshalCode, 3), "m."+field.Name)
		_, _ = fmt.Fprint(code, "\n		}\n	}\n")

		return nil
	}))

	methodName := "makePatch"
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
