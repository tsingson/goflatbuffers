package flatbuffers

// Struct wraps a byte slice and provides read access to its data.
//
// Structs do not have a VTable.
type Struct struct {
	Table
}
