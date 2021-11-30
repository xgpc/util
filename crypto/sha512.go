package crypto

import (
	"bytes"
	"crypto/sha512"
	"errors"
)

func Sha512WithSalt(key []byte, salt []byte) (string, error) {
	saltLen := len(salt)
	if saltLen > 16 {
		return "", errors.New("sha512 salt max length is 16")
	}

	keyLen := len(key)
	h := sha512.New()

	// compute sumB
	// step 4-8
	h.Write(key)
	h.Write(salt)
	h.Write(key)
	sumB := h.Sum(nil)

	// Compute sumA
	// step 1-3, 9-12
	h.Reset()
	h.Write(key)
	h.Write(salt)
	h.Write(repeatByteSequence(sumB, keyLen))
	for i := keyLen; i > 0; i >>= 1 {
		if i%2 == 0 {
			h.Write(key)
		} else {
			h.Write(sumB)
		}
	}
	sumA := h.Sum(nil)
	cleanSensitiveData(sumB)

	// Compute seqP
	// step 13-16
	h.Reset()
	for i := 0; i < keyLen; i++ {
		h.Write(key)
	}
	seqP := repeatByteSequence(h.Sum(nil), keyLen)

	// Compute seqS
	// step 17-20
	h.Reset()
	for i := 0; i < 16+int(sumA[0]); i++ {
		h.Write(salt)
	}
	seqS := repeatByteSequence(h.Sum(nil), saltLen)

	// step 21
	for i := 0; i < 5000; i++ {
		h.Reset()

		if i&1 != 0 {
			h.Write(seqP)
		} else {
			h.Write(sumA)
		}
		if i%3 != 0 {
			h.Write(seqS)
		}
		if i%7 != 0 {
			h.Write(seqP)
		}
		if i&1 != 0 {
			h.Write(sumA)
		} else {
			h.Write(seqP)
		}
		copy(sumA, h.Sum(nil))
	}
	cleanSensitiveData(seqP)
	cleanSensitiveData(seqS)

	// make output
	buf := bytes.Buffer{}
	buf.Grow(107 + len(salt))
	buf.Write([]byte("$6$"))

	buf.Write(salt)
	buf.WriteByte('$')
	buf.Write(base64For24Bit([]byte{
		sumA[42], sumA[21], sumA[0],
		sumA[1], sumA[43], sumA[22],
		sumA[23], sumA[2], sumA[44],
		sumA[45], sumA[24], sumA[3],
		sumA[4], sumA[46], sumA[25],
		sumA[26], sumA[5], sumA[47],
		sumA[48], sumA[27], sumA[6],
		sumA[7], sumA[49], sumA[28],
		sumA[29], sumA[8], sumA[50],
		sumA[51], sumA[30], sumA[9],
		sumA[10], sumA[52], sumA[31],
		sumA[32], sumA[11], sumA[53],
		sumA[54], sumA[33], sumA[12],
		sumA[13], sumA[55], sumA[34],
		sumA[35], sumA[14], sumA[56],
		sumA[57], sumA[36], sumA[15],
		sumA[16], sumA[58], sumA[37],
		sumA[38], sumA[17], sumA[59],
		sumA[60], sumA[39], sumA[18],
		sumA[19], sumA[61], sumA[40],
		sumA[41], sumA[20], sumA[62],
		sumA[63],
	}))
	return buf.String(), nil
}
