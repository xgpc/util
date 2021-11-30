package util

func InArrayString(elem string, arr []string) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

/**
 * 判断目标值ID是否在枚举值数组中
 * @param	elem	目标值
 * @param	arr		uint64数组
 */
func InArrayUint64(elem uint64, arr []uint64) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

/**
 * 判断目标值是否在枚举值数组中
 * @param	elem	目标值
 * @param	arr		uint8数组
 */
func InArrayUint8(elem uint8, arr []uint8) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

/**
 * 判断目标值是否在枚举值数组中
 * @param	elem	目标值
 * @param	arr		int数组
 */
func InArrayInt(elem int, arr []int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

/**
 * 判断目标值是否在枚举值数组中
 * @param	elem	目标值
 * @param	arr		byte数组
 */
func InArrayByte(elem byte, arr []byte) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

/**
 * 字符串数组去重 - 用于ID数组去重
 * @param	arr		字符串数组
 */
func UniqueArrayString(arr []string) []string {
	var m = map[string]bool{}
	for _, v := range arr {
		m[v] = true
	}
	a := []string{}
	for v := range m {
		a = append(a, v)
	}
	return a
}

/**
 * 整型数组去重 - 用于ID数组去重 todo:排序错乱
 * @param	arr		字符串数组
 */
func UniqueArrayUint64(arr []uint64) []uint64 {
	var m = map[uint64]bool{}
	for _, v := range arr {
		m[v] = true
	}
	a := []uint64{}
	for v := range m {
		a = append(a, v)
	}
	return a
}

/**
 * 整型数组去重 - 用于ID数组去重
 * @param	arr		字符串数组
 */
func UniqueArrayUint8(arr []uint8) []uint8 {
	var m = map[uint8]bool{}
	for _, v := range arr {
		m[v] = true
	}
	a := []uint8{}
	for v := range m {
		a = append(a, v)
	}
	return a
}

/**
 * 整型数组去重 - 用于ID数组去重
 * @param	arr		字符串数组
 */
func UniqueArrayInt(arr []int) []int {
	var m = map[int]bool{}
	for _, v := range arr {
		m[v] = true
	}
	a := []int{}
	for v := range m {
		a = append(a, v)
	}
	return a
}

/**
 * 大数组分割成多个小数组 - 用于批量处理数据的时候，防止SQL参数过大
 * @param	arr		大数组
 * @param	per		每个小数组的长度
 * @return			小数组
 */
func CutArrayString(arr []string, per int) [][]string {

	array := [][]string{}
	temp := []string{}
	var index = 0

	for _, v := range arr {
		temp = append(temp, v)
		index++
		if index > per {
			array = append(array, temp)
			index = 0
			temp = []string{}
		}
	}

	if len(temp) > 0 {
		array = append(array, temp)
	}

	return array
}

/**
 * 大数组分割成多个小数组 - 用于批量处理数据的时候，防止SQL参数过大
 * @param	arr		大数组
 * @param	per		每个小数组的长度
 * @return			小数组
 */
func CutArrayUint64(arr []uint64, per int) [][]uint64 {
	array := [][]uint64{}
	temp := []uint64{}
	var index = 0

	for _, v := range arr {
		temp = append(temp, v)
		index++
		if index >= per {
			array = append(array, temp)
			index = 0
			temp = []uint64{}
		}
	}

	if len(temp) > 0 {
		array = append(array, temp)
	}

	return array
}

/**
 * 切片数组中移除指定的值
 * @param	arr		原始切片
 * @param	exclude	被排除的切片
 * @return			切片
 */
func ExcludeArr(arr []uint64, exclude []uint64) []uint64 {
	if len(arr) == 0 || len(exclude) == 0 {
		return arr
	}

	for _, v1 := range exclude {
		for i, v2 := range arr {
			if v1 != v2 {
				continue
			}
			arr = append(arr[0:i], arr[i+1:]...)
		}
	}

	return arr
}
