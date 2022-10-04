package cj_db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Futures_list_min_enable struct {
	Id     int    `sql:"id"`
	Market string `sql:"market"`
	Symbol string `sql:"symbol"`
}

type Futures_min_select_no_ma struct {
	Id         int `sql:"id"`
	S_unixtime int `sql:"s_unixtime"`
}

type Futures_min_select_close_price struct {
	Id          int     `sql:"id"`
	Symbol      string  `sql:"symbol"`
	Int_date    int     `sql:"int_date"`
	S_close     float64 `sql:"s_close"`
	Unchanged   float64 `sql:"unchanged"`
	Daily_limit float64 `sql:"daily_limit"`
}

func DB_futures_min_truncate() {

	Exec("TRUNCATE TABLE `stock_day4`")

}

func DB_get_futures_enable_min_list() []Futures_list_min_enable {
	fmt.Println("stock_list db select")

	// Execute the query
	results, err := G_db.Query("SELECT id, market, symbol FROM stock_list WHERE is_update_day = 1 AND update_day < NOW() ORDER BY priority DESC")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	data_list := []Futures_list_min_enable{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var futures_list_min_enable Futures_list_min_enable
		err = results.Scan(&futures_list_min_enable.Id,
			&futures_list_min_enable.Market,
			&futures_list_min_enable.Symbol)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, futures_list_min_enable)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

func DB_get_futures_enable_min_limit_list() []Futures_list_min_enable {
	fmt.Println("stock_list db select day limit")

	// Execute the query
	results, err := G_db.Query("SELECT id, market, symbol FROM stock_list WHERE is_update_day = 1 AND update_day < NOW() AND symbol_type IN ('上市', '上櫃', '上市ETF') ORDER BY priority DESC")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	data_list := []Futures_list_min_enable{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var futures_list_min_enable Futures_list_min_enable
		err = results.Scan(&futures_list_min_enable.Id,
			&futures_list_min_enable.Market,
			&futures_list_min_enable.Symbol)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, futures_list_min_enable)
	}
	results.Close()
	fmt.Println("query db day limit success")
	return data_list
}

func DB_futures_min_insert_day_data(s_high float64, s_low float64, s_open float64, s_close float64, s_vol int, s_time int, s_date int) {

	//results, err := g_db.Query("INSERT INTO `test` (`s_high`) VALUES (?)", s_high)
	dt_date := Change_datetime(s_date, s_time)
	fmt.Printf("insert futures_min, date =%s\r\n", dt_date)
	Exec("INSERT INTO `futures_min4` (`dt_date`, `s_open`, `s_high`, `s_close`, `s_low`, `s_vol`, `s_time`, `int_date`, `int_time`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", dt_date, s_open, s_high, s_close, s_low, s_vol, s_time, s_date, s_time)

}

func DB_futures_min_set_update() {
	Exec("DELETE a FROM `futures_min4` AS a, `futures_min` AS b WHERE a.int_date = b.int_date AND a.int_time = b.int_time ")

	Exec("UPDATE `futures_min4` SET  `s_unixtime` = UNIX_TIMESTAMP(`dt_date`) * 1000")

	Exec("INSERT INTO futures_min (symbol, market, symbol_combine, dt_date, s_open, s_close, s_high, s_low, s_time, s_unixtime, int_date, int_time, s_vol) SELECT symbol, market, symbol_combine, dt_date, s_open, s_close, s_high, s_low, s_time, s_unixtime, int_date, int_time, s_vol FROM futures_min4")

}

func DB_futures_min_get_ma(ma_limit int, s_unixtime int) float64 {

	var ma float64
	var ma_count int
	var sql, sql_select string
	//sql_select := "SELECT * FROM `stock_day` WHERE `symbol` = '" + symbol + "' AND `int_date` < " + int_date + " ORDER BY `int_date` DESC LIMIT " + ma_limit
	sql_select = fmt.Sprintf("SELECT * FROM `futures_min` WHERE `s_unixtime` < %d ORDER BY `s_unixtime` DESC LIMIT %d", s_unixtime, ma_limit)
	sql = fmt.Sprintf("SELECT IFNULL(SUM(s_close) / COUNT(s_close),0) AS ma, COUNT(s_close) AS ma_count FROM (%s) as totals", sql_select)

	row := G_db.QueryRow(sql)
	err := row.Scan(&ma, &ma_count)
	if err != nil {
		log.Fatal(err)
	}
	if ma_count < ma_limit {
		ma = 0
	}
	return ma
}

func DB_futures_min_update_ma(id int, ma5 float64, ma10 float64, ma20 float64, ma60 float64) {

	Exec("UPDATE `futures_min` SET `is_ma` = 1, `ma5` = ?, `ma10` = ?, `ma20` = ?, `ma60` = ? WHERE `id` = ?", ma5, ma10, ma20, ma60, id)
}

func DB_get_Futures_min_no_ma_list() []Futures_min_select_no_ma {
	fmt.Println("stock_list db select")

	// Execute the query is no ma
	results, err := G_db.Query("SELECT id, s_unixtime FROM futures_min WHERE is_ma = 0 LIMIT 100")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Futures_min_select_no_ma{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var futures_min_select_no_ma Futures_min_select_no_ma
		err = results.Scan(&futures_min_select_no_ma.Id,
			&futures_min_select_no_ma.S_unixtime)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, futures_min_select_no_ma)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

// 取得收盤價
func DB_get_futures_min_get_close_price() []Futures_min_select_close_price {
	fmt.Println("stock_list db select")

	// Execute the query is no ma
	results, err := G_db.Query("SELECT id, symbol, int_date, s_close, unchanged, daily_limit FROM futures_min ORDER by int_date ASC, int_time ASC")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Futures_min_select_close_price{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var futures_min_select_close_price Futures_min_select_close_price
		err = results.Scan(&futures_min_select_close_price.Id,
			&futures_min_select_close_price.Symbol,
			&futures_min_select_close_price.Int_date,
			&futures_min_select_close_price.S_close,
			&futures_min_select_close_price.Unchanged,
			&futures_min_select_close_price.Daily_limit)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, futures_min_select_close_price)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

// 更新 台股漲停板跟跌停板
func DB_futures_min_update_limit(id int, unchanged float64, daily_limit float64, lower_limit float64) {
	Exec("UPDATE `stock_day` SET unchanged = ?, `daily_limit` = ?, `lower_limit` = ? WHERE `id` = ?", unchanged, daily_limit, lower_limit, id)
}
