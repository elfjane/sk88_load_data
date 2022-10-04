package main

import (
	"cj_db"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 下載台灣證卷交易所 每日收盤行情
func Statistics_monthly_revenue_increase_start() {
	fmt.Println("statistics_monthly_revenue_increase_start start")
	cj_db.Connect()

	statistics_monthly_revenue_increase_run()

	cj_db.Close()

}

// 統計財務每個月都有增加
func statistics_monthly_revenue_increase_count(symbol string) int64 {
	//fmt.Println("statistics_monthly_revenue_increase_count")

	data_list := cj_db.DB_stock_monthly_revenue_data_get_mr_desc(symbol)
	var count int64
	var before_value int64
	first := true

	before_value = 0
	count = 0

	for _, this := range data_list {
		//fmt.Println(before_value)
		if first {
			before_value = this.Oi_mr
			first = false
			continue
		}

		if this.Oi_mr < 1 {
			break
		}

		if before_value < this.Oi_mr {
			break
		}
		count = count + 1
		before_value = this.Oi_mr
	}
	return count
}

// 統計財務每個月都有增加 分月
func statistics_monthly_revenue_increase_monthly_spilt(node_type string, symbol string, symbol_name string) {
	dt_start := cj_db.DB_statistics_mr_increase_monthly_get_company_desc(symbol)
	fmt.Println(dt_start)
	dt_date_last := cj_db.DB_stock_monthly_revenue_data_get_company_last_date(symbol)
	diffMonth := cj_db.DBF_diffTwoMonth(dt_start, dt_date_last)
	//fmt.Printf("%d \r\n", diffMonth)

	dtstr1 := dt_start + " 00:00:00"
	dt, _ := time.Parse("2006-01-02 15:04:05", dtstr1)
	if diffMonth == 0 {
		return
	}
	i := 1               // i預設為 0
	for i <= diffMonth { // i<=10 為真就執行 {} 內的咒語，否則不執行
		dt_add := dt.AddDate(0, i, 0)
		//fmt.Println(dt_add)
		//dt_day := fmt.Sprintf("%04d%02d%02d", dt_add.Year(), dt_add.Month(), dt_add.Day())
		dt_id := fmt.Sprintf("%04d-%02d-%02d", dt_add.Year(), dt_add.Month(), dt_add.Day())
		fmt.Println(dt_id)
		count := statistics_monthly_revenue_increase_monthly_count(symbol, dt_id)
		// count mr
		cj_db.DB_statistics_mr_increase_monthly_insert_data(node_type, symbol, symbol_name, dt_id, count)
		i = i + 1
	}
}

// 統計財務每個月都有增加 分月
func statistics_monthly_revenue_increase_monthly_count(symbol string, dt_date string) int64 {
	//fmt.Println("statistics_monthly_revenue_increase_count")

	data_list := cj_db.DB_stock_monthly_revenue_data_get_mr_desc_monthly(symbol, dt_date)
	var count int64
	var before_value int64
	first := true

	before_value = 0
	count = 0

	for _, this := range data_list {
		//fmt.Println(before_value)
		if first {
			before_value = this.Oi_mr
			first = false
			continue
		}
		if this.Oi_mr < 1 {
			break
		}
		if before_value < this.Oi_mr {
			break
		}
		count = count + 1
		before_value = this.Oi_mr
	}
	return count
}

// 統計財務每個月都有增加
func statistics_monthly_revenue_increase_run() {
	fmt.Println("statistics_monthly_revenue_increase_run")

	stock_list_enable := cj_db.DB_stock_List_select_enable()

	// by monthly
	for _, this := range stock_list_enable {

		fmt.Printf("[%s] %s monthly\r\n", this.Symbol, this.Symbol_name)
		statistics_monthly_revenue_increase_monthly_spilt(this.Symbol_type, this.Symbol, this.Symbol_name)

	}

	// 清除 DB
	cj_db.DB_statistics_mr_increase_truncate()

	for _, this := range stock_list_enable {

		count := statistics_monthly_revenue_increase_count(this.Symbol)
		// save
		cj_db.DB_statistics_mr_increase_insert_data(this.Symbol_type, this.Symbol, this.Symbol_name, count)

	}

}
