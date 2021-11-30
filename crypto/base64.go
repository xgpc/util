package crypto

import (
	"encoding/base64"
)

func Base64Decode(str string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
