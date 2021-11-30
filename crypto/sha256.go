package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
