package flatbuffers

import "github.com/tsingson/goflatbuffers/go/bytepool"

// StructBuffers scalar vec
type StructBuffers struct {
	parent int64

	b    *bytepool.ByteBuffer // encoded vec into binary slice
	name string               // vec name

	minalign   int
	byteSize   int //  how many byte in a vec
	objectSize int
	padding    int
	slot       int // how many field in a vec
	slotend    int // write end

	t FieldType

	list []*Scalar // list of scalars in a vec

	alignSize int
	release   bool
	nested    bool
	finished  bool
	vector    bool
}

func (b *StructBuffers) Press() error {
	return nil
}

// Type return field type
func (b *StructBuffers) Type() FieldType {
	return b.t
}

// Finished return serialize result
func (b *StructBuffers) IsFinished() bool {
	return b.finished
}

// IsVector is vector or not
func (b *StructBuffers) IsVector() bool {
	return b.vector
}

// Size  return byte size
func (b *StructBuffers) Size() int {
	return b.byteSize
}

// NewScalarStruct new scalar vec
func NewStruct() *StructBuffers {
	b := &StructBuffers{}
	b.t = FieldTypeStruct
	return b
}

// Pack pack
func (b *StructBuffers) Release() bool {
	b.release = false
	if b.b != nil {
		bytepool.Put(b.b)
		b.b = nil
	}
	return b.release
}

// Init  initial  new scalar struct
func Init(minalign, bytesize, slot int) *StructBuffers {
	b := &StructBuffers{}
	return b.setup(minalign, bytesize, slot)
}

func (b *StructBuffers) setup(minalign, fixted, slot int) *StructBuffers {
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
func (b *StructBuffers) StructStart(minalign, bytesize, slot int) *StructBuffers {
	b.setup(minalign, bytesize, slot)
	return b
}

// TypeName name of vec
func (b *StructBuffers) TypeName() string {
	return b.name
}

// Pad filled the pad via "0" byte
func (b *StructBuffers) Pad(n int) {
	// do nothing
}

// Bool insert a field into vec
func (b *StructBuffers) Bool(x bool, n VOffsetT) *StructBuffers {
	WriteBool(b.b.B[int(n):], x)
	return b
}

// Int8 insert a field into vec
func (b *StructBuffers) Int8(x int8, n VOffsetT) *StructBuffers {
	WriteInt8(b.b.B[int(n):], x)
	return b
}

// Int8 insert a field into vec
func (b *StructBuffers) Byte(x byte, n VOffsetT) *StructBuffers {
	WriteByte(b.b.B[int(n):], x)
	return b
}

// Int16 insert a field into vec
func (b *StructBuffers) Int16(x int16, n VOffsetT) *StructBuffers {
	WriteInt16(b.b.B[int(n):], x)
	return b
}

// Int32 insert a field into vec
func (b *StructBuffers) Int32(x int32, n VOffsetT) *StructBuffers {
	WriteInt32(b.b.B[int(n):], x)
	return b
}

// Int64 insert a field into vec
func (b *StructBuffers) Int64(x int64, n VOffsetT) *StructBuffers {
	WriteInt64(b.b.B[int(n):], x)
	return b
}

// Uint8 insert a field into vec
func (b *StructBuffers) Uint8(x uint8, n VOffsetT) *StructBuffers {
	WriteUint8(b.b.B[int(n):], x)
	return b
}

// Uint16 insert a field into vec
func (b *StructBuffers) Uint16(x uint16, n VOffsetT) *StructBuffers {
	WriteUint16(b.b.B[int(n):], x)
	return b
}

// Uint32 insert a field into vec
func (b *StructBuffers) Uint32(x uint32, n VOffsetT) *StructBuffers {
	WriteUint32(b.b.B[int(n):], x)
	return b
}

// Uint64 insert a field into vec
func (b *StructBuffers) Uint64(x uint64, n VOffsetT) *StructBuffers {
	WriteUint64(b.b.B[int(n):], x)
	return b
}

// Float32 insert a field into vec
func (b *StructBuffers) Float32(x float32, n VOffsetT) *StructBuffers {
	WriteFloat32(b.b.B[int(n):], x)
	return b
}

// Float64 insert a field into vec
func (b *StructBuffers) Float64(x float64, n VOffsetT) *StructBuffers {
	WriteFloat64(b.b.B[int(n):], x)
	return b
}

// SetName
func (b *StructBuffers) SetName(name string) *StructBuffers {
	b.name = name
	return b
}

// StructEnd finish scalar vec
func (b *StructBuffers) StructEnd() VField {
	b.finished = true
	return b
}

// FinishByte return []byte
func (b *StructBuffers) FinishByte() []byte {
	if !b.finished {
		_ = b.Press()
	}
	// header.Release()
	return b.b.Bytes()
}

// B return []byte
func (b *StructBuffers) ByteSize() int {
	return b.byteSize
}

// B return []byte
func (b *StructBuffers) ObjectSize() int {
	return b.objectSize
}