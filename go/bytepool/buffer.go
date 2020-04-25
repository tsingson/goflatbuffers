// This is a derivitive work of Go's bytes.ByteBuffer implementation.
//
// Originally copyright 2009 The Go Authors. All rights reserved.
//
// Modifications copyright 2018 Steven Allen. All rights reserved.
//
// Use of this source code is governed by both a BSD-style and an MIT-style
// license that can be found in the LICENSE_BSD and LICENSE files.

package bytepool

type ByteBuffer struct {
	// Pool is the buffer pool to use. If nil, this ByteBuffer will use the
	// global buffer pool.
	Pool *BufferPool
	B    []byte
	Pos  uint32
	size int
	// Preallocated slice for samll reads/writes.
	// This is *really* important for performance and only costs 8 words.
	// bootstrap [64]byte
}

func (b *ByteBuffer) getPool() *BufferPool {
	if b.Pool == nil {
		return GlobalPool
	}
	return b.Pool
}

func (b *ByteBuffer) returnBuf() {
	b.getPool().PutClean(b.B)
	b.B = nil
}

func (b *ByteBuffer) getBuf(n int) []byte {
	return b.getPool().Get(n)
}

func (b *ByteBuffer) grow(n int) int {
	wOff := len(b.B)
	bCap := cap(b.B)

	if bCap >= wOff+n {
		b.B = b.B[:wOff+n]
		b.size = wOff + n
		return b.size
	}

	minCap := b.Len() + n

	if minCap >= bCap {
		// Needs new buffer.
		newBuf := b.getBuf(minCap)
		copy(newBuf, b.B)
		b.returnBuf()
		b.B = newBuf
		b.size = minCap
	}
	return minCap
}

// NewByteBuffer initial a new byte buffer
func NewByteBuffer(n int) *ByteBuffer {
	b := new(ByteBuffer)
	if n > 0 {
		b.B = b.getBuf(n)
		b.size = n
	}
	return b
}

// Len returns the number of bytes that can be read from this buffer.
func (b *ByteBuffer) Len() int {
	return len(b.B)
}

// Cap returns the current capacity of the buffer.
func (b *ByteBuffer) Cap() int {
	return cap(b.B)
}

// Bytes returns the slice of bytes currently buffered in the ByteBuffer.
func (b *ByteBuffer) Bytes() []byte {
	return b.B
}

// Head return head position
func (b *ByteBuffer) Head() uint32 {
	return b.Pos
}

// ByteSize return size of byte slice
func (b *ByteBuffer) ByteSize() int {
	return b.size
}

// Grow grows the internal buffer such that `n` bytes
func (b *ByteBuffer) Grow(n int) int {
	return b.grow(n)
}

func (b *ByteBuffer) shrink() {
	c := b.Cap()
	l := b.Len()
	if l == 0 {
		// Shortcut if empty.
		b.returnBuf()
		b.B = b.getBuf(l)
		b.Pos = 0
	} else if l*8 < c {
		// Only shrink when capacity > 8x length. Avoids shrinking too aggressively.
		newBuf := b.getBuf(l)
		copy(newBuf, b.B)
		b.returnBuf()
		b.Pos = 0
		b.B = newBuf[:l]
	}
}

// Reset makes ByteBuffer.B empty.
func (b *ByteBuffer) Reset(n int) *ByteBuffer {
	if n == 0 {
		n = 64
	}
	b.returnBuf()
	b.Pos = 0
	b.B = b.getPool().Get(n)
	b.size = n
	return b
}

// WriteAt write at
func (b *ByteBuffer) WriteAt(off int, data []byte) {
	n := off + len(data)
	if n > b.Len() {
		b.Grow(n - b.Len())
		b.size = n
	}
	copy(b.B[off:n], data)
}

// PrePad pre pad
func (b *ByteBuffer) PrePad(align, add int) *ByteBuffer {
	pad, total := prePad(align, add+b.Len())
	add1 := total - b.Len()
	if add1 > 0 {
		newSlice := b.getPool().Get(total)
		for i := total - 1; i >= (total - pad); i-- {
			writeInt8(newSlice[i:], 0x0)
		}
		copy(newSlice, b.B)
		b.returnBuf()
		b.B = newSlice
	}
	return b
}

func (b *ByteBuffer) Pad(pad int) *ByteBuffer {
	if pad > 0 && pad <= b.Len() {
		for i := b.Len() - 1; i >= (b.Len() - pad); i-- {
			writeInt8(b.B[i:], 0x0)
		}
	}
	return b
}

func (b *ByteBuffer) Release() *ByteBuffer {
	b.returnBuf()
	return b
}

// prePad pre pad
func prePad(size, additionalBytes int) (int, int) {
	alignSize := (^(additionalBytes)) + 1
	alignSize &= size - 1
	return alignSize, additionalBytes + alignSize
}
