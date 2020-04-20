package testcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

func TestColor_String(t *testing.T) {
	as := assert.New(t)
	s := "123"
	v := flatbuffers.NewVector().String(0, s)
	_, total := flatbuffers.Prepad(flatbuffers.SizeSOffsetT, len(s)+1+flatbuffers.SizeSOffsetT)
	// fmt.Printf("%b\n", v.FinishByte())
	// fmt.Println(total, len(s))
	as.Equal(v.ObjectSize(), len(s))
	as.Equal(v.ByteSize(), total)
}

func TestColor_Strings(t *testing.T) {
	as := assert.New(t)

	v1 := flatbuffers.NewVector().String(0, "111")
	v2 := flatbuffers.NewVector().String(0, "12345")
	v3 := flatbuffers.NewVector().String(0, "333")

	v := flatbuffers.NewVector().VectcorArray(0, []flatbuffers.VField{v1, v2, v3}...)
	// fmt.Printf("%b\n", v.FinishByte())
	as.Equal(v.ObjectSize(), 3)
	// as.Equal(v.ByteSize(), 20)
	as.Equal(v.ByteSize(), len(v.Payload()))

	v = flatbuffers.NewVector().VectcorArray(0, []flatbuffers.VField{v3, v2, v1}...)
	v = flatbuffers.NewVector().VectcorArray(0, []flatbuffers.VField{v3, v2, v1}...)
	// fmt.Printf("%b\n", v.FinishByte())
	as.Equal(v.ObjectSize(), 3)
	// as.Equal(v.ByteSize(), 20)
	as.Equal(v.ByteSize(), len(v.Payload()))
}

func BenchmarkColor_String(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			flatbuffers.NewVector().String(0, "1").Payload()
		}
	})
}
