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

// filter - num
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

// filter - date
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
	return startTimeObj.Unix(), endTimeObj.Unix()
}

func Order(sortType, sortField string) string {
	if sortType != "" && sortField != "" {
		return sortType + " " + sortField
	} else if sortType != "" && sortField == "" {
		return sortType
	}
	return "id desc"
}

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

func Page(num, size int) (pageNum int, pageSize int) {
	return PageNum(num), PageSize(size)
}

func PageNum(num int) int {
	if num <= 0 {
		return 1
	}
	return num
}

func PageSize(size int) int {
	if size <= 0 {
		return 15
	}
	return size
}
