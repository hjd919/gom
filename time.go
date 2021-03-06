package gom

import "time"

const DATETIME = "2006-01-02 15:03:04"
const DATE = "2006-01-02"
const DATELOG = "20060102"
const TIME = "15:03:04"

// 获取iso8601time = "2006-01-02T15:04:05-0700"
// 返回如：2021-01-08T20:52:55+0800
func Iso8601(t time.Time) string {
	return t.Format(time.RFC3339)
}

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

// 获取昨天日期
func Yesterday() string {
	yesterdayStr := time.Now().AddDate(0, 0, -1).Format(DATE)
	return ParseDate(yesterdayStr).Format(DATE)
}

// 获取昨天0点时间戳
func YesterdaySecond() int64 {
	yesterdayStr := time.Now().AddDate(0, 0, -1).Format(DATE)
	return ParseDate(yesterdayStr).Unix()
}

// 获取当天日期
func NowDate() string {
	return time.Now().Format(DATE)
}

// 获取当天0点时间戳
func NowDateSecond() int64 {
	return ParseDate(NowDate()).Unix()
}

// 获取当前秒时间戳
func NowSecond() int64 {
	return time.Now().Unix()
}

// 获取当前毫秒时间戳
func NowMillis() int64 {
	return Millis(time.Now())
}

// 获取某个时间obj的毫秒时间戳
func Millis(timeObj time.Time) int64 {
	return timeObj.UnixNano() / 1e6
}

// 获取上个月第一天0点
func GetFirstDateOfLastMonth() time.Time {
	d := time.Now()

	d = d.AddDate(0, -1, -d.Day()+1)
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取本周周一的日期
func GetFirstDateOfWeek() (weekStartDate time.Time) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return
}

// 获取上周的周一日期
func GetLastWeekFirstDate() (lastWeekMonday time.Time) {
	thisWeekMonday := GetFirstDateOfWeek()
	thisWeekMondayStr := thisWeekMonday.Format("2006-01-02")
	TimeMonday, _ := time.ParseInLocation("2006-01-02", thisWeekMondayStr, time.Local)
	lastWeekMonday = TimeMonday.AddDate(0, 0, -7)
	return
}
