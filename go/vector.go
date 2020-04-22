package flatbuffers

import (
	"github.com/tsingson/goflatbuffers/go/bytepool"
)

// IsVector vector vec
type Vector struct {
	objectSize int
	b          *bytepool.ByteBuffer
	t          FieldType
	FieldName  string // field name for json object, come from flatc compiler

	head     int
	nested   bool
	finished bool
	VOffset  VOffsetT // for scalar vec only, vec's scalar field VOffset be pre-defined in flatc compiler
	release  bool
	shared   bool
}

func (b *Vector) Release() bool {
	b.b.Release()
	b.release = true
	return b.release
}

func (b *Vector) ByteSize() int {
	return b.b.Len()
}

func (b *Vector) ObjectSize() int {
	return b.objectSize
}

// FinishByte  finish serialize and return []byte
func (b *Vector) FinishByte() []byte {
	if b.finished {
		b.release = true
		return b.b.Bytes()
	}
	return nil
}

// Type field type
func (b *Vector) Type() FieldType {
	return b.t
}

// IsVector bool
func (b *Vector) IsVector() bool {
	return true
}

// Finished bool
func (b *Vector) IsFinished() bool {
	return b.finished
}

// TypeName type name
func (b *Vector) TypeName() string {
	return b.FieldName
}

// NewVector  new vector
func NewVector() *Vector {
	b := &Vector{}
	b.b = bytepool.NewByteBuffer(64)
	b.head = SizeSOffsetT
	return b
}

// ByteSlice add a byte slice vector
func (b *Vector) ByteSlice(off VOffsetT, s []byte) VField {
	b.VOffset = off

	b.t = FieldTypeByteSlice
	b.objectSize = len(s)

	pad, total := Prepad(int(SizeUOffsetT), (b.objectSize+1)*SizeByte)

	if b.b == nil {
		b.b = bytepool.NewByteBuffer(64)
	}
	b.b.Reset(total + SizeSOffsetT).Pad(pad)

	for i := 0; i < pad; i++ {
		WriteByte(b.b.B[int(b.head)+i+b.objectSize:int(b.head)+i+b.objectSize+1], uint8(0))
	}
	copy(b.b.B[b.head:], s)
	b.EndVector(b.objectSize)
	b.finished = true

	return b
}

// Int8Slice add a int8 slice vector
func (b *Vector) Int8Slice(off VOffsetT, s []int8) VField {
	b.t = FieldTypeInt8Slice

	b.VOffset = off

	b.objectSize = len(s)

	pad, total := Prepad(int(SizeUOffsetT), (b.objectSize)*SizeByte)

	if b.b == nil {
		b.b = bytepool.NewByteBuffer(64)
	}
	b.b.Reset(total + SizeSOffsetT)

	for i := 0; i < pad; i++ {
		WriteByte(b.b.B[int(b.head)+b.objectSize+i:int(b.head)+b.objectSize+i+1], uint8(0))
	}

	b.objectSize = len(s)
	b.head = SizeSOffsetT
	for j := 0; j < b.objectSize; j++ {
		WriteInt8(b.b.B[int(b.head)+j:], s[j])
	}

	b.EndVector(b.objectSize)
	b.finished = true

	return b

	// return header.EndVector(len(s))
}

// EndVector writes data necessary to finish vector construction.
func (b *Vector) EndVector(vectorNumElems int) VField {
	// we already made space for this, so write without PrependUint32
	b.placeUOffsetT(UOffsetT(vectorNumElems))
	b.nested = false
	return b
}

// PlaceUOffsetT prepends a UOffsetT to the Builder, without checking for space.
func (b *Vector) placeUOffsetT(x UOffsetT) {
	WriteUOffsetT(b.b.B[0:], x)
}

// Serialize work
func (b *Vector) Press() error {
	b.finished = true
	return nil
}

func (b *Vector) assertNotNested() {
	// If you hit this, you're trying to construct a Table/IsVector/String
	// during the construction of its parent table (between the MyTableBuilder
	// and builder.Finish()).
	// Move the creation of these sub-objects to above the MyTableBuilder to
	// not get this assert.
	// Ignoring this assert may appear to work in simple cases, but the reason
	// it is here is that storing objects in-line may cause VTable offsets
	// to not fit anymore. It also leads to VTable duplication.
	if b.nested {
		panic("Incorrect creation order: object must not be nested.")
	}
}
