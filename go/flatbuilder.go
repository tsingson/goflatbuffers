package flatbuffers

// FlatBuilder flatbuffers builder
type FlatBuilder struct {
	header Header
	VirtualTable
	vtables []*VirtualTable // registed VTable to
	root    bool
	vactors []VField
}

// NewFlatBuilder new flatbuffers builder
func NewFlatBuilder() *FlatBuilder {
	f := &FlatBuilder{}
	f.root = false
	f.header = Header{}
	return f
}
