package flatbuffers

// Header root B
type Header struct {
	identifier []byte   // file identifier
	soffset    SOffsetT // point to root table
}

func (h *Header) SetIdentifier(name []byte) {
	copy(h.identifier, name)
}
