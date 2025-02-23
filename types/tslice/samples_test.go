package tslice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgSlice{}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(1, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestEmpty(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgSlice{
		Value: []bool{},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(1, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(&pkg1.MsgSlice{}, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgSlice{
		Value: []bool{true},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(2, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x01, 0x01}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func Test2(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgSlice{
		Value: []bool{false},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(2, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x01, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func Test3(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgSlice{
		Value: []bool{true, false},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x02, 0x01, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func Test127Items(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgSlice{
		Value: []bool{true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(128, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x7F,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func Test128Items(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgSlice{
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

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(130, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x80, 0x01,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x01, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgSliceCustom{
		Value: custom.Slice{1, 10, 30, 100, 5000, 45543, 4323432432, 32849023843243, 432748732984732},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(29, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x09, 0x01, 0x0A, 0x1E, 0x64, 0x88, 0x27, 0xE7, 0xE3, 0x02, 0xF0, 0xAF, 0xC9, 0x8D, 0x10,
		0xAB, 0xF7, 0x96, 0x93, 0x84, 0xBC, 0x07, 0x9C, 0xDB, 0x86, 0xD4, 0xD2, 0xB2, 0x62}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}
