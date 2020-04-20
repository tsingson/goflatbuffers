package testcase

import (
	"bytes"
	"github.com/tsingson/goflatbuffers/go"
	"testing"
)

func TestBuilder_CreateByteString(t *testing.T) {
	CheckCreateByteVector(t.Fatalf)
}
func TestBuilder_CreateByteVector(t *testing.T) {
	raw := [30]byte{}
	for i := 0; i < len(raw); i++ {
		raw[i] = byte(i)
	}
	size := len(raw)
	for s := 1; s < size; s++ {
		b2 := flatbuffers.NewBuilder(0)
		b2.CreateByteVector(raw[:s])
		b1 := flatbuffers.NewBuilder(0)
		b1.StartVector(1, s, 1)
		for i := s - 1; i >= 0; i-- {
			b1.PrependByte(raw[i])
		}
		b1.EndVector(s)
		CheckByteEquality(b1.Bytes, b2.Bytes, t.Fatalf, "go CreateByteVector")
	}
}
func CheckCreateByteVector(fail func(string, ...interface{})) {
	raw := [30]byte{}
	for i := 0; i < len(raw); i++ {
		raw[i] = byte(i)
	}
	size := len(raw)
	for s := 1; s < size; s++ {
		b2 := flatbuffers.NewBuilder(0)
		b2.CreateByteVector(raw[:s])
		b1 := flatbuffers.NewBuilder(0)
		b1.StartVector(1, s, 1)
		for i := s - 1; i >= 0; i-- {
			b1.PrependByte(raw[i])
		}
		b1.EndVector(s)
		CheckByteEquality(b1.Bytes, b2.Bytes, fail, "go CreateByteVector")
	}
}

// CheckByteEquality verifies that two byte buffers are the same.
func CheckByteEquality(a, b []byte, fail func(string, ...interface{}), str string) {
	if !bytes.Equal(a, b) {
		fail("objects are not byte-wise equal in " + str)
	}
}
