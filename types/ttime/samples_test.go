package ttime_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func TestZero(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgTime{
		Value: time.Time{}.Local(), //nolint:gosmopolitan
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(2, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x00, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)

	id2, err := m.ID(msg2)
	requireT.NoError(err)
	requireT.Equal(id, id2)
}

func TestPositive(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgTime{
		Value: time.Unix(1, 99),
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(7, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x82, 0xdc, 0x8f, 0xf9, 0xce, 0x3, 0x63}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)

	id2, err := m.ID(msg2)
	requireT.NoError(err)
	requireT.Equal(id, id2)
}

func TestNegative(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgTime{
		Value: time.Unix(-1, 9999),
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(8, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xfe, 0xdb, 0x8f, 0xf9, 0xce, 0x3, 0x8f, 0x4e}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)

	id2, err := m.ID(msg2)
	requireT.NoError(err)
	requireT.Equal(id, id2)
}
