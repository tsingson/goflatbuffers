package stack

import (
	"sync"
)

// Stacker is an interface describing the behaviour of a FILO (first in, last out) stack. It allows concurrency-safe
// stacks to be used in the same places as regular stacks, if performance or concurrency safety are specific
// requirements.
type Stacker interface {
	Len() int          // Return the number of elements in the stack.
	Push(interface{})  // Push an object of unknown type onto the stack.
	Pop() interface{}  // Remove an object from the top of the stack and return it.
	Peek() interface{} // Return an object from the top of the stack without removing it.
}

// Stack is an implementation of a FILO stack structure using a linked list. The reason behind using a linked list
// rather than a dynamic array is because we guarantee any operation on the stack can be completed in O(1) time,
// disregarding overhead. While compiler optimisation might mean dynamic arrays are better in some circumstances, the
// linked list gives us a more general guarantee.
type Stack struct {
	topPtr *stackElement
	size   int
}

// stackElement holds one element from a Stack and is equivalent to a node in a linked list.
type stackElement struct {
	value interface{}
	next  *stackElement
}

// Len returns the number of elements in the stack.
func (s Stack) Len() int {
	return s.size
}

// Push pushes a new element on to the stack.
func (s *Stack) Push(v interface{}) {
	s.topPtr = &stackElement{
		value: v,
		next:  s.topPtr,
	}
	s.size++
}

// Pop removes the top element from the stack and returns it. If the stack is empty then this function will return nil.
func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		retVal := s.topPtr.value
		s.topPtr = s.topPtr.next
		s.size--
		return retVal
	}
	return nil
}

// Peek returns a copy of the top element on the stack (the one which will be popped first) without removing it from the
// underlying stack. If the stack is empty, it will return nil.
func (s Stack) Peek() interface{} {
	if s.size > 0 {
		return s.topPtr.value
	}
	return nil
}

// ConcurrentStack is a concurrency-safe implementation of the Stacker interface. It has a slight performance hit when
// compared to the other implementation (Stack), but the trade-off is that ConcurrentStack can be safely used between
// different goroutines, while the object is kept synchronised.
type ConcurrentStack struct {
	internalStack Stack
	lock          sync.RWMutex
}

// Len returns the number of elements in the stack. Unlike a regular Stack, this function operates on the pointer to cs
// so that the mutex is not duplicated.
func (cs *ConcurrentStack) Len() int {
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	return cs.internalStack.size
}

// Push pushes a new element onto the stack.
func (cs *ConcurrentStack) Push(v interface{}) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	cs.internalStack.Push(v)
}

// Pop removes an element from the top of the stack and returns it. If the stack is empty, it will return nil.
func (cs *ConcurrentStack) Pop() interface{} {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	return cs.internalStack.Pop()
}

// Peek returns a copy of the top element on the stack (the one which will be popped first) without removing it from the
// underlying stack. If the stack is empty, it will return nil.
func (cs *ConcurrentStack) Peek() interface{} {
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	return cs.internalStack.Peek()
}
