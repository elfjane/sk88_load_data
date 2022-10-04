package cj_db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Stock_list_day_enable struct {
	Id     int    `sql:"id"`
	Market string `sql:"market"`
	Symbol string `sql:"symbol"`
}

type Stock_day_select_no_ma struct {
	Id       int    `sql:"id"`
	Symbol   string `sql:"symbol"`
	Int_date int    `sql:"int_date"`
}

type Stock_day_select_close_price struct {
	Id          int     `sql:"id"`
	Symbol      string  `sql:"symbol"`
	Int_date    int     `sql:"int_date"`
	S_close     float64 `sql:"s_close"`
	Unchanged   float64 `sql:"unchanged"`
	Daily_limit float64 `sql:"daily_limit"`
}

func DB_stock_day_truncate() {

	Exec("TRUNCATE TABLE `stock_day4`")

}

func DB_get_stock_enable_day_list() []Stock_list_day_enable {
	fmt.Println("stock_list db select")

	// Execute the query
	results, err := G_db.Query("SELECT id, market, symbol FROM stock_list WHERE is_update_day = 1 AND update_day < NOW() ORDER BY priority DESC")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	data_list := []Stock_list_day_enable{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_list_day_enable Stock_list_day_enable
		err = results.Scan(&stock_list_day_enable.Id,
			&stock_list_day_enable.Market,
			&stock_list_day_enable.Symbol)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_list_day_enable)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

func DB_get_stock_enable_day_limit_list() []Stock_list_day_enable {
	fmt.Println("stock_list db select day limit")

	// Execute the query
	results, err := G_db.Query("SELECT id, market, symbol FROM stock_list WHERE is_update_day = 1 AND update_day < NOW() AND symbol_type IN ('上市', '上櫃', '上市ETF') ORDER BY priority DESC")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	data_list := []Stock_list_day_enable{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_list_day_enable Stock_list_day_enable
		err = results.Scan(&stock_list_day_enable.Id,
			&stock_list_day_enable.Market,
			&stock_list_day_enable.Symbol)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_list_day_enable)
	}
	results.Close()
	fmt.Println("query db day limit success")
	return data_list
}

func DB_stock_day_insert_day_data(id int, market string, symbol string, symbol_combine string, s_high float64, s_low float64, s_open float64, s_close float64, s_vol int, s_time int, s_date int) {

	//results, err := g_db.Query("INSERT INTO `test` (`s_high`) VALUES (?)", s_high)
	dt_date := Change_date(s_date)
	fmt.Printf("insert data symbol = %s, date =%s\r\n", symbol_combine, dt_date)
	Exec("INSERT INTO `stock_day4` (`symbol`, `market`, `symbol_combine`, `dt_date`, `s_open`, `s_high`, `s_close`, `s_low`, `s_vol`, `s_time`, `int_date`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", symbol, market, symbol_combine, dt_date, s_open, s_high, s_close, s_low, s_vol, s_time, s_date)

}

func DB_stock_day_set_update(id int) {

	//Exec("DELETE a FROM `stock_day4` AS a, `stock_day` AS b WHERE a.symbol = b.symbol AND a.int_date = b.int_date")

	Exec("DELETE a FROM `stock_day4` AS a, `stock_day` AS b WHERE a.symbol = b.symbol AND a.int_date = b.int_date")

	Exec("UPDATE `stock_day4` SET  `s_unixtime` = UNIX_TIMESTAMP(`dt_date`) * 1000")

	Exec("INSERT INTO stock_day (symbol, market, symbol_combine, dt_date, s_open, s_close, s_high, s_low, s_time, s_unixtime, int_date, s_vol) SELECT symbol, market, symbol_combine, dt_date, s_open, s_close, s_high, s_low, s_time, s_unixtime, int_date, s_vol FROM stock_day4")

	Exec("UPDATE `stock_list` SET `update_day` = DATE(DATE_ADD(NOW(), INTERVAL 1 DAY)) WHERE `id` = ?", id)

}

func DB_stock_day_get_ma(symbol string, ma_limit int, int_date int) float64 {

	var ma float64
	var ma_count int
	var sql, sql_select string
	//sql_select := "SELECT * FROM `stock_day` WHERE `symbol` = '" + symbol + "' AND `int_date` < " + int_date + " ORDER BY `int_date` DESC LIMIT " + ma_limit
	sql_select = fmt.Sprintf("SELECT * FROM `stock_day` WHERE `symbol` = '%s' AND `int_date` < %d ORDER BY `int_date` DESC LIMIT %d", symbol, int_date, ma_limit)
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

func DB_stock_day_update_ma(id int, ma5 float64, ma10 float64, ma20 float64, ma60 float64) {

	Exec("UPDATE `stock_day` SET `is_ma` = 1, `ma5` = ?, `ma10` = ?, `ma20` = ?, `ma60` = ? WHERE `id` = ?", ma5, ma10, ma20, ma60, id)
}

func DB_get_stock_day_no_ma_list() []Stock_day_select_no_ma {
	fmt.Println("stock_list db select")

	// Execute the query is no ma
	results, err := G_db.Query("SELECT id, symbol, int_date FROM stock_day WHERE is_ma = 0 LIMIT 100")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Stock_day_select_no_ma{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_day_select_no_ma Stock_day_select_no_ma
		err = results.Scan(&stock_day_select_no_ma.Id,
			&stock_day_select_no_ma.Symbol,
			&stock_day_select_no_ma.Int_date)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_day_select_no_ma)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

// 取得收盤價
func DB_get_stock_day_get_close_price(symbol string) []Stock_day_select_close_price {
	fmt.Println("stock_list db select")

	// Execute the query is no ma
	results, err := G_db.Query("SELECT id, symbol, int_date, s_close, unchanged, daily_limit FROM stock_day WHERE symbol = ? ORDER by int_date ASC", symbol)
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Stock_day_select_close_price{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_day_select_close_price Stock_day_select_close_price
		err = results.Scan(&stock_day_select_close_price.Id,
			&stock_day_select_close_price.Symbol,
			&stock_day_select_close_price.Int_date,
			&stock_day_select_close_price.S_close,
			&stock_day_select_close_price.Unchanged,
			&stock_day_select_close_price.Daily_limit)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_day_select_close_price)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

// 更新 台股漲停板跟跌停板
func DB_stock_day_update_limit(id int, unchanged float64, daily_limit float64, lower_limit float64) {
	Exec("UPDATE `stock_day` SET unchanged = ?, `daily_limit` = ?, `lower_limit` = ? WHERE `id` = ?", unchanged, daily_limit, lower_limit, id)
}

func DB_stock_list_set_update_limit(id int) {
	Exec("UPDATE `stock_list` SET `update_day_limit` = DATE(DATE_ADD(NOW(), INTERVAL 1 DAY)) WHERE `id` = ?", id)
}
