package types


func GetBitWidth(b uint8) uint8 {
	return uint8((b << 6) >> 6)
}

func GetFBType(b uint8) uint8 {
	return uint8(b >> 2)
}

func SetFBType(w, t uint8) uint8 {
	if t > 3 {
		return uint8(254)
	}
	return uint8(t + w<<2)
}

func WriteUint8(buf []byte, n uint8) {
	buf[0] = byte(n)
}

func GetInt8(buf []byte) (n int8) {
	n = int8(buf[0])
	return
}


func IsInline(t FBT) bool {
	return (t <= FBTFLOAT) || (t == FBTBOOL)
}

func IsTypedVectorElementType(t FBT) bool {
	return ((t >= FBTINT) && (t <= FBTSTRING)) || (t == FBTBOOL)
}

func IsTypedVector(t FBT) bool {
	return ((t >= FBTVECTOR_INT) && (t <= FBTVECTOR_STRING_DEPRECATED)) || (t == FBTVECTOR_BOOL)
}

func IsFixedTypedVector(t FBT) bool {
	return (t >= FBTVECTOR_INT2) && (t <= FBTVECTOR_FLOAT4)
}
