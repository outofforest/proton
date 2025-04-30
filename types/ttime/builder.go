package ttime

import (
	"fmt"
	"reflect"

	"github.com/outofforest/proton/types"
)

// New returns new code builder.
func New(tm *types.TypeMap) Builder {
	return Builder{
		tm: tm,
	}
}

// Builder generates the code.
type Builder struct {
	tm *types.TypeMap
}

// Dependencies returns the list of other types which code must be generated for.
func (b Builder) Dependencies() []reflect.Type {
	return nil
}

// ConstantSize returns the amount of bytes data will always need to be marshaled, independent of actual content.
func (b Builder) ConstantSize() uint64 {
	return 1 // size always takes at least one byte
}

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(_ *uint64) (string, bool) {
	return `helpers.Int64Size({{ . }}.Unix(), &n)`, true
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	return `helpers.Int64Marshal({{ . }}.Unix(), b, &o)`
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	time := b.tm.Import("time")
	return fmt.Sprintf(`var vi int64
helpers.Int64Unmarshal(&vi, b, &o)
{{ . }} = %[1]s.Unix(vi, 0)`, time)
}
