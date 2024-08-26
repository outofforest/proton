package tsliceint8_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceInt8{}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgSliceInt8{
		Value: []int8{-128, 127, -1, 0, 1},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgSliceInt8{
		Value: []int8{},
	}, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceInt8{
		Value: []int8{-128, 127, -1, 0, 1},
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0x80, 0x7F, 0xFF, 0x00, 0x01}, b)

	var msg2 pkg1.MsgSliceInt8
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceInt8Custom{
		Value: custom.SliceInt8{-128, 127, -1, 0, 1},
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0x80, 0x7F, 0xFF, 0x00, 0x01}, b)

	var msg2 pkg1.MsgSliceInt8Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
