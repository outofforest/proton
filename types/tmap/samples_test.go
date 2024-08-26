package tmap_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMap{}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgMap{
		Value: map[uint64]uint64{2: 1, 3: 2, 128: 128},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgMap{Value: map[uint64]uint64{}}, msg2)
}

func TestEmpty(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMap{
		Value: map[uint64]uint64{},
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgMap{
		Value: map[uint64]uint64{1: 1, 2: 2},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMap{
		Value: map[uint64]uint64{1: 2},
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01, 0x01, 0x02}, b)

	var msg2 pkg1.MsgMap
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test2(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMap{
		Value: map[uint64]uint64{1: 2, 3: 4, 5: 6},
	}

	requireT.EqualValues(7, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x03}, b[0:1])

	var msg2 pkg1.MsgMap
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test127(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMap{
		Value: map[uint64]uint64{},
	}

	for i := uint64(1); i <= 127; i++ {
		msg1.Value[i] = i
	}

	requireT.EqualValues(255, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x7F}, b[0:1])

	var msg2 pkg1.MsgMap
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test128(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMap{
		Value: map[uint64]uint64{},
	}

	for i := uint64(1); i <= 128; i++ {
		msg1.Value[i] = i
	}

	requireT.EqualValues(260, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x01}, b[0:2])

	var msg2 pkg1.MsgMap
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestString(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMapString{
		Value: map[string]string{},
	}

	for i := range 128 {
		msg1.Value[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("value-%d", i)
	}

	requireT.EqualValues(2086, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x01}, b[0:2])

	var msg2 pkg1.MsgMapString
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMapCustom{
		Value: custom.Map{},
	}

	for i := custom.Int32(1); i <= 128; i++ {
		msg1.Value[custom.String(fmt.Sprintf("key%d", i))] = i
	}

	requireT.EqualValues(983, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x01}, b[0:2])

	var msg2 pkg1.MsgMapCustom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
