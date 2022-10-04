package cj_db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 取得要下載的列表
type Stock_mr_company_desc struct {
	Node_type            string `sql:"node_type"`
	Company_code         string `sql:"company_code"`
	company_abbreviation string `sql:"company_abbreviation"`
	Dt_date              string `sql:"dt_date"`
	Oi_mr                int64  `sql:"oi_mr"`
}

func DB_stock_monthly_revenue_data_get_mr_desc(company_code string) []Stock_mr_company_desc {
	fmt.Println("DB_stock_monthly_revenue_data_get_mr_desc db select")

	// Execute the query is no ma
	results, err := G_db.Query("SELECT node_type, company_code, company_abbreviation, dt_date, oi_mr FROM `stock_monthly_revenue_data` WHERE company_code = ? ORDER BY dt_date DESC", company_code)
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//fmt.Println("query db start")
	var data_list = []Stock_mr_company_desc{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_mr_company_desc Stock_mr_company_desc
		err = results.Scan(&stock_mr_company_desc.Node_type,
			&stock_mr_company_desc.Company_code,
			&stock_mr_company_desc.company_abbreviation,
			&stock_mr_company_desc.Dt_date,
			&stock_mr_company_desc.Oi_mr)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_mr_company_desc)
	}
	results.Close()
	//fmt.Println("query db success")
	return data_list
}

// monthly
func DB_stock_monthly_revenue_data_get_mr_desc_monthly(company_code string, dt_date string) []Stock_mr_company_desc {
	fmt.Println("DB_stock_monthly_revenue_data_get_mr_desc db select")

	// Execute the query is no ma
	fmt.Println(company_code)
	fmt.Println(dt_date)
	results, err := G_db.Query("SELECT node_type, company_code, company_abbreviation, dt_date, oi_mr FROM `stock_monthly_revenue_data` WHERE company_code = ? AND dt_date <= ? ORDER BY dt_date DESC", company_code, dt_date)
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//fmt.Println("query db start")
	var data_list = []Stock_mr_company_desc{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_mr_company_desc Stock_mr_company_desc
		err = results.Scan(&stock_mr_company_desc.Node_type,
			&stock_mr_company_desc.Company_code,
			&stock_mr_company_desc.company_abbreviation,
			&stock_mr_company_desc.Dt_date,
			&stock_mr_company_desc.Oi_mr)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_mr_company_desc)
	}
	results.Close()
	//fmt.Println("query db success")
	return data_list
}

func DB_stock_monthly_revenue_data_get_mr_desc_date(company_code string, dt_date string) []Stock_mr_company_desc {
	fmt.Println("DB_stock_monthly_revenue_data_get_mr_desc db select")

	// Execute the query is no ma

	results, err := G_db.Query("SELECT node_type, company_code, company_abbreviation, dt_date, oi_mr FROM `stock_monthly_revenue_data` WHERE company_code = ? AND dt_date <= ? ORDER BY dt_date DESC", company_code, dt_date)
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Stock_mr_company_desc{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var stock_mr_company_desc Stock_mr_company_desc
		err = results.Scan(&stock_mr_company_desc.Node_type,
			&stock_mr_company_desc.Company_code,
			&stock_mr_company_desc.company_abbreviation,
			&stock_mr_company_desc.Dt_date,
			&stock_mr_company_desc.Oi_mr)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, stock_mr_company_desc)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

// 取得最後一筆資料
func DB_stock_monthly_revenue_data_get_company_last_date(company_code string) string {
	var dt_date string
	row := G_db.QueryRow("SELECT dt_date FROM `stock_monthly_revenue_data` WHERE company_code = ? ORDER BY dt_date DESC LIMIT 1;", company_code)
	err := row.Scan(&dt_date)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		return dt_date
	}
}
