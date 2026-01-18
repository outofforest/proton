package test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/proton"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
	spkg1 "github.com/outofforest/proton/test/pkg1/spkg"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()
	msg1 := &pkg1.MsgMixed{}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(18, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)

	requireT.Equal(&pkg1.MsgMixed{
		Value1: nil,
		Value2: nil,
		Value3: nil,
		Value4: [12]map[int8]float32{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		Value5: nil,
	}, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{
			"aa": {
				Value: 1,
			},
			"afdsfdsa": {
				Value: 14325,
			},
		},
		Value2: map[uint8][]string{
			4:  {"aa", "bb"},
			10: {"rsedfs", "fdsfsd", "gfdgfdgdf"},
		},
		Value3: [][32]uint16{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
				29, 30, 31},
			{32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
				58, 59, 60, 61, 62, 63},
		},
		Value4: [12]map[int8]float32{
			{0: 1.}, {1: 1.1, 3: 45.78}, {2: 1.2}, {3: 1.3, 5: 4390.12, 8: 8945.154}, {4: 1.4}, {5: 1.5}, {6: 1.6},
			{7: 1.7}, {8: 1.8}, {9: 1.9}, {10: 2.}, {11: 2.1},
		},
		Value5: [][3][]map[int16][2]int64{
			{
				{
					{2: {1, 2}, 4: {435, 545}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}, 6: {543, 3232}},
				},
				{
					{2: {1, 2}, 9: {42343242, 432432}},
					{2: {1, 2}},
				},
			},
			{
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
			},
		},
		Value6: true,
		Value7: false,
		Value8: "fdfsfdsfkds;fdsfd;skfl;sdkd",
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgMixedCustom{
		Value: custom.Mixed{
			{
				{
					{2: {1, 2}, 4: {435, 545}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}, 6: {543, 3232}},
				},
				{
					{2: {1, 2}, 9: {42343242, 432432}},
					{2: {1, 2}},
				},
			},
			{
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
			},
		},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)

	requireT.Equal(msg1, msg2)
}

func TestMarshalBufferTooSmall(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{
			"aa": {
				Value: 1,
			},
			"afdsfdsa": {
				Value: 14325,
			},
		},
		Value2: map[uint8][]string{
			4:  {"aa", "bb"},
			10: {"rsedfs", "fdsfsd", "gfdgfdgdf"},
		},
		Value3: [][32]uint16{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
				29, 30, 31},
			{32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
				58, 59, 60, 61, 62, 63},
		},
		Value4: [12]map[int8]float32{
			{0: 1.}, {1: 1.1, 3: 45.78}, {2: 1.2}, {3: 1.3, 5: 4390.12, 8: 8945.154}, {4: 1.4}, {5: 1.5}, {6: 1.6},
			{7: 1.7}, {8: 1.8}, {9: 1.9}, {10: 2.}, {11: 2.1},
		},
		Value5: [][3][]map[int16][2]int64{
			{
				{
					{2: {1, 2}, 4: {435, 545}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}, 6: {543, 3232}},
				},
				{
					{2: {1, 2}, 9: {42343242, 432432}},
					{2: {1, 2}},
				},
			},
			{
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
			},
		},
		Value6: true,
		Value7: false,
		Value8: "fdfsfdsfkds;fdsfd;skfl;sdkd",
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	b := make([]byte, size-1)
	_, _, err = m.Marshal(msg1, b)
	requireT.ErrorIs(err, proton.ErrBufferFailure)
}

func TestUnmarshalBufferTooSmall(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{
			"aa": {
				Value: 1,
			},
			"afdsfdsa": {
				Value: 14325,
			},
		},
		Value2: map[uint8][]string{
			4:  {"aa", "bb"},
			10: {"rsedfs", "fdsfsd", "gfdgfdgdf"},
		},
		Value3: [][32]uint16{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
				29, 30, 31},
			{32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
				58, 59, 60, 61, 62, 63},
		},
		Value4: [12]map[int8]float32{
			{0: 1.}, {1: 1.1, 3: 45.78}, {2: 1.2}, {3: 1.3, 5: 4390.12, 8: 8945.154}, {4: 1.4}, {5: 1.5}, {6: 1.6},
			{7: 1.7}, {8: 1.8}, {9: 1.9}, {10: 2.}, {11: 2.1},
		},
		Value5: [][3][]map[int16][2]int64{
			{
				{
					{2: {1, 2}, 4: {435, 545}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}, 6: {543, 3232}},
				},
				{
					{2: {1, 2}, 9: {42343242, 432432}},
					{2: {1, 2}},
				},
			},
			{
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
				{
					{2: {1, 2}},
					{2: {1, 2}},
				},
			},
		},
		Value6: true,
		Value7: false,
		Value8: "fdfsfdsfkds;fdsfd;skfl;sdkd",
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)

	b2 := make([]byte, len(b)-1)
	copy(b2, b)

	_, _, err = m.Unmarshal(id, b2)
	requireT.ErrorIs(err, proton.ErrBufferFailure)
}

