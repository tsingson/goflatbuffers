package flatbuffers

import (
	"sync"
	"sync/atomic"
)

// Generic queue interfacelist mathing all implementations
type Register interface {
	Pop() (VField, bool)
	Size() int
	Len() int
	Push(VField)
	Close()
}

// Stack   single publish, single consume
type FieldRegister struct {
	_   [7]int64
	sc  int64 // read position
	_   [7]int64
	wc  int64 // write position
	_   [7]int64
	len int
	ring
	registers  []bool
	_          [4]byte
	lock       *sync.Mutex
	registered int
}

func (r *FieldRegister) Size() int {
	return int(r.size)
}

func NewFieldRegister(n uint32) Register {
	s := new(FieldRegister)
	s.Init(n)
	s.registers = make([]bool, roundUp2(n)) // initial ring data slice
	s.len = int(n)
	s.registered = 0
	s.lock = &sync.Mutex{}
	return s
}

func (r *FieldRegister) Len() int {
	return r.len
}

// Get check item iextis or net
func (r *FieldRegister) Pop() (v VField, b bool) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.registered < r.len {
		return nil, false
	}

	rc := atomic.LoadInt64(&r.sc) // get the write cursor
	wc := atomic.LoadInt64(&r.wc) // get the write cursor
	if wc == 0 {
		return nil, false
	}

	if (rc + 1) == wc {
		v = r.data[rc]
		switch {
		case rc > 0:
			atomic.StoreInt64(&r.sc, rc-1)
			atomic.StoreInt64(&r.wc, wc-1)
		default:
			atomic.StoreInt64(&r.wc, 0)
			atomic.StoreInt64(&r.sc, 0)
		}
	}
	return v, true
}

// Put put item into queue
func (r *FieldRegister) Push(v VField) {
	r.lock.Lock()
	defer r.lock.Unlock()

	wc := atomic.LoadInt64(&r.wc) // get the write cursor
	r.data[wc] = v
	atomic.StoreInt64(&r.sc, int64(wc))
	wc = atomic.AddInt64(&r.wc, 1) // set write cursor add 1
	if wc == r.size {
		r.Grow()
	}
}

// Put put item into queue
func (r *FieldRegister) PushAt(off int, v VField) {
	r.lock.Lock()
	defer r.lock.Unlock()

	wc := atomic.LoadInt64(&r.wc) // get the write cursor
	copy(r.registers[off:], []bool{true})
	r.data[off] = v
	atomic.StoreInt64(&r.sc, int64(wc))
	wc = atomic.AddInt64(&r.wc, 1) // set write cursor add 1
	if wc == r.size {
		r.Grow()
	}
}
