package cj_db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 取得要下載的列表
type Stock_list_symbol_data struct {
	Symbol      string `sql:"symbol"`
	Symbol_name string `sql:"symbol_name"`
	Symbol_type string `sql:"symbol_type"`
}

func DB_stock_List_select_enable() []Stock_list_symbol_data {
	fmt.Println("stock_list db enable select")

	// Execute the query is no ma
	results, err := G_db.Query("SELECT symbol, symbol_name, symbol_type FROM `stock_list`  WHERE is_enable = 1 AND is_company = 1")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Stock_list_symbol_data{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_list_symbol_data Stock_list_symbol_data
		err = results.Scan(&stock_list_symbol_data.Symbol,
			&stock_list_symbol_data.Symbol_name,
			&stock_list_symbol_data.Symbol_type)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_list_symbol_data)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}
