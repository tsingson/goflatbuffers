package testcase

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVtable_Init(t *testing.T) {
	as := assert.New(t)
	v := flatbuffers.NewVTable().Init(4, 16)
	v.SetVOffset(0, 0+4)
	v.SetVOffset(1, 4+4)
	v.SetVOffset(2, 8+4)
	v.SetVOffset(3, 12+4)
	as.Equal(v.ByteSize(), 12)
	as.Equal(v.Head(), uint16(0))
}

/**
func TestVtable_Init2(t *testing.T) {
	as := assert.New(t)
	v := flatbuffers.NewVTable().Init(4, 16)
	v.SetVOffset(0, 0+4)
	v.SetVOffset(1, 4+4)
	v.SetVOffset(2, 6+4)
	v.SetVOffset(3, 7+4)
	v.SetVOffset(4, 7+4)
	as.Equal(v.ByteSize(), 12)
	as.Equal(v.Head(), uint16(0))

	payload := CreateVec3(1.1, 1.1, 1.2, ColorBlue).Payload()

	buf := bytepool.NewByteBuffer(0).FixedLength(4 + 12 + 4 + 16)
	flatbuffers.WriteSOffsetT(buf.B[0:4], 16)
	copy(buf.B[4:], v.Payload())
	flatbuffers.WriteSOffsetT(buf.B[16:20], 12)
	copy(buf.B[20:], payload)

	fmt.Printf("%b\n", buf.Bytes())
}
*/
