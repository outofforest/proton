package tslicefloat32_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/mass"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceFloat32{}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgSliceFloat32{
		Value: []float32{-128.23, 127.78, -1.2, 0., 1.1},
	}
	l = msg2.Unmarshal(b, mass.New[float32](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgSliceFloat32{}, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceFloat32{
		Value: []float32{-128.23, 127.78, -1.2, 0., 1.1},
	}

	requireT.EqualValues(21, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0xE1, 0x3A, 0x00, 0xC3, 0x5C, 0x8F, 0xFF, 0x42, 0x9A, 0x99, 0x99, 0xBF, 0x00, 0x00,
		0x00, 0x00, 0xCD, 0xCC, 0x8C, 0x3F}, b)

	var msg2 pkg1.MsgSliceFloat32
	l = msg2.Unmarshal(b, mass.New[float32](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceFloat32Custom{
		Value: custom.SliceFloat32{-128.23, 127.78, -1.2, 0., 1.1},
	}

	requireT.EqualValues(21, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0xE1, 0x3A, 0x00, 0xC3, 0x5C, 0x8F, 0xFF, 0x42, 0x9A, 0x99, 0x99, 0xBF, 0x00, 0x00,
		0x00, 0x00, 0xCD, 0xCC, 0x8C, 0x3F}, b)

	var msg2 pkg1.MsgSliceFloat32Custom
	l = msg2.Unmarshal(b, mass.New[custom.Float32](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
