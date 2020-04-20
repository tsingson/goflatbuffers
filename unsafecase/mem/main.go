package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

type Foo struct {
	a int
}

func main() {
	debug.SetGCPercent(-1)

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)
	var g sync.WaitGroup
	g.Add(2)
	go func() {
		fmt.Println("-------------------------------------------")
		for i := 0; i < 8; i++ {
			f := NewFoo(i)
			_ = fmt.Sprintf("%d", f.a)
		}
		time.Sleep(1 * time.Second)
		g.Done()
	}()

	go func() {
		fmt.Println("=========================================")
		for i := 0; i < 8; i++ {
			f := NewFoo(i)
			_ = fmt.Sprintf("%d", f.a)
		}
		time.Sleep(1 * time.Second)
		g.Done()
	}()

	fmt.Printf("-----------: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(300 * time.Millisecond)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(300 * time.Millisecond)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1024), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)
	g.Wait()
}

//go:noinline
func NewFoo(i int) *Foo {
	f := &Foo{a: rand.Intn(50)}
	runtime.SetFinalizer(f, func(f *Foo) {
		fmt.Println("foo >>>>>>> " + strconv.Itoa(i) + " has been garbage collected")
	})

	return f
}
