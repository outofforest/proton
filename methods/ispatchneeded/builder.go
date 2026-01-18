package ispatchneeded

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/samber/lo"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/methods"
	"github.com/outofforest/proton/types"
)

const (
	header = `func {{ .FuncName}}(m, mSrc *{{ .TypeName }}) bool {
`
)

// Build generates code of Marshal method.
func Build(cfg methods.Config, tm *types.TypeMap) []byte {
	code := &bytes.Buffer{}

	boolOffset := methods.BitMapLength(cfg.NumOfFields)

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

		if m.%[1]s != mSrc.%[1]s {
			return true
		}
	}
`, field.Name, boolOffset+byteIndex, 0x01<<bitIndex, 0xFF^(0x01<<bitIndex))
			return nil
		}

		reflect := tm.Import("reflect")
		_, _ = fmt.Fprintf(code, `	{
		// %[2]s

		if !%[1]s.DeepEqual(m.%[2]s, mSrc.%[2]s) {
			return true
		}
`, reflect, field.Name)

		_, _ = fmt.Fprint(code, "\n	}\n")

		return nil
	}))

	methodName := "isPatchNeeded"
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

	_, _ = fmt.Fprint(b, "\n	return false\n}")
	return b.Bytes()
}
