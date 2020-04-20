package testcase

import (
	"fmt"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector_String(t *testing.T) {
	as := assert.New(t)

	testCases := []struct {
		off    flatbuffers.VOffsetT
		str    string
		name   string
		size   int
		vector bool
		finish bool
	}{
		{0, "1", "13:31", 8, true, true},              // incorrect location name
		{0, "1234", "13:31", 12, true, true},          // incorrect location name
		{0, "1234567", "13:31", 12, true, true},       // incorrect location name
		{0, "12345678123456", "7:31", 20, true, true}, // should be 07:31
		{0, "12345678123456781", "18:08", 24, true, true},
		{0, "Australia/Sydney123456722", "18:08", 32, true, true},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("%s in %s", tc.str, tc.name), func(t *testing.T) {
			t.Parallel()
			v := flatbuffers.NewVector().String(tc.off, tc.str)
			as.Equal(v.ByteSize(), tc.size)
			as.Equal(v.IsVector(), tc.vector)
			as.Equal(v.IsFinished(), tc.finish)

			// fmt.Printf("%b\n", v.Byte())
			// fmt.Println(v.TypeName())
		})

	}
}

func TestVector_Int8Slice(t *testing.T) {
	as := assert.New(t)
	sl := []int8{0x1, 0x1, 0x3, 0x3, 0x3}
	v := flatbuffers.NewVector().Int8Slice(0, sl)

	as.Equal(v.ByteSize(), 12)
	as.Equal(v.ObjectSize(), len(sl))
	as.True(v.IsVector())
	as.True(v.IsFinished())
	// fmt.Printf("%b\n", v.Byte())
	// fmt.Println(v.TypeName())
}

func TestVector_ByteSlice(t *testing.T) {
	as := assert.New(t)

	v := flatbuffers.NewVector()
	v.ByteSlice(0, []byte("1234"))

	// fmt.Println(v.Size())
	as.Equal(v.ObjectSize(), 4)
	as.True(v.IsVector())
	as.True(v.IsFinished())
	// fmt.Printf("%b\n", v.Byte())
	// fmt.Println(v.TypeName())
}

func TestVector_Array_Scalar_Bool(t *testing.T) {
	as := assert.New(t)
	v1 := flatbuffers.NewScalar().Bool(true)
	v2 := flatbuffers.NewScalar().Bool(true)
	v3 := flatbuffers.NewScalar().Bool(true)
	v4 := flatbuffers.NewScalar().Bool(true)
	// v5 := flatbuffers.NewScalar().Bool(true)
	l := []flatbuffers.VField{v1, v2, v3, v4}
	v := flatbuffers.NewVector().ScalarArray(0, l...)
	// if !v.IsFinished() {
	// 	v.Press()
	// }
	_, total := flatbuffers.Prepad(4, 4+len(l))
	as.Equal(v.ByteSize(), total)
	as.Equal(v.ObjectSize(), len(l))
	as.True(v.IsVector())
	as.True(v.IsFinished())
	// fmt.Printf("%04b\n", v.Payload())
}

func TestVector_Array_Scalar_int32(t *testing.T) {
	as := assert.New(t)
	v1 := flatbuffers.NewScalar().Int32(1)

	// v5 := flatbuffers.NewScalar().Bool(true)
	l := []flatbuffers.VField{v1}
	v := flatbuffers.NewVector().ScalarArray(0, l...)

	_, total := flatbuffers.Prepad(4, 4+len(l)*4)
	as.Equal(v.ByteSize(), total)
	as.Equal(v.ObjectSize(), len(l))
	as.True(v.IsVector())
	as.True(v.IsFinished())
	// fmt.Printf("%0b\n", v.Payload())
}

func TestVector_Array_Scalar_int64(t *testing.T) {
	as := assert.New(t)
	v1 := flatbuffers.NewScalar().Int64(1)

	// v5 := flatbuffers.NewScalar().Bool(true)
	l := []flatbuffers.VField{v1, v1}
	v := flatbuffers.NewVector().ScalarArray(0, l...)

	_, total := flatbuffers.Prepad(4, 4+len(l)*8)
	as.Equal(v.ByteSize(), total)
	as.Equal(v.ObjectSize(), len(l))
	as.True(v.IsVector())
	as.True(v.IsFinished())
	// fmt.Printf("%0b\n", v.Payload())
}

func TestVector_Array_Scalar_Byte(t *testing.T) {
	as := assert.New(t)
	v1 := flatbuffers.NewScalar().Byte(1)
	// v2 := flatbuffers.NewScalar().Byte(1)
	// v3 := flatbuffers.NewScalar().Byte(1)
	// v4 := flatbuffers.NewScalar().Byte(1)
	// v5 := flatbuffers.NewScalar().Bool(true)
	l := []flatbuffers.VField{v1, v1, v1, v1} //   v2, v3, v4}
	v := flatbuffers.NewVector().ScalarArray(0, l...)

	_, total := flatbuffers.Prepad(4, 4+len(l)*1)
	as.Equal(v.ByteSize(), total)
	as.Equal(v.ObjectSize(), len(l))
	as.True(v.IsVector())
	as.True(v.IsFinished())
	// fmt.Printf("%04b\n", v.Payload())
}

func TestVector_Array_Int8Slice(t *testing.T) {
	as := assert.New(t)

	v1 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
	v2 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
	v3 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
	v4 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
	v5 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})

	v := flatbuffers.NewVector().VectcorArray(0, []flatbuffers.VField{v1, v2, v3, v4, v5}...)
	if !v.IsFinished() {
		v.Press()
	}

	as.Equal(v.ByteSize(), 44)
	as.Equal(v.ObjectSize(), 5)
	as.True(v.IsVector())
	as.True(v.IsFinished())
}

