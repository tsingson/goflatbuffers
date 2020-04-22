package flatbuffers

import (
	"errors"
	"sync"

	"github.com/tsingson/goflatbuffers/go/bytepool"
	"github.com/willf/bitset"
)

// VirtualTable virtual table
type VirtualTable struct {
	header    *bytepool.ByteBuffer
	payload   *bytepool.ByteBuffer
	head      int
	soff      int
	tableSize int
	soffset   SOffsetT //
	vfields   []VField
	voffsets  []VOffsetT

	isFixed bool

	fixedSize   int
	numsElement int
	r           *bitset.BitSet
	once        sync.Once

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
}

// NewVirtualTable initial virtual table
func NewVirtualTable() *VirtualTable {
	v := &VirtualTable{}
	v.once = sync.Once{}
	v.header = bytepool.NewByteBuffer(64)
	/**
	if n > 0 {
		v.tableSize = VOffsetT((n + 2) * SizeVOffsetT)
		pad, _ := Prepad(SizeSOffsetT, int(v.tableSize))
		v.head = pad
		v.soffset = SOffsetT(v.tableSize)

		// WriteVOffsetT(v.B[(v.head+SizeVOffsetT):], VOffsetT((n+2)*SizeVOffsetT))

		if fixed > 0 {
			v.isFixed = true
			v.objectSize = VOffsetT(fixed + SizeSOffsetT)
			total := v.head + (n+2)*SizeVOffsetT + SizeSOffsetT + fixed
			v.header.Reset()
			v.header.FixedLength(total)
			// table size in VTable
			WriteVOffsetT(v.header.B[(v.head):], VOffsetT(v.tableSize))
			// object size in VTable
			WriteVOffsetT(v.header.B[(v.head+SizeVOffsetT):], VOffsetT(v.objectSize))
			// SOffset
			WriteSOffsetT(v.header.B[(v.head+int(v.tableSize)):], SOffsetT(v.tableSize))
		} else {
			v.objectSize = VOffsetT(fixed + SizeSOffsetT)
			total := v.head + (n+2)*SizeVOffsetT + SizeSOffsetT
			v.header.Reset()
			v.header.FixedLength(total)
			// table size in VTable
			WriteVOffsetT(v.header.B[(v.head):], VOffsetT(v.tableSize))
			// object size in VTable
			WriteVOffsetT(v.header.B[(v.head+SizeVOffsetT):], VOffsetT(v.objectSize))
			// SOffset
			WriteSOffsetT(v.header.B[(v.head+int(v.tableSize)):], SOffsetT(v.tableSize))
		}


	} else {
		v.head = 0
		v.objectSize = 0
		v.tableSize = 4
		v.soffset = SizeSOffsetT

		v.header.Reset()
		v.header.FixedLength(8)
		// table size in VTable
		WriteVOffsetT(v.header.B[(v.head):], VOffsetT(v.tableSize))
		// object size in VTable
		WriteVOffsetT(v.header.B[(v.head+SizeVOffsetT):], VOffsetT(v.objectSize))
		// SOffset
		WriteSOffsetT(v.header.B[(v.head+int(v.tableSize)):], SOffsetT(v.soffset))
	}
	*/
	return v
}

func (b *VirtualTable) Init(minalign, bytesize, slot int) *VirtualTable {
	b.setup(minalign, bytesize, slot)

	return b
}

func (b *VirtualTable) setup(minalign, fixted, slot int) *VirtualTable {
	b.minalign = minalign
	b.byteSize = fixted // ( slot + 2) * SizeVOffsetT
	b.slot = slot
	b.finished = false

	if b.header == nil {
		b.header = bytepool.NewByteBuffer(64)
	}
	b.header.Reset(fixted)

	return b
}

// Press do the serialize job
func (b *VirtualTable) Press() error {
	if !b.r.All() {
		return errors.New("please fill all field")
	}
	return nil
}

// IsVector check this VField is vector or not
func (b *VirtualTable) IsVector() bool {
	return true
}

// IsFinished  check serialize job done or not
func (b *VirtualTable) IsFinished() bool {
	return b.finished
}

// Release set cache back to pool
func (b *VirtualTable) Release() bool {
	panic("implement me")
}

// Type  virtual field type
func (b *VirtualTable) Type() FieldType {
	return FieldTypeTable
}

// Head  return head offset
func (b *VirtualTable) Head() int {
	return b.head
}

// SOffset return SOffset
func (b *VirtualTable) SOffset() int {
	return int(b.soffset)
}

// Scalar register scalar
func (b *VirtualTable) Scalar(slot int, vf VOffsetT, in, df VField) *VirtualTable {
	return b.scalar(slot, vf, in, df)
}

// Struct register string
func (b *VirtualTable) Struct(slot int, vf VOffsetT, in VField) *VirtualTable {
	return b.vector(slot, vf, in)
}

// String register string
func (b *VirtualTable) String(slot int, vf VOffsetT, in, df VField) *VirtualTable {
	return b.vector(slot, vf, in)
}

// SharedString register string
func (b *VirtualTable) vector(slot int, vf VOffsetT, in VField) *VirtualTable {
	b.r.Set(uint(slot))
	b.vfields[slot] = in
	b.voffsets[slot] = vf
	return b
}

// SharedString register string
func (b *VirtualTable) scalar(slot int, vf VOffsetT, in, df VField) *VirtualTable {
	b.r.Set(uint(slot))
	b.vfields[slot] = in
	b.voffsets[slot] = vf
	return b
}

// Vector register vector
func (b *VirtualTable) Vector(slot int, vf VOffsetT, in VField) *VirtualTable {
	return b.vector(slot, vf, in)
}

// Array register array
func (b *VirtualTable) Array(slot int, vf VOffsetT, in VField) *VirtualTable {
	return b.vector(slot, vf, in)
}

// Table register table
func (b *VirtualTable) Table(slot int, vf VOffsetT, in VField) *VirtualTable {
	return b.vector(slot, vf, in)
}

// UnionType register table  // scalar.byte
func (b *VirtualTable) UnionType(slot int, vf VOffsetT, in, df VField) *VirtualTable {
	return b.scalar(slot, vf, in, df)
}

// Union register table
func (b *VirtualTable) Union(slot int, vf VOffsetT, in VField) *VirtualTable {
	return b.vector(slot, vf, in)
}

// FlexTable register flex table
func (b *VirtualTable) FlexTable(slot int, vf VOffsetT, in VField) *VirtualTable {
	return b.vector(slot, vf, in)
}

// ObjectSize object size in VTable
func (b *VirtualTable) ObjectSize() int {
	return int(b.objectSize)
}

// TableSize table size
func (b *VirtualTable) TableSize() int {
	return int(b.tableSize)
}

// ByteSize  size of serialized binary slice
func (b *VirtualTable) ByteSize() int {
	return b.header.Len()
}

// FinishByte return []byte
func (b *VirtualTable) FinishByte() []byte {
	if !b.finished {
		_ = b.Press()
	}
	return b.header.Bytes()
}

// AllSet check all field been registered or not
func (b *VirtualTable) AllSet() bool {
	return b.r.All()
}

// StartTable start table
func (b *VirtualTable) StartTable(i int) *VirtualTable {
	b.once.Do(func() {
		b.r = bitset.New(uint(i))
		b.r.ClearAll()

		b.numsElement = i
	})
	return b
}

// EndTable end table
func (b *VirtualTable) EndTable() *VirtualTable {
	if !b.r.All() {
		return nil
	}
	return b
}

// Finish do serialize job
func (b *VirtualTable) Finish() VField {
	return b
}
