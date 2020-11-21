package gom

func Order(sortType, sortField string) string {
	if sortType != "" && sortField != "" {
		return sortType + " " + sortField
	} else if sortType != "" && sortField == "" {
		return sortType
	}
	return "id desc"
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
