package tfloat32_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func TestZero(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgFloat32{
		Value: 0x00,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgFloat32{
		Value: 1.,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgFloat32{
		Value: -math.MaxFloat32,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x7F, 0xFF}, b)

	var msg2 pkg1.MsgFloat32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgFloat32{
		Value: math.MaxFloat32,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x7F, 0x7F}, b)

	var msg2 pkg1.MsgFloat32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgFloat32Custom{
		Value: math.MaxFloat32,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x7F, 0x7F}, b)

	var msg2 pkg1.MsgFloat32Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
