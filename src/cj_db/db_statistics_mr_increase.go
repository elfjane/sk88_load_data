package cj_db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DB_statistics_mr_increase_truncate() {

	Exec("TRUNCATE TABLE `statistics_mr_increase`")

}

func DB_statistics_mr_increase_insert_data(node_type string, symbol string, symbol_name string, increase_monthly int64) {

	fmt.Printf("insert data symbol = %s, increase_monthly =%d\r\n", symbol, increase_monthly)
	Exec("INSERT INTO `stock`.`statistics_mr_increase` (`node_type`, `symbol`, `symbol_name`, `increase_monthly`) VALUES (?, ?, ?, ?)", node_type, symbol, symbol_name, increase_monthly)

}
