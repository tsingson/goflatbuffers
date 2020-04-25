package flatbuffers

import (
	"github.com/cespare/xxhash/v2"
	"github.com/tsingson/goflatbuffers/go/bytepool"
)

// VTable virtual table
type VTable struct {
	b          *bytepool.ByteBuffer
	head       int
	tableSize  int
	objectSize int
	size       int
}

func NewVTable() *VTable {
	return &VTable{}
}

func (v *VTable) Init(n, fixed int) *VTable {
	ts := (n + 2) * SizeVOffsetT
	v.head, v.size = Prepad(SizeSOffsetT, ts)
	v.b = bytepool.NewByteBuffer(v.size)
	// WriteVOffsetT(v.B[(v.head+SizeVOffsetT):], VOffsetT((n+2)*SizeVOffsetT))
	if fixed > 0 {
		v.objectSize = fixed
	}
	// table size in VTable
	WriteVOffsetT(v.b.B[(v.head):], VOffsetT(v.tableSize))
	// object size in VTable
	WriteVOffsetT(v.b.B[(v.head+SizeVOffsetT):], VOffsetT(v.objectSize))
	// SOffset
	// WriteSOffsetT(v.b.B[(v.head+int(v.tableSize)):], SOffsetT(v.tableSize))
	return v
}

// SetObjectSize set object size
func (v *VTable) SetObjectSize(n int) *VTable {
	WriteVOffsetT(v.b.B[(v.head+SizeVOffsetT):], VOffsetT(v.objectSize))
	return v
}

func (v *VTable) SetVOffset(id int, off uint16) *VTable {
	WriteVOffsetT(v.b.B[(v.head+SizeVOffsetT*(2+id)):], VOffsetT(off))
	return v
}

// Head return head
func (v *VTable) Head() uint16 {
	return uint16(v.head)
}

// ByteSize return
func (v *VTable) ByteSize() int {
	return v.size
}

// Hash return vtable hash
func (v *VTable) Hash() uint64 {
	return xxhash.Sum64(v.b.B)
}

// Payload return
func (v *VTable) Payload() []byte {
	return v.b.Bytes()
}
