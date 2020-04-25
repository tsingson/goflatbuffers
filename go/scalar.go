// Scalar scalar in go
// bool
// uint8, uint16, uint32, uint64 (unsigned)
// int8, int16, int32, int64     (two's complement)
// float32, float64              (IEEE-754)

package flatbuffers

import (
	"github.com/tsingson/goflatbuffers/go/bytepool"
)

// Scalar vec for scalar in flatc
type Scalar struct {
	b         *bytepool.ByteBuffer
	bytesize  int
	release   bool
	FieldName string    // field name for json object, come from flatc compiler
	t         FieldType // field type for json object , come from flatc compiler
	VOffset   VOffsetT  // for scalar vec only, vec's scalar field VOffset be pre-defined in flatc compiler
}

// Serialize process
func (b *Scalar) Press() error {
	return nil
}

// Type return field type
func (b *Scalar) Type() FieldType {
	return b.t
}

// TypeName return type name
func (b *Scalar) TypeName() string {
	return "scalar"
}

// NewScalar new
func NewScalar() *Scalar {
	return &Scalar{
		release:  false,
		bytesize: 0,
	}
}

// Pack pack
func (b *Scalar) Release() bool {
	b.release = true
	b.b.Release()
	return b.release
}

// ObjectSize return scalar encoded binary size
func (b *Scalar) ObjectSize() int {
	return b.bytesize
}

func (b *Scalar) ByteSize() int {
	return b.b.Len()
}

// Finished is finish
func (b *Scalar) IsFinished() bool {
	return true
}

// IsVector is vector or not
func (b *Scalar) IsVector() bool {
	return false
}

// Bool encoded
func (b *Scalar) Bool(x bool) VField {
	b.t = FieldTypeBool
	b.bytesize = SizeBool

	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteBool(b.b.B, x)
	return b
}

// Int8 encoded
func (b *Scalar) Int8(x int8) VField {
	b.t = FieldTypeInt8
	b.bytesize = SizeInt8
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}

	WriteInt8(b.b.B, x)
	return b
}

func (b *Scalar) Byte(x byte) VField {
	b.t = FieldTypeInt8
	b.bytesize = SizeInt8
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}

	WriteByte(b.b.B, x)
	return b
}

func (b *Scalar) Ubyte(x int8) VField {
	b.t = FieldTypeInt8
	b.bytesize = SizeInt8
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteInt8(b.b.B, x)
	return b
}

// Int16 encoded
func (b *Scalar) Int16(x int16) VField {
	b.t = FieldTypeInt16
	b.bytesize = SizeInt16
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}

	WriteInt16(b.b.B, x)
	return b
}

// Int32 encoded
func (b *Scalar) Int32(x int32) VField {
	b.t = FieldTypeInt32
	b.bytesize = SizeInt32
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}

	WriteInt32(b.b.B, x)
	return b
}

// Int64 encoded
func (b *Scalar) Int64(x int64) VField {
	b.t = FieldTypeInt64
	b.bytesize = SizeInt64
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteInt64(b.b.B, x)
	return b
}

// Uint8 encoded
func (b *Scalar) Uint8(x uint8) VField {
	b.t = FieldTypeUint8
	b.bytesize = SizeUint8
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteUint8(b.b.B, x)
	return b
}

// Uint16 encoded
func (b *Scalar) Uint16(x uint16) VField {
	b.t = FieldTypeUint16
	b.bytesize = SizeUint16
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteUint16(b.b.B, x)
	return b
}

// Uint32 encoded
func (b *Scalar) Uint32(x uint32) VField {
	b.t = FieldTypeUint32
	b.bytesize = SizeUint32
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteUint32(b.b.B, x)
	return b
}

// Uint64 encoded
func (b *Scalar) Uint64(x uint64) VField {
	b.t = FieldTypeUint64
	b.bytesize = SizeUint64
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteUint64(b.b.B, x)
	return b
}

// Float32 encoded
func (b *Scalar) Float32(x float32) VField {
	b.t = FieldTypeFloat32
	b.bytesize = SizeFloat32
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteFloat32(b.b.B, x)
	return b
}

// Float64 encoded
func (b *Scalar) Float64(x float64) VField {
	b.t = FieldTypeFloat64
	b.bytesize = SizeFloat64
	if b.b == nil {
		b.b = bytepool.NewByteBuffer(b.bytesize)
	} else {
		b.b.Reset(b.bytesize)
	}
	WriteFloat64(b.b.B, x)
	return b
}

// FinishedBytes return []byte
func (b *Scalar) FinishedBytes() []byte {
	return b.b.Bytes()
}