func TestIDsAreGeneratedInSequence(t *testing.T) {
	requireT := require.New(t)
	m := pkg1.NewMarshaller()

	for i, v := range pkg1.List {
		id, err := m.ID(reflect.New(reflect.TypeOf(v)).Elem().Addr().Interface())
		requireT.NoError(err)
		requireT.EqualValues(i+1, id)
	}
}

func TestMessages(t *testing.T) {
	require.Equal(t, pkg1.List, pkg1.NewMarshaller().Messages())
}

func TestFieldsInMessageAreIgnored(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgIgnore{
		Value1:        true,
		Value2Ignored: true,
		Value3:        "A",
		Value4Ignored: "B",
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(3, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x01, 0x01, 0x41}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(&pkg1.MsgIgnore{
		Value1:        true,
		Value2Ignored: false,
		Value3:        "A",
		Value4Ignored: "",
	}, msg2)

	id2, err := m.ID(msg2)
	requireT.NoError(err)
	requireT.Equal(id, id2)
}

func TestFieldsInSubmessageAreNotIgnored(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msg1 := &pkg1.MsgNotIgnore{
		SubMsg: pkg1.MsgIgnore{
			Value1:        true,
			Value2Ignored: true,
			Value3:        "A",
			Value4Ignored: "B",
		},
	}

	size, err := m.Size(msg1)
	requireT.NoError(err)
	requireT.EqualValues(5, size)

	b := make([]byte, size)
	id, l, err := m.Marshal(msg1, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal([]byte{0x03, 0x01, 0x41, 0x01, 0x42}, b)

	msg2, l, err := m.Unmarshal(id, b)
	requireT.NoError(err)
	requireT.Equal(size, l)
	requireT.Equal(&pkg1.MsgNotIgnore{
		SubMsg: pkg1.MsgIgnore{
			Value1:        true,
			Value2Ignored: true,
			Value3:        "A",
			Value4Ignored: "B",
		},
	}, msg2)

	id2, err := m.ID(msg2)
	requireT.NoError(err)
	requireT.Equal(id, id2)
}

const maxPatchSize = 128

func TestPatch3Bools1(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: true,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: true,
		Value2: true,
		Value3: false,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x05}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools2(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: true,
		Value2: true,
		Value3: false,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: true,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x05}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools3(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: false,
		Value2: false,
		Value3: false,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: true,
		Value2: true,
		Value3: true,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x07}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools4(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: true,
		Value2: true,
		Value3: true,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: false,
		Value2: false,
		Value3: false,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x07}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools5(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: true,
		Value2: false,
		Value3: false,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: false,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x03}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools6(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: false,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: false,
		Value2: false,
		Value3: true,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x06}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools7(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: false,
		Value2: false,
		Value3: true,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: true,
		Value2: false,
		Value3: false,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x05}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools8(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: false,
		Value2: false,
		Value3: false,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: false,
		Value2: false,
		Value3: false,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x00}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools9(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: true,
		Value2: true,
		Value3: true,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: true,
		Value2: true,
		Value3: true,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x00}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools10(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: true,
		Value2: false,
		Value3: true,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: true,
		Value2: false,
		Value3: true,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x00}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch3Bools11(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: false,
	}
	msgDst := &pkg1.MsgBool3{
		Value1: false,
		Value2: true,
		Value3: false,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x00}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch10Bools1(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool10{
		Value1:  false,
		Value2:  true,
		Value3:  false,
		Value4:  true,
		Value5:  false,
		Value6:  true,
		Value7:  false,
		Value8:  true,
		Value9:  false,
		Value10: true,
	}
	msgDst := &pkg1.MsgBool10{
		Value1:  true,
		Value2:  false,
		Value3:  true,
		Value4:  false,
		Value5:  true,
		Value6:  false,
		Value7:  true,
		Value8:  false,
		Value9:  true,
		Value10: false,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0xff, 0x03}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch10Bools2(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool10{
		Value1:  true,
		Value2:  false,
		Value3:  true,
		Value4:  false,
		Value5:  true,
		Value6:  false,
		Value7:  true,
		Value8:  false,
		Value9:  true,
		Value10: false,
	}
	msgDst := &pkg1.MsgBool10{
		Value1:  false,
		Value2:  true,
		Value3:  false,
		Value4:  true,
		Value5:  false,
		Value6:  true,
		Value7:  false,
		Value8:  true,
		Value9:  false,
		Value10: true,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0xff, 0x03}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch10Bools3(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool10{
		Value1:  true,
		Value2:  false,
		Value3:  true,
		Value4:  false,
		Value5:  true,
		Value6:  false,
		Value7:  true,
		Value8:  false,
		Value9:  true,
		Value10: false,
	}
	msgDst := &pkg1.MsgBool10{
		Value1:  true,
		Value2:  false,
		Value3:  true,
		Value4:  false,
		Value5:  true,
		Value6:  false,
		Value7:  true,
		Value8:  false,
		Value9:  true,
		Value10: false,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x00, 0x00}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatch10Bools4(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgBool10{
		Value1:  true,
		Value2:  false,
		Value3:  true,
		Value4:  false,
		Value5:  true,
		Value6:  false,
		Value7:  true,
		Value8:  false,
		Value9:  true,
		Value10: false,
	}
	msgDst := &pkg1.MsgBool10{
		Value1:  false,
		Value2:  false,
		Value3:  true,
		Value4:  false,
		Value5:  true,
		Value6:  false,
		Value7:  false,
		Value8:  true,
		Value9:  true,
		Value10: true,
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0xc1, 0x02}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchIgnored1(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgIgnore{
		Value1:        false,
		Value2Ignored: false,
		Value3:        "",
		Value4Ignored: "",
	}
	msgDst := &pkg1.MsgIgnore{
		Value1:        true,
		Value2Ignored: true,
		Value3:        "A",
		Value4Ignored: "B",
	}
	msgExpected := &pkg1.MsgIgnore{
		Value1:        true,
		Value2Ignored: false,
		Value3:        "A",
		Value4Ignored: "",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	requireT.False(m.IsPatchNeeded(msgSrc, &pkg1.MsgIgnore{
		Value1:        false,
		Value2Ignored: false,
		Value3:        "",
		Value4Ignored: "A",
	}))
	requireT.False(m.IsPatchNeeded(msgDst, &pkg1.MsgIgnore{
		Value1:        true,
		Value2Ignored: true,
		Value3:        "A",
		Value4Ignored: "A",
	}))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x01, 0x01, 0x01, 0x41}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgExpected, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchIgnored2(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgIgnore{
		Value1:        false,
		Value2Ignored: false,
		Value3:        "",
		Value4Ignored: "",
	}
	msgDst := &pkg1.MsgIgnore{
		Value1:        false,
		Value2Ignored: true,
		Value3:        "",
		Value4Ignored: "B",
	}
	msgExpected := &pkg1.MsgIgnore{
		Value1:        false,
		Value2Ignored: false,
		Value3:        "",
		Value4Ignored: "",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.False(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.False(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x00, 0x00}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgExpected, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchIgnored3(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgIgnore{
		Value1:        true,
		Value2Ignored: false,
		Value3:        "A",
		Value4Ignored: "",
	}
	msgDst := &pkg1.MsgIgnore{
		Value1:        false,
		Value2Ignored: true,
		Value3:        "C",
		Value4Ignored: "B",
	}
	msgExpected := &pkg1.MsgIgnore{
		Value1:        false,
		Value2Ignored: false,
		Value3:        "C",
		Value4Ignored: "",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	requireT.False(m.IsPatchNeeded(msgSrc, &pkg1.MsgIgnore{
		Value1:        true,
		Value2Ignored: false,
		Value3:        "A",
		Value4Ignored: "A",
	}))
	requireT.False(m.IsPatchNeeded(msgDst, &pkg1.MsgIgnore{
		Value1:        false,
		Value2Ignored: true,
		Value3:        "C",
		Value4Ignored: "A",
	}))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x01, 0x01, 0x01, 0x43}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgExpected, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchMixed1(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{"a": {Value: 1}},
		Value2: map[uint8][]string{1: {"b"}},
		Value3: [][32]uint16{{1, 2}},
		Value4: [12]map[int8]float32{{1: 2}},
		Value5: [][3][]map[int16][2]int64{{{{1: {2}}}}},
		Value6: true,
		Value7: true,
		Value8: "A",
	}
	msgDst := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{"a": {Value: 1}},
		Value2: map[uint8][]string{1: {"b"}},
		Value3: [][32]uint16{{1, 2}},
		Value4: [12]map[int8]float32{{1: 2}},
		Value5: [][3][]map[int16][2]int64{{{{1: {2}}}}},
		Value6: true,
		Value7: true,
		Value8: "A",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgDst))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x0, 0x0}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchMixed2(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{"a": {Value: 1}},
		Value2: map[uint8][]string{1: {"b"}},
		Value3: [][32]uint16{{1, 2}},
		Value4: [12]map[int8]float32{{1: 2}},
		Value5: [][3][]map[int16][2]int64{{{{1: {2}}}}},
		Value6: true,
		Value7: true,
		Value8: "A",
	}
	msgDst := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{"b": {Value: 2}},
		Value2: map[uint8][]string{1: {"c"}},
		Value3: [][32]uint16{{1}},
		Value4: [12]map[int8]float32{{2: 2}},
		Value5: [][3][]map[int16][2]int64{{{{2: {3}}}}},
		Value6: false,
		Value7: false,
		Value8: "B",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{
		0x3f, 0x03, 0x01, 0x01, 0x62, 0x04, 0x01, 0x01, 0x01, 0x01, 0x63, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x04, 0x06, 0x00, 0x00, 0x00, 0x01, 0x42,
	}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchMixed3(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{"a": {Value: 1}},
		Value2: map[uint8][]string{1: {"b"}},
		Value3: [][32]uint16{{1, 2}},
		Value4: [12]map[int8]float32{{1: 2}},
		Value5: [][3][]map[int16][2]int64{{{{1: {2}}}}},
		Value6: true,
		Value7: true,
		Value8: "A",
	}
	msgDst := &pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{"a": {Value: 1}},
		Value2: map[uint8][]string{1: {"c"}},
		Value3: [][32]uint16{{1, 2}},
		Value4: [12]map[int8]float32{{1: 2}},
		Value5: [][3][]map[int16][2]int64{{{{2: {2}}}}},
		Value6: true,
		Value7: false,
		Value8: "A",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x12, 0x02, 0x01, 0x01, 0x01, 0x01, 0x63, 0x01, 0x01, 0x01, 0x04, 0x04, 0x00, 0x00, 0x00},
		b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchStrings1(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgStrings{
		Value1:  "",
		Value2:  "",
		Value3:  "",
		Value4:  "",
		Value5:  "",
		Value6:  "",
		Value7:  "",
		Value8:  "",
		Value9:  "",
		Value10: "",
	}
	msgDst := &pkg1.MsgStrings{
		Value1:  "",
		Value2:  "",
		Value3:  "",
		Value4:  "",
		Value5:  "",
		Value6:  "",
		Value7:  "",
		Value8:  "",
		Value9:  "",
		Value10: "",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgDst))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0x0, 0x0}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchStrings2(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgStrings{
		Value1:  "",
		Value2:  "",
		Value3:  "",
		Value4:  "",
		Value5:  "",
		Value6:  "",
		Value7:  "",
		Value8:  "",
		Value9:  "",
		Value10: "",
	}
	msgDst := &pkg1.MsgStrings{
		Value1:  "A",
		Value2:  "B",
		Value3:  "C",
		Value4:  "D",
		Value5:  "E",
		Value6:  "F",
		Value7:  "G",
		Value8:  "H",
		Value9:  "I",
		Value10: "J",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{
		0xff, 0x03, 0x01, 0x41, 0x01, 0x42, 0x01, 0x43, 0x01, 0x44, 0x01, 0x45, 0x01, 0x46, 0x01, 0x47, 0x01, 0x48,
		0x01, 0x49, 0x01, 0x4a,
	}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}

func TestPatchStrings3(t *testing.T) {
	requireT := require.New(t)

	m := pkg1.NewMarshaller()

	msgSrc := &pkg1.MsgStrings{
		Value1:  "A",
		Value2:  "",
		Value3:  "C",
		Value4:  "",
		Value5:  "E",
		Value6:  "",
		Value7:  "G",
		Value8:  "",
		Value9:  "I",
		Value10: "",
	}
	msgDst := &pkg1.MsgStrings{
		Value1:  "A",
		Value2:  "B",
		Value3:  "C",
		Value4:  "D",
		Value5:  "E",
		Value6:  "F",
		Value7:  "G",
		Value8:  "H",
		Value9:  "I",
		Value10: "J",
	}

	requireT.False(m.IsPatchNeeded(msgSrc, msgSrc))
	requireT.False(m.IsPatchNeeded(msgDst, msgDst))
	requireT.True(m.IsPatchNeeded(msgSrc, msgDst))
	requireT.True(m.IsPatchNeeded(msgDst, msgSrc))

	msgSrc2 := *msgSrc
	b := make([]byte, maxPatchSize)
	id, l, err := m.MakePatch(msgDst, &msgSrc2, b)
	requireT.NoError(err)
	requireT.Equal([]byte{0xaa, 0x02, 0x01, 0x42, 0x01, 0x44, 0x01, 0x46, 0x01, 0x48, 0x01, 0x4a}, b[:l])

	l2, err := m.ApplyPatch(msgSrc, b[:l])
	requireT.NoError(err)
	requireT.Equal(l, l2)
	requireT.Equal(msgDst, msgSrc)

	id2, err := m.ID(msgDst)
	requireT.NoError(err)
	requireT.Equal(id2, id)
}
