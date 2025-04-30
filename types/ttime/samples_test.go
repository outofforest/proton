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
		Value: time.Unix(0, 0),
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
	requireT.Equal(msg1, msg2)

	id2, err := m.ID(msg2)
	requireT.NoError(err)
	requireT.Equal(id, id2)
}

func TestPositive(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgTime{
		Value: time.Unix(1, 0),
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(1, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x02}, b)

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
		Value: time.Unix(-1, 0),
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(1, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x01}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)

	id2, err := m.ID(msg2)
	requireT.NoError(err)
	requireT.Equal(id, id2)
}
