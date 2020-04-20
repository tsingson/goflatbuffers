package testcase

import (
	"strconv"

	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type Color byte

const (
	ColorRed    Color = 1
	ColorGreen  Color = 2
	ColorBlue   Color = 4
	ColorPurple Color = 8
)

var EnumNamesColor = map[Color]string{
	ColorRed:    "Red",
	ColorGreen:  "Green",
	ColorBlue:   "Blue",
	ColorPurple: "Purple",
}

var EnumValuesColor = map[string]Color{
	"Red":    ColorRed,
	"Green":  ColorGreen,
	"Blue":   ColorBlue,
	"Purple": ColorPurple,
}

// String get enum's name
func (v Color) String() string {
	if s, ok := EnumNamesColor[v]; ok {
		return s
	}
	return "Color(" + strconv.FormatInt(int64(v), 10) + ")"
}

// ColorTypeTable return type table
func ColorTypeTable() flatbuffers.TypeTable {
	return flatbuffers.TypeTable{
		SequenceType: flatbuffers.FieldTypeEnum,
		Name:         []byte("Color"),
		NumsElement:  4,
		Fields: []flatbuffers.TypeCode{
			{Name: "Red", BaseType: flatbuffers.FieldTypeByte},
			{Name: "Green", BaseType: flatbuffers.FieldTypeByte},
			{Name: "Blue", BaseType: flatbuffers.FieldTypeByte},
			{Name: "Purple", BaseType: flatbuffers.FieldTypeByte},
		},
	}
}
