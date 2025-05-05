package proton

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"

	"github.com/outofforest/proton/helpers"
	"github.com/outofforest/proton/methods"
	"github.com/outofforest/proton/methods/applypatch"
	"github.com/outofforest/proton/methods/makepatch"
	"github.com/outofforest/proton/methods/marshal"
	"github.com/outofforest/proton/methods/size"
	"github.com/outofforest/proton/methods/unmarshal"
	"github.com/outofforest/proton/types"
	"github.com/outofforest/proton/types/factory"
	"github.com/outofforest/run"
)

// Msg represents message to generate code for.
type Msg struct {
	MsgType      reflect.Type
	IgnoreFields map[string]bool
}

// Message defines message to generate code for.
func Message(msg any, ignoreFields ...string) Msg {
	ifMap := map[string]bool{}
	for _, f := range ignoreFields {
		ifMap[f] = true
	}

	return Msg{
		MsgType:      reflect.TypeOf(msg),
		IgnoreFields: ifMap,
	}
}

// Generate generates proton methods for provided types and stores them in a file.
func Generate(filePath string, msgs ...Msg) {
	run.New().Run(context.Background(), "generator", func(ctx context.Context) error {
		return generate(filePath, msgs)
	})
}

func generate(filePath string, rootMsgs []Msg) error {
	if len(rootMsgs) == 0 {
		return nil
	}

	b := &bytes.Buffer{}
	pkg, err := findPkg(filePath)
	if err != nil {
		return err
	}

	rootProcessed := map[reflect.Type]bool{}
	subProcessed := map[reflect.Type]bool{}
	stack := make([]Msg, 0, len(rootMsgs))
	for _, msg := range rootMsgs {
		stack = append(stack, msg)
		if _, exists := rootProcessed[msg.MsgType]; exists {
			return errors.Errorf("duplicate message type %s", msg.MsgType)
		}
		rootProcessed[msg.MsgType] = true
		if len(msg.IgnoreFields) == 0 {
			subProcessed[msg.MsgType] = true
		}
	}

	tm := types.NewTypeMap(pkg)
	tm.Import("github.com/pkg/errors")
	tm.Import("github.com/outofforest/proton")
	tm.Import("github.com/outofforest/proton/helpers")

	for len(stack) > 0 {
		msg := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		code, dependencies, err := generateMsg(msg, tm, rootProcessed[msg.MsgType])
		if err != nil {
			return err
		}

		for _, dep := range dependencies {
			if !subProcessed[dep.MsgType] {
				stack = append(stack, dep)
				subProcessed[dep.MsgType] = true
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

	if _, err := fmt.Fprint(out, "package "+path.Base(pkg)+"\n"); err != nil {
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

		if _, err := fmt.Fprint(out, "\nimport (\n"); err != nil {
			return errors.WithStack(err)
		}

		if err := writeImports(out, sdkPkgs, imports); err != nil {
			return err
		}
		if len(sdkPkgs) > 0 && len(pkgs) > 0 {
			if _, err := fmt.Fprint(out, "\n"); err != nil {
				return errors.WithStack(err)
			}
		}
		if err := writeImports(out, pkgs, imports); err != nil {
			return err
		}

		if _, err := fmt.Fprint(out, ")\n"); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := writeMsgConsts(out, rootMsgs, tm); err != nil {
		return errors.WithStack(err)
	}

	if err := writeMarshaller(out, rootMsgs, tm); err != nil {
		return errors.WithStack(err)
	}

	if _, err := out.Write(b.Bytes()); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func generateMsg(msg Msg, tm *types.TypeMap, isRoot bool) ([]byte, []Msg, error) {
	if msg.MsgType.Kind() != reflect.Struct {
		return nil, nil, errors.Errorf("type %s is not a struct", msg.MsgType)
	}
	for f := range msg.IgnoreFields {
		if _, exists := msg.MsgType.FieldByName(f); !exists {
			return nil, nil, errors.Errorf("field %s does not exist in message %s", f, msg.MsgType)
		}
	}

	cfg := methods.Config{
		Type:         msg.MsgType,
		IgnoreFields: msg.IgnoreFields,
	}

	dependencies := []Msg{}
	dependenciesMap := map[reflect.Type]struct{}{}

	err := helpers.ForEachField(msg.MsgType, func(field reflect.StructField) error {
		builder, err := factory.Get(field.Type, tm)
		if err != nil {
			return err
		}

		for _, d := range builder.Dependencies() {
			if _, exists := dependenciesMap[d]; !exists {
				dependenciesMap[d] = struct{}{}
				dependencies = append(dependencies, Msg{MsgType: d})
			}
		}

		if !cfg.IgnoreFields[field.Name] {
			if field.Type.Kind() == reflect.Bool {
				cfg.NumOfBooleanFields++
			} else {
				cfg.NumOfFields++
			}
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
	if isRoot {
		b.WriteString("\n\n")
		b.Write(makepatch.Build(cfg, tm))
		b.WriteString("\n\n")
		b.Write(applypatch.Build(cfg, tm))
	}

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

func writeImports(out io.Writer, pkgs []string, aliases map[string]string) error {
	sort.Strings(pkgs)
	for _, pkg := range pkgs {
		if _, err := fmt.Fprint(out, "	"); err != nil {
			return errors.WithStack(err)
		}
		alias := aliases[pkg]
		if path.Base(pkg) == alias {
			if _, err := fmt.Fprintf(out, "\"%s\"\n", pkg); err != nil {
				return errors.WithStack(err)
			}
		} else {
			if _, err := fmt.Fprintf(out, "%s \"%s\"\n", alias, pkg); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}

func writeMsgConsts(out io.Writer, msgs []Msg, tm *types.TypeMap) error {
	const header = `
const (
`

	if _, err := fmt.Fprint(out, header); err != nil {
		return errors.WithStack(err)
	}

	for i, msg := range msgs {
		varName := tm.VarName(msg.MsgType, "id")
		if i == 0 {
			if _, err := fmt.Fprintf(out, "\t%s uint64 = iota + 1\n", varName); err != nil {
				return errors.WithStack(err)
			}
		} else {
			if _, err := fmt.Fprintf(out, "\t%s\n", varName); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	if _, err := fmt.Fprintf(out, ")\n"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//nolint:gocyclo
func writeMarshaller(out io.Writer, msgs []Msg, tm *types.TypeMap) error {
	const constructor = `
var _ proton.Marshaller = Marshaller{}

// NewMarshaller creates marshaller.
func NewMarshaller() Marshaller {
	return Marshaller{}
}
`
	const typeHeader = `
// Marshaller marshals and unmarshals messages.
type Marshaller struct {
`

	const messagesHeader = `
// Messages returns list of the message types supported by marshaller.
func (m Marshaller) Messages() []any {
	return []any {
`
	const messagesFooter = `	}
}
`

	const messagesTemplate = `		%[1]s{},
`

	const idHeader = `
// ID returns ID of message type.
func (m Marshaller) ID(msg any) (uint64, error) {
	switch msg.(type) {
`
	const idFooter = `	default:
		return 0, errors.Errorf("unknown message type %T", msg)
	}
}
`

	const idTemplate = `	case *%[1]s:
		return %[2]s, nil
`

	const sizeHeader = `
// Size computes the size of marshalled message.
func (m Marshaller) Size(msg any) (uint64, error) {
	switch msg2 := msg.(type) {
`
	const sizeFooter = `	default:
		return 0, errors.Errorf("unknown message type %T", msg)
	}
}
`

	const sizeTemplate = `	case *%[1]s:
		return %[2]s(msg2), nil
`

	const marshalHeader = `
// Marshal marshals message.
func (m Marshaller) Marshal(msg any, buf []byte) (retID, retSize uint64, retErr error) {
	defer helpers.RecoverMarshal(&retErr)

	switch msg2 := msg.(type) {
`
	const marshalFooter = `	default:
		return 0, 0, errors.Errorf("unknown message type %T", msg)
	}
}
`

	const marshalTemplate = `	case *%[1]s:
		return %[2]s, %[3]s(msg2, buf), nil
`

	const unmarshalHeader = `
// Unmarshal unmarshals message.
func (m Marshaller) Unmarshal(id uint64, buf []byte) (retMsg any, retSize uint64, retErr error) {
	defer helpers.RecoverUnmarshal(&retErr)

	switch id {
`

	const unmarshalTemplate = `	case %[2]s:
		msg := &%[3]s{}
		return msg, %[1]s(msg, buf), nil
`

	const unmarshalFooter = `	default:
		return nil, 0, errors.Errorf("unknown ID %d", id)
	}
}
`

	const makePatchHeader = `
// MakePatch creates a patch.
func (m Marshaller) MakePatch(msgDst, msgSrc any, buf []byte) (retID, retSize uint64, retErr error) {
	defer helpers.RecoverMakePatch(&retErr)

	switch msg2 := msgDst.(type) {
`
	const makePatchFooter = `	default:
		return 0, 0, errors.Errorf("unknown message type %T", msgDst)
	}
}
`

	const makePatchTemplate = `	case *%[1]s:
		return %[2]s, %[3]s(msg2, msgSrc.(*%[1]s), buf), nil
`

	const applyPatchHeader = `
// ApplyPatch applies patch.
func (m Marshaller) ApplyPatch(msg any, buf []byte) (retSize uint64, retErr error) {
	defer helpers.RecoverUnmarshal(&retErr)

	switch msg2 := msg.(type) {
`

	const applyPatchTemplate = `	case *%[1]s:
		return %[2]s(msg2, buf), nil
`

	const applyPatchFooter = `	default:
		return 0, errors.Errorf("unknown message type %T", msg)
	}
}
`

	if _, err := fmt.Fprint(out, constructor); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, typeHeader); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, "}\n"); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, messagesHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msg := range msgs {
		if _, err := fmt.Fprintf(out, messagesTemplate, tm.TypeName(msg.MsgType)); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := fmt.Fprint(out, messagesFooter); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, idHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msg := range msgs {
		if _, err := fmt.Fprintf(out, idTemplate, tm.TypeName(msg.MsgType), tm.VarName(msg.MsgType, "id")); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.Write([]byte(idFooter)); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, sizeHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msg := range msgs {
		methodName := "size"
		if len(msg.IgnoreFields) > 0 {
			methodName += "i"
		}

		if _, err := fmt.Fprintf(out, sizeTemplate, tm.TypeName(msg.MsgType),
			tm.VarName(msg.MsgType, methodName)); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.Write([]byte(sizeFooter)); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, marshalHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msg := range msgs {
		methodName := "marshal"
		if len(msg.IgnoreFields) > 0 {
			methodName += "i"
		}

		if _, err := fmt.Fprintf(out, marshalTemplate, tm.TypeName(msg.MsgType), tm.VarName(msg.MsgType, "id"),
			tm.VarName(msg.MsgType, methodName)); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.Write([]byte(marshalFooter)); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, unmarshalHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msg := range msgs {
		methodName := "unmarshal"
		if len(msg.IgnoreFields) > 0 {
			methodName += "i"
		}

		if _, err := fmt.Fprintf(out, unmarshalTemplate, tm.VarName(msg.MsgType, methodName),
			tm.VarName(msg.MsgType, "id"), tm.TypeName(msg.MsgType)); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.Write([]byte(unmarshalFooter)); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, makePatchHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msg := range msgs {
		methodName := "makePatch"
		if len(msg.IgnoreFields) > 0 {
			methodName += "i"
		}

		if _, err := fmt.Fprintf(out, makePatchTemplate, tm.TypeName(msg.MsgType), tm.VarName(msg.MsgType, "id"),
			tm.VarName(msg.MsgType, methodName)); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.Write([]byte(makePatchFooter)); err != nil {
		return errors.WithStack(err)
	}

	if _, err := fmt.Fprint(out, applyPatchHeader); err != nil {
		return errors.WithStack(err)
	}

	for _, msg := range msgs {
		methodName := "applyPatch"
		if len(msg.IgnoreFields) > 0 {
			methodName += "i"
		}

		if _, err := fmt.Fprintf(out, applyPatchTemplate, tm.TypeName(msg.MsgType),
			tm.VarName(msg.MsgType, methodName)); err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err := out.Write([]byte(applyPatchFooter)); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func findPkg(filePath string) (string, error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	dir := filepath.Dir(absFilePath)
	origDir := dir

	var module string
loop:
	for {
		modFile := filepath.Join(dir, "go.mod")
		modData, err := os.ReadFile(modFile)
		switch {
		case errors.Is(err, os.ErrNotExist):
		case err != nil:
			return "", errors.WithStack(err)
		default:
			modF, err := modfile.Parse(modFile, modData, nil)
			if err != nil {
				return "", errors.WithStack(err)
			}
			module = modF.Module.Mod.Path
			break loop
		}

		if dir == "/" {
			return "", errors.New("go.mod not found")
		}
		dir = filepath.Dir(dir)
	}

	return module + strings.TrimPrefix(origDir, dir), nil
}
