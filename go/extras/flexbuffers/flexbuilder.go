package flexbuffers

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
	"github.com/tsingson/goflatbuffers/go/bytepool"
)

// FlexBuilder scalar vec
type FlexBuilder struct {
	parent int64

	b    *bytepool.ByteBuffer // encoded vec into binary slice
	name string               // vec name

	minalign   int
	byteSize   int //  how many byte in a vec
	objectSize int
	padding    int
	slot       int // how many field in a vec
	slotend    int // write end

	t flatbuffers.FieldType

	list []*flatbuffers.Scalar // list of scalars in a vec

	alignSize int
	release   bool
	nested    bool
	finished  bool
	vector    bool
}

func (b *FlexBuilder) Press() error {
	return nil
}

// Type return field type
func (b *FlexBuilder) Type() flatbuffers.FieldType {
	return b.t
}

// Finished return serialize result
func (b *FlexBuilder) IsFinished() bool {
	return b.finished
}

// IsVector is vector or not
func (b *FlexBuilder) IsVector() bool {
	return b.vector
}

// Size  return byte size
func (b *FlexBuilder) Size() int {
	return b.byteSize
}

// NewScalarStruct new scalar vec
func NewStruct() *FlexBuilder {
	b := &FlexBuilder{}
	b.t = flatbuffers.FieldTypeStruct
	return b
}

// Pack pack
func (b *FlexBuilder) Release() bool {
	b.release = false
	if b.b != nil {
		bytepool.Put(b.b)
		b.b = nil
	}
	return b.release
}

// Init  initial  new scalar struct
func Init(minalign, bytesize, slot int) *FlexBuilder {
	b := &FlexBuilder{}
	return b.setup(minalign, bytesize, slot)
}

func (b *FlexBuilder) setup(minalign, fixted, slot int) *FlexBuilder {
	b.minalign = minalign
	b.byteSize = fixted // ( slot + 2) * SizeVOffsetT
	b.slot = slot
	b.finished = false
	b.vector = false

	if b.b == nil {
		b.b = bytepool.Get()
	}
	b.b.Reset()
	b.b.FixedLength(fixted)
	return b
}

// Init initial scalar vec
func (b *FlexBuilder) StructStart(minalign, bytesize, slot int) *FlexBuilder {
	b.setup(minalign, bytesize, slot)
	return b
}

// TypeName name of vec
func (b *FlexBuilder) TypeName() string {
	return b.name
}

// Pad filled the pad via "0" byte
func (b *FlexBuilder) Pad(n int) {
	// do nothing
}

// Bool insert a field into vec
func (b *FlexBuilder) Bool(x bool, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteBool(b.b.B[int(n):], x)
	return b
}

// Int8 insert a field into vec
func (b *FlexBuilder) Int8(x int8, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteInt8(b.b.B[int(n):], x)
	return b
}

// Int8 insert a field into vec
func (b *FlexBuilder) Byte(x byte, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteByte(b.b.B[int(n):], x)
	return b
}

// Int16 insert a field into vec
func (b *FlexBuilder) Int16(x int16, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteInt16(b.b.B[int(n):], x)
	return b
}

// Int32 insert a field into vec
func (b *FlexBuilder) Int32(x int32, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteInt32(b.b.B[int(n):], x)
	return b
}

// Int64 insert a field into vec
func (b *FlexBuilder) Int64(x int64, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteInt64(b.b.B[int(n):], x)
	return b
}

// Uint8 insert a field into vec
func (b *FlexBuilder) Uint8(x uint8, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteUint8(b.b.B[int(n):], x)
	return b
}

// Uint16 insert a field into vec
func (b *FlexBuilder) Uint16(x uint16, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteUint16(b.b.B[int(n):], x)
	return b
}

// Uint32 insert a field into vec
func (b *FlexBuilder) Uint32(x uint32, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteUint32(b.b.B[int(n):], x)
	return b
}

// Uint64 insert a field into vec
func (b *FlexBuilder) Uint64(x uint64, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteUint64(b.b.B[int(n):], x)
	return b
}

// Float32 insert a field into vec
func (b *FlexBuilder) Float32(x float32, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteFloat32(b.b.B[int(n):], x)
	return b
}

// Float64 insert a field into vec
func (b *FlexBuilder) Float64(x float64, n flatbuffers.VOffsetT) *FlexBuilder {
	flatbuffers.WriteFloat64(b.b.B[int(n):], x)
	return b
}

// SetName
func (b *FlexBuilder) SetName(name string) *FlexBuilder {
	b.name = name
	return b
}

// StructEnd finish scalar vec
func (b *FlexBuilder) StructEnd() flatbuffers.VField {
	b.finished = true
	return b
}

// FinishByte return []byte
func (b *FlexBuilder) FinishByte() []byte {
	if !b.finished {
		_ = b.Press()
	}
	// header.Release()
	return b.b.Bytes()
}

// B return []byte
func (b *FlexBuilder) ByteSize() int {
	return b.byteSize
}

// B return []byte
func (b *FlexBuilder) ObjectSize() int {
	return b.objectSize
}
