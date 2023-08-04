package proton

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"

	"github.com/pkg/errors"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/methods"
	"github.com/outofforest/proton/methods/marshal"
	"github.com/outofforest/proton/methods/size"
	"github.com/outofforest/proton/methods/unmarshal"
	"github.com/outofforest/proton/types"
	"github.com/outofforest/proton/types/factory"
)

// Generate generates proton methods for provided types and stores them in a file
func Generate(filePath string, msgs ...interface{}) error {
	if len(msgs) == 0 {
		return nil
	}

	b := &bytes.Buffer{}
	var pkg string

	processed := map[reflect.Type]bool{}
	stack := make([]reflect.Type, 0, len(msgs))
	for _, msg := range msgs {
		msgType := reflect.TypeOf(msg)
		stack = append(stack, msgType)
		processed[msgType] = true
	}

	tm := types.NewTypeMap()

	for len(stack) > 0 {
		msgType := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		code, dependencies, err := generateMsg(msgType, tm)

		if err != nil {
			return err
		}
		switch {
		case pkg == "":
			pkg = msgType.PkgPath()
		case pkg != msgType.PkgPath():
			return errors.New("all the types must belong to the same package")
		}

		for dep := range dependencies {
			if !processed[dep] {
				stack = append(stack, dep)
				processed[dep] = true
			}
		}

		b.WriteString("\n")
		b.Write(code)
		b.WriteString("\n")
	}

	out := &bytes.Buffer{}
	out.WriteString("package " + path.Base(pkg) + "\n")

	if imports := tm.Imports(); len(imports) > 0 {
		sdkPkgs := make([]string, 0, len(imports))
		pkgs := make([]string, 0, len(imports))

		for pkg := range imports {
			if isSDKPkg(pkg) {
				sdkPkgs = append(sdkPkgs, pkg)
			} else {
				pkgs = append(pkgs, pkg)
			}
		}

		out.WriteString("\nimport (\n")

		writeImports(out, sdkPkgs, imports)
		if len(sdkPkgs) > 0 && len(pkgs) > 0 {
			out.WriteString("\n")
		}
		writeImports(out, pkgs, imports)

		out.WriteString(")\n")
	}

	out.Write(b.Bytes())

	return errors.WithStack(os.WriteFile(filePath, out.Bytes(), 0o600))
}

func generateMsg(msgType reflect.Type, tm types.TypeMap) ([]byte, map[reflect.Type]bool, error) {
	pkg := msgType.PkgPath()
	if msgType.Kind() != reflect.Struct {
		return nil, nil, errors.Errorf("type %s is not a struct", msgType)
	}

	cfg := methods.Config{
		Type: msgType,
	}

	dependencies := map[reflect.Type]bool{}

	err := helpers.ForEachField(msgType, func(field reflect.StructField) error {
		builder, err := factory.Get(msgType, field.Type, tm)
		if err != nil {
			return err
		}

		for _, d := range builder.Dependencies() {
			if d.PkgPath() != pkg {
				continue
			}
			dependencies[d] = true
		}

		if field.Type.Kind() == reflect.Bool {
			cfg.NumOfBooleanFields++
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	b := &bytes.Buffer{}
	b.Write(size.Build(cfg, tm))
	b.WriteString("\n\n")
	b.Write(marshal.Build(cfg, tm))
	b.WriteString("\n\n")
	b.Write(unmarshal.Build(cfg, tm))

	return b.Bytes(), dependencies, nil
}

var sdkPkgs = map[string]bool{
	"archive": true, "bufio": true, "builtin": true, "bytes": true, "compress": true, "container": true, "context": true,
	"crypto": true, "database": true, "debug": true, "embed": true, "encoding": true, "errors": true, "expvar": true,
	"flag": true, "fmt": true, "hash": true, "html": true, "image": true, "index": true, "io": true, "log": true,
	"math": true, "mime": true, "net": true, "os": true, "path": true, "plugin": true, "reflect": true, "regexp": true,
	"runtime": true, "sort": true, "strconv": true, "strings": true, "sync": true, "syscall": true, "testdata": true,
	"testing": true, "text": true, "time": true, "unicode": true, "unsafe": true,
}

func isSDKPkg(pkg string) bool {
	if i := strings.Index(pkg, "/"); i >= 0 {
		pkg = pkg[:i]
	}
	return sdkPkgs[pkg]
}

func writeImports(out *bytes.Buffer, pkgs []string, aliases map[string]string) {
	sort.Strings(pkgs)
	for _, pkg := range pkgs {
		out.WriteString("	")
		alias := aliases[pkg]
		if path.Base(pkg) == alias {
			out.WriteString(fmt.Sprintf("\"%s\"\n", pkg))
		} else {
			out.WriteString(fmt.Sprintf("%s \"%s\"\n", alias, pkg))
		}
	}
}
