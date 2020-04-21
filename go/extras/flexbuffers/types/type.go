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
