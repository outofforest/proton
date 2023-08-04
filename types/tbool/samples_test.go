package tbool

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
)

func Test3Bools1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgBool3{
		Value1: false,
		Value2: false,
		Value3: false,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00}, b)

	msg2 := pkg1.MsgBool3{
		Value1: true,
		Value2: true,
		Value3: true,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test3Bools2(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: false,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x02}, b)

	msg2 := pkg1.MsgBool3{
		Value1: true,
		Value2: false,
		Value3: true,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test3Bools3(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgBool3{
		Value1: true,
		Value2: false,
		Value3: true,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x05}, b)

	msg2 := pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: false,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test10Bools1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgBool10{
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

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00}, b)

	msg2 := pkg1.MsgBool10{
		Value1:  true,
		Value2:  true,
		Value3:  true,
		Value4:  true,
		Value5:  true,
		Value6:  true,
		Value7:  true,
		Value8:  true,
		Value9:  true,
		Value10: true,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test10Bools2(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgBool10{
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

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01, 0x01}, b)

	msg2 := pkg1.MsgBool10{
		Value1:  false,
		Value2:  true,
		Value3:  true,
		Value4:  true,
		Value5:  true,
		Value6:  true,
		Value7:  true,
		Value8:  true,
		Value9:  false,
		Value10: true,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test10Bools3(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgBool10{
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

	requireT.EqualValues(2, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x81, 0x03}, b)

	msg2 := pkg1.MsgBool10{
		Value1:  false,
		Value2:  true,
		Value3:  true,
		Value4:  true,
		Value5:  true,
		Value6:  true,
		Value7:  true,
		Value8:  false,
		Value9:  false,
		Value10: false,
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgBoolCustom{
		Value: true,
	}

	requireT.EqualValues(1, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x01}, b)

	var msg2 pkg1.MsgBoolCustom
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
