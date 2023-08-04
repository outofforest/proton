package tarrayuint8

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgArrayUint8{}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgArrayUint8{
		Value: [5]uint8{255, 255, 255, 255, 255},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgArrayUint8{
		Value: [5]uint8{255, 254, 253, 0, 1},
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFE, 0xFD, 0x00, 0x01}, b)

	var msg2 pkg1.MsgArrayUint8
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgArrayUint8Custom{
		Value: test.CustomArrayUint8{255, 254, 253, 0, 1},
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFE, 0xFD, 0x00, 0x01}, b)

	var msg2 pkg1.MsgArrayUint8Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom2(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgArrayUint8Custom2{
		Value: test.CustomArrayCustomUint8{255, 254, 253, 0, 1},
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0xFF, 0xFE, 0xFD, 0x00, 0x01}, b)

	var msg2 pkg1.MsgArrayUint8Custom2
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
