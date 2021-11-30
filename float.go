package util

import (
	"math"
)

// 保留两位小数
func Round(f float64) float64 {
	return math.Round(f*100) / 100
}

// 保留六位小数（经纬度）
func Round6(f float64) float64 {
	return math.Round(f*1000000) / 1000000
}
