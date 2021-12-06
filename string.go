package util

import (
	"encoding/json"
	"github.com/axgle/mahonia"
	"net/url"
	"strconv"
	"strings"
)

func Split(str *string, sep string) []string {
	return strings.Split(*str, sep)
}

func StrLR(str *string, subStr string) string {
	pot := strings.Index(*str, subStr)
	if pot == -1 {
		return ""
	}

	subStrPot := pot + len(subStr)
	return (*str)[subStrPot:]
}

func StrLL(str *string, subStr string) string {

	pot := strings.Index(*str, subStr)
	if pot == -1 {
		return ""
	}

	return (*str)[0:pot]
}

func StrRL(str *string, subStr string) string {

	pot := strings.LastIndex(*str, subStr)
	if pot == -1 {
		return ""
	}

	return (*str)[0:pot]
}

func StrRR(str *string, subStr string) string {
	pot := strings.LastIndex(*str, subStr)
	if pot == -1 {
		return ""
	}

	pot += len(subStr)
	return (*str)[pot:]
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func StrIndex(str *string, subStr string) int {
	res := strings.Index(*str, subStr)
	if res >= 0 {
		prefix := []byte(*str)[0:res]
		rs := []rune(string(prefix))
		res = len(rs)
	}
	return res
}

func ConvertToByte(src *string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(*src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func Substr(str *string, start int, length ...int) string {
	rs := []rune(*str)
	var strLen = len(rs)

	if start < 0 {
		start = 0
	}

	var end = strLen
	if len(length) != 0 {
		end = start + length[0]
		if end > strLen {
			end = strLen
		}
	}

	return string(rs[start:end])
}

func Replace(str *string, old string, new string) string {
	return strings.Replace(*str, old, new, -1)
}

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' {
			data = append(data, '_')
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func CamelCaseString(s string) string {
	res := ""
	arr := strings.Split(s, "_")
	for _, v := range arr {
		vv := []rune(v)
		if bool(vv[0] >= 'a' && vv[0] <= 'z') {
			vv[0] -= 32
		}
		res += string(vv)
	}
	return res
}

func StrCut(str string, length int) string {
	rs := []rune(str)
	if len(rs) <= length {
		return str
	}

	return string(rs[0:length]) + "..."
}

func StrCutDirect(str string, length int) string {
	rs := []rune(str)
	if len(rs) <= length {
		return str
	}

	return string(rs[0:length])
}

func UrlEncode(value string) string {
	u := url.Values{}
	u.Set("tmp", value)
	res := u.Encode()
	return StrLR(&res, "tmp=")
}

func IsMobile(str string) bool {
	if len(str) != 11 {
		return false
	}
	for _, v := range str {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// Strval
//convert interface to string
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
