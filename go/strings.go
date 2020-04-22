package flatbuffers

import "github.com/tsingson/goflatbuffers/go/bytepool"

// VectorString create string vector
func VectorString(s string) VField {
	return NewVector().string(s)
}

// String add a string vector
func (b *Vector) String(off VOffsetT, s string) VField {
	b.VOffset = off
	return b.string(s)
}

func (b *Vector) string(s string) VField {
	b.t = FieldTypeString
	b.objectSize = len(s)

	pad, total := Prepad(int(SizeUOffsetT), (b.objectSize+1)*SizeByte)
	//
	// header.B = make([]byte, int(header.head)+total)
	// l := len(s)
	// for i := 0; i < pad; i++ {
	// 	// WriteInt8(header.B[i:], int8(0))
	// 	WriteByte(header.B[int(header.head)+l+i:int(header.head)+l+i+1], uint8(0))
	// }
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(64)
	}
	b.b.Reset(total + SizeSOffsetT).Pad(pad)
	copy(b.b.B[b.head:], s)
	b.EndVector(b.objectSize)
	b.finished = true

	return b
}

// SharedString  add a shared string vector
func (b *Vector) SharedString(off VOffsetT, s string) VField {
	b.shared = true
	b.String(off, s)
	return b
}
