package bytepool

import (
	"io"
)

// ByteBuffer provides byte buffer, which can be used for minimizing
// memory allocations.
//
// ByteBuffer may be used with functions appending data to the given []byte
// slice. See hand code for details.
//
// Use Get for obtaining an empty byte buffer.
type ByteBuffer struct {

	// B is a byte buffer to use in append-like workloads.
	// See hand code for details.
	B []byte
}

// Len returns the size of the byte buffer.
func (b *ByteBuffer) Len() int {
	return len(b.B)
}

// Cap returns the size of the byte buffer.
func (b *ByteBuffer) Cap() int {
	return cap(b.B)
}

// NewByteBuffer new
func NewByteBuffer(n int) *ByteBuffer {
	if n <= 0 {
		n = 1
	}
	return &ByteBuffer{
		B: make([]byte, 0, n),
	}
}

// ReadFrom implements io.ReaderFrom.
//
// The function appends all the data read from r to b.
func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) {
	p := b.B
	nStart := int64(len(p))
	nMax := int64(cap(p))
	n := nStart
	if nMax == 0 {
		nMax = 64
		p = make([]byte, nMax)
	} else {
		p = p[:nMax]
	}
	for {
		if n == nMax {
			nMax *= 2
			bNew := make([]byte, nMax)
			copy(bNew, p)
			p = bNew
		}
		nn, err := r.Read(p[n:])
		n += int64(nn)
		if err != nil {
			b.B = p[:n]
			n -= nStart
			if err == io.EOF {
				return n, nil
			}
			return n, err
		}
	}
}

// WriteTo implements io.WriterTo.
func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.B)
	return int64(n), err
}

// B returns b.B, i.e. all the bytes accumulated in the buffer.
//
// The purpose of this function is bytes.Buffer compatibility.
func (b *ByteBuffer) Bytes() []byte {
	return b.B
}

// Write implements io.Writer - it appends p to ByteBuffer.B
func (b *ByteBuffer) Write(p []byte) (int, error) {
	b.B = append(b.B, p...)
	return len(p), nil
}

// WriteByte appends the byte c to the buffer.
//
// The purpose of this function is bytes.Buffer compatibility.
//
// The function always returns nil.
func (b *ByteBuffer) WriteByte(c byte) error {
	b.B = append(b.B, c)
	return nil
}

// WriteString appends s to ByteBuffer.B.
func (b *ByteBuffer) WriteString(s string) (int, error) {
	b.B = append(b.B, s...)
	return len(s), nil
}

// Set sets ByteBuffer.B to p.
func (b *ByteBuffer) Set(p []byte) {
	b.B = append(b.B[:0], p...)
}

// SetString sets ByteBuffer.B to s.
func (b *ByteBuffer) SetString(s string) {
	b.B = append(b.B[:0], s...)
}

// String returns string representation of ByteBuffer.B.
func (b *ByteBuffer) String() string {
	return string(b.B)
}

// Reset makes ByteBuffer.B empty.
func (b *ByteBuffer) Reset() *ByteBuffer {
	if len(b.B) > 0 {
		for i := 0; i < len(b.B); i++ {
			WriteInt8(b.B[i:], 0x0)
		}
	}
	b.B = b.B[:0]
	return b
}

func WriteInt8(buf []byte, n int8) {
	buf[0] = byte(n)
}

// WriteAt write at
func (b *ByteBuffer) WriteAt(off int, data []byte) {
	n := off + len(data)
	if n > b.Len() {
		b.Grow(n - b.Len())
	}
	copy(b.B[off:n], data)
}

// FixedLength set  byte slice in fixed length
func (b *ByteBuffer) FixedLength(n int) *ByteBuffer {
	if n > 0 && n > b.Len() {
		b.Grow(n - b.Len())
	} else {
		b.B = b.B[:n]
	}
	return b
}

// Grow grow
func (b *ByteBuffer) Grow(add int) *ByteBuffer {
	total := len(b.B) + add
	if add > 0 && total > b.Len() {
		newSlice := make([]byte, total, total)
		for i := total - 1; i >= (total - add); i-- {
			WriteInt8(newSlice[i:], 0x0)
		}
		copy(newSlice, b.B)
		b.B = newSlice
	}
	return b
}

// Prepad prepad
func (b *ByteBuffer) Prepad(align, add int) *ByteBuffer {
	pad, total := prepad(align, add+b.Len())
	add1 := total - b.Len()
	if add1 > 0 {
		newSlice := make([]byte, total, total)
		for i := total - 1; i >= (total - pad); i-- {
			WriteInt8(newSlice[i:], 0x0)
		}
		copy(newSlice, b.B)
		b.B = newSlice
	}
	return b
}

func (b *ByteBuffer) Truncate(n int) *ByteBuffer {
	b.B = b.B[:n]
	return b
}

func (b *ByteBuffer) Pad(pad int) *ByteBuffer {
	if pad > 0 && pad <= b.Len() {
		for i := b.Len() - 1; i >= (b.Len() - pad); i-- {
			WriteInt8(b.B[i:], 0x0)
		}
	}
	return b
}

// Pad prepad
func prepad(size, additionalBytes int) (int, int) {
	alignSize := (^(additionalBytes)) + 1
	alignSize &= (size - 1)
	return alignSize, additionalBytes + alignSize
}
