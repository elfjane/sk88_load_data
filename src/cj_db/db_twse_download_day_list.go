package cj_db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 取得最後一天
func DB_twse_download_day_list_lastday() string {

	var dt_start string

	sql := "SELECT dt_id as dt_start FROM `twse_download_day_list` ORDER BY `dt_id` DESC LIMIT 1"
	row := G_db.QueryRow(sql)
	err := row.Scan(&dt_start)
	if err != nil {
		return ""
	}

	return dt_start
}

// 新增要下載的日期
func DB_twse_download_day_list_insert_init(dt_id string, csv_filename string) {
	fmt.Printf("insert data dt_id = %s, filename =%s\r\n", dt_id, csv_filename)
	Exec("INSERT INTO `stock`.`twse_download_day_list` (`dt_id`, `csv_filename`) VALUES (?, ?);", dt_id, csv_filename)
}

// 刪除星期日天數
func DB_twse_download_day_list_delete_week() {
	Exec("DELETE FROM `stock`.`twse_download_day_list` WHERE DAYOFWEEK(`dt_id`) = 1")
}

// 取得要下載的列表
type Twse_select_no_download struct {
	Dt_id        string `sql:"dt_id"`
	Csv_filename string `sql:"csv_filename"`
}

func DB_twse_download_day_list_get() []Twse_select_no_download {
	fmt.Println("twse_download_day_list db select")

	// Execute the query is no ma
	results, err := G_db.Query("SELECT dt_id, csv_filename FROM `twse_download_day_list` WHERE is_download = 0")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Twse_select_no_download{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var twse_select_no_download Twse_select_no_download
		err = results.Scan(&twse_select_no_download.Dt_id,
			&twse_select_no_download.Csv_filename)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, twse_select_no_download)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

// 更新 download
func DB_twse_download_day_list_download_update(dt_id string, file_size int64) {
	Exec("UPDATE `stock`.`twse_download_day_list` SET `csv_size` = ? , `is_download` = is_download + 1 WHERE `dt_id` = ?; ", file_size, dt_id)
}

// 更新 download
func DB_twse_download_day_list_is_change_update(dt_id string) {
	Exec("UPDATE `stock`.`twse_download_day_list` SET `is_change` = is_change + 1 WHERE `dt_id` = ?; ", dt_id)
}

// 更新 download
func DB_twse_download_day_list_is_insert_update(dt_id string) {
	Exec("UPDATE `stock`.`twse_download_day_list` SET `is_insert` = is_insert + 1 WHERE `dt_id` = ?; ", dt_id)
}

// 取得尚未讀取的列表列表
type Twse_day_no_insert_list struct {
	Dt_id        string `sql:"dt_id"`
	Csv_filename string `sql:"csv_filename"`
}

func DB_twse_download_no_insert_list() []Twse_day_no_insert_list {
	fmt.Println("twse_download_day_list no insert db select")

	// 取得尚未新增的列表
	results, err := G_db.Query("SELECT dt_id, csv_filename FROM `twse_download_day_list` WHERE is_change = 0 AND csv_size > 0 LIMIT 100")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Twse_day_no_insert_list{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var twse_day_no_insert_list Twse_day_no_insert_list
		err = results.Scan(&twse_day_no_insert_list.Dt_id,
			&twse_day_no_insert_list.Csv_filename)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, twse_day_no_insert_list)
	}
	results.Close()
	fmt.Println("query db success")
	return data_list
}

// 取得尚未讀取的列表列表
type Twse_day_is_change_list struct {
	Dt_id        string `sql:"dt_id"`
	Csv_filename string `sql:"csv_filename"`
}

func DB_twse_download_is_change_list() []Twse_day_is_change_list {
	Connect()
	fmt.Println("twse_download_day_list is_change db select")

	// 取得尚未新增的列表
	results, err := G_db.Query("SELECT dt_id, csv_filename FROM `twse_download_day_list` WHERE is_change > 0 AND is_insert = 0 LIMIT 100")
	if err != nil {
		fmt.Println("query db error")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("query db start")
	var data_list = []Twse_day_is_change_list{}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		var twse_day_is_change_list Twse_day_is_change_list
		err = results.Scan(&twse_day_is_change_list.Dt_id,
			&twse_day_is_change_list.Csv_filename)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data_list = append(data_list, twse_day_is_change_list)
	}
	results.Close()
	fmt.Println("query db success")
	Close()
	return data_list
}
