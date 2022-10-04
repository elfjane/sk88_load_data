package main

import (
	"cj_db"
	"cj_func"
	"cj_html"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 下載台灣證卷交易所 每日收盤行情
func Twse_day_csv_start() {
	fmt.Println("twsw day csv down start")
	cj_db.Connect()

	twse_day_csv_init()
	twse_day_csv_download_start()

	cj_db.Close()
}

// 初始化下載資料
func twse_day_csv_init() {
	fmt.Println("twsw day csv down init")

	dt_start := cj_db.DB_twse_download_day_list_lastday()
	if dt_start == "" {
		dt_start = cj_func.G_save_twse_day_init_date
	}
	fmt.Println(dt_start)

	diffday := cj_db.DBF_diffDay(dt_start)
	fmt.Println(diffday)

	dtstr1 := dt_start + " 00:00:00"
	dt, _ := time.Parse("2006-01-02 15:04:05", dtstr1)
	fmt.Println(dt)

	i := 1             // i預設為 0
	for i <= diffday { // i<=10 為真就執行 {} 內的，否則不執行
		dt_add := dt.AddDate(0, 0, i)
		fmt.Println(dt_add)
		dt_day := fmt.Sprintf("%04d%02d%02d", dt_add.Year(), dt_add.Month(), dt_add.Day())
		dt_id := fmt.Sprintf("%04d-%02d-%02d", dt_add.Year(), dt_add.Month(), dt_add.Day())
		path_filename := fmt.Sprintf("%04d/%s.csv", dt_add.Year(), dt_day)
		cj_db.DB_twse_download_day_list_insert_init(dt_id, path_filename)
		i = i + 1
	}
	cj_db.DB_twse_download_day_list_delete_week()
}

// 下載台灣證卷交易所 每日收盤行情
func twse_day_csv_download_start() {
	fmt.Println("twsw day csv download start")
	twse_no_download_list := cj_db.DB_twse_download_day_list_get()

	for _, this := range twse_no_download_list {

		dtstr1 := this.Dt_id + " 00:00:00"
		dt, _ := time.Parse("2006-01-02 15:04:05", dtstr1)
		dt_day := fmt.Sprintf("%04d%02d%02d", dt.Year(), dt.Month(), dt.Day())
		pause_time := rand.Int31n(5) + 5
		fmt.Printf("pause time wait  %d sec download = %s\r\n", pause_time, this.Dt_id)
		twse_day_csv_download(this.Csv_filename, this.Dt_id, dt_day)
		time.Sleep(time.Second * time.Duration(pause_time))
	}
}

// 下載台灣證卷交易所 每日收盤行情
func twse_day_csv_download(csv_filename string, dt_date string, dt_day string) {

	url := cj_func.G_save_twse_day_url + "?response=csv&date=" + dt_day + "&type=ALL"
	fmt.Println(url)

	//os.Exit(1)
	body, filename := cj_html.Twse_Get_CVS(url)

	path_filename := cj_func.G_save_twse_day_csv + "/" + csv_filename
	err := os.WriteFile(path_filename, body, 0644)
	if err != nil {
		panic(err)
	}

	fi, err2 := os.Stat(path_filename)
	if err2 != nil {
		panic(err)
	}
	// get the size
	size := fi.Size()
	//fmt.Println(body)
	fmt.Println(filename)
	//fmt.Println(size)
	cj_db.DB_twse_download_day_list_download_update(dt_date, size)
}
