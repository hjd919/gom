package gom

import (
	"fmt"
	"reflect"

	"github.com/tealeg/xlsx/v3"
)

func New() *Excel {
	return &Excel{}
}

type Excel struct {
	//file
	file *xlsx.File
	// 标题
	titles []string
}

// 导出
type ExportParams struct {
	//sheet
	SheetName string
	//导出的文件全路径
	FilePath string
	Data     []interface{}
}

func (e *Excel) getTitles(stru interface{}) {
	t := reflect.TypeOf(stru).Elem()
	for i := 0; i < t.NumField(); i++ {
		e.titles = append(e.titles, t.Field(i).Tag.Get("comment"))
	}
}

func (e *Excel) Export(params *ExportParams) (err error) {
	if len(params.Data) == 0 {
		return fmt.Errorf("list is empty")
	}
	// 创建excel
	if e.file == nil {
		e.file = xlsx.NewFile()
	}
	file := e.file

	// sheet
	sheet, err := file.AddSheet(params.SheetName)
	if err != nil {
		return
	}

	//title
	e.getTitles(params.Data[0])
	row := sheet.AddRow()
	row.WriteSlice(e.titles, -1)

	// data
	for _, vorRow := range params.Data {
		row := sheet.AddRow()
		row.WriteStruct(vorRow, -1)
	}

	//save file
	err = file.Save(params.FilePath)
	if err != nil {
		return
	}
	return
}

// 导入
type ImportParams struct {
	//sheet
	SheetName string
	//导出的文件全路径
	FilePath  string
	RowHandle func(r *xlsx.Row) error
}

func (e *Excel) Import(params *ImportParams) (err error) {
	// 创建excel
	if e.file == nil {
		e.file, err = xlsx.OpenFile(params.FilePath)
		if err != nil {
			return
		}
	}
	file := e.file

	// sheet
	sh, ok := file.Sheet[params.SheetName]
	if !ok {
		return fmt.Errorf("sheet not found")
	}

	// data
	err = sh.ForEachRow(params.RowHandle)
	if err != nil {
		return
	}
	return
}
