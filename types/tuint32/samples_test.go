package tuint32

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func TestZero(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x00,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgUint32{
		Value: 1,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1Min(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x01,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte1Max(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x7F,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x7F}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2Min(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x80,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x01}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte2Max(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x3FFF,
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0x7F}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3Min(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x4000,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte3Max(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x1FFFFF,
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte4Min(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x200000,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte4Max(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0xFFFFFFF,
	}

	requireT.EqualValues(4, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0x7F}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte5Min(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0x10000000,
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x80, 0x80, 0x80, 0x01}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestByte5Max(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32{
		Value: 0xFFFFFFFF,
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x0F}, b)

	var msg2 pkg1.MsgUint32
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgUint32Custom{
		Value: 0xFFFFFFFF,
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x0F}, b)

	var msg2 pkg1.MsgUint32Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
