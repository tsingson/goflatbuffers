package flatbuffers

import "github.com/tsingson/goflatbuffers/go/bytepool"

// ScalarStruct scalar vec
type ScalarStruct struct {
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

func (b *ScalarStruct) Press() error {
	return nil
}

// Type return field type
func (b *ScalarStruct) Type() FieldType {
	return b.t
}

// Finished return serialize result
func (b *ScalarStruct) IsFinished() bool {
	return b.finished
}

// IsVector is vector or not
func (b *ScalarStruct) IsVector() bool {
	return b.vector
}

// Size  return byte size
func (b *ScalarStruct) Size() int {
	return b.byteSize
}

// NewScalarStruct new scalar vec
func NewStruct() *ScalarStruct {
	b := &ScalarStruct{}
	b.t = FieldTypeStruct
	return b
}

// Pack pack
func (b *ScalarStruct) Release() bool {
	b.release = false
	if b.b != nil {
		bytepool.Put(b.b)
		b.b = nil
	}
	return b.release
}

// Init  initial  new scalar struct
func Init(minalign, bytesize, slot int) *ScalarStruct {
	b := &ScalarStruct{}
	return b.setup(minalign, bytesize, slot)
}

func (b *ScalarStruct) setup(minalign, fixted, slot int) *ScalarStruct {
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
func (b *ScalarStruct) StartStruct(minalign, bytesize, slot int) *ScalarStruct {
	b.setup(minalign, bytesize, slot)
	return b
}

// TypeName name of vec
func (b *ScalarStruct) TypeName() string {
	return b.name
}

// Pad filled the pad via "0" byte
func (b *ScalarStruct) Pad(n int) {
	// do nothing
}

// Bool insert a field into vec
func (b *ScalarStruct) Bool(x bool, n VOffsetT) *ScalarStruct {
	WriteBool(b.b.B[int(n):], x)
	return b
}

// Int8 insert a field into vec
func (b *ScalarStruct) Int8(x int8, n VOffsetT) *ScalarStruct {
	WriteInt8(b.b.B[int(n):], x)
	return b
}

// Int8 insert a field into vec
func (b *ScalarStruct) Byte(x byte, n VOffsetT) *ScalarStruct {
	WriteByte(b.b.B[int(n):], x)
	return b
}

// Int16 insert a field into vec
func (b *ScalarStruct) Int16(x int16, n VOffsetT) *ScalarStruct {
	WriteInt16(b.b.B[int(n):], x)
	return b
}

// Int32 insert a field into vec
func (b *ScalarStruct) Int32(x int32, n VOffsetT) *ScalarStruct {
	WriteInt32(b.b.B[int(n):], x)
	return b
}

// Int64 insert a field into vec
func (b *ScalarStruct) Int64(x int64, n VOffsetT) *ScalarStruct {
	WriteInt64(b.b.B[int(n):], x)
	return b
}

// Uint8 insert a field into vec
func (b *ScalarStruct) Uint8(x uint8, n VOffsetT) *ScalarStruct {
	WriteUint8(b.b.B[int(n):], x)
	return b
}

// Uint16 insert a field into vec
func (b *ScalarStruct) Uint16(x uint16, n VOffsetT) *ScalarStruct {
	WriteUint16(b.b.B[int(n):], x)
	return b
}

// Uint32 insert a field into vec
func (b *ScalarStruct) Uint32(x uint32, n VOffsetT) *ScalarStruct {
	WriteUint32(b.b.B[int(n):], x)
	return b
}

// Uint64 insert a field into vec
func (b *ScalarStruct) Uint64(x uint64, n VOffsetT) *ScalarStruct {
	WriteUint64(b.b.B[int(n):], x)
	return b
}

// Float32 insert a field into vec
func (b *ScalarStruct) Float32(x float32, n VOffsetT) *ScalarStruct {
	WriteFloat32(b.b.B[int(n):], x)
	return b
}

// Float64 insert a field into vec
func (b *ScalarStruct) Float64(x float64, n VOffsetT) *ScalarStruct {
	WriteFloat64(b.b.B[int(n):], x)
	return b
}

// SetName
func (b *ScalarStruct) SetName(name string) *ScalarStruct {
	b.name = name
	return b
}

// FinishStruct finish scalar vec
func (b *ScalarStruct) EndStruct() VField {
	b.finished = true
	return b
}

// FinishByte return []byte
func (b *ScalarStruct) Payload() []byte {
	if !b.finished {
		_ = b.Press()
	}
	// header.Release()
	return b.b.Bytes()
}

// B return []byte
func (b *ScalarStruct) ByteSize() int {
	return b.byteSize
}

// B return []byte
func (b *ScalarStruct) ObjectSize() int {
	return b.objectSize
}
