package testcase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tsingson/goflatbuffers/go"
)

func TestScalar_Int64(t *testing.T) {
	as := assert.New(t)
	testCases := []struct {
		in   int64
		size int
	}{
		{100, 8},
		{101, 8},
		{102, 8},
		{103, 8},
		{1003, 8},
		{200, 8},
		{300, 8},
		{400, 8},
		{5500, 8},
		{6600, 8},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("%v  in %v ", tc.in, tc.size), func(t *testing.T) {
			t.Parallel()
			v := flatbuffers.NewScalar().Int64(tc.in)
			as.False(v.IsVector())
			as.Equal(v.ByteSize(), tc.size)
			as.True(v.IsFinished())
			as.Equal(flatbuffers.GetInt64(v.FinishedBytes()), int64(tc.in))
		})

	}
}

func TestScalar_Float64(t *testing.T) {
	as := assert.New(t)

	v := flatbuffers.NewScalar().Float64(100.01)
	as.False(v.IsVector())
	as.True(v.IsFinished())
	as.Equal(v.ByteSize(), 8)

	as.Equal(flatbuffers.GetFloat64(v.FinishedBytes()), float64(100.01))
}

func BenchmarkScalar_Float64(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 100; i++ {
				v := flatbuffers.NewScalar().Float64(100.01)
				v.Release()
			}
		}
	})
}

func BenchmarkScalar_Float64_1(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// for i := 0; i < 100; i++ {
			flatbuffers.NewScalar().Float64(100.01).FinishedBytes()
			// }
			// v.Release()
		}
	})
}

func BenchmarkScalar_WriteFloat64(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := make([]byte, flatbuffers.SizeFloat64, flatbuffers.SizeFloat64)
			for i := 0; i < 100; i++ {
				flatbuffers.WriteFloat64(buf, 100.01)
			}
		}
	})
}

func TestNewScalar(t *testing.T) {
	as := assert.New(t)

	v := flatbuffers.NewScalar().Float64(100.01)
	as.Equal(flatbuffers.GetFloat64(v.FinishedBytes()), float64(100.01))
	as.False(v.IsVector())
	as.True(v.IsFinished())
	as.Equal(v.ByteSize(), 8)
}