func TestVector_Array_Srruct(t *testing.T) {
	as := assert.New(t)

	v1 := CreateVec3(1.1, 1.1, 1.2, ColorGreen)
	// v2 := vec.CreateVec3(1, 1.1, 1, true, vec.ColorBlue)
	// v3 := vec.CreateVec3(1, 1.1, 1, true, vec.ColorBlue)
	// v4 := vec.CreateVec3(1, 1.1, 1, true, vec.ColorBlue)
	// v5 := vec.CreateVec3(1, 1.1, 1, true, vec.ColorBlue)
	l := []flatbuffers.VField{v1, v1} //  v2, v3, v4, v5}
	// fmt.Printf("%4b\n", v1.FinishByte())
	v := flatbuffers.NewVector().StructArray(0, l...)

	as.Equal(v.ByteSize(), 4+len(l)*v1.ByteSize())

	// fmt.Printf("%b\n", v.Payload())

	as.True(v.IsVector())
	as.True(v.IsFinished())
}

func BenchmarkVector_ByteSlice(b *testing.B) {
	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			flatbuffers.NewVector().ByteSlice(0, []byte("1234"))
		}
	})
}

func BenchmarkVector_String(b *testing.B) {
	testCases := []struct {
		off    flatbuffers.VOffsetT
		str    string
		name   string
		size   int
		vector bool
		finish bool
	}{
		{0, "Europe/Zuri", "13:31", 16, true, true},     // incorrect location name
		{0, "America/New_York", "7:31", 24, true, true}, // should be 07:31
		{0, "Australia/Sydney", "18:08", 24, true, true},
		{0, "Australia/Sydney", "18:08", 24, true, true},
		{0, "Australia/Sydney1234", "18:08", 28, true, true},
		{0, "Australia/Sydney123456722", "18:08", 32, true, true},
	}
	for _, tc := range testCases {
		tc := tc
		b.Run(fmt.Sprintf("%s in %s", tc.str, tc.name), func(b *testing.B) {
			b.SetParallelism(runtime.NumCPU() * 4)
			b.StartTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				flatbuffers.NewVector().String(tc.off, tc.str)
			}
		})
	}

	// b.RunParallel(func(pb *testing.PB) {
	// 	for pb.Next() {
	// 		flatbuffers.NewVector().ByteSlice(0, []byte("1234"))
	// 	}
	// })
}

func BenchmarkVector_String1(b *testing.B) {
	testCases := []struct {
		off    flatbuffers.VOffsetT
		str    string
		name   string
		size   int
		vector bool
		finish bool
	}{
		{0, "Europe/Zuri", "13:31", 16, true, true},     // incorrect location name
		{0, "America/New_York", "7:31", 24, true, true}, // should be 07:31
		{0, "Australia/Sydney", "18:08", 24, true, true},
		{0, "Australia/Sydney", "18:08", 24, true, true},
		{0, "Australia/Sydney1234", "18:08", 28, true, true},
		{0, "Australia/Sydney123456722", "18:08", 32, true, true},
	}
	for _, tc := range testCases {
		tc := tc
		b.Run(fmt.Sprintf("%s in %s", tc.str, tc.name), func(b *testing.B) {
			b.SetParallelism(runtime.NumCPU() * 4)
			b.StartTimer()
			b.ReportAllocs()

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					flatbuffers.NewVector().String(tc.off, tc.str)
				}
			})
		})

	}

	// b.RunParallel(func(pb *testing.PB) {
	// 	for pb.Next() {
	// 		flatbuffers.NewVector().ByteSlice(0, []byte("1234"))
	// 	}
	// })
}

func BenchmarkVector_Array_String(b *testing.B) {
	v1 := flatbuffers.NewVector().String(0, "1234567")
	v2 := flatbuffers.NewVector().String(0, "1234567")
	v3 := flatbuffers.NewVector().String(0, "1234567")
	v4 := flatbuffers.NewVector().String(0, "1234567")
	v5 := flatbuffers.NewVector().String(0, "1234567")

	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			flatbuffers.NewVector().VectcorArray(0, []flatbuffers.VField{v1, v2, v3, v4, v5}...)
		}
	})
}

func BenchmarkVector_Array_ByteSlice(b *testing.B) {
	v1 := flatbuffers.NewVector().ByteSlice(0, []byte{0x1, 0x1, 0x3, 0x1})
	v2 := flatbuffers.NewVector().ByteSlice(0, []byte{0x1, 0x1, 0x3, 0x1})
	v3 := flatbuffers.NewVector().ByteSlice(0, []byte{0x1, 0x1, 0x3, 0x1})
	v4 := flatbuffers.NewVector().ByteSlice(0, []byte{0x1, 0x1, 0x3, 0x1})
	v5 := flatbuffers.NewVector().ByteSlice(0, []byte{0x1, 0x1, 0x3, 0x1})

	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			flatbuffers.NewVector().VectcorArray(0, []flatbuffers.VField{v1, v2, v3, v4, v5}...)
		}
	})
}

func BenchmarkVector_Array_Int8Slice(b *testing.B) {
	v1 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
	v2 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
	v3 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
	v4 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
	v5 := flatbuffers.NewVector().Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})

	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			flatbuffers.NewVector().VectcorArray(0, []flatbuffers.VField{v1, v2, v3, v4, v5}...)
		}
	})
}

func BenchmarkVector_StringParallel(b *testing.B) {
	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			flatbuffers.NewVector().String(0, "1234567")
		}
	})
}

func BenchmarkVector_Int8Slice(b *testing.B) {
	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v := flatbuffers.NewVector()
			v.Int8Slice(0, []int8{0x1, 0x1, 0x3, 0x1})
		}
	})
}
