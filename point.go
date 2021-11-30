package util

import (
	"math"
)

const (
	PI     = math.Pi * 3000.0 / 180.0
	OFFSET = 0.00669342162296594323
	AXIS   = 6378245.0
)

//BD09toGCJ02 百度坐标系->火星坐标系
func BD09toGCJ02(lng, lat float64) (float64, float64) {
	x := lng - 0.0065
	y := lat - 0.006

	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*PI)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*PI)

	gLng := z * math.Cos(theta)
	gLat := z * math.Sin(theta)

	return gLng, gLat
}

//GCJ02toBD09 火星坐标系->百度坐标系
func GCJ02toBD09(lng, lat float64) (float64, float64) {
	z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*PI)
	theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*PI)

	bdLng := z*math.Cos(theta) + 0.0065
	bdLat := z*math.Sin(theta) + 0.006

	return bdLng, bdLat
}

//WGS84toGCJ02 WGS84坐标系->火星坐标系
func WGS84toGCJ02(lng, lat float64) (float64, float64) {
	if isOutOFChina(lng, lat) {
		return lng, lat
	}

	return delta(lng, lat)
}

//GCJ02toWGS84 火星坐标系->WGS84坐标系
func GCJ02toWGS84(lng, lat float64) (float64, float64) {
	if isOutOFChina(lng, lat) {
		return lng, lat
	}

	mgLng, mgLat := delta(lng, lat)

	return lng*2 - mgLng, lat*2 - mgLat
}

//BD09toWGS84 百度坐标系->WGS84坐标系
func BD09toWGS84(lng, lat float64) (float64, float64) {
	lng, lat = BD09toGCJ02(lng, lat)
	return GCJ02toWGS84(lng, lat)
}

//WGS84toBD09 WGS84坐标系->百度坐标系
func WGS84toBD09(lng, lat float64) (float64, float64) {
	lng, lat = WGS84toGCJ02(lng, lat)
	return GCJ02toBD09(lng, lat)
}

func delta(lng, lat float64) (float64, float64) {
	dLat, dLng := transform(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - OFFSET*magic*magic
	sqrtMagic := math.Sqrt(magic)

	dLat = (dLat * 180.0) / ((AXIS * (1 - OFFSET)) / (magic * sqrtMagic) * math.Pi)
	dLng = (dLng * 180.0) / (AXIS / sqrtMagic * math.Cos(radLat) * math.Pi)

	return lng + dLng, lat + dLat
}

func transform(lng, lat float64) (x, y float64) {
	var lngLat = lng * lat
	var absX = math.Sqrt(math.Abs(lng))
	var lngPi, latPi = lng * math.Pi, lat * math.Pi
	var d = 20.0*math.Sin(6.0*lngPi) + 20.0*math.Sin(2.0*lngPi)
	x, y = d, d
	x += 20.0*math.Sin(latPi) + 40.0*math.Sin(latPi/3.0)
	y += 20.0*math.Sin(lngPi) + 40.0*math.Sin(lngPi/3.0)
	x += 160.0*math.Sin(latPi/12.0) + 320*math.Sin(latPi/30.0)
	y += 150.0*math.Sin(lngPi/12.0) + 300.0*math.Sin(lngPi/30.0)
	x *= 2.0 / 3.0
	y *= 2.0 / 3.0
	x += -100.0 + 2.0*lng + 3.0*lat + 0.2*lat*lat + 0.1*lngLat + 0.2*absX
	y += 300.0 + lng + 2.0*lat + 0.1*lng*lng + 0.1*lngLat + 0.1*absX
	return
}

func isOutOFChina(lng, lat float64) bool {
	return !(lng > 72.004 && lng < 135.05 && lat > 3.86 && lat < 53.55)
}

// 两点之间的距离，单位（米）
func Distance(lng1, lat1, lng2, lat2 float64) float64 {
	pi := math.Pi / 180

	a1 := math.Sin(lat1*pi) * math.Sin(lat2*pi)
	a2 := math.Cos(lat1*pi) * math.Cos(lat2*pi) * math.Cos((lng1-lng2)*pi)
	a := a1 + a2

	b := math.Atan(-a/math.Sqrt(-a*a+1)) + 2*math.Atan(1)
	d := b * 3437.74677 * 1.1508 * 1.6093470878864446 * 1000
	if d == math.NaN() {
		return 0
	}
	return d
}
