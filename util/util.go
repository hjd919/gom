package util

import "encoding/json"

func ConvJsonStr(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
