package ttime

import (
	"fmt"
	"reflect"
	"time"

	"github.com/outofforest/proton/types"
)

var secondsOffset = time.Time{}.Unix()

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
	return 2
}

// SizeCodeTemplate returns code template computing the required size of buffer
// (above constant size) required to marshal the data.
func (b Builder) SizeCodeTemplate(_ *uint64) (string, bool) {
	return fmt.Sprintf(`helpers.Int64Size({{ . }}.Unix() - %[1]d, &n)
helpers.UInt32Size(uint32({{ . }}.Nanosecond()), &n)`, secondsOffset), true
}

// MarshalCodeTemplate returns code template marshaling the data.
func (b Builder) MarshalCodeTemplate(_ *uint64) string {
	return fmt.Sprintf(`helpers.Int64Marshal({{ . }}.Unix() - %[1]d, b, &o)
helpers.UInt32Marshal(uint32({{ . }}.Nanosecond()), b, &o)`, secondsOffset)
}

// UnmarshalCodeTemplate returns code template unmarshaling the data.
func (b Builder) UnmarshalCodeTemplate(_ *uint64) string {
	time := b.tm.Import("time")
	return fmt.Sprintf(`var seconds int64
var nanoseconds uint32
helpers.Int64Unmarshal(&seconds, b, &o)
helpers.UInt32Unmarshal(&nanoseconds, b, &o)
{{ . }} = %[1]s.Unix(seconds + %[2]d, int64(nanoseconds))`, time, secondsOffset)
}
