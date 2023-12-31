package types

import (
	"fmt"
	"math/rand"
	"path"
	"reflect"
)

// BuilderFactory is the interface every type builder must implement
type BuilderFactory interface {
	Dependencies() []reflect.Type
	ConstantSize() uint64
	SizeCodeTemplate() (string, bool)
	MarshalCodeTemplate() string
	UnmarshalCodeTemplate() string
}

// BuilderFactoryConstant is the relaxed interface implemented by builders representing types requiring only constant size of buffer
type BuilderFactoryConstant interface {
	Dependencies() []reflect.Type
	ConstantSize() uint64
	MarshalCodeTemplate() string
	UnmarshalCodeTemplate() string
}

// BuilderFactoryNonConstant is the relaxed interface implemented by builders representing types requiring only non-constant size of buffer
type BuilderFactoryNonConstant interface {
	Dependencies() []reflect.Type
	SizeCodeTemplate() string
	MarshalCodeTemplate() string
	UnmarshalCodeTemplate() string
}

// NewTypeMap creates new type map
func NewTypeMap() TypeMap {
	return TypeMap{
		imports: map[string]string{},
		aliases: map[string]bool{},
	}
}

// TypeMap implements type mapping required by go code
type TypeMap struct {
	imports map[string]string
	aliases map[string]bool
}

// TypeName generates a type name for type
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
			return fmt.Sprintf("[]%s", tm.TypeName(parentType, childType.Elem()))
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

// Import adds package to the list of imports
func (tm TypeMap) Import(pkg string) string {
	if alias := tm.imports[pkg]; alias != "" {
		return alias
	}

	base := path.Base(pkg)
	pkgBase := base
	for tm.aliases[pkgBase] {
		pkgBase = fmt.Sprintf("%s%d", base, rand.Uint64())
	}
	tm.aliases[pkgBase] = true
	tm.imports[pkg] = pkgBase
	return pkgBase
}

// Imports returns the list of collected imports
func (tm TypeMap) Imports() map[string]string {
	return tm.imports
}
