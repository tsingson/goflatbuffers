// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package reflection

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type FieldT struct {
	Name           string
	Type           *TypeT
	Id             uint16
	Offset         uint16
	DefaultInteger int64
	DefaultReal    float64
	Deprecated     bool
	Required       bool
	Key            bool
	Attributes     []*KeyValueT
	Documentation  []string
}

// FieldT object pack function
func (t *FieldT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	nameOffset := flatbuffers.UOffsetT(0)
	if len(t.Name) > 0 {
		nameOffset = builder.CreateString(t.Name)
	}
	typeOffset := t.Type.Pack(builder)
	attributesOffset := flatbuffers.UOffsetT(0)
	if t.Attributes != nil {
		attributesLength := len(t.Attributes)
		attributesOffsets := make([]flatbuffers.UOffsetT, attributesLength)
		for j := attributesLength - 1; j >= 0; j-- {
			attributesOffsets[j] = t.Attributes[j].Pack(builder)
		}
		FieldStartAttributesVector(builder, attributesLength)
		for j := attributesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(attributesOffsets[j])
		}
		attributesOffset = FieldEndAttributesVector(builder, attributesLength)
	}
	documentationOffset := flatbuffers.UOffsetT(0)
	if t.Documentation != nil {
		documentationOffset = builder.StringsVector(t.Documentation...)
	}

	// pack process all field

	FieldStart(builder)
	FieldAddName(builder, nameOffset)
	FieldAddType(builder, typeOffset)
	FieldAddId(builder, t.Id)
	FieldAddOffset(builder, t.Offset)
	FieldAddDefaultInteger(builder, t.DefaultInteger)
	FieldAddDefaultReal(builder, t.DefaultReal)
	FieldAddDeprecated(builder, t.Deprecated)
	FieldAddRequired(builder, t.Required)
	FieldAddKey(builder, t.Key)
	FieldAddAttributes(builder, attributesOffset)
	FieldAddDocumentation(builder, documentationOffset)
	return FieldEnd(builder)
}

// FieldT object unpack function
func (rcv *Field) UnPackTo(t *FieldT) {
	t.Name = string(rcv.Name())
	t.Type = rcv.Type(nil).UnPack()
	t.Id = rcv.Id()
	t.Offset = rcv.Offset()
	t.DefaultInteger = rcv.DefaultInteger()
	t.DefaultReal = rcv.DefaultReal()
	t.Deprecated = rcv.Deprecated()
	t.Required = rcv.Required()
	t.Key = rcv.Key()
	attributesLength := rcv.AttributesLength()
	t.Attributes = make([]*KeyValueT, attributesLength)
	for j := 0; j < attributesLength; j++ {
		x := KeyValue{}
		rcv.Attributes(&x, j)
		t.Attributes[j] = x.UnPack()
	}
	documentationLength := rcv.DocumentationLength()
	t.Documentation = make([]string, documentationLength)
	for j := 0; j < documentationLength; j++ {
		t.Documentation[j] = string(rcv.Documentation(j))
	}
}

func (rcv *Field) UnPack() *FieldT {
	if rcv == nil {
		return nil
	}
	t := &FieldT{}
	rcv.UnPackTo(t)
	return t
}

type Field struct {
	_tab flatbuffers.Table
}

// GetRootAsField shortcut to access root table
func GetRootAsField(buf []byte, offset flatbuffers.UOffsetT) *Field {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Field{}
	x.Init(buf, n+offset)
	return x
}

// GetTableVectorAsField shortcut to access table in vector of  unions
func GetTableVectorAsField(table *flatbuffers.Table) *Field {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &Field{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetTableAsField shortcut to access table in single union field
func GetTableAsField(table *flatbuffers.Table) *Field {
	x := &Field{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *Field) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Field) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Field) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Field) Type(obj *Type) *Type {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Type)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Field) Id() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Field) Offset() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Field) DefaultInteger() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Field) DefaultReal() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Field) Deprecated() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Field) Required() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Field) Key() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Field) Attributes(obj *KeyValue, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Field) AttributesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Field) Documentation(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Field) DocumentationLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func FieldStart(builder *flatbuffers.Builder) {
	builder.StartObject(11)
}

func FieldAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}

func FieldAddType(builder *flatbuffers.Builder, type_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(type_), 0)
}

func FieldAddId(builder *flatbuffers.Builder, id uint16) {
	builder.PrependUint16Slot(2, id, 0)
}

func FieldAddOffset(builder *flatbuffers.Builder, offset uint16) {
	builder.PrependUint16Slot(3, offset, 0)
}

func FieldAddDefaultInteger(builder *flatbuffers.Builder, defaultInteger int64) {
	builder.PrependInt64Slot(4, defaultInteger, 0)
}

func FieldAddDefaultReal(builder *flatbuffers.Builder, defaultReal float64) {
	builder.PrependFloat64Slot(5, defaultReal, 0.0)
}

func FieldAddDeprecated(builder *flatbuffers.Builder, deprecated bool) {
	builder.PrependBoolSlot(6, deprecated, false)
}

func FieldAddRequired(builder *flatbuffers.Builder, required bool) {
	builder.PrependBoolSlot(7, required, false)
}

func FieldAddKey(builder *flatbuffers.Builder, key bool) {
	builder.PrependBoolSlot(8, key, false)
}

func FieldAddAttributes(builder *flatbuffers.Builder, attributes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(attributes), 0)
}

func FieldStartAttributesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}

func FieldEndAttributesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func FieldAddDocumentation(builder *flatbuffers.Builder, documentation flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(documentation), 0)
}

func FieldStartDocumentationVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}

func FieldEndDocumentationVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func FieldEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}