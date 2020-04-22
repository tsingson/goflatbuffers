package flatbuffers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	as := assert.New(t)

	s := NewStack(32)

	v := NewScalar().Float64(100.01)
	s.Push(v)
	v1 := NewScalar().Float64(200.01)
	s.Push(v1)
	fmt.Println(s.NumsElement())

	as.Equal(s.Size(), 32)
	v1, chk := s.Pop()
	if chk {
		as.Equal(v1.ByteSize(), 8)
		fmt.Println(v1.ByteSize())
		fmt.Println(GetFloat64(v1.FinishByte()))
	}
	fmt.Println(s.NumsElement())
	v2, chk2 := s.Pop()
	if chk2 {
		as.Equal(v2.ByteSize(), 8)
		fmt.Println(v2.ByteSize())
		fmt.Println(GetFloat64(v2.FinishByte()))
	}
	fmt.Println(s.NumsElement())
}

func BenchmarkNewStack(b *testing.B) {
	s := NewStack(32)
	v1 := NewScalar().Float64(200.01)
	b.SetParallelism(8)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Push(v1)
			s.Pop()
		}
	})
}
