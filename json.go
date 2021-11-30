package util

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
)

// replace encoding/json with jsoniter

func JsonIter() jsoniter.API {
	return jsoniter.ConfigCompatibleWithStandardLibrary
}

// JsonDecode unmarshal
func JsonDecode(s []byte, data interface{}) (err error) {
	if len(s) < 1 {
		return errors.New("nil")
	}
	err = JsonIter().Unmarshal(s, &data)
	if err != nil {
		return
	}
	return
}

// JsonEncode marshall
func JsonEncode(data interface{}) (marshal []byte, err error) {
	marshal, err = JsonIter().Marshal(&data)
	if err != nil {
		return
	}
	return
}
