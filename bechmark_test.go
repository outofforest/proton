package proton

import (
	"testing"

	"github.com/outofforest/proton/test/pkg1"
	spkg1 "github.com/outofforest/proton/test/pkg1/spkg"
)

var _ iface = implValue{}
var _ iface = &implPointer{}

type iface interface {
	Method(arg string) string
}

type implValue struct {
}

func (i implValue) Method(arg string) string {
	return arg
}

type implPointer struct {
}

func (i *implPointer) Method(arg string) string {
	return arg
}

func method(arg string) string {
	return arg
}

var methodVar = func(arg string) string {
	return arg
}

var str = ""

func BenchmarkValue(b *testing.B) {
	s := implValue{}

	r := str

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = s.Method(r)
	}
	b.StopTimer()
}

func BenchmarkPointer(b *testing.B) {
	s := &implPointer{}

	r := str

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = s.Method(r)
	}
	b.StopTimer()
}

func BenchmarkInterfaceValue(b *testing.B) {
	var s iface = implValue{}

	r := str

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = s.Method(r)
	}
	b.StopTimer()
}

func BenchmarkInterfacePointer(b *testing.B) {
	var s iface = &implPointer{}

	r := str

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = s.Method(r)
	}
	b.StopTimer()
}

func BenchmarkVarFunction(b *testing.B) {
	r := str

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = methodVar(r)
	}
	b.StopTimer()
}

func BenchmarkFunction(b *testing.B) {
	r := str

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = method(r)
	}
	b.StopTimer()
}

func BenchmarkFunctionInVar(b *testing.B) {
	r := str
	f := method

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = f(r)
	}
	b.StopTimer()
}

func BenchmarkValueMethodInVar(b *testing.B) {
	r := str

	s := implValue{}
	f := s.Method

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = f(r)
	}
	b.StopTimer()
}

func BenchmarkPointerMethodInVar(b *testing.B) {
	r := str

	s := &implPointer{}
	f := s.Method

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = f(r)
	}
	b.StopTimer()
}

func BenchmarkInterfaceValueMethodInVar(b *testing.B) {
	r := str

	var s iface = implValue{}
	f := s.Method

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = f(r)
	}
	b.StopTimer()
}

func BenchmarkInterfacePointerMethodInVar(b *testing.B) {
	r := str

	var s iface = &implPointer{}
	f := s.Method

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		r = f(r)
	}
	b.StopTimer()
}

var msg = pkg1.MsgMixed{
	Value1: map[string]spkg1.SubMsg{
		"aa": {
			Value: 143443,
		},
		"fdfsdsd": {
			Value: 433,
		},
		"frewrwerwerwedfsdsd": {
			Value: -433,
		},
		"dfsdfsd": {
			Value: 100000,
		},
		"dsjklfjdsklfjdslkjflsdkfsfjlskdjflskdjfkldsjklfsdjklfjsdlkslioisufsdlf": {
			Value: 43423423,
		},
	},
	Value2: map[uint8][]string{
		4:   {"aafsdfsd", "fdsfsdfdsbb", "fdsfsdfsdfstwsfsdfsdfdsfdsfsd", "fdfsdfsdffdsfds"},
		1:   {"fdsfsdfsdfsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdsfsdfsddfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsdfsdfsdfsdffdsfds"},
		8:   {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
		10:  {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
		41:  {"aafsdfsd", "fdsfsdfdsbb", "fdsfsdfsdfstwsfsdfsdfdsfdsfsd", "fdfsdfsdffdsfds"},
		11:  {"fdsfsdfsdfsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdsfsdfsddfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsdfsdfsdfsdffdsfds"},
		81:  {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
		110: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
		42:  {"aafsdfsd", "fdsfsdfdsbb", "fdsfsdfsdfstwsfsdfsdfdsfdsfsd", "fdfsdfsdffdsfds"},
		13:  {"fdsfsdfsdfsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdsfsdfsddfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsdfsdfsdfsdffdsfds"},
		84:  {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
		105: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
		234: {"aafsdfsd", "fdsfsdfdsbb", "fdsfsdfsdfstwsfsdfsdfdsfdsfsd", "fdfsdfsdffdsfds"},
		12:  {"fdsfsdfsdfsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdsfsdfsddfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsdfsdfsdfsdffdsfds"},
		85:  {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
		15:  {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd", "fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
	},
	Value3: [][32]uint16{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
	},
	Value4: [12]map[int8]float32{
		{0: 1., 2: 5., 10: 11., 12: 56, 89: 100, 54: 43, 125: 435},
		{0: 1., 2: 5., 10: 11., 12: 56, 89: 100, 54: 43, 125: 435},
		{0: 1., 2: 5., 10: 11., 12: 56, 89: 100, 54: 43, 125: 435},
		{0: 1., 2: 5., 10: 11., 12: 56, 89: 100, 54: 43, 125: 435},
		{0: 1., 2: 5., 10: 11., 12: 56, 89: 100, 54: 43, 125: 435},
		{0: 1., 2: 5., 10: 11., 12: 56, 89: 100, 54: 43, 125: 435},
	},
	Value5: [][3][]map[int16][2]int64{
		{
			{
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
			},
			{
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
			},
			{
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
				{2: {1, 2}, 4: {14321, 553234234232}, 40: {14321, 553234234232}},
			},
		},
	},
	Value6: true,
	Value7: true,
	Value8: "fdfsd",
}

func BenchmarkMarshaling(b *testing.B) {
	buf := make([]byte, msg.Size())

	var msg2 pkg1.MsgMixed

	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		msg.Marshal(buf)
		msg2.Unmarshal(buf)
	}
	b.StopTimer()
}
