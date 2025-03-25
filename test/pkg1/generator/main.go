package main

import (
	"github.com/outofforest/proton"
	"github.com/outofforest/proton/test/pkg1"
)

//go:generate go run .

func main() {
	proton.Generate("../types.proton.go", pkg1.List...)
}
