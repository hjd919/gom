package gom

import (
	"reflect"
	"strings"
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
