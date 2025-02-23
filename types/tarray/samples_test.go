package tarray_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgArray{}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x00, 0x00, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgArray{
		Value: [3]bool{true, true, true},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x01, 0x01, 0x01}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func Test2(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgArray{
		Value: [3]bool{false, true, false},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x00, 0x01, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgArrayCustom{
		Value: custom.Array{53454, 89898},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(6, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xCE, 0xA1, 0x03, 0xAA, 0xBE, 0x05}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}
