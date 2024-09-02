package tslice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSlice{}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgSlice{
		Value: []bool{true, true, true},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgSlice{}, msg2)
}

func TestEmpty(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSlice{
		Value: []bool{},
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgSlice{
		Value: []bool{true, true, true},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgSlice{}, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSlice{
		Value: []bool{true},
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01, 0x01}, b)

	msg2 := pkg1.MsgSlice{
		Value: []bool{false, false},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test2(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSlice{
		Value: []bool{false},
	}

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01, 0x00}, b)

	msg2 := pkg1.MsgSlice{
		Value: []bool{true, true},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test3(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSlice{
		Value: []bool{true, false},
	}

	requireT.EqualValues(3, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x02, 0x01, 0x00}, b)

	msg2 := pkg1.MsgSlice{
		Value: []bool{false, true, false},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test127Items(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSlice{
		Value: []bool{true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true},
	}

	requireT.EqualValues(128, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x7F,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01}, b)

	msg2 := pkg1.MsgSlice{
		Value: []bool{false, true, false},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test128Items(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSlice{
		Value: []bool{true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false},
	}

	requireT.EqualValues(130, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x80, 0x01,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00}, b)

	msg2 := pkg1.MsgSlice{
		Value: []bool{false, true, false},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgSliceCustom{
		Value: custom.Slice{1, 10, 30, 100, 5000, 45543, 4323432432, 32849023843243, 432748732984732},
	}

	requireT.EqualValues(29, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x09, 0x01, 0x0A, 0x1E, 0x64, 0x88, 0x27, 0xE7, 0xE3, 0x02, 0xF0, 0xAF, 0xC9, 0x8D, 0x10,
		0xAB, 0xF7, 0x96, 0x93, 0x84, 0xBC, 0x07, 0x9C, 0xDB, 0x86, 0xD4, 0xD2, 0xB2, 0x62}, b)

	var msg2 pkg1.MsgSliceCustom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
