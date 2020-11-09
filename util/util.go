package util

import (
	"encoding/json"
	"log"
	"time"
)

func ConvJsonStr(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func Log(v interface{}) {
	log.Println(v)
	return
}

func Parse(ymdhis string) time.Time {
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", ymdhis, time.Local)
	return timeObj
}
