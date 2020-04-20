package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/sanity-io/litter"

	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

func main() {
	u := int64(1)

	b := Int16ToByte(u)

	fmt.Printf("%b\n", b[0:])
	// fmt.Printf("%v\n", b)
	fmt.Printf("%4b\n", u)
	fmt.Println(u, flatbuffers.GetInt16(b), Byte2Int16(b[:2]))
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println(unsafe.Pointer(&u))

	a := &MyStruct{
		A: 1,
		B: 2,
	}
	c := BytesToMyStruct(MyStructToBytes(a))
	litter.Dump(a, c)
}

func Byte2Int16(b []byte) int64 {
	// b[0], b[1] = b[1], b[0]
	return *(*int64)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}

func Int16ToByte(u int64) []byte {
	sizeOfMyStruct := 8
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(&u))
	return *(*[]byte)(unsafe.Pointer(&x))
}

type MyStruct struct {
	A int32
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}
