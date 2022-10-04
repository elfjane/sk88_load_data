package main

import (
	"cj_db"
	"fmt"
)

func Futures_min_update_ma() {
	fmt.Println("stock day update ma start.")

	cj_db.Connect()
	for {
		data_list := cj_db.DB_get_Futures_min_no_ma_list()
		if len(data_list) < 1 {
			break
		}
		for _, this := range data_list {
			//fmt.Println(this)
			ma5 := cj_db.DB_futures_min_get_ma(5, this.S_unixtime)
			ma10 := cj_db.DB_futures_min_get_ma(10, this.S_unixtime)
			ma20 := cj_db.DB_futures_min_get_ma(20, this.S_unixtime)
			ma60 := cj_db.DB_futures_min_get_ma(60, this.S_unixtime)

			fmt.Printf("id = %d, ma5 = %f, ma10 = %f, ma20 = %f, ma60 = %f\r\n", this.Id, ma5, ma10, ma20, ma60)

			cj_db.DB_futures_min_update_ma(this.Id, ma5, ma10, ma20, ma60)
		}
	}
	cj_db.Close()
	fmt.Println("stock day update ma end.")
}
