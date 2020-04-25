package main

import (
	"fmt"
	"github.com/tsingson/goflatbuffers/go/bytepool"
	"math/bits"
)

func main() {
	g := bytepool.Get(66)
	fmt.Println(cap(g))
	fmt.Println(len(g))
	copy(g[0:], []byte("1238"))
	fmt.Printf("%b\n", g)
	// bytepool.CleanPut(g )
	bytepool.Put(g)
	g = bytepool.Get(6)
	fmt.Println(cap(g))
	fmt.Println(len(g))
	// copy ( g[0:], []byte("12345678"))
	fmt.Printf("%b\n", g)
	buf := bytepool.NewByteBuffer(129)
	buf = bytepool.NewByteBuffer(245)
	copy(buf.B[0:], []byte("1238"))
	fmt.Println(buf.Cap())
	fmt.Println(buf.Len())
	buf.Reset(111)
	copy(buf.B[0:], []byte("1238"))
	// copy(buf.B[0:], []byte("1238"))
	fmt.Printf("%b\n", buf.Bytes())
	buf1 := bytepool.NewByteBuffer(64)

	// copy(buf1.B[0:], []byte("1238"))
	fmt.Println(buf1.Cap())
	fmt.Println(buf1.Len())
	fmt.Printf("%b\n", buf1.Bytes())

	v := uint32(bits.Len32(512 - 1))
	fmt.Println("v -> ", v, " = ", 1<<9)
}
