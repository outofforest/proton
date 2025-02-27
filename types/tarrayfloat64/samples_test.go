package tarrayfloat64_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgArrayFloat64{}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(40, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x000, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x000, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)

	id2, err := m.ID(msg2)
	requireT.NoError(err)
	requireT.Equal(id, id2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgArrayFloat64{
		Value: [5]float64{-128.23, 127.78, -1.2, 0., 1.1},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(40, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x8F, 0xC2, 0xF5, 0x28, 0x5C, 0x07, 0x60, 0xC0, 0x52, 0xB8, 0x1E, 0x85, 0xEB, 0xF1, 0x5F,
		0x40, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xF3, 0xBF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9A,
		0x99, 0x99, 0x99, 0x99, 0x99, 0xF1, 0x3F}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgArrayFloat64Custom{
		Value: custom.ArrayFloat64{-128.23, 127.78, -1.2, 0., 1.1},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(40, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x8F, 0xC2, 0xF5, 0x28, 0x5C, 0x07, 0x60, 0xC0, 0x52, 0xB8, 0x1E, 0x85, 0xEB, 0xF1, 0x5F,
		0x40, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xF3, 0xBF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x9A, 0x99, 0x99, 0x99, 0x99, 0x99, 0xF1, 0x3F}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}
