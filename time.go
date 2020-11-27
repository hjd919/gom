package gom

import "time"

const DATETIME = "2006-01-02 15:03:04"
const DATE = "2006-01-02"
const TIME = "15:03:04"

func Parse(ymdhis string) time.Time {
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", ymdhis, time.Local)
	return timeObj
}

func ParseDatetime(ymdhis string) time.Time {
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", ymdhis, time.Local)
	return timeObj
}

func ParseDate(ymd string) time.Time {
	timeObj, _ := time.ParseInLocation("2006-01-02", ymd, time.Local)
	return timeObj
}
