package types

import (
	"fmt"
	"path"
	"reflect"
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// BuilderFactory is the interface every type builder must implement.
type BuilderFactory interface {
	Dependencies() []reflect.Type
	ConstantSize() uint64
	SizeCodeTemplate(varIndex *uint64) (string, bool)
	MarshalCodeTemplate(varIndex *uint64) string
	UnmarshalCodeTemplate(varIndex *uint64) string
}

// BuilderFactoryConstant is the relaxed interface implemented by builders representing types requiring only
// constant size of buffer.
type BuilderFactoryConstant interface {
	Dependencies() []reflect.Type
	ConstantSize() uint64
	MarshalCodeTemplate(varIndex *uint64) string
	UnmarshalCodeTemplate(varIndex *uint64) string
}

// BuilderFactoryNonConstant is the relaxed interface implemented by builders representing types requiring only
// non-constant size of buffer.
type BuilderFactoryNonConstant interface {
	Dependencies() []reflect.Type
	SizeCodeTemplate(varIndex *uint64) string
	MarshalCodeTemplate(varIndex *uint64) string
	UnmarshalCodeTemplate(varIndex *uint64) string
}

// TypeMap implements type mapping required by go code.
type TypeMap struct {
	pkg          string
	imports      map[string]string
	aliases      map[string]struct{}
	varIndexes   map[reflect.Type]uint64
	varNameCaser cases.Caser
}

// NewTypeMap creates new type map.
func NewTypeMap(pkg string) *TypeMap {
	return &TypeMap{
		pkg:          pkg,
		imports:      map[string]string{},
		aliases:      map[string]struct{}{},
		varIndexes:   map[reflect.Type]uint64{},
		varNameCaser: cases.Title(language.English, cases.NoLower),
	}
}

// TypeName generates a type name for type.
func (tm *TypeMap) TypeName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Array:
		switch {
		case t.PkgPath() == tm.pkg:
			return t.Name()
		case t.PkgPath() != "":
			return tm.Import(t.PkgPath()) + "." + t.Name()
		default:
			return fmt.Sprintf("[%d]%s", t.Len(), tm.TypeName(t.Elem()))
		}
	case reflect.Slice:
		switch {
		case t.PkgPath() == tm.pkg:
			return t.Name()
		case t.PkgPath() != "":
			return tm.Import(t.PkgPath()) + "." + t.Name()
		default:
			return "[]" + tm.TypeName(t.Elem())
		}
	case reflect.Map:
		switch {
		case t.PkgPath() == tm.pkg:
			return t.Name()
		case t.PkgPath() != "":
			return tm.Import(t.PkgPath()) + "." + t.Name()
		default:
			return fmt.Sprintf("map[%s]%s", tm.TypeName(t.Key()), tm.TypeName(t.Elem()))
		}
	default:
		if t.PkgPath() == "" || t.PkgPath() == tm.pkg {
			return t.Name()
		}
		return tm.Import(t.PkgPath()) + "." + t.Name()
	}
}

// VarName generates deterministic variable name based on its type.
func (tm *TypeMap) VarName(t reflect.Type, prefix string) string {
	if _, exists := tm.varIndexes[t]; !exists {
		tm.varIndexes[t] = uint64(len(tm.varIndexes))
	}
	return prefix + strconv.FormatUint(tm.varIndexes[t], 10)
}

// Import adds package to the list of imports.
func (tm *TypeMap) Import(pkg string) string {
	if alias := tm.imports[pkg]; alias != "" {
		return alias
	}

	base := path.Base(pkg)
	pkgBase := base
	if _, exists := tm.aliases[pkgBase]; exists {
		pkgBase = fmt.Sprintf("%s%d", base, len(tm.aliases))
	}
	tm.aliases[pkgBase] = struct{}{}
	tm.imports[pkg] = pkgBase
	return pkgBase
}

// Imports returns the list of collected imports.
func (tm *TypeMap) Imports() map[string]string {
	return tm.imports
}
