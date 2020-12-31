package gom

import (
	"encoding/json"

	"github.com/gogf/gf/util/gconv"
)

func ConvJsonStr(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// go结构转字符串
func JsonEncode(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// 字符串转go结构
func JsonDecode(str string, v interface{}) error {
	err := json.Unmarshal(gconv.Bytes(str), v)
	return err
}
