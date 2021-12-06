package util

import (
	"reflect"
	"sort"
)

// 将传入的变量转换成 - map[string]interface{}
func ReflectToMap(m interface{}, filter []string) map[string]interface{} {
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)

	var data = map[string]interface{}{}
	for k := 0; k < t.NumField(); k++ {

		var key = t.Field(k).Name

		if key[0] >= 'a' && key[0] <= 'z' {
			continue
		}

		var value = v.Field(k).Interface()

		var skip = false
		for _, f := range filter {
			if key == f {
				skip = true
				continue
			}
		}
		if skip {
			continue
		}

		data[key] = value
	}
	return data
}

// 将传入的变量转换成接口签名校验需要的格式
func ReflectToApiSignData(p interface{}) (res []string) {
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	var data = map[string]string{}
	for k := 0; k < t.NumField(); k++ {
		var key = t.Field(k).Name
		if key[0] >= 'a' && key[0] <= 'z' {
			continue
		}
		if key == "Base" || key == "ResSign" || key == "ResPayWx" {
			continue
		}
		var value = v.Field(k).Interface()
		switch vv := value.(type) {
		case string:
			data[key] = vv
		case int:
			data[key] = IntToStr(vv)
		case uint64:
			data[key] = Uint64ToStr(vv)
		case uint8:
			data[key] = Uint8ToStr(vv)
		case float64:
			data[key] = FloatToStr(vv)
		case bool:
			if vv {
				data[key] = "true"
			} else {
				data[key] = "false"
			}
		default:

			str, err := JsonEncode(vv)
			if err != nil {
				panic(err)
			}
			data[key] = string(str)
		}
	}

	var keyArr []string
	for k := range data {
		keyArr = append(keyArr, k)
	}
	sort.Strings(keyArr)

	for _, k := range keyArr {
		res = append(res, k+"="+data[k])
	}

	return
}
