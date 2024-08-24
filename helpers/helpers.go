package helpers

import (
	"io"
	"reflect"
	"text/template"

	"github.com/samber/lo"
)

// ForEachField iterates over fields in the type and calls a function for each public one
func ForEachField(msgType reflect.Type, fn func(field reflect.StructField) error) error {
	numOfFields := msgType.NumField()
	for i := range numOfFields {
		field := msgType.Field(i)

		if !field.IsExported() {
			continue
		}

		if err := fn(field); err != nil {
			return err
		}
	}
	return nil
}

// Execute executes a template using provided data
func Execute(b io.Writer, code string, data any) {
	lo.Must0(template.Must(template.New("").Parse(code)).Execute(b, data))
}
