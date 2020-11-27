package gom

import (
	"fmt"

	"github.com/gogf/gf/util/gconv"
)

func GetRate(part, total int64) float64 {
	if total == 0 {
		return 0
	}
	resultFl := gconv.Float64(part) / gconv.Float64(total) * 100
	return gconv.Float64(fmt.Sprintf("%.2f", resultFl))
}

// true => 1 false => 0
func BoolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}
