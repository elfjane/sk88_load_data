package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

type Excel_data struct {
	Name     string
	Row_name string
	Type     string
	Default  string
}

// 取讀 中文英對照表 excel資料
func Read_excel(filename string) []Excel_data {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// 获取 Sheet1 上所有单元格

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}
	excel_datas := []Excel_data{}
	var excel_data Excel_data
	var first bool = true
	for _, row := range rows {
		if first {
			first = false
			continue

		}
		excel_data.Name = row[0]
		excel_data.Row_name = row[1]
		excel_data.Type = row[2]
		excel_data.Default = row[3]
		excel_datas = append(excel_datas, excel_data)
	}
	//record :=  "123"

	return excel_datas
}
