package main

import "fmt"

type MyStruct struct {
	a int
	b int
}

type Decorator func(m *MyStruct) *MyStruct

func (f Decorator) WithA(a int) Decorator {
	return func(m *MyStruct) *MyStruct {
		f(m).a = a
		return m
	}
}

func (f Decorator) WithB(b int) Decorator {
	return func(m *MyStruct) *MyStruct {
		f(m).b = b
		return m
	}
}

func NewMyType() Decorator {
	return func(m *MyStruct) *MyStruct {
		return m
	}
}

func (f Decorator) Build() *MyStruct {
	return f(&MyStruct{})
}

func main() {
	instance := NewMyType().
		WithA(42).
		WithB(123).
		Build()
	fmt.Println(instance)
}
