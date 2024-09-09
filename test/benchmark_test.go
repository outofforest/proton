package test

import (
	"testing"

	"github.com/outofforest/mass"
	"github.com/outofforest/proton/test/pkg1"
	spkg1 "github.com/outofforest/proton/test/pkg1/spkg"
)

// go test -bench=. -run=^$ -cpuprofile profile.out
// go tool pprof -http="localhost:8000" pprofbin ./profile.out

func BenchmarkValue(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	s := implValue{}

	r := str

	for bi := 0; bi < b.N; bi++ {
		b.StartTimer()
		for range 100000 {
			r = s.Method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkPointer(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	s := &implPointer{}

	r := str

	for bi := 0; bi < b.N; bi++ {
		b.StartTimer()
		for range 100000 {
			r = s.Method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkInterfaceValue(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var s iface = implValue{}

	r := str

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = s.Method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkInterfacePointer(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var s iface = &implPointer{}

	r := str

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = s.Method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkFunction(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	r := str

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkVarFunction(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	r := str

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = methodVar(r)
		}
		b.StopTimer()
	}
}

func BenchmarkFunctionInVar(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	r := str
	f := method

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkValueMethodInVar(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	r := str

	s := implValue{}
	f := s.Method

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkPointerMethodInVar(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	r := str

	s := &implPointer{}
	f := s.Method

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkInterfaceValueMethodInVar(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	r := str

	var s iface = implValue{}
	f := s.Method

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkInterfacePointerMethodInVar(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	r := str

	var s iface = &implPointer{}
	f := s.Method

	for range b.N {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkMarshalingMixed(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var msg2 pkg1.MsgMixed
	buf := make([]byte, msgMixed.Size())

	mass1 := mass.New[string](100)
	mass2 := mass.New[[32]uint16](100)
	mass3 := mass.New[[3][]map[int16][2]int64](100)
	mass4 := mass.New[map[int16][2]int64](100)

	b.StartTimer()
	for range b.N {
		msgMixed.Marshal(buf)
		msg2.Unmarshal(buf, mass1, mass2, mass3, mass4)
	}
	b.StopTimer()
}

func BenchmarkMarshalingSlices(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var msg2 pkg1.MsgSlice
	buf := make([]byte, msgSlice.Size())

	mass := mass.New[bool](1000)

	b.StartTimer()
	for range b.N {
		msgSlice.Marshal(buf)
		msg2.Unmarshal(buf, mass)
	}
	b.StopTimer()
}

func BenchmarkMarshalingByteSlices(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var msg2 pkg1.MsgSliceUint8
	buf := make([]byte, msgBytes.Size())

	b.StartTimer()
	for range b.N {
		msgBytes.Marshal(buf)
		msg2.Unmarshal(buf)
	}
	b.StopTimer()
}

func BenchmarkMarshalingStrings(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var msg2 pkg1.MsgString
	buf := make([]byte, msgStrings.Size())

	b.StartTimer()
	for range b.N {
		msgStrings.Marshal(buf)
		msg2.Unmarshal(buf)
	}
	b.StopTimer()
}

func BenchmarkMarshalingEmptySlices(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var msg pkg1.MsgSlice
	buf := make([]byte, msg.Size())

	mass := mass.New[bool](100)

	b.StartTimer()
	for range b.N {
		msg.Marshal(buf)
		msg.Unmarshal(buf, mass)
	}
	b.StopTimer()
}

func BenchmarkMarshalingEmptyMaps(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var msg pkg1.MsgMap
	buf := make([]byte, msg.Size())

	b.StartTimer()
	for range b.N {
		msg.Marshal(buf)
		msg.Unmarshal(buf)
	}
	b.StopTimer()
}

var (
	msgMixed = pkg1.MsgMixed{
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
			4: {"aafsdfsd", "fdsfsdfdsbb", "fdsfsdfsdfstwsfsdfsdfdsfdsfsd", "fdfsdfsdffdsfds"},
			1: {"fdsfsdfsdfsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdsfsdfsddfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsdfsdfsdfsdffdsfds"},
			8: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
			10: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
			41: {"aafsdfsd", "fdsfsdfdsbb", "fdsfsdfsdfstwsfsdfsdfdsfdsfsd", "fdfsdfsdffdsfds"},
			11: {"fdsfsdfsdfsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdsfsdfsddfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsdfsdfsdfsdffdsfds"},
			81: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
			110: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
			42: {"aafsdfsd", "fdsfsdfdsbb", "fdsfsdfsdfstwsfsdfsdfdsfdsfsd", "fdfsdfsdffdsfds"},
			13: {"fdsfsdfsdfsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdsfsdfsddfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsdfsdfsdfsdffdsfds"},
			84: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
			105: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
			234: {"aafsdfsd", "fdsfsdfdsbb", "fdsfsdfsdfstwsfsdfsdfdsfdsfsd", "fdfsdfsdffdsfds"},
			12: {"fdsfsdfsdfsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdsfsdfsddfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsdfsdfsdfsdffdsfds"},
			85: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
			15: {"fdsfsdsdaafsdfsd", "fdsfsdffsdfsdfdsdsbb", "fdsfsfdfsdfstwsfsdfsdfdsfdsfsd",
				"fdfdsfsfsdfsdfsdfsfsfsdsddfsdfsdfsdffdsfds"},
		},
		Value3: [][32]uint16{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31},
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
	msgBytes = pkg1.MsgSliceUint8{
		Value: []uint8{
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		},
	}
	msgStrings = pkg1.MsgString{
		//nolint:lll
		Value: "fdfsdfdsgfdlghfdkghkdfhkdfhkjghdfkhgkjdfhgkjdfhkjdgfhkjhdfkjhgkjdfhkjdfhkjdfhkjhfdkjhkjdfhkjdfhkjfdghkjhfdkjhgkjdfhkjdfhkjghfdkjhgdfkjhkjfdhkjfhkjghdfkjhdfkhgkjsfysdfydgdfkghdfkjghkfd",
	}
	msgSlice = pkg1.MsgSlice{
		Value: make([]bool, 10),
	}
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
