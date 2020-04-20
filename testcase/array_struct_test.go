package testcase

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

func TestVector_Array_Struct(t *testing.T) {
	v1 := CreateVec3(1.1, 1.1, 1.2, ColorGreen)
	v2 := CreateVec3(1.2, 1.1, 1.2, ColorGreen)
	v3 := CreateVec3(1.3, 1.1, 1.2, ColorGreen)

	vf := flatbuffers.NewVector().StructArray(0, []flatbuffers.VField{v1, v2, v3}...)
	buf := vf.Payload()

	vec := &Vec3{}
	vec.Init(buf[4:], 0)
	as := assert.New(t)
	as.Equal(vec.X(), float32(1.1))
	vec.Init(buf[4+v1.ByteSize():], 0)
	as.Equal(vec.X(), float32(1.2))
	vec.Init(buf[4+v1.ByteSize()*2:], 0)
	as.Equal(vec.X(), float32(1.3))
}

func BenchmarkVector_ArrayStruct(b *testing.B) {
	v1 := CreateVec3(1.1, 1.1, 1.2, ColorGreen)

	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			flatbuffers.NewVector().StructArray(0, []flatbuffers.VField{v1, v1, v1, v1, v1, v1, v1, v1}...)
		}
	})
}
