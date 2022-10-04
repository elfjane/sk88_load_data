package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"cj_db"
	"cj_html"

	"cj_func"
)

//var g_url_string = "http://stock.elfjane.com/data.php"

type Futures_min_list_struct struct {
	Field struct {
		High  int `json:"High"`
		Vol   int `json:"Vol"`
		Low   int `json:"Low"`
		Time  int `json:"Time"`
		Close int `json:"Close"`
		Date  int `json:"Date"`
		Open  int `json:"Open"`
	} `json:"Field"`
	Sum  int         `json:"sum"`
	Type string      `json:"type"`
	List [][]float64 `json:"list"`
}

func Sk88_load_futures_min() {
	fmt.Println("sk888_load_day......")

	//test_post()
	cj_func.Load_file("sk88.ini")

	cj_db.G_db_dsn = cj_func.G_db_dsn

	cj_db.Connect()

	// get list data
	data := url.Values{}
	data.Set("f", "j")
	data.Set("Type", "TDMIN5")
	data.Set("Symbol", "1-WTX&")
	body := cj_html.Post(data)

	t := time.Now()
	filename := fmt.Sprintf("sk88_futures_json/futures_min_%04d%02d%02d%02d%02d%02d.json", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	err := os.WriteFile(filename, body, 0644)
	if err != nil {
		panic(err)
	}

	var futures_min_list_struct Futures_min_list_struct
	json.Unmarshal(body, &futures_min_list_struct)

	//stock_list_show(s)
	fmt.Println(futures_min_list_struct)

	if futures_min_list_struct.Sum == 0 {
		return
	}
	v_high := futures_min_list_struct.Field.High
	v_close := futures_min_list_struct.Field.Close
	v_date := futures_min_list_struct.Field.Date
	v_low := futures_min_list_struct.Field.Low
	v_open := futures_min_list_struct.Field.Open
	v_time := futures_min_list_struct.Field.Time
	v_vol := futures_min_list_struct.Field.Vol

	cj_db.DB_stock_day_truncate()
	for _, list := range futures_min_list_struct.List {
		s_high := list[v_high]
		s_close := list[v_close]
		s_date := int(list[v_date])
		s_low := list[v_low]
		s_open := list[v_open]
		s_time := int(list[v_time])
		s_vol := int(list[v_vol])

		cj_db.DB_futures_min_insert_day_data(s_high, s_low, s_open, s_close, s_vol, s_time, s_date)
	}
	cj_db.DB_futures_min_set_update()
	Futures_min_update_ma()
	cj_db.Close()
}
