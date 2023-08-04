package tint8

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func TestZero(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt8{
		Value: 0x00,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgInt8{
		Value: 1,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt8{
		Value: 0x01,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01}, b)

	var msg2 pkg1.MsgInt8
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt8{
		Value: 0x7F,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x7F}, b)

	var msg2 pkg1.MsgInt8
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt8{
		Value: -0x01,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF}, b)

	var msg2 pkg1.MsgInt8
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt8{
		Value: -0x80,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80}, b)

	var msg2 pkg1.MsgInt8
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt8Custom{
		Value: -0x80,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80}, b)

	var msg2 pkg1.MsgInt8Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
