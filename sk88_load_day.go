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

type Stock_day_list_struct struct {
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

func Sk88_load_day() {
	fmt.Println("sk888_load_day......")
	//test_post()
	cj_func.Load_file("sk88.ini")
	uuid := cj_func.G_uuid
	cj_db.G_db_dsn = cj_func.G_db_dsn

	cj_db.Connect()
	var data_list = cj_db.DB_get_stock_enable_day_list()
	//os.Exit(1)
	//fmt.Println(data_list)
	//change_date(100)
	//

	var i = 0
	total_symbol := len(data_list)
	var s_total_symbol = fmt.Sprintf("total symbol = %d/%d", i, total_symbol)
	fmt.Println(s_total_symbol)

	//os.Exit(1)

	for _, this := range data_list {
		// get list data
		var symbol_combine = this.Market + "-" + this.Symbol
		fmt.Printf("id = %d, market = %s, symbol = %s, symbol_combine = %s\r\n", this.Id, this.Market, this.Symbol, symbol_combine)
		data := url.Values{}

		data.Set("f", "j")
		data.Set("Type", "DAY")
		data.Set("UUID", uuid)
		data.Set("Symbol", symbol_combine)
		body := cj_html.Post(data)

		t := time.Now()
		filename := fmt.Sprintf("sk88_json/%s_%04d%02d%02d%02d%02d%02d.json", this.Symbol, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

		err := os.WriteFile(filename, body, 0644)
		if err != nil {
			panic(err)
		}

		var stock_day_list_struct Stock_day_list_struct
		json.Unmarshal(body, &stock_day_list_struct)

		//stock_list_show(s)
		fmt.Println(stock_day_list_struct)

		if stock_day_list_struct.Sum == 0 {
			continue
		}
		v_high := stock_day_list_struct.Field.High
		v_close := stock_day_list_struct.Field.Close
		v_date := stock_day_list_struct.Field.Date
		v_low := stock_day_list_struct.Field.Low
		v_open := stock_day_list_struct.Field.Open
		v_time := stock_day_list_struct.Field.Time
		v_vol := stock_day_list_struct.Field.Vol

		cj_db.DB_stock_day_truncate()

		for _, list := range stock_day_list_struct.List {
			s_high := list[v_high]
			s_close := list[v_close]
			s_date := int(list[v_date])
			s_low := list[v_low]
			s_open := list[v_open]
			s_time := int(list[v_time])
			s_vol := int(list[v_vol])

			cj_db.DB_stock_day_insert_day_data(this.Id, this.Market, this.Symbol, symbol_combine, s_high, s_low, s_open, s_close, s_vol, s_time, s_date)
		}
		cj_db.DB_stock_day_set_update(this.Id)
		//os.Exit(1)
	}
	cj_db.Close()
}
