package util

import (
	"time"
)

const (
	DateFormat    = "2006-01-02 15:04:05"
	DayFormatFile = "2006_01_02"
	DayFormat     = "2006-01-02"
	DayIntFormat  = "20060102"
	DayFormatYmdH = "2006-01-02-15"
)

var monthsMd = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func TimeUnix() uint32 {
	return uint32(time.Now().Unix())
}

func TimeYmdNow() string {
	return time.Now().Format(DayFormatFile)
}

func TimeYmd_Now2() string {
	return time.Now().Format(DayFormat)
}

func TimeDataToString(t int64) string {
	return time.Unix(t, 0).Format(DayFormat)
}

func TimeToString(t int64) string {
	return time.Unix(t, 0).Format(DayFormatYmdH)
}

func DayToUnix2359(day string) (int64, error) {
	t, err := time.ParseInLocation(DateFormat, day+" 23:59:59", time.Local)
	return t.Unix(), err
}

func DayToUnix0000(day string) (int64, error) {
	t, err := time.ParseInLocation(DateFormat, day+" 00:00:00", time.Local)
	return t.Unix(), err
}

func DayUint64(t *time.Time) (uint64, error) {
	s := t.Format(DayIntFormat)
	return StrToUint64(s)
}

func GetMonth(t *time.Time) int {
	return monthsMd[t.Month().String()]
}

func GetYmdInt(t *time.Time) (int, int, int) {
	return t.Year(), GetMonth(t), t.Day()
}

// TimeToYm 时间戳转年月
func TimeToYm(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01")
}

func TimeToY(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006")
}

func TimeTom(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("01")
}

func BeijingTime() time.Time {
	local, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = local
	return time.Now()
}

// TimeMPeriod 月份时间区间
func TimeMPeriod(timestamp int64) (start, end uint32) {
	ts := time.Unix(timestamp, 0)
	//time.Time格式
	firstOfMonth := time.Date(ts.Year(), ts.Month(), 1, 0, 0, 0, 0, ts.Location())
	last := time.Date(ts.Year(), ts.Month(), 1, 23, 59, 59, 0, ts.Location())
	lastOfMonth := last.AddDate(0, 1, -1)
	start = uint32(firstOfMonth.Unix())
	end = uint32(lastOfMonth.Unix())
	return
}

// TimeYPeriod 年区间
func TimeYPeriod(timestamp int64) (start, middle, end uint32) {
	local, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = local
	//	return time.Now()
	ts := time.Unix(timestamp, 0)
	ts.Location()
	//time.Time格式
	firstOfYear := time.Date(ts.Year(), time.January, 1, 0, 0, 0, 0, ts.Location())
	lastOfYear := time.Date(ts.Year(), time.January, 1, 23, 59, 59, 0, ts.Location()).
		AddDate(1, 0, -1)
	middleYear := time.Date(ts.Year(), time.January, 1, 23, 59, 59, 0, ts.Location()).
		AddDate(0, 6, -1)

	start = uint32(firstOfYear.Unix())
	middle = uint32(middleYear.Unix())
	end = uint32(lastOfYear.Unix())
	return
}
