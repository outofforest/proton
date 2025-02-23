package tint16_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func TestZero(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: 0x00,
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
}

func TestByte1PositiveMin(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: 0x01,
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
}

func TestByte1PositiveMax(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: 0x3F,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(1, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x7E}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte2PositiveMin(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: 0x40,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(2, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x80, 0x01}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte2PositiveMax(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: 0x1FFF,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(2, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xFE, 0x7F}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte3PositiveMin(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: 0x2000,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x80, 0x80, 0x01}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte3PositiveMax(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: 0x7FFF,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xFE, 0xFF, 0x03}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte1NegativeMin(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: -0x01,
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
}

func TestByte1NegativeMax(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: -0x40,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(1, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x7F}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte2NegativeMin(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: -0x41,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(2, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x81, 0x01}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte2NegativeMax(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: -0x2000,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(2, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xFF, 0x7F}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte3NegativeMin(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: -0x2001,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x81, 0x80, 0x01}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestByte3NegativeMax(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16{
		Value: -0x8000,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x03}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller(100)

	msg1 := &pkg1.MsgInt16Custom{
		Value: -0x8000,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0xFF, 0xFF, 0x03}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}
