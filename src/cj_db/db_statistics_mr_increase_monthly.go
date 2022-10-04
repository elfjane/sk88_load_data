package cj_db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DB_statistics_mr_increase_monthly_truncate() {

	Exec("TRUNCATE TABLE `statistics_mr_increase_monthly`")

}

// 取得最後一筆資料
func DB_statistics_mr_increase_monthly_get_company_desc(company_code string) string {
	var dt_date string
	row := G_db.QueryRow("SELECT dt_date FROM `statistics_mr_increase_monthly` WHERE symbol = ? ORDER BY dt_date DESC LIMIT 1", company_code)
	err := row.Scan(&dt_date)

	if err != nil {
		return "2013-02-01"
	} else {
		return dt_date
	}
}

func DB_statistics_mr_increase_monthly_insert_data(node_type string, symbol string, symbol_name string, dt_date string, increase_monthly int64) {

	fmt.Printf("insert data symbol = %s, increase_monthly_monthly =%d\r\n", symbol, increase_monthly)
	Exec("INSERT INTO `stock`.`statistics_mr_increase_monthly` (`node_type`, `symbol`, `symbol_name`, `dt_date`, `increase_monthly`) VALUES (?, ?, ?, ?, ?)", node_type, symbol, symbol_name, dt_date, increase_monthly)

}
