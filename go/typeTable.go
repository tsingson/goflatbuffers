package flatbuffers

/**
TypeCode define in c++, redesign for go

 Basic type info cost just 16bits per field!
struct TypeCode {
  uint16_t base_type : 4;  // ElementaryType
  uint16_t is_vector : 1;
  int16_t sequence_ref : 11;  // Index into type_refs below, or -1 for none.
};
*/

// TypeCode define type code  for scalar / enum in IDL
type TypeCode struct {
	Id             uint16
	Offset         uint16
	Name           string
	BaseType       FieldType
	FieldType      FieldType
	DefaultInteger int64
	DefaultReal    float64
	Element        FieldType
	Index          int32
	FixedLength    uint16
	IsVector       bool // set true if vector field
	IsDeprecated   bool // set true if deprecated field
	SequenceRef    TypeTable
}

// TypeTable define type table, for enum / struct / union / table in IDL
// SequenceType : table / struct / union / enum , only 4 type in SequenceType
type TypeTable struct {
	SequenceType FieldType
	IsRootType   bool // set true if set root_type
	Minalign     int32
	Bytesize     int32
	NumsElement  int // how many field inside
	FixedLength  int // default is 0,  set if fixed size like struct
	Name         []byte
	RootType     []byte // root_type name
	Identifier   []byte // set if root table with file_identifier

	Fields []TypeCode
}

// TypeFunction a func return TypeTable
type TypeFunction func() TypeTable
