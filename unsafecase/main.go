package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	// i 就是后面说的 数据对象
	i := 3
	// 这里的func 就是后面一直说的 finalizer
	runtime.GC()

	runtime.SetFinalizer(&i, func(i *int) {
		fmt.Println(i, *i, "set finalizer")
	})
	time.Sleep(time.Second * 5)
}

func case4() {
	// case3()

	buf := []byte{1}

	n := uint8(buf[0])

	fmt.Println("-- ", unsafe.Pointer(&buf))
	fmt.Println("-- ", unsafe.Pointer(&n))
	fmt.Println("-- ", buf)
	fmt.Printf("-- %b\n", buf)
	fmt.Println("-- ", n)
	fmt.Println("-- ", *(*int8)(unsafe.Pointer((&buf))))
}

func case3() {
	var u1 user1

	fmt.Println(unsafe.Offsetof(u1.i)) // 4
	fmt.Println(unsafe.Offsetof(u1.j)) // 8

	fmt.Println(unsafe.Offsetof(u1.b)) // 0
	fmt.Println(unsafe.Sizeof(u1))     // 16
}

type user1 struct {
	i int32
	j int64
	b byte
}

func case2() {
	var b bool
	var i8 int8
	var i16 int16
	var i64 int64

	var f32 float32

	var s string

	var m map[string]string

	var p *int32

	fmt.Println(unsafe.Alignof(b))   // 1
	fmt.Println(unsafe.Alignof(i8))  // 1
	fmt.Println(unsafe.Alignof(i16)) // 2
	fmt.Println(unsafe.Alignof(i64)) // 8
	fmt.Println(unsafe.Alignof(f32)) // 4
	fmt.Println(unsafe.Alignof(s))   // 8
	fmt.Println(unsafe.Alignof(m))   // 8
	fmt.Println(unsafe.Alignof(p))   // 8
}

func case1() {
	var cr creator
	pointer := cr.CreateInt()
	fmt.Printf("%[1]T %[1]p\n", pointer) // *int
}

type creator struct {
	buf []interface{}
}

func (cr *creator) Create() *interface{} {
	if len(cr.buf) == 0 {
		cr.buf = make([]interface{}, 256)
	}
	current := &cr.buf[0]
	cr.buf = cr.buf[1:]
	return current
}

func (cr *creator) CreateInt() *int {
	pointer := cr.Create()
	fmt.Printf("%[1]T %[1]p\n", pointer)   // *interface {}
	return (*int)(unsafe.Pointer(pointer)) // *interface {} to *int
}
