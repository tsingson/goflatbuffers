package main

/*
#include <stdlib.h>
#include <float.h>

float *reverse(float *f, int len) {
    float *g = calloc(len, sizeof(float));
    for (int i = 0; i < len; i++) {
        g[i] = f[len-1-i];
    }
    return g;
}
*/
import "C"

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	a := []float32{3.14159, 2.718, 1, 2}
	r := make([]float32, len(a))
	fmt.Println("a:", a, "r:", r)
	c := C.reverse((*C.float)(&a[0]), C.int(len(a)))
	copy(r, (*[1 << 20]float32)(unsafe.Pointer(c))[:])
	C.free(unsafe.Pointer(c))
	fmt.Println("a:", a, "r:", r)
}

var okCFloat = func() bool {
	if C.sizeof_float != unsafe.Sizeof(float32(0)) {
		panic("C float != Go float32")
	}
	if C.FLT_MAX != math.MaxFloat32 {
		panic("C float != Go float32")
	}
	return true
}()
