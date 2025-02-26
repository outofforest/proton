package tarrayfloat32

import (
	"fmt"
	"reflect"

	"github.com/outofforest/proton/types"
)

// New returns new code builder.
func New(fieldType reflect.Type, tm *types.TypeMap) Builder {
	return Builder{
		fieldType: fieldType,
		tm:        tm,
	}
}

// Builder generates the code.
type Builder struct {
	fieldType reflect.Type
	tm        *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return nil
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return uint64(b.fieldType.Len()) * 4
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	unsafe := b.tm.Import("unsafe")
	return fmt.Sprintf(`copy(b[o:o+%[2]d], %[1]s.Slice((*byte)(%[1]s.Pointer(&{{ . }}[0])), %[2]d))
o += %[2]d`, unsafe, b.fieldType.Len()*4)
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	unsafe := b.tm.Import("unsafe")
	return fmt.Sprintf(`copy(%[1]s.Slice((*byte)(%[1]s.Pointer(&{{ . }}[0])), %[2]d), b[o:o+%[2]d])
o += %[2]d`, unsafe, b.fieldType.Len()*4)
}
