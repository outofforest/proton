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

// Generate generates proton methods for provided types and stores them in a file.
func Generate(filePath string, msgs ...any) error {
	if len(msgs) == 0 {
		return nil
	}

	b := &bytes.Buffer{}
	var pkg string

	msgTypes := make([]reflect.Type, 0, len(msgs))
	processed := map[reflect.Type]bool{}
	roots := map[reflect.Type]bool{}
	stack := make([]reflect.Type, 0, len(msgs))
	for _, msg := range msgs {
		msgType := reflect.TypeOf(msg)
		stack = append(stack, msgType)
		if !processed[msgType] {
			msgTypes = append(msgTypes, msgType)
		}
		processed[msgType] = true
		roots[msgType] = true
	}

	tm := types.NewTypeMap()
	tm.Import("github.com/pkg/errors")
	tm.Import("github.com/outofforest/proton")
	tm.Import("github.com/outofforest/mass")

	msgInfos := make([]msgInfo, 0, len(msgs))
	for len(stack) > 0 {
		msgType := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		code, dependencies, allocators, err := generateMsg(msgType, tm)
		if err != nil {
			return err
		}

		switch {
		case pkg == "":
			pkg = msgType.PkgPath()
		case pkg != msgType.PkgPath():
			return errors.New("all the msgTypes must belong to the same package")
		}

		if roots[msgType] {
			msgInfos = append(msgInfos, msgInfo{
				Type:       msgType,
				Allocators: allocators,
			})
		}

		for _, dep := range dependencies {
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

	//nolint:nestif
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

	if err := writeMsgConsts(out, msgTypes, tm); err != nil {
		return errors.WithStack(err)
	}

	if err := writeMarshaller(out, msgInfos, tm); err != nil {
		return errors.WithStack(err)
	}

	if _, err := out.Write(b.Bytes()); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func generateMsg(msgType reflect.Type, tm types.TypeMap) ([]byte, []reflect.Type, []reflect.Type, error) {
	pkg := msgType.PkgPath()
	if msgType.Kind() != reflect.Struct {
		return nil, nil, nil, errors.Errorf("type %s is not a struct", msgType)
	}

	cfg := methods.Config{
		Type: msgType,
	}

	dependencies := []reflect.Type{}
	dependenciesMap := map[reflect.Type]struct{}{}

	err := helpers.ForEachField(msgType, func(field reflect.StructField) error {
		builder, err := factory.Get(msgType, field.Type, tm)
		if err != nil {
			return err
		}

		for _, d := range builder.Dependencies() {
			if d.PkgPath() != pkg {
				continue
			}
			if _, exists := dependenciesMap[d]; !exists {
				dependenciesMap[d] = struct{}{}
				dependencies = append(dependencies, d)
			}
		}

		if field.Type.Kind() == reflect.Bool {
			cfg.NumOfBooleanFields++
		}
		return nil
	})
	if err != nil {
		return nil, nil, nil, err
	}

	b := &bytes.Buffer{}
	b.WriteString(fmt.Sprintf("var _ proton.Message = &%s{}\n\n", msgType.Name()))
	b.Write(size.Build(cfg, tm))
	b.WriteString("\n\n")
	b.Write(marshal.Build(cfg, tm))
	b.WriteString("\n\n")
	unmarshalBytes, allocators := unmarshal.Build(cfg, tm)
	b.Write(unmarshalBytes)

	return b.Bytes(), dependencies, allocators, nil
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

func writeMsgConsts(out io.StringWriter, msgTypes []reflect.Type, tm types.TypeMap) error {
	const header = `
const (
`

	if _, err := out.WriteString(header); err != nil {
		return errors.WithStack(err)
	}

	for i, msgType := range msgTypes {
		varName := tm.VarName(msgType, msgType, "ID")
		if _, err := out.WriteString(fmt.Sprintf("\t// %[1]s is the ID of %[2]s message.\n",
			varName, msgType.Name())); err != nil {
			return errors.WithStack(err)
		}
		if i == 0 {
			if _, err := out.WriteString(fmt.Sprintf("\t%s uint64 = iota + 1\n", varName)); err != nil {
				return errors.WithStack(err)
			}
		} else {
			if _, err := out.WriteString(fmt.Sprintf("\t%s\n", varName)); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	if _, err := out.WriteString(")\n"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func writeMarshaller(out io.StringWriter, msgInfos []msgInfo, tm types.TypeMap) error {
	anyType := msgInfos[0].Type

	allocators := []reflect.Type{}
	for _, msgInfo := range msgInfos {
		allocators = append(allocators, msgInfo.Type)
	}
	for _, msgInfo := range msgInfos {
		allocators = types.MergeTypes(allocators, msgInfo.Allocators)
	}

	const constructorHeader = `
var _ proton.Marshaller = Marshaller{}

// NewMarshaller creates marshaller.
func NewMarshaller(capacity uint64) Marshaller {
	return Marshaller{
`
	const typeHeader = `
// Marshaller marshals and unmarshals messages.
type Marshaller struct {
`

	const marshalHeader = `
// Marshal marshals message.
func (m Marshaller) Marshal(msg proton.Marshallable, buf []byte) (retID, retSize uint64, retErr error) {
	defer func() {
		if res := recover(); res != nil {
			retErr = errors.Errorf("marshaling message failed: %s", res)
		}
	}()

	switch msg2 := msg.(type) {
`
	const marshalFooter = `	default:
		return 0, 0, errors.Errorf("unknown message type %T", m)
	}
}
`

	const marshalTemplate = `	case *%[1]s:
		return %[2]s, msg2.Marshal(buf), nil
`

	const unmarshalHeader = `
// Unmarshal unmarshals message.
func (m Marshaller) Unmarshal(id uint64, buf []byte) (retMsg any, retSize uint64, retErr error) {
	defer func() {
		if res := recover(); res != nil {
			retErr = errors.Errorf("unmarshaling message failed: %s", res)
		}
	}()

	switch id {
`
	const unmarshalFooter = `	default:
		return nil, 0, errors.Errorf("unknown ID %d", id)
	}
}
`

	var longestName int
	for _, t := range allocators {
		varName := tm.VarName(anyType, t, "mass")
		if len(varName) > longestName {
			longestName = len(varName)
		}
	}

	if _, err := out.WriteString(constructorHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, t := range allocators {
		varName := tm.VarName(anyType, t, "mass")
		if _, err := out.WriteString(fmt.Sprintf("		%[1]s:%[2]s mass.New[%[3]s](capacity),\n",
			varName, types.Align(varName, longestName),
			tm.TypeName(anyType, t))); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.WriteString("	}\n}\n"); err != nil {
		return errors.WithStack(err)
	}

	if _, err := out.WriteString(typeHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, t := range allocators {
		varName := tm.VarName(anyType, t, "mass")
		if _, err := out.WriteString(fmt.Sprintf("	%[1]s%[2]s *mass.Mass[%[3]s]\n",
			varName, types.Align(varName, longestName),
			tm.TypeName(anyType, t))); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.WriteString("}\n"); err != nil {
		return errors.WithStack(err)
	}

	if _, err := out.WriteString(marshalHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msgInfo := range msgInfos {
		if _, err := out.WriteString(fmt.Sprintf(marshalTemplate, msgInfo.Type.Name(),
			tm.VarName(anyType, msgInfo.Type, "ID"))); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.WriteString(marshalFooter); err != nil {
		return errors.WithStack(err)
	}

	if _, err := out.WriteString(unmarshalHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msgInfo := range msgInfos {
		if _, err := out.WriteString(fmt.Sprintf(`	case %[1]s:
		msg := m.%[2]s.New()
		return msg, msg.Unmarshal(
			buf,
`, tm.VarName(anyType, msgInfo.Type, "ID"), tm.VarName(anyType, msgInfo.Type, "mass"))); err != nil {
			return errors.WithStack(err)
		}

		for _, t := range msgInfo.Allocators {
			if _, err := out.WriteString(fmt.Sprintf("			m.%[1]s,\n",
				tm.VarName(anyType, t, "mass"))); err != nil {
				return errors.WithStack(err)
			}
		}

		if _, err := out.WriteString("		), nil\n"); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.WriteString(unmarshalFooter); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

type msgInfo struct {
	Type       reflect.Type
	Allocators []reflect.Type
}
