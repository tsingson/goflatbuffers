package flatbuffers

// VField field of VTable
type VField interface {
	// ID() uint
	// IsDeprecated() bool

	IsVector() bool   // is vector or not
	IsFinished() bool // check is serialize operation done or not
	Payload() []byte
	Release() bool
	ByteSize() int
	ObjectSize() int
	Type() FieldType //  scalar / union / vector
	Press() error    // worker for goroutine
}

// TableHeader define table header
type TableHeader interface {
	TableSize() int
	FinishByte() []byte
}
