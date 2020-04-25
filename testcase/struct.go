package testcase

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type Vec3T struct {
	X float32
	Y float32
	Z float32
	C Color
}

//
// func (t *Vec3T) Pack(builder *flatbuffers.StructBuffers) flatbuffers.VField {
// 	if t == nil {
// 		return 0
// 	}
// 	return CreateVec3(builder, t.X, t.Y, t.Z, t.C)
// }
func (rcv *Vec3) UnPackTo(t *Vec3T) {
	t.X = rcv.X()
	t.Y = rcv.Y()
	t.Z = rcv.Z()
	t.C = rcv.C()
}

func (rcv *Vec3) UnPack() *Vec3T {
	if rcv == nil {
		return nil
	}
	t := &Vec3T{}
	rcv.UnPackTo(t)
	return t
}

type Vec3 struct {
	_tab flatbuffers.Struct
}

func (rcv *Vec3) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec3) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Vec3) X() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}

func (rcv *Vec3) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Vec3) Y() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}

func (rcv *Vec3) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *Vec3) Z() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}

func (rcv *Vec3) MutateZ(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func (rcv *Vec3) C() Color {
	return Color(rcv._tab.GetByte(rcv._tab.Pos + flatbuffers.UOffsetT(12)))
}

func (rcv *Vec3) MutateC(n Color) bool {
	return rcv._tab.MutateByte(rcv._tab.Pos+flatbuffers.UOffsetT(12), byte(n))
}

func CreateVec3(builder *flatbuffers.StructBuffers, x float32, y float32, z float32, c Color) flatbuffers.VField {
	builder.Prep(4, 16)
	builder.Pad(3, 13)
	builder.Byte(byte(c), 12)
	builder.Float32(z, 8)
	builder.Float32(y, 4)
	builder.Float32(x, 0)
	return builder.StructEnd()
}

// Vec3TypeTable return type table
func Vec3TypeTable() flatbuffers.TypeTable {
	return flatbuffers.TypeTable{
		NumsElement:  4,
		FixedLength:  16,
		SequenceType: flatbuffers.FieldTypeStruct,
		Name:         []byte("Vec3"),
		Fields: []flatbuffers.TypeCode{
			// scalar
			{Id: 0, Offset: 0, Name: "X", IsDeprecated: false, BaseType: flatbuffers.FieldTypeFloat32},
			// scalar
			{Id: 1, Offset: 4, Name: "Y", IsDeprecated: false, BaseType: flatbuffers.FieldTypeFloat32},
			// scalar
			{Id: 2, Offset: 8, Name: "Z", IsDeprecated: false, BaseType: flatbuffers.FieldTypeFloat32},
			//  enum
			{Id: 3, Offset: 12, Name: "C", IsDeprecated: false, BaseType: flatbuffers.FieldTypeEnum, SequenceRef: ColorTypeTable()},
		},
	}
}
