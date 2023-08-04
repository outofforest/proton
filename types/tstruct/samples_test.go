package tstruct

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton/test/pkg1"
	spkg1 "github.com/outofforest/proton/test/pkg1/spkg"
	"github.com/outofforest/proton/test/pkg2"
	spkg2 "github.com/outofforest/proton/test/pkg2/spkg"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgStruct{}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgStruct{
		Value1: pkg1.SubMsg{
			Value1: []float32{14.2, 15.6},
			Value2: []spkg1.SubMsg{
				{
					Value: -32,
				},
			},
			Value3: []spkg2.SubMsg{
				{
					Value: -12,
				},
			},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{false: "bb"},
		},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgStruct{
		Value1: pkg1.SubMsg{
			Value1: []float32{},
			Value2: []spkg1.SubMsg{},
			Value3: []spkg2.SubMsg{},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{},
		},
	}, msg2)
}

func TestEmpty(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgStruct{
		Value1: pkg1.SubMsg{
			Value1: []float32{},
			Value2: []spkg1.SubMsg{},
			Value3: []spkg2.SubMsg{},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{},
		},
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgStruct{
		Value1: pkg1.SubMsg{
			Value1: []float32{14.2, 15.6},
			Value2: []spkg1.SubMsg{},
			Value3: []spkg2.SubMsg{},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{true: "aa"},
		},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgStruct{
		Value1: pkg1.SubMsg{
			Value1: []float32{345.7, 892.1},
			Value2: []spkg1.SubMsg{
				{
					Value: 439438,
				},
				{
					Value: -4338,
				},
				{
					Value: 139438,
				},
			},
			Value3: []spkg2.SubMsg{
				{
					Value: 12,
				},
				{
					Value: -321,
				},
			},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{true: "fsfsd"},
			Value2: "refdsfds",
			Value3: 54560,
		},
	}

	requireT.EqualValues(42, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x02, 0x9A, 0xD9, 0xAC, 0x43, 0x66, 0x06, 0x5F, 0x44, 0x03, 0x9C, 0xD2, 0x35, 0xE3, 0x43, 0xDC,
		0x82, 0x11, 0x02, 0x18, 0x81, 0x05, 0x01, 0x01, 0x05, 0x66, 0x73, 0x66, 0x73, 0x64, 0x08, 0x72, 0x65, 0x66, 0x64,
		0x73, 0x66, 0x64, 0x73, 0xA0, 0xAA, 0x03}, b)

	var msg2 pkg1.MsgStruct
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestAnonymousDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgStructAnonymous{}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgStructAnonymous{
		SubMsg: pkg1.SubMsg{
			Value1: []float32{14.2, 15.6},
			Value2: []spkg1.SubMsg{
				{
					Value: -32,
				},
			},
			Value3: []spkg2.SubMsg{
				{
					Value: -12,
				},
			},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{false: "bb"},
		},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgStructAnonymous{
		SubMsg: pkg1.SubMsg{
			Value1: []float32{},
			Value2: []spkg1.SubMsg{},
			Value3: []spkg2.SubMsg{},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{},
		},
	}, msg2)
}

func TestAnonymousEmpty(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgStructAnonymous{
		SubMsg: pkg1.SubMsg{
			Value1: []float32{},
			Value2: []spkg1.SubMsg{},
			Value3: []spkg2.SubMsg{},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{},
		},
	}

	requireT.EqualValues(6, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgStructAnonymous{
		SubMsg: pkg1.SubMsg{
			Value1: []float32{14.2, 15.6},
			Value2: []spkg1.SubMsg{},
			Value3: []spkg2.SubMsg{},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{true: "aa"},
		},
	}
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestAnonymous1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgStructAnonymous{
		SubMsg: pkg1.SubMsg{
			Value1: []float32{345.7, 892.1},
			Value2: []spkg1.SubMsg{
				{
					Value: 439438,
				},
				{
					Value: -4338,
				},
				{
					Value: 139438,
				},
			},
			Value3: []spkg2.SubMsg{
				{
					Value: 12,
				},
				{
					Value: -321,
				},
			},
		},
		Value2: pkg2.SubMsg{
			Value1: map[bool]string{false: "rsfsd"},
			Value2: "refdsfds",
			Value3: 54560,
		},
	}

	requireT.EqualValues(42, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x02, 0x9A, 0xD9, 0xAC, 0x43, 0x66, 0x06, 0x5F, 0x44, 0x03, 0x9C, 0xD2, 0x35, 0xE3, 0x43, 0xDC,
		0x82, 0x11, 0x02, 0x18, 0x81, 0x05, 0x01, 0x00, 0x05, 0x72, 0x73, 0x66, 0x73, 0x64, 0x08, 0x72, 0x65, 0x66, 0x64,
		0x73, 0x66, 0x64, 0x73, 0xA0, 0xAA, 0x03}, b)

	var msg2 pkg1.MsgStructAnonymous
	l = msg2.Unmarshal(b)
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
