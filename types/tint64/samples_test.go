package tint64_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func TestZero(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x00,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgInt64{
		Value: 1,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x01,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x02}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x3F,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x7E}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x40,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x1FFF,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x2000,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0xFFFFF,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte4PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x100000,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte4PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x7FFFFFF,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte5PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x8000000,
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte5PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x3FFFFFFFF,
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte6PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x400000000,
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte6PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x1FFFFFFFFFF,
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte7PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x20000000000,
	}

	requireT.EqualValues(7, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte7PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0xFFFFFFFFFFFF,
	}

	requireT.EqualValues(7, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte8PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x1000000000000,
	}

	requireT.EqualValues(8, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte8PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x7FFFFFFFFFFFFF,
	}

	requireT.EqualValues(8, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte9PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x80000000000000,
	}

	requireT.EqualValues(9, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte9PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x3FFFFFFFFFFFFFFF,
	}

	requireT.EqualValues(9, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte10PositiveMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x4000000000000000,
	}

	requireT.EqualValues(10, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte10PositiveMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: 0x7FFFFFFFFFFFFFFF,
	}

	requireT.EqualValues(10, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x01,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x40,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x41,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x2000,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x2001,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x100000,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte4NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x100001,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte4NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x8000000,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte5NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x8000001,
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte5NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x400000000,
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte6NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x400000001,
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte6NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x20000000000,
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte7NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x20000000001,
	}

	requireT.EqualValues(7, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte7NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x1000000000000,
	}

	requireT.EqualValues(7, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte8NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x1000000000001,
	}

	requireT.EqualValues(8, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte8NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x80000000000000,
	}

	requireT.EqualValues(8, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte9NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x80000000000001,
	}

	requireT.EqualValues(9, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte9NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x4000000000000000,
	}

	requireT.EqualValues(9, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte10NegativeMin(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x4000000000000001,
	}

	requireT.EqualValues(10, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte10NegativeMax(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64{
		Value: -0x8000000000000000,
	}

	requireT.EqualValues(10, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01}, b)

	var msg2 pkg1.MsgInt64
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgInt64Custom{
		Value: -0x8000000000000000,
	}

	requireT.EqualValues(10, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01}, b)

	var msg2 pkg1.MsgInt64Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
