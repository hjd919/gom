package gom

import (
	"fmt"

	"github.com/gogf/gf/util/gconv"
)

// 获取百分数
func GetRate(part, total int64) float64 {
	if total == 0 {
		return 0
	}
	resultFl := gconv.Float64(part) / gconv.Float64(total) * 100
	return gconv.Float64(fmt.Sprintf("%.2f", resultFl))
}

// 布尔值转整数
func BoolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}

// 整数转布尔值
func IntToBool(num int) bool {
	if num > 0 {
		return true
	}
	return false
}
