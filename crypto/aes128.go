package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesCBCDecrypto(text, key, iv string) string {
	keyBytes := []byte(key)
	ivBytes := []byte(iv)
	textBytes := []byte(text)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	if len(textBytes) < blockSize {
		panic("密文太短")
	}
	if len(textBytes)%blockSize != 0 {
		panic("密文长度不合法")
	}

	if len(ivBytes) == 0 {
		ivBytes = textBytes[:aes.BlockSize]
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(textBytes, textBytes)
	res := PKCS7UnPadding(textBytes)
	return string(res)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
