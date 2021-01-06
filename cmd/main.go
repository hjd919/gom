package main

import (
	"github.com/hjd919/gom"
)

type ExcelRow struct {
	Name      string `xlsx:"0" comment:"达人账号名称"`
	ExpiresIn string `xlsx:"1" comment:"授权到期时间"`
}

// 导出
func ExcelExport() {
	var exportRows = make([]interface{}, 0, 2) // 导入的数据
	exportRows = append(exportRows, &ExcelRow{Name: "a", ExpiresIn: "b"})
	exportRows = append(exportRows, &ExcelRow{Name: "c", ExpiresIn: "v"})
	err := gom.NewExcel().Export(&gom.ExportParams{
		FilePath:  "./text2.xlsx", // 导入文件名
		SheetName: "生活",           // 页签名
		Rows:      exportRows,     // 导入数据，二维数组
	})
	if err != nil {
		return
	}
}

func main() {
	ExcelExport()
}
