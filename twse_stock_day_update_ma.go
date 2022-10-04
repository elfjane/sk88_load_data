package main

import (
	"cj_db"
	"fmt"
)

func Twse_stock_day_update_ma() {
	fmt.Println("stock day update ma start.")

	cj_db.Connect()
	for {
		data_list := cj_db.DB_get_twse_stock_day_no_ma_list()
		if len(data_list) < 1 {
			break
		}
		for _, this := range data_list {
			//fmt.Println(this)
			ma5 := cj_db.DB_twse_stock_day_get_ma(this.Symbol, 5, this.Dt_date)
			ma10 := cj_db.DB_twse_stock_day_get_ma(this.Symbol, 10, this.Dt_date)
			ma20 := cj_db.DB_twse_stock_day_get_ma(this.Symbol, 20, this.Dt_date)
			ma60 := cj_db.DB_twse_stock_day_get_ma(this.Symbol, 60, this.Dt_date)

			fmt.Printf("id = %d, symbol = %s, ma5 = %f, ma10 = %f, ma20 = %f, ma60 = %f\r\n", this.Id, this.Symbol, ma5, ma10, ma20, ma60)

			cj_db.DB_twse_stock_day_update_ma(this.Id, ma5, ma10, ma20, ma60)
		}
	}
	cj_db.Close()
	fmt.Println("stock day update ma end.")
}
