package crypto

const (
	cleanBytesLen = 64
	alphabet      = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var (
	cleanBytes = make([]byte, cleanBytesLen)
)

func cleanSensitiveData(b []byte) {
	l := len(b)

	for ; l > cleanBytesLen; l -= cleanBytesLen {
		copy(b[l-cleanBytesLen:l], cleanBytes)
	}

	if l > 0 {
		copy(b[0:l], cleanBytes[0:l])
	}
}

func repeatByteSequence(input []byte, length int) []byte {
	var (
		sequence = make([]byte, length)
		unit     = len(input)
	)

	j := length / unit * unit
	for i := 0; i < j; i += unit {
		copy(sequence[i:length], input)
	}
	if j < length {
		copy(sequence[j:length], input[0:length-j])
	}

	return sequence
}

func base64For24Bit(src []byte) []byte {
	if len(src) == 0 {
		return []byte{}
	}

	dstLen := (len(src)*8 + 5) / 6
	dst := make([]byte, dstLen)

	di, si := 0, 0
	n := len(src) / 3 * 3
	for si < n {
		val := uint(src[si+2])<<16 | uint(src[si+1])<<8 | uint(src[si])
		dst[di+0] = alphabet[val&0x3f]
		dst[di+1] = alphabet[val>>6&0x3f]
		dst[di+2] = alphabet[val>>12&0x3f]
		dst[di+3] = alphabet[val>>18]
		di += 4
		si += 3
	}

	rem := len(src) - si
	if rem == 0 {
		return dst
	}

	val := uint(src[si+0])
	if rem == 2 {
		val |= uint(src[si+1]) << 8
	}

	dst[di+0] = alphabet[val&0x3f]
	dst[di+1] = alphabet[val>>6&0x3f]
	if rem == 2 {
		dst[di+2] = alphabet[val>>12]
	}
	return dst
}
