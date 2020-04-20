package testcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec3_Init(t *testing.T) {
	as := assert.New(t)
	v := CreateVec3(1.1, 1.1, 1.2, ColorGreen)
	v1 := new(Vec3)
	v1.Init(v.Payload(), 0)
	as.Equal(v1.X(), float32(1.1))
	as.Equal(v1.C(), ColorGreen)
	as.Equal(v.ObjectSize(), 0)
	as.Equal(v.ByteSize(), 16)
}

func BenchmarkCreateVec3(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			CreateVec3(1.1, 1.1, 1.2, ColorBlue).Payload()
		}
	})
}
