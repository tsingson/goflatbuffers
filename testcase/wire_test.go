package testcase

import (
	"runtime"
	"testing"
	"unsafe"
)

func BenchmarkColor_Wire(b *testing.B) {

	b.ReportAllocs()
	b.SetParallelism(runtime.NumCPU() * 2)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 1000; i++ {
				for j := 0; j < 16; j++ {
					buf := []byte{0x0, 0x0, 0x0, 0x0}
					WriteUint8(buf, uint8(j))
				}
			}
		}
	})
}
func BenchmarkColor_WireUnsafe(b *testing.B) {

	b.ReportAllocs()
	b.SetParallelism(runtime.NumCPU() * 2)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 1000; i++ {
				for j := 0; j < 16; j++ {
					buf := []byte{0x0, 0x0, 0x0, 0x0}
					WriteUint8Unsafe(buf, uint8(j))
				}

			}
		}
	})
}

func WriteUint8Unsafe(buf []byte, n uint8) {
	*(*uint8)(unsafe.Pointer(&buf[0])) = n
}

// WriteUint8 encodes a little-endian uint8 into a byte slice.
func WriteUint8(buf []byte, n uint8) {
	buf[0] = byte(n)
}
