package main

import (
	"reflect"

	"github.com/outofforest/proton"
	"github.com/outofforest/proton/test/pkg1"
)

//go:generate go run .

func main() {
	msgs := make([]proton.Msg, 0, len(pkg1.List))
	for _, m := range pkg1.List {
		if reflect.TypeOf(m) == reflect.TypeOf(pkg1.MsgIgnore{}) {
			msgs = append(msgs, proton.Message(m, "Value2Ignored", "Value4Ignored"))
			continue
		}
		msgs = append(msgs, proton.Message(m))
	}
	proton.Generate("../types.proton.go", msgs...)
}
