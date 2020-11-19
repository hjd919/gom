package gom

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gogf/gf/util/gconv"
)

func PageNum(num int) int {
	if num <= 0 {
		return 1
	}
	return num
}

func PageSize(size int) int {
	if size <= 0 {
		return 10
	}
	return size
}

func ConvJsonStr(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func JsonEncode(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func JsonDecode(str string, v interface{}) error {
	err := json.Unmarshal(gconv.Bytes(str), v)
	return err
}

func Log(v interface{}) {
	log.Println(v)
	return
}

func Parse(ymdhis string) time.Time {
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", ymdhis, time.Local)
	return timeObj
}
