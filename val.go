package util

import (
	"encoding/json"
	"errors"
	"strconv"
)

func ToJsonString(v interface{}) (string, error) {
	bt, err := json.Marshal(v)
	return string(bt), err
}

func FloatToStr(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func IntToStr(v int) string {
	return strconv.Itoa(v)
}

func Int64ToStr(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Uint64ToStr(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func Uint8ToStr(v uint8) string {
	return strconv.FormatUint(uint64(v), 10)
}

func FloatToInt(v float64) int {
	return int(v)
}

func IntToFloat(v int) float64 {
	return float64(v)
}

func StrToInt(v string) (int, error) {
	if v == "" {
		return 0, nil
	}
	res, err := strconv.Atoi(v)
	return res, err
}

func StrToUint8(v string) (uint8, error) {
	if v == "" {
		return 0, nil
	}
	res, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	if res < 0 || res > 255 {
		return 0, errors.New("[error] 超出范围")
	}
	return uint8(res), err
}

func StrToUint64(v string) (uint64, error) {
	if v == "" {
		return 0, nil
	}
	res, err := strconv.ParseUint(v, 10, 64)
	return res, err
}

func StrToInt64(v string) (int64, error) {
	if v == "" {
		return 0, nil
	}
	res, err := strconv.ParseInt(v, 10, 64)
	return res, err
}

func StrToFloat(v string) (float64, error) {
	if v == "" {
		return 0, nil
	}
	res, err := strconv.ParseFloat(v, 64)
	return res, err
}
