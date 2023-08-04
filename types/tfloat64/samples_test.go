package tfloat64

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func TestZero(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgFloat64{
		Value: 0x00,
	}

	requireT.EqualValues(8, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgFloat64{
		Value: 1.,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgFloat64{
		Value: -math.MaxFloat64,
	}

	requireT.EqualValues(8, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEF, 0xFF}, b)

	var msg2 pkg1.MsgFloat64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgFloat64{
		Value: math.MaxFloat64,
	}

	requireT.EqualValues(8, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEF, 0x7F}, b)

	var msg2 pkg1.MsgFloat64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgFloat64Custom{
		Value: math.MaxFloat64,
	}

	requireT.EqualValues(8, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEF, 0x7F}, b)

	var msg2 pkg1.MsgFloat64Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
