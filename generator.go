package proton

import (
	"bytes"
	"fmt"
	"io"
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

	msgTypes := make([]reflect.Type, 0, len(msgs))
	processed := map[reflect.Type]bool{}
	stack := make([]reflect.Type, 0, len(msgs))
	for _, msg := range msgs {
		msgType := reflect.TypeOf(msg)
		stack = append(stack, msgType)
		if !processed[msgType] {
			msgTypes = append(msgTypes, msgType)
		}
		processed[msgType] = true
	}

	tm := types.NewTypeMap()
	tm.Import("github.com/pkg/errors")

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
			return errors.New("all the msgTypes must belong to the same package")
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

	out, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return errors.WithStack(err)
	}
	defer out.Close()

	if _, err := out.WriteString("package " + path.Base(pkg) + "\n"); err != nil {
		return errors.WithStack(err)
	}

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

		if _, err := out.WriteString("\nimport (\n"); err != nil {
			return errors.WithStack(err)
		}

		if err := writeImports(out, sdkPkgs, imports); err != nil {
			return err
		}
		if len(sdkPkgs) > 0 && len(pkgs) > 0 {
			if _, err := out.WriteString("\n"); err != nil {
				return errors.WithStack(err)
			}
		}
		if err := writeImports(out, pkgs, imports); err != nil {
			return err
		}

		if _, err := out.WriteString(")\n"); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := writeMsgConsts(out, msgTypes); err != nil {
		return errors.WithStack(err)
	}

	if err := writeMsgToIDMapper(out, msgTypes); err != nil {
		return errors.WithStack(err)
	}

	if err := writeIDToMsgMapper(out, msgTypes); err != nil {
		return errors.WithStack(err)
	}

	if _, err := out.Write(b.Bytes()); err != nil {
		return errors.WithStack(err)
	}

	return nil
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

func writeImports(out io.StringWriter, pkgs []string, aliases map[string]string) error {
	sort.Strings(pkgs)
	for _, pkg := range pkgs {
		if _, err := out.WriteString("	"); err != nil {
			return errors.WithStack(err)
		}
		alias := aliases[pkg]
		if path.Base(pkg) == alias {
			if _, err := out.WriteString(fmt.Sprintf("\"%s\"\n", pkg)); err != nil {
				return errors.WithStack(err)
			}
		} else {
			if _, err := out.WriteString(fmt.Sprintf("%s \"%s\"\n", alias, pkg)); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}

func writeMsgConsts(out io.StringWriter, msgTypes []reflect.Type) error {
	const header = `
// ID represents the ID of the message.
type ID uint64

const (
`

	if _, err := out.WriteString(header); err != nil {
		return errors.WithStack(err)
	}

	for i, msgType := range msgTypes {
		if _, err := out.WriteString(fmt.Sprintf("\t// ID%[1]s is the ID of %[1]s message.\n", msgType.Name())); err != nil {
			return errors.WithStack(err)
		}
		if i == 0 {
			if _, err := out.WriteString(fmt.Sprintf("\tID%s ID = iota + 1\n", msgType.Name())); err != nil {
				return errors.WithStack(err)
			}
		} else {
			if _, err := out.WriteString(fmt.Sprintf("\tID%s\n", msgType.Name())); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	if _, err := out.WriteString(")\n"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func writeMsgToIDMapper(out io.StringWriter, msgTypes []reflect.Type) error {
	const header = `
// MsgToID maps message to its ID.
func MsgToID(m interface{}) (ID, error) {
	switch m.(type) {
`
	const footer = `	default:
		return 0, errors.Errorf("unknown message type %T", m)
	}
}
`

	const template = `	case *%[1]s:
		return ID%[1]s, nil
`

	if _, err := out.WriteString(header); err != nil {
		return errors.WithStack(err)
	}

	for _, msgType := range msgTypes {
		if _, err := out.WriteString(fmt.Sprintf(template, msgType.Name())); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.WriteString(footer); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func writeIDToMsgMapper(out io.StringWriter, msgTypes []reflect.Type) error {
	const header = `
// IDToMsg maps ID to the corresponding message.
func IDToMsg(id ID) (interface{}, error) {
	switch id {
`
	const footer = `	default:
		return nil, errors.Errorf("unknown ID %d", id)
	}
}
`

	const template = `	case ID%[1]s:
		return &%[1]s{}, nil
`

	if _, err := out.WriteString(header); err != nil {
		return errors.WithStack(err)
	}

	for _, msgType := range msgTypes {
		if _, err := out.WriteString(fmt.Sprintf(template, msgType.Name())); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.WriteString(footer); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
