package tsliceuint8_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceUint8{}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgSliceUint8{
		Value: []uint8{255, 255, 255, 255, 255},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgSliceUint8{}, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceUint8{
		Value: []uint8{255, 254, 253, 0, 1},
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0xFF, 0xFE, 0xFD, 0x00, 0x01}, b)

	var msg2 pkg1.MsgSliceUint8
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustomDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceUint8Custom2{
		Value: custom.SliceCustomUint8{},
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgSliceUint8Custom2{
		Value: custom.SliceCustomUint8{255, 255, 255, 255, 255},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgSliceUint8Custom2{}, msg2)
}

func TestCustom1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceUint8Custom{
		Value: custom.SliceUint8{255, 254, 253, 0, 1},
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0xFF, 0xFE, 0xFD, 0x00, 0x01}, b)

	var msg2 pkg1.MsgSliceUint8Custom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom2(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceUint8Custom2{
		Value: custom.SliceCustomUint8{255, 254, 253, 0, 1},
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05, 0xFF, 0xFE, 0xFD, 0x00, 0x01}, b)

	var msg2 pkg1.MsgSliceUint8Custom2
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
