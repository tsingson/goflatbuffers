package testcase

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type GunT struct {
	Color Color
}

type Gun struct {
	_tab flatbuffers.Table
}

func GetRootAsGun(buf []byte, offset flatbuffers.UOffsetT) *Gun {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Gun{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Gun) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Gun) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Gun) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Color(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Gun) MutateColor(n Color) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func GunStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}

func GunAddColor(builder *flatbuffers.Builder, color Color) {
	builder.PrependByteSlot(0, byte(color), 0)
}

func GunEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
