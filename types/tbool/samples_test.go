package tbool_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func Test3Bools1(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgBool3{
		Value1: false,
		Value2: false,
		Value3: false,
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

func Test3Bools2(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: false,
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

func Test3Bools3(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgBool3{
		Value1: true,
		Value2: false,
		Value3: true,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(1, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x05}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func Test10Bools1(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgBool10{
		Value1:  false,
		Value2:  false,
		Value3:  false,
		Value4:  false,
		Value5:  false,
		Value6:  false,
		Value7:  false,
		Value8:  false,
		Value9:  false,
		Value10: false,
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
}

func Test10Bools2(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgBool10{
		Value1:  true,
		Value2:  false,
		Value3:  false,
		Value4:  false,
		Value5:  false,
		Value6:  false,
		Value7:  false,
		Value8:  false,
		Value9:  true,
		Value10: false,
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

func Test10Bools3(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgBool10{
		Value1:  true,
		Value2:  false,
		Value3:  false,
		Value4:  false,
		Value5:  false,
		Value6:  false,
		Value7:  false,
		Value8:  true,
		Value9:  true,
		Value10: true,
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(2, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x81, 0x03}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgBoolCustom{
		Value: true,
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
