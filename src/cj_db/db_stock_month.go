package cj_db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Stock_list_month_enable struct {
	Id     int    `sql:"id"`
	Market string `sql:"market"`
	Symbol string `sql:"symbol"`
}

func DB_stock_month_truncate() {

	Exec("TRUNCATE TABLE `stock_month4`")

}

func DB_get_stock_enable_month_list() []Stock_list_month_enable {
	fmt.Println("stock_list db select")

	// Execute the query
	results, err := G_db.Query("SELECT id, market, symbol FROM stock_list WHERE is_update_month = 1 AND update_month < NOW()")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	data_list := []Stock_list_month_enable{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_list_month_enable Stock_list_month_enable
		err = results.Scan(&stock_list_month_enable.Id,
			&stock_list_month_enable.Market,
			&stock_list_month_enable.Symbol)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_list_month_enable)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

func DB_stock_month_insert_month_data(id int, market string, symbol string, symbol_combine string, s_high float64, s_low float64, s_open float64, s_close float64, s_vol int, s_time int, s_date int) {

	//results, err := g_db.Query("INSERT INTO `test` (`s_high`) VALUES (?)", s_high)
	dt_date := Change_date(s_date)
	fmt.Printf("insert data symbol = %s, date =%s\r\n", symbol_combine, dt_date)
	Exec("INSERT INTO `stock_month4` (`symbol`, `market`, `symbol_combine`, `dt_date`, `s_open`, `s_high`, `s_close`, `s_low`, `s_vol`, `s_time`, `int_date`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", symbol, market, symbol_combine, dt_date, s_open, s_high, s_close, s_low, s_vol, s_time, s_date)

}

func DB_stock_month_set_update(id int) {

	//Exec("DELETE a FROM `stock_month4` AS a, `stock_month` AS b WHERE a.symbol = b.symbol AND a.int_date = b.int_date")
	Exec("TRUNCATE TABLE `stock_month4`")
	Exec("UPDATE `stock_month4` SET  `s_unixtime` = UNIX_TIMESTAMP(`dt_date`) * 1000")

	Exec("INSERT INTO stock_month (symbol, market, symbol_combine, dt_date, s_open, s_close, s_high, s_low, s_time, s_unixtime, int_date, s_vol) SELECT symbol, market, symbol_combine, dt_date, s_open, s_close, s_high, s_low, s_time, s_unixtime, int_date, s_vol FROM stock_month4")

	Exec("DELETE a FROM `stock_month4` AS a, `stock_month` AS b WHERE a.symbol = b.symbol AND a.int_date = b.int_date")

	Exec("UPDATE `stock_list` SET `update_month` = DATE_FORMAT(update_month + INTERVAL 1 MONTH,'%Y-%m-01') WHERE `id` = ?", id)

}

func DB_stock_month_get_ma(symbol string, ma_limit int, int_date int) float64 {

	var ma float64
	var ma_count int
	var sql, sql_select string
	sql_select = fmt.Sprintf("SELECT * FROM `stock_month` WHERE `symbol` = '%s' AND `int_date` < %d ORDER BY `int_date` DESC LIMIT %d", symbol, int_date, ma_limit)
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

func DB_stock_month_update_ma(id int, ma5 float64, ma10 float64, ma20 float64, ma60 float64) {
	Exec("UPDATE `stock_month` SET `is_ma` = 1, `ma5` = ?, `ma10` = ?, `ma20` = ?, `ma60` = ? WHERE `id` = ?", ma5, ma10, ma20, ma60, id)
}

type Stock_month_select_no_ma struct {
	Id       int    `sql:"id"`
	Symbol   string `sql:"symbol"`
	Int_date int    `sql:"int_date"`
}

func DB_get_stock_month_no_ma_list() []Stock_month_select_no_ma {
	fmt.Println("stock_month db select")

	// Execute the query is no ma
	results, err := G_db.Query("SELECT id, symbol, int_date FROM stock_month WHERE is_ma = 0 LIMIT 100")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Stock_month_select_no_ma{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_month_select_no_ma Stock_month_select_no_ma
		err = results.Scan(&stock_month_select_no_ma.Id,
			&stock_month_select_no_ma.Symbol,
			&stock_month_select_no_ma.Int_date)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_month_select_no_ma)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}
