package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/tsingson/goflatbuffers/go/bytepool"
)

func main() {
	b := bytepool.Get()
	b.FixedLength(4).WriteAt(4, []byte("1234"))
	fmt.Printf("%4b\n", b.Bytes())
	fmt.Println(b.Len())
	fmt.Printf("%v\n", b.String())
	b.Reset()
	b.FixedLength(8).WriteAt(2, []byte("1234"))
	b.WriteAt(7, []byte("9"))
	fmt.Printf("%4b\n", b.Bytes())
	fmt.Println(b.Len())
	fmt.Println(b.Cap())
	fmt.Printf("%v\n", b.String())
	bytepool.Put(b)
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	time.Sleep(1 * time.Second)
	fmt.Println("-------------------------------------")
	b1 := bytepool.Get()

	b1.FixedLength(1).WriteAt(0, []byte("1"))
	b1.WriteAt(4, []byte("12341"))
	fmt.Println(b1.Len())
	b1.Pad(1)
	// b.WriteAt(2, []byte("0"))
	// b.WriteAt(7, []byte("0"))
	fmt.Printf("%4b\n", b1.Bytes())
	fmt.Println(b1.Len())
	fmt.Println(b1.Cap())
}
