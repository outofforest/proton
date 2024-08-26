package tarrayint8_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgArrayInt8{}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgArrayInt8{
		Value: [5]int8{-128, 127, -1, 0, 1},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgArrayInt8{
		Value: [5]int8{-128, 127, -1, 0, 1},
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x7F, 0xFF, 0x00, 0x01}, b)

	var msg2 pkg1.MsgArrayInt8
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgArrayInt8Custom{
		Value: custom.ArrayInt8{-128, 127, -1, 0, 1},
	}

	requireT.EqualValues(5, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x7F, 0xFF, 0x00, 0x01}, b)

	var msg2 pkg1.MsgArrayInt8Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
