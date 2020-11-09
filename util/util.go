package util

import (
	"encoding/json"
	"log"
)

func ConvJsonStr(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func Log(v interface{}) {
	log.Println(v)
	return
}
