package bytepool

import (
	"bytes"
	"fmt"

	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestByteBuffer_Grow2(t *testing.T) {
	as := assert.New(t)
	var b ByteBuffer
	b.FixedLength(4).Grow(3)
	as.Equal(b.Len(), 7)
	// fmt.Printf("%b\n", b.B)
}

func TestByteBuffer_Prepad(t *testing.T) {
	as := assert.New(t)
	var b ByteBuffer
	fixed := 8
	pre := 1
	_, ex := prepad(4, pre+fixed)
	b.FixedLength(fixed).Prepad(4, pre)
	as.Equal(b.Len(), ex)
	// fmt.Printf("%b\n", b.B)
}

func TestByteBuffer_Prepad2(t *testing.T) {
	as := assert.New(t)
	var b ByteBuffer
	fixed := 8
	pre := 1
	_, ex := prepad(4, pre+fixed)
	b.FixedLength(fixed).Grow(fixed).Prepad(4, pre)
	as.Equal(b.Len(), ex+fixed)
	// fmt.Printf("%b\n", b.B)
}

func TestByteBuffer_Pad(t *testing.T) {
	as := assert.New(t)
	var b ByteBuffer
	s := []byte("1")
	pad, total := prepad(4, len(s))
	b.FixedLength(total)
	copy(b.B[0:], s)
	b.Pad(pad)
	as.Equal(b.Len(), total)
	// fmt.Printf("%b\n", b.B)
}

func TestByteBuffer_Prepad1(t *testing.T) {
	as := assert.New(t)
	var b ByteBuffer

	old := []byte("1111")
	b.FixedLength(len(old))
	copy(b.B[0:], old)
	off := b.Len()
	s := []byte("55555")

	b.Prepad(4, len(s))
	// fmt.Printf("%b\n", b.B)
	copy(b.B[off:], s)
	as.Equal(b.Len(), 12)
	// fmt.Printf("%b\n", b.B)
}

func TestByteBuffer_WriteAt(t *testing.T) {
	as := assert.New(t)
	var b ByteBuffer
	s := []byte("12345678")
	off := 4
	b.WriteAt(off, s)
	as.Equal(b.Len(), off+len(s))
	// fmt.Printf("%4b\n", b.B)
	// fmt.Println(off + len(s))
}

func TestByteBuffer_Grow(t *testing.T) {
	as := assert.New(t)
	fixed := 8
	grow := 8
	v := Get().FixedLength(fixed).Grow(grow)
	as.Equal(v.Len(), fixed+grow)
}

func TestByteBufferReadFrom(t *testing.T) {
	prefix := "foobar"
	expectedS := "asadfsdafsadfasdfisdsdfa"
	prefixLen := int64(len(prefix))
	expectedN := int64(len(expectedS))

	var bb ByteBuffer
	bb.WriteString(prefix)

	rf := (io.ReaderFrom)(&bb)
	for i := 0; i < 20; i++ {
		r := bytes.NewBufferString(expectedS)
		n, err := rf.ReadFrom(r)
		if n != expectedN {
			t.Fatalf("unexpected n=%d. Expecting %d. iteration %d", n, expectedN, i)
		}
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		bbLen := int64(bb.Len())
		expectedLen := prefixLen + int64(i+1)*expectedN
		if bbLen != expectedLen {
			t.Fatalf("unexpected byteBuffer length: %d. Expecting %d", bbLen, expectedLen)
		}
		for j := 0; j < i; j++ {
			start := prefixLen + int64(j)*expectedN
			b := bb.B[start : start+expectedN]
			if string(b) != expectedS {
				t.Fatalf("unexpected byteBuffer contents: %q. Expecting %q", b, expectedS)
			}
		}
	}
}

func TestByteBufferWriteTo(t *testing.T) {
	expectedS := "foobarbaz"
	var bb ByteBuffer
	bb.WriteString(expectedS[:3])
	bb.WriteString(expectedS[3:])

	wt := (io.WriterTo)(&bb)
	var w bytes.Buffer
	for i := 0; i < 10; i++ {
		n, err := wt.WriteTo(&w)
		if n != int64(len(expectedS)) {
			t.Fatalf("unexpected n returned from WriteTo: %d. Expecting %d", n, len(expectedS))
		}
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		s := string(w.Bytes())
		if s != expectedS {
			t.Fatalf("unexpected string written %q. Expecting %q", s, expectedS)
		}
		w.Reset()
	}
}

func TestByteBufferGetPutSerial(t *testing.T) {
	testByteBufferGetPut(t)
}

func TestByteBufferGetPutConcurrent(t *testing.T) {
	concurrency := 10
	ch := make(chan struct{}, concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			testByteBufferGetPut(t)
			ch <- struct{}{}
		}()
	}

	for i := 0; i < concurrency; i++ {
		select {
		case <-ch:
		case <-time.After(time.Second):
			t.Fatalf("timeout!")
		}
	}
}

func testByteBufferGetPut(t *testing.T) {
	for i := 0; i < 10; i++ {
		expectedS := fmt.Sprintf("num %d", i)
		b := Get()
		b.B = append(b.B, "num "...)
		b.B = append(b.B, fmt.Sprintf("%d", i)...)
		if string(b.B) != expectedS {
			t.Fatalf("unexpected result: %q. Expecting %q", b.B, expectedS)
		}
		Put(b)
	}
}

func testByteBufferGetString(t *testing.T) {
	for i := 0; i < 10; i++ {
		expectedS := fmt.Sprintf("num %d", i)
		b := Get()
		b.SetString(expectedS)
		if b.String() != expectedS {
			t.Fatalf("unexpected result: %q. Expecting %q", b.B, expectedS)
		}
		Put(b)
	}
}

func TestByteBufferGetStringSerial(t *testing.T) {
	testByteBufferGetString(t)
}

func TestByteBufferGetStringConcurrent(t *testing.T) {
	concurrency := 10
	ch := make(chan struct{}, concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			testByteBufferGetString(t)
			ch <- struct{}{}
		}()
	}

	for i := 0; i < concurrency; i++ {
		select {
		case <-ch:
		case <-time.After(time.Second):
			t.Fatalf("timeout!")
		}
	}
}
