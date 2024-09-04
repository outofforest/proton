package types

import (
	"fmt"
	"path"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// BuilderFactory is the interface every type builder must implement.
type BuilderFactory interface {
	Dependencies() []reflect.Type
	Allocators() []reflect.Type
	ConstantSize() uint64
	SizeCodeTemplate(varIndex *uint64) (string, bool)
	MarshalCodeTemplate(varIndex *uint64) string
	UnmarshalCodeTemplate(varIndex *uint64) string
}

// BuilderFactoryConstant is the relaxed interface implemented by builders representing types requiring only
// constant size of buffer.
type BuilderFactoryConstant interface {
	Dependencies() []reflect.Type
	Allocators() []reflect.Type
	ConstantSize() uint64
	MarshalCodeTemplate(varIndex *uint64) string
	UnmarshalCodeTemplate(varIndex *uint64) string
}

// BuilderFactoryNonConstant is the relaxed interface implemented by builders representing types requiring only
// non-constant size of buffer.
type BuilderFactoryNonConstant interface {
	Dependencies() []reflect.Type
	Allocators() []reflect.Type
	SizeCodeTemplate(varIndex *uint64) string
	MarshalCodeTemplate(varIndex *uint64) string
	UnmarshalCodeTemplate(varIndex *uint64) string
}

// NewTypeMap creates new type map.
func NewTypeMap() TypeMap {
	return TypeMap{
		imports:      map[string]string{},
		aliases:      map[string]bool{},
		importIndex:  new(uint64),
		varNameCaser: cases.Title(language.English, cases.NoLower),
	}
}

// TypeMap implements type mapping required by go code.
type TypeMap struct {
	imports      map[string]string
	aliases      map[string]bool
	importIndex  *uint64
	varNameCaser cases.Caser
}

// TypeName generates a type name for type.
func (tm TypeMap) TypeName(parentType, childType reflect.Type) string {
	switch childType.Kind() {
	case reflect.Array:
		switch {
		case childType.PkgPath() == parentType.PkgPath():
			return childType.Name()
		case childType.PkgPath() != "":
			return tm.Import(childType.PkgPath()) + "." + childType.Name()
		default:
			return fmt.Sprintf("[%d]%s", childType.Len(), tm.TypeName(parentType, childType.Elem()))
		}
	case reflect.Slice:
		switch {
		case childType.PkgPath() == parentType.PkgPath():
			return childType.Name()
		case childType.PkgPath() != "":
			return tm.Import(childType.PkgPath()) + "." + childType.Name()
		default:
			return "[]" + tm.TypeName(parentType, childType.Elem())
		}
	case reflect.Map:
		switch {
		case childType.PkgPath() == parentType.PkgPath():
			return childType.Name()
		case childType.PkgPath() != "":
			return tm.Import(childType.PkgPath()) + "." + childType.Name()
		default:
			return fmt.Sprintf("map[%s]%s", tm.TypeName(parentType, childType.Key()), tm.TypeName(parentType, childType.Elem()))
		}
	default:
		if childType.PkgPath() == "" || childType.PkgPath() == parentType.PkgPath() {
			return childType.Name()
		}
		return tm.Import(childType.PkgPath()) + "." + childType.Name()
	}
}

// VarName generates deterministic variable name based on its type.
func (tm TypeMap) VarName(parentType, childType reflect.Type, prefix string) string {
	typeName := tm.TypeName(parentType, childType)
	parts := strings.Split(typeName, ".")
	for i, part := range parts {
		parts[i] = tm.varNameCaser.String(part)
	}

	varName := strings.Join(parts, "")
	varName = strings.ReplaceAll(varName, "[", "A")
	varName = strings.ReplaceAll(varName, "]", "B")

	return prefix + varName
}

// Import adds package to the list of imports.
func (tm TypeMap) Import(pkg string) string {
	if alias := tm.imports[pkg]; alias != "" {
		return alias
	}

	base := path.Base(pkg)
	pkgBase := base
	if tm.aliases[pkgBase] {
		*tm.importIndex++
		pkgBase = fmt.Sprintf("%s%d", base, *tm.importIndex)
	}
	tm.aliases[pkgBase] = true
	tm.imports[pkg] = pkgBase
	return pkgBase
}

// Imports returns the list of collected imports.
func (tm TypeMap) Imports() map[string]string {
	return tm.imports
}
