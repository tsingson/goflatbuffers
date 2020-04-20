package flatbuffers

import (
	"sync"

	"github.com/tsingson/goflatbuffers/go/bytepool"
)

// VectcorArray add a string vector
func (b *Vector) VectcorArray(off VOffsetT, s ...VField) VField {
	b.VOffset = off
	b.t = FieldTypeArray
	b.objectSize = len(s)

	add := 0
	p := make([]VField, 0)
	for _, v := range s {
		add += v.ByteSize()
		if !v.IsFinished() {
			p = append(p, v)
		}
	}
	// TODO: tsingson need performance check
	if len(p) > 0 {
		var wg sync.WaitGroup
		wg.Add(len(p))
		for _, v1 := range s {
			go func(v VField) {
				// TODO: handle error  ?
				_ = v.Press()
				wg.Done()
			}(v1)
		}
		wg.Wait()
	}

	if b.b == nil {
		b.b = bytepool.Get()
	}

	b.b.Reset()
	b.b.FixedLength(SizeSOffsetT + add)
	offset := SizeSOffsetT + add
	for j := b.objectSize - 1; j >= 0; j-- {
		offset = offset - s[j].ByteSize()
		copy(b.b.B[offset:], s[j].Payload())
		// s[j].Release()
	}
	b.EndVector(b.objectSize)
	b.finished = true

	return b
}

// StructArray add a string vector
func (b *Vector) StructArray(off VOffsetT, s ...VField) VField {
	b.VOffset = off
	b.t = FieldTypeArray
	b.objectSize = len(s)

	add := 0
	for _, v := range s {
		add += v.ByteSize()
	}
	if b.b == nil {
		b.b = bytepool.Get()
	}

	b.b.Reset()
	b.b.FixedLength(SizeSOffsetT + add)
	offset := SizeSOffsetT + add
	for j := b.objectSize - 1; j >= 0; j-- {
		offset = offset - s[j].ByteSize()

		//	fmt.Println(" off :", offset , " j: ", j, " bytesize: ", s[j].ByteSize())

		copy(b.b.B[offset:], s[j].Payload())
		// s[j].Release()
	}
	b.EndVector(b.objectSize)
	b.finished = true

	return b
}

// ScalarArray add a string vector
func (b *Vector) ScalarArray(off VOffsetT, s ...VField) VField {
	b.VOffset = off
	b.t = FieldTypeArray
	b.objectSize = len(s)

	add := 0
	for _, v := range s {
		add += v.ByteSize()
	}
	pad, total := Prepad(SizeSOffsetT, add+SizeSOffsetT)
	// fmt.Println(" pad: ", pad, " total: ", total)
	b.b.Reset()
	b.b.FixedLength(total).Pad(pad)
	offset := total - pad
	for j := b.objectSize - 1; j >= 0; j-- {
		offset = offset - s[j].ByteSize()
		// fmt.Println(" off :", offset, " j: ", j, " bytesize: ", s[j].ByteSize())
		copy(b.b.B[offset:], s[j].Payload())
		// s[j].Release()
	}

	b.EndVector(b.objectSize)
	b.finished = true

	return b
}
