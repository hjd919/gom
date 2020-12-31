package gom

import (
	"log"
	"reflect"
	"strings"
	"time"
)

/*
orderArr := []string{"id desc"}
orderStr := ""
if req.Sort != "" {
	switch req.Sort {
	case "createtime":
		orderArr = append(orderArr, "ys_team.id desc")
}
orderStr = strings.Join(orderArr, ",")
*/

// 数字范围，数组转化为开始和结束数字
func NumRange(numRange []int64) (startNum, endNum int64) {
	if len(numRange) == 0 {
		return
	}
	// eq
	if len(numRange) == 1 {
		startNum = numRange[0]
		endNum = numRange[0]
		return
	}
	// between
	if len(numRange) == 2 {
		startNum = numRange[0]
		endNum = numRange[1]
		return
	}
	return
}

// 日期范围，数组转化为开始和结束时间戳
func DateRange(dateRange []string) (startTime, endTime int64) {
	if len(dateRange) == 0 {
		return
	}
	startTimeObj, err := time.ParseInLocation(DATE, dateRange[0], time.Local)
	if err != nil {
		log.Println("DateRange=", err)
		return
	}
	endTimeObj, err := time.ParseInLocation(DATE, dateRange[1], time.Local)
	if err != nil {
		log.Println("DateRange=", err)
		return
	}
	return startTimeObj.Unix(), endTimeObj.AddDate(0, 0, 1).Unix()
}

// 获取select字段（通过结构体tag获取所需的字段）
func Fields(point interface{}) (str string) {
	t := reflect.TypeOf(point).Elem()
	var orms []string
	for i := 0; i < t.NumField(); i++ {
		orms = append(orms, t.Field(i).Tag.Get("orm"))
	}
	if len(orms) == 0 {
		panic("GetFieldsByOrm err")
	}
	str = strings.Join(orms, ",")
	return
}

// 获取分页参数
func PageParam(num, size int) (pageOffset int, pageSize int) {
	pageSize = PageSize(size)
	pageOffset = (PageNum(num) - 1) * pageSize
	return
}

// 格式化分页
func Page(num, size int) (pageNum int, pageSize int) {
	return PageNum(num), PageSize(size)
}

// 获取分页页码
func PageNum(num int) int {
	if num <= 0 {
		return 1
	}
	return num
}

// 获取分页长度
func PageSize(size int) int {
	if size <= 0 {
		return 15
	}
	return size
}
