package test

import (
	"fmt"
	"io"
	"testing"

	"github.com/outofforest/proton/test/pkg1"
	spkg1 "github.com/outofforest/proton/test/pkg1/spkg"
)

// go test -bench=. -run=^$ -cpuprofile profile.out
// go tool pprof -http="localhost:8000" pprofbin ./profile.out

func BenchmarkValue(b *testing.B) {
	s := implValue{}

	r := str

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = s.Method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkPointer(b *testing.B) {
	s := &implPointer{}

	r := str

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = s.Method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkInterfaceValue(b *testing.B) {
	var s iface = implValue{}

	r := str

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = s.Method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkInterfacePointer(b *testing.B) {
	var s iface = &implPointer{}

	r := str

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = s.Method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkFunction(b *testing.B) {
	r := str

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = method(r)
		}
		b.StopTimer()
	}
}

func BenchmarkVarFunction(b *testing.B) {
	r := str

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = methodVar(r)
		}
		b.StopTimer()
	}
}

func BenchmarkFunctionInVar(b *testing.B) {
	r := str
	f := method

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkValueMethodInVar(b *testing.B) {
	r := str

	s := implValue{}
	f := s.Method

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkPointerMethodInVar(b *testing.B) {
	r := str

	s := &implPointer{}
	f := s.Method

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkInterfaceValueMethodInVar(b *testing.B) {
	r := str

	var s iface = implValue{}
	f := s.Method

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkInterfacePointerMethodInVar(b *testing.B) {
	r := str

	var s iface = &implPointer{}
	f := s.Method

	for b.Loop() {
		b.StartTimer()
		for range 100000 {
			r = f(r)
		}
		b.StopTimer()
	}
}

func BenchmarkMarshalingMixed(b *testing.B) {
	m := pkg1.NewMarshaller()

	var msg2 any

	size, _ := m.Size(msgMixed)
	buf := make([]byte, size)

	for b.Loop() {
		id, _, _ := m.Marshal(msgMixed, buf)
		msg2, _, _ = m.Unmarshal(id, buf)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, msg2)
}

func BenchmarkMarshalingSlices(b *testing.B) {
	m := pkg1.NewMarshaller()

	var msg2 any
	size, _ := m.Size(msgSlice)
	buf := make([]byte, size)

	for b.Loop() {
		id, _, _ := m.Marshal(msgSlice, buf)
		msg2, _, _ = m.Unmarshal(id, buf)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, msg2)
}

func BenchmarkMarshalingByteSlices(b *testing.B) {
	m := pkg1.NewMarshaller()

	var msg2 any
	size, _ := m.Size(msgBytes)
	buf := make([]byte, size)

	for b.Loop() {
		id, _, _ := m.Marshal(msgBytes, buf)
		msg2, _, _ = m.Unmarshal(id, buf)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, msg2)
}

func BenchmarkMarshalingStrings(b *testing.B) {
	m := pkg1.NewMarshaller()

	var msg2 any
	size, _ := m.Size(msgStrings)
	buf := make([]byte, size)

	for b.Loop() {
		id, _, _ := m.Marshal(msgStrings, buf)
		msg2, _, _ = m.Unmarshal(id, buf)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, msg2)
}

func BenchmarkMarshalingEmptySlices(b *testing.B) {
	m := pkg1.NewMarshaller()

	var msg2 any

	msg := &pkg1.MsgSlice{}
	size, _ := m.Size(msg)
	buf := make([]byte, size)

	for b.Loop() {
		id, _, _ := m.Marshal(msg, buf)
		msg2, _, _ = m.Unmarshal(id, buf)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, msg2)
}

func BenchmarkMarshalingEmptyMaps(b *testing.B) {
	m := pkg1.NewMarshaller()

	var msg2 any

	msg := &pkg1.MsgMap{}
	size, _ := m.Size(msg)
	buf := make([]byte, size)

	for b.Loop() {
		id, _, _ := m.Marshal(msg, buf)
		msg2, _, _ = m.Unmarshal(id, buf)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, msg2)
}

var (
	msgMixed = &pkg1.MsgMixed{
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
	msgBytes = &pkg1.MsgSliceUint8{
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
	msgStrings = &pkg1.MsgString{
		//nolint:lll
		Value: "fdfsdfdsgfdlghfdkghkdfhkdfhkjghdfkhgkjdfhgkjdfhkjdgfhkjhdfkjhgkjdfhkjdfhkjdfhkjhfdkjhkjdfhkjdfhkjfdghkjhfdkjhgkjdfhkjdfhkjghfdkjhgdfkjhkjfdhkjfhkjghdfkjhdfkhgkjsfysdfydgdfkghdfkjghkfd",
	}
	msgSlice = &pkg1.MsgSlice{
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
