package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/mass"
	"github.com/outofforest/proton/test/custom"
	"github.com/outofforest/proton/test/pkg1"
	spkg1 "github.com/outofforest/proton/test/pkg1/spkg"
)

func TestDefault(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMixed{}

	requireT.EqualValues(18, msg1.Size())

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)
	requireT.Equal([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00}, b)

	msg2 := pkg1.MsgMixed{
		Value1: map[string]spkg1.SubMsg{
			"aa": {
				Value: 1,
			},
		},
		Value2: map[uint8][]string{
			4: {"aa", "bb"},
		},
		Value3: [][32]uint16{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
				29, 30, 31},
		},
		Value4: [12]map[int8]float32{
			{0: 1.}, {0: 1.}, {0: 1.}, {0: 1.}, {0: 1.}, {0: 1.}, {0: 1.}, {0: 1.}, {0: 1.}, {0: 1.}, {0: 1.}, {0: 1.},
		},
		Value5: [][3][]map[int16][2]int64{
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
		Value7: true,
		Value8: "fdfsd",
	}
	l = msg2.Unmarshal(b, mass.New[string](10), mass.New[[32]uint16](10),
		mass.New[[3][]map[int16][2]int64](10), mass.New[map[int16][2]int64](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(pkg1.MsgMixed{
		Value1: nil,
		Value2: nil,
		Value3: nil,
		Value4: [12]map[int8]float32{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		Value5: nil,
	}, msg2)
}

func Test1(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMixed{
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

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)

	var msg2 pkg1.MsgMixed
	l = msg2.Unmarshal(b, mass.New[string](10), mass.New[[32]uint16](10),
		mass.New[[3][]map[int16][2]int64](10), mass.New[map[int16][2]int64](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}

func TestCustom(t *testing.T) {
	requireT := require.New(t)

	msg1 := pkg1.MsgMixedCustom{
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

	b := make([]byte, msg1.Size())
	l := msg1.Marshal(b)
	requireT.Equal(msg1.Size(), l)

	var msg2 pkg1.MsgMixedCustom
	l = msg2.Unmarshal(b, mass.New[[3][]map[int16]custom.Array](10),
		mass.New[map[int16]custom.Array](10))
	requireT.Equal(msg1.Size(), l)

	requireT.Equal(msg1, msg2)
}
