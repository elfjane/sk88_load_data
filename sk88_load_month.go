package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"

	"cj_db"
	"cj_func"
	"cj_html"
)

//var g_url_string = "http://stock.elfjane.com/data.php"

type Stock_month_list_struct struct {
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

func Sk88_load_month() {
	fmt.Println("sk888_load_month......")
	//test_post()

	cj_db.Connect()
	var data_list = cj_db.DB_get_stock_enable_month_list()
	//os.Exit(1)
	//fmt.Println(data_list)
	//change_date(100)
	//

	//os.Exit(1)
	for _, this := range data_list {
		// get list data
		var symbol_combine = this.Market + "-" + this.Symbol
		fmt.Printf("id = %d, market = %s, symbol = %s, symbol_combine = %s\r\n", this.Id, this.Market, this.Symbol, symbol_combine)
		data := url.Values{}

		data.Set("f", "j")
		data.Set("Type", "MONTH")
		data.Set("UUID", cj_func.G_uuid)
		data.Set("Symbol", symbol_combine)
		body := cj_html.Post(data)
		var stock_month_list_struct Stock_month_list_struct
		json.Unmarshal(body, &stock_month_list_struct)

		//stock_list_show(s)
		fmt.Println(stock_month_list_struct)
		if stock_month_list_struct.Sum == 0 {
			continue
		}
		v_high := stock_month_list_struct.Field.High
		v_close := stock_month_list_struct.Field.Close
		v_date := stock_month_list_struct.Field.Date
		v_low := stock_month_list_struct.Field.Low
		v_open := stock_month_list_struct.Field.Open
		v_time := stock_month_list_struct.Field.Time
		v_vol := stock_month_list_struct.Field.Vol

		for _, list := range stock_month_list_struct.List {
			s_high := list[v_high]
			s_close := list[v_close]
			s_date := int(list[v_date])
			s_low := list[v_low]
			s_open := list[v_open]
			s_time := int(list[v_time])
			s_vol := int(list[v_vol])

			cj_db.DB_stock_month_insert_month_data(this.Id, this.Market, this.Symbol, symbol_combine, s_high, s_low, s_open, s_close, s_vol, s_time, s_date)
		}
		cj_db.DB_stock_month_set_update(this.Id)
		//os.Exit(1)
	}
	cj_db.Close()
}
