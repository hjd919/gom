package test

import (
	"fmt"
	"testing"
	"time"
)

// 上个月
func TestDate(t *testing.T) {
	// fmt.Println(int(time.Now().AddDate(0, -1, 0).Month()))
	d := time.Now()
	d = d.AddDate(0, -1, -d.Day()+1)
	res := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	fmt.Println(res)
}
