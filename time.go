package gom

import "time"

const DATETIME = "2006-01-02 15:03:04"
const DATE = "2006-01-02"
const DATELOG = "20060102"
const TIME = "15:03:04"

// 废除
func Parse(ymdhis string) time.Time {
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", ymdhis, time.Local)
	return timeObj
}

// 获取ymdhis格式的obj
func ParseDatetime(ymdhis string) time.Time {
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", ymdhis, time.Local)
	return timeObj
}

// 获取ymd格式的obj
func ParseDate(ymd string) time.Time {
	timeObj, _ := time.ParseInLocation("2006-01-02", ymd, time.Local)
	return timeObj
}

// 获取当前毫秒时间戳
func NowMillis() int64 {
	return Millis(time.Now())
}

// 获取某个时间obj的毫秒时间戳
func Millis(timeObj time.Time) int64 {
	return timeObj.UnixNano() / 1e6
}

// 获取上个月第一天
func GetFirstDateOfLastMonth(d time.Time) time.Time {
	d = d.AddDate(0, -1, -d.Day()+1)
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
