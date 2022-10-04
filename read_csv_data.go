package main

import (
	"encoding/csv"
	"log"
	"os"
)

// 讀取 csv 資料
func Read_csv_data(filename string) [][]string {
	/*
	   	in := `first_name;last_name;username
	   "Rob";"Pike";rob
	   # lines beginning with a # character are ignored
	   Ken;Thompson;ken
	   "Robert";"Griesemer";"gri"
	   `
	   	//f = strings.NewReader(in)
	*/

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	r := csv.NewReader(f)
	//r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Print(records)
	return records
}
