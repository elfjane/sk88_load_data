package main

import (
	"cj_db"
	"cj_func"
	"fmt"
	"os"
	"strings"
)

// 程式開始
func Twse_day_stock_csv_insert() {
	fmt.Println("Twse_day_stock_csv_insert start.")

	cj_db.Connect()
	// 將 下載的 csv 轉成股票資料
	for {
		data_list := cj_db.DB_twse_download_no_insert_list()
		if len(data_list) < 1 {
			break
		}
		for _, this := range data_list {
			fmt.Println(this)

			//fmt.Printf("dt_id = %d, symbol = %s, ma5 = %f, ma10 = %f, ma20 = %f, ma60 = %f\r\n", this.Dt_id, this.Csv_filename)

			//cj_db.DB_stock_day_update_ma(this.Id, ma5, ma10, ma20, ma60)
			dt_id := this.Dt_id
			csv_filename := "./csv/day/" + this.Csv_filename
			csv_save_filename := "./csv/day_change/" + this.Csv_filename
			Twse_day_stock_csv_insert_change_csv(dt_id, csv_filename, csv_save_filename)
		}
	}
	cj_db.Close()

	// 將 下載的 csv 轉成股票資料
	for {

		data_list := cj_db.DB_twse_download_is_change_list()
		if len(data_list) < 1 {
			break
		}

		for _, this := range data_list {
			fmt.Println(this)

			//fmt.Printf("dt_id = %d, symbol = %s, ma5 = %f, ma10 = %f, ma20 = %f, ma60 = %f\r\n", this.Dt_id, this.Csv_filename)

			//cj_db.DB_stock_day_update_ma(this.Id, ma5, ma10, ma20, ma60)
			dt_id := this.Dt_id
			csv_change_filename := "./csv/day_change/" + this.Csv_filename
			csv_save_sql_filename := "./csv/day_sql/" + this.Csv_filename + ".sql"
			Twse_day_stock_csv_insert_data(dt_id, csv_change_filename, csv_save_sql_filename)

		}

	}

	fmt.Println("Twse_day_stock_csv_insert end.")
}

// 新增資料
func Twse_day_stock_csv_insert_change_csv(dt_id string, csv_filename string, csv_save_filename string) {
	fmt.Printf("dt_id = %s, csv_filename = %s, csv_filename = %s\r\n", dt_id, csv_filename, csv_save_filename)

	text := cj_func.Func_read_file(csv_filename)

	// 取得 CVS 股市的每日資料
	is_start := false
	i := 0
	f, err := os.Create(csv_save_filename)
	if err != nil {
		panic(err)
	}
	for _, each_ln := range text {
		idx := strings.Index(each_ln, "每日收盤行情(全部)")
		if is_start {
			i = i + 1
		}

		if !is_start && idx != -1 {
			is_start = true
			fmt.Println(each_ln)
		}
		if i > 0 {
			//last := len(each_ln)
			idx_miss := strings.Index(each_ln, "(元,交易單位)")
			if idx_miss != -1 {
				continue
			}
			save_string := each_ln
			if last := len(each_ln) - 1; last >= 0 && save_string[last] == ',' {
				save_string = save_string[:last]
			}
			if i == 1 {
				f.WriteString(save_string + "\r\n")
			} else if each_ln[0] != 61 && i > 1 { // 如果開頭 文字是 = 則過濾掉
				f.WriteString(save_string + "\r\n")
			}
		}

	}
	f.Sync()
	cj_db.DB_twse_download_day_list_is_change_update(dt_id)
}

// 建立 stock_base_info 新增資料
// save_sql_insert_name 存檔檔案名稱
// records csv 的來源資料
// excel_records 中英文對照資料
// node_type sql 欄位存放資料
// add_data_first 是否建立 insert 與 清除資料 sql 語法
/*
sqlStr := "INSERT INTO test(n1, n2, n3) VALUES "
vals := []interface{}{}

for _, row := range data {
    sqlStr += "(?, ?, ?),"
    vals = append(vals, row["v1"], row["v2"], row["v3"])
}
//trim the last ,
sqlStr = sqlStr[0:len(sqlStr)-1]
//prepare the statement
stmt, _ := db.Prepare(sqlStr)

//format all vals at once
res, _ := stmt.Exec(vals...)
*/
func Twse_day_stock_csv_insert_data(dt_id string, csv_change_filename string, csv_save_sql_filename string) {

	fmt.Printf("dt_id = %s, csv_change_filename = %s, csv_save_sql_filename = %s\r\n", dt_id, csv_change_filename, csv_save_sql_filename)

	f, err := os.OpenFile(csv_save_sql_filename, os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	records := Read_csv_data(csv_change_filename)
	defer f.Close()
	var first bool = true
	vals := []interface{}{}
	sqlStr := "INSERT INTO `stock`.`twse_stock_day4` (`dt_date`,`symbol`, `s_open`, s_close, s_high, s_low, s_vol, s_count, s_money) VALUES "
	var data_first bool = true
	cj_db.Connect()
	var sql_insert string
	for _, record := range records {
		if first {
			if _, err = f.WriteString(sqlStr); err != nil {
				panic(err)
			}
			first = false
			continue
		}
		sqlStr += "(?,?,?,?,?,?,?,?,?),"
		symbol := record[0]
		s_open := cj_func.Func_base_replace_string_dot(record[5]) //
		s_close := cj_func.Func_base_replace_string_dot(record[8])
		s_high := cj_func.Func_base_replace_string_dot(record[6])
		s_low := cj_func.Func_base_replace_string_dot(record[7])
		s_vol := cj_func.Func_base_replace_string_dot(record[2])
		s_count := cj_func.Func_base_replace_string_dot(record[3])
		s_money := cj_func.Func_base_replace_string_dot(record[4])
		//fmt.Printf(s_open)

		if data_first {
			sql_insert = fmt.Sprintf("('%s', '%s', %s, %s, %s, %s, %s, %s, %s)", dt_id, symbol, s_open, s_close, s_high, s_low, s_vol, s_count, s_money)
			data_first = false
		} else {
			sql_insert = fmt.Sprintf(",('%s', '%s', %s, %s, %s, %s, %s, %s, %s)", dt_id, symbol, s_open, s_close, s_high, s_low, s_vol, s_count, s_money)
		}
		if _, err = f.WriteString(sql_insert); err != nil {
			panic(err)
		}

		//fmt.Printf("dt_id=%s, symbol=%s, s_open=%s, %s, %s, %s, %s, %s, %s \r\n", dt_id, symbol, s_open, s_close, s_high, s_low, s_vol, s_count, s_money)
		vals = append(vals, dt_id, symbol, s_open, s_close, s_high, s_low, s_vol, s_count, s_money)
	}
	f.Close()
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	//fmt.Println(sqlStr)
	//prepare the statement
	stmt, _ := cj_db.G_db.Prepare(sqlStr)

	//format all vals at once

	res, err := stmt.Exec(vals...)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.LastInsertId())
	cj_db.DB_twse_download_day_list_is_insert_update(dt_id)
	cj_db.DB_twse_stock_day_set_update()
	stmt.Close()
	cj_db.Close()

	//os.Exit(1)
}
