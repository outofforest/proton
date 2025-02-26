package tarrayfloat32_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgArrayFloat32{}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(20, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x000, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgArrayFloat32{
		Value: [5]float32{-128.23, 127.78, -1.2, 0., 1.1},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(20, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xE1, 0x3A, 0x00, 0xC3, 0x5C, 0x8F, 0xFF, 0x42, 0x9A, 0x99, 0x99, 0xBF, 0x00, 0x00, 0x00,
		0x00, 0xCD, 0xCC, 0x8C, 0x3F}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgArrayFloat32Custom{
		Value: custom.ArrayFloat32{-128.23, 127.78, -1.2, 0., 1.1},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(20, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xE1, 0x3A, 0x00, 0xC3, 0x5C, 0x8F, 0xFF, 0x42, 0x9A, 0x99, 0x99, 0xBF, 0x00, 0x00, 0x00,
		0x00, 0xCD, 0xCC, 0x8C, 0x3F}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}
