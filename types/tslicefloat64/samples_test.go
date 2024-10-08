package tslicefloat64_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/mass"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceFloat64{}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgSliceFloat64{
		Value: []float64{-128.23, 127.78, -1.2, 0., 1.1},
	}
	l = msg2.Unmarshal(b, mass.New[float64](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgSliceFloat64{}, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceFloat64{
		Value: []float64{-128.23, 127.78, -1.2, 0., 1.1},
	}

	requireT.EqualValues(41, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0x8F, 0xC2, 0xF5, 0x28, 0x5C, 0x07, 0x60, 0xC0, 0x52, 0xB8, 0x1E, 0x85, 0xEB, 0xF1,
		0x5F, 0x40, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xF3, 0xBF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x9A, 0x99, 0x99, 0x99, 0x99, 0x99, 0xF1, 0x3F}, b)

	var msg2 pkg1.MsgSliceFloat64
	l = msg2.Unmarshal(b, mass.New[float64](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceFloat64Custom{
		Value: custom.SliceFloat64{-128.23, 127.78, -1.2, 0., 1.1},
	}

	requireT.EqualValues(41, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0x8F, 0xC2, 0xF5, 0x28, 0x5C, 0x07, 0x60, 0xC0, 0x52, 0xB8, 0x1E, 0x85, 0xEB, 0xF1,
		0x5F, 0x40, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xF3, 0xBF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x9A, 0x99, 0x99, 0x99, 0x99, 0x99, 0xF1, 0x3F}, b)

	var msg2 pkg1.MsgSliceFloat64Custom
	l = msg2.Unmarshal(b, mass.New[custom.Float64](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
