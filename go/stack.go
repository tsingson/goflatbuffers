package flatbuffers

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// default max batch size for batched consumers
const DefaultMaxBatch = (1 << 8) - 1

// Generic queue interfacelist mathing all implementations
type Queue interface {
	Pop() (VField, bool)
	Size() int
	Len() int
	NumsElement() int
	Push(VField)
	Close()
}

type ring struct {
	_        [7]int64
	wp       int64 //  write position
	_        [7]int64
	rp       int64 // read position
	_        [7]int64
	rc       int64 // read cursor
	_        [7]int64
	data     []VField
	mask     int64 // ring mask
	size     int64 // size of ring data slice
	maxbatch int64 //
	done     int32 // set true to disable ring buffer
}

func (r *ring) Init(n uint32) {
	r.data = make([]VField, roundUp2(n)) // initial ring data slice
	r.mask = int64(len(r.data) - 1)
	r.size = int64(len(r.data))

	var bs uint32 = DefaultMaxBatch
	r.maxbatch = int64(roundUp2(bs) - 1)
}

// Close  close
func (r *ring) Close() {
	atomic.AddInt32(&r.done, 1)
}

func (r *ring) Grow() {
	data := make([]VField, r.size*2)
	copy(data[0:], r.data)
	r.data = data
	r.size = int64(len(r.data))
}

// Done  done
func (r *ring) Done() bool {
	return atomic.LoadInt64(&r.wp) <= atomic.LoadInt64(&r.rp) && atomic.LoadInt32(&r.done) > 0
}

func (r *ring) wait() {
	runtime.Gosched()
}

func (r *ring) waitForEq(data *int64, val int64) (keep bool) {
	for keep = true; keep && atomic.LoadInt64(data) != val; runtime.Gosched() {
		keep = atomic.LoadInt64(&r.wp) > atomic.LoadInt64(&r.rp) || atomic.LoadInt32(&r.done) == 0
	}
	return
}

// Stack   single publish, single consume
type Stack struct {
	_  [7]int64
	sc int64 // read position
	_  [7]int64
	wc int64 // write position
	_  [7]int64
	ring
	_           [4]byte
	lock        *sync.Mutex
	len         int
	numsElement int
}

func NewStack(n uint32) Queue {
	s := new(Stack)
	s.Init(n)
	s.len = int(n)

	s.lock = &sync.Mutex{}
	return s
}

func (r *Stack) Len() int {
	return r.len
}

func (r *Stack) Size() int {
	return int(r.size)
}

func (r *Stack) NumsElement() int {
	return r.numsElement
}

// Get check item iextis or net
func (r *Stack) Pop() (v VField, b bool) {
	r.lock.Lock()
	defer r.lock.Unlock()

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
	r.numsElement--
	return v, true
}

// Put put item into queue
func (r *Stack) Push(v VField) {
	r.lock.Lock()
	defer r.lock.Unlock()

	wc := atomic.LoadInt64(&r.wc) // get the write cursor
	r.data[wc] = v
	atomic.StoreInt64(&r.sc, int64(wc))
	wc = atomic.AddInt64(&r.wc, 1) // set write cursor add 1
	if wc == r.size {
		r.Grow()
	}
	r.numsElement++
}

// type iface struct {
// 	t, d unsafe.Pointer
// }
//
// func extractptr(i interface{}) unsafe.Pointer {
// 	return (*iface)(unsafe.Pointer(&i)).d
//
// }
//
// func inject(i interface{}, ptr unsafe.Pointer) {
// 	v := (*unsafe.Pointer)((*iface)(unsafe.Pointer(&i)).d)
// 	*v = ptr
// }

func roundUp2(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	return v + 1
}
