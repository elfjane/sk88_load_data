package main

import (
	"fmt"
	"math"

	_ "github.com/go-sql-driver/mysql"

	"cj_db"
)

func get_taiwan_gear(price float64) float64 {
	var gear float64
	if price < 10 {
		gear = 0.01
	} else if price >= 10 && price < 50 {
		gear = 0.05
	} else if price >= 50 && price < 100 {
		gear = 0.1
	} else if price >= 100 && price < 500 {
		gear = 0.5
	} else if price >= 500 && price < 1000 {
		gear = 1
	} else if price >= 1000 {
		gear = 5
	} else {
		panic("gear error")
	}

	return gear
}

func taiwan_stock_get_gear_limit(s_close float64) (float64, float64) {
	daily_limit := s_close + (s_close * 0.1)                   // 取得漲停價試算
	daily_limit_gear := get_taiwan_gear(daily_limit)           // 取得檔位
	daily_limit_mod := math.Mod(daily_limit, daily_limit_gear) // 取得檔位上限
	daily_limit_ret := daily_limit - daily_limit_mod           // 漲停價

	lower_limit := s_close - (s_close * 0.1)                   // 取得跌停價試算
	lower_limit_gear := get_taiwan_gear(lower_limit)           // 取得跌停價檔位
	lower_limit_mod := math.Mod(lower_limit, lower_limit_gear) // 取得跌停價下限
	var lower_limit_ret float64                                // 跌停價
	// 跌停溢位算法
	if lower_limit_mod > 0 {
		lower_limit_ret = lower_limit - lower_limit_mod + lower_limit_gear
	} else {
		lower_limit_ret = lower_limit
	}

	/*
		fmt.Println(daily_limit)
		fmt.Println(daily_limit_gear)
		fmt.Println(daily_limit_mod)
		fmt.Println(daily_limit_ret)

		fmt.Println(lower_limit)
		fmt.Println(lower_limit_gear)
		fmt.Println(lower_limit_mod)
		fmt.Println(lower_limit_ret)
	*/
	return daily_limit_ret, lower_limit_ret
}
func Stock_day_limit() {
	fmt.Println("sk888_load_day......")
	//test_post()
	cj_db.Connect()
	var data_list = cj_db.DB_get_stock_enable_day_limit_list()
	//os.Exit(1)
	fmt.Println(data_list)
	//change_date(100)
	//

	var i = 0
	total_symbol := len(data_list)
	var s_total_symbol = fmt.Sprintf("total symbol = %d/%d", i, total_symbol)
	fmt.Println(s_total_symbol)
	taiwan_stock_get_gear_limit(75.1)

	//daily_limit, lower_limit := taiwan_stock_get_gear_limit(72)
	//fmt.Println(daily_limit)
	//fmt.Println(lower_limit)
	//daily_limit2, lower_limit2 := taiwan_stock_get_gear_limit(50.1)
	//os.Exit(1)

	for _, this := range data_list {

		stock_day_select_close_price := cj_db.DB_get_stock_day_get_close_price(this.Symbol)
		fmt.Println(stock_day_select_close_price)
		len_stock_day := len(stock_day_select_close_price) - 1
		if len_stock_day < 1 {
			continue

		}
		fmt.Printf("total symbol = %d\r\n", len_stock_day)

		for j := 0; j < len_stock_day; j = j + 1 {
			stock_day_data_today := stock_day_select_close_price[j]
			stock_day_data_nextday := stock_day_select_close_price[j+1]
			if stock_day_data_nextday.Unchanged != -1 {
				continue
			}
			fmt.Println(stock_day_data_today)
			daily_limit, lower_limit := taiwan_stock_get_gear_limit(stock_day_data_today.S_close)
			fmt.Println(daily_limit, lower_limit)
			cj_db.DB_stock_day_update_limit(stock_day_data_nextday.Id, stock_day_data_today.S_close, daily_limit, lower_limit)

		}
		cj_db.DB_stock_list_set_update_limit(this.Id)

		//os.Exit(1)
	}
	cj_db.Close()
}
