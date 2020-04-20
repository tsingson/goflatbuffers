package flatbuffers

import "strconv"

type FieldType int8

const (
	FieldTypeNone        FieldType = 0
	FieldTypeUType       FieldType = 1
	FieldTypeBool        FieldType = 2
	FieldTypeInt8        FieldType = 3
	FieldTypeUint8       FieldType = 4
	FieldTypeByte        FieldType = 5
	FieldTypeInt16       FieldType = 6
	FieldTypeUint16      FieldType = 7
	FieldTypeInt32       FieldType = 8
	FieldTypeUint32      FieldType = 9
	FieldTypeInt64       FieldType = 10
	FieldTypeUint64      FieldType = 11
	FieldTypeFloat32     FieldType = 12
	FieldTypeFloat64     FieldType = 13
	FieldTypeString      FieldType = 14
	FieldTypeByteSlice   FieldType = 15
	FieldTypeInt8Slice   FieldType = 16
	FieldTypeVectorPoint FieldType = 17
	FieldTypeScalar      FieldType = 18
	FieldTypeEnum        FieldType = 19
	FieldTypeStruct      FieldType = 20
	FieldTypeTable       FieldType = 21
	FieldTypeUnion       FieldType = 22
	FieldTypeArray       FieldType = 23
	FieldTypeFlexTable   FieldType = 24
	FieldTypeCustomized  FieldType = 25
)

var EnumNamesFieldType = map[FieldType]string{
	FieldTypeNone:        "None",
	FieldTypeUType:       "UType",
	FieldTypeBool:        "Bool",
	FieldTypeInt8:        "Int8",
	FieldTypeUint8:       "Uint8",
	FieldTypeByte:        "Byte",
	FieldTypeInt16:       "Int16",
	FieldTypeUint16:      "Uint16",
	FieldTypeInt32:       "Int32",
	FieldTypeUint32:      "Uint32",
	FieldTypeInt64:       "Int64",
	FieldTypeUint64:      "Uint64",
	FieldTypeFloat32:     "Float32",
	FieldTypeFloat64:     "Float64",
	FieldTypeString:      "String",
	FieldTypeByteSlice:   "ByteSlice",
	FieldTypeInt8Slice:   "Int8Slice",
	FieldTypeVectorPoint: "VectorPoint",
	FieldTypeScalar:      "Scalar",
	FieldTypeEnum:        "Enum",
	FieldTypeStruct:      "Struct",
	FieldTypeTable:       "Table",
	FieldTypeUnion:       "Union",
	FieldTypeArray:       "Array",
	FieldTypeFlexTable:   "FlexTable",
	FieldTypeCustomized:  "Customized",
}

var EnumValuesFieldType = map[string]FieldType{
	"None":        FieldTypeNone,
	"UType":       FieldTypeUType,
	"Bool":        FieldTypeBool,
	"Int8":        FieldTypeInt8,
	"Uint8":       FieldTypeUint8,
	"Byte":        FieldTypeByte,
	"Int16":       FieldTypeInt16,
	"Uint16":      FieldTypeUint16,
	"Int32":       FieldTypeInt32,
	"Uint32":      FieldTypeUint32,
	"Int64":       FieldTypeInt64,
	"Uint64":      FieldTypeUint64,
	"Float32":     FieldTypeFloat32,
	"Float64":     FieldTypeFloat64,
	"String":      FieldTypeString,
	"ByteSlice":   FieldTypeByteSlice,
	"Int8Slice":   FieldTypeInt8Slice,
	"VectorPoint": FieldTypeVectorPoint,
	"Scalar":      FieldTypeScalar,
	"Enum":        FieldTypeEnum,
	"Struct":      FieldTypeStruct,
	"Table":       FieldTypeTable,
	"Union":       FieldTypeUnion,
	"Array":       FieldTypeArray,
	"FlexTable":   FieldTypeFlexTable,
	"Customized":  FieldTypeCustomized,
}

func (v FieldType) String() string {
	if s, ok := EnumNamesFieldType[v]; ok {
		return s
	}
	return "FieldType(" + strconv.FormatInt(int64(v), 10) + ")"
}
