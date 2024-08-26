package tint16_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func TestZero(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: 0x00,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgInt16{
		Value: 1,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: 0x01,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x02}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: 0x3F,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x7E}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: 0x40,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x01}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: 0x1FFF,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0x7F}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: 0x2000,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: 0x7FFF,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0x03}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: -0x01,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: -0x40,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x7F}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: -0x41,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x01}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: -0x2000,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: -0x2001,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16{
		Value: -0x8000,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x03}, b)

	var msg2 pkg1.MsgInt16
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt16Custom{
		Value: -0x8000,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x03}, b)

	var msg2 pkg1.MsgInt16Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
