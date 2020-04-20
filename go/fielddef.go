package flatbuffers

type StructDef struct {
	fixed             bool
	predecl           bool
	sortbysize        bool
	has_key           bool
	minalign          int
	bytesize          int
	original_location []byte
}

type FieldDef struct {
	deprecated        bool
	required          bool
	key               bool
	shared            bool
	native_inline     bool
	flexbuffers       bool
	nested_flatbuffer *StructDef
	padding           int
}
