package flatbuffers

// Pad prepad
func Prepad(size, additionalBytes int) (int, int) {
	alignSize := (^(additionalBytes)) + 1
	alignSize &= (size - 1)
	return alignSize, additionalBytes + alignSize
}
