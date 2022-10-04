package main

import (
	"cj_db"
	"cj_func"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

// 下載台灣證卷交易所 每日收盤行情
func Twse_cis_season_csv_start() {
	fmt.Println("twsw day csv down start")
	cj_db.Connect()

	twse_cis_season_csv_init()
	twse_cis_season_csv_download_start()

	cj_db.Close()
}

// 初始化下載資料
func twse_cis_season_csv_init() {
	fmt.Println("twsw consolidated income statement csv down init")

	dt_start := cj_db.DB_twse_download_day_list_lastday()
	fmt.Println(dt_start)
	if dt_start == "" {
		dt_start = cj_func.G_save_twse_cis_season_init_year
	}
	fmt.Println(dt_start)
	return
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
func twse_cis_season_csv_download_start() {
	fmt.Println("twsw day csv download start")
	twse_no_download_list := cj_db.DB_twse_download_day_list_get()
	dt := time.Now()
	dt_day := fmt.Sprintf("%04d%02d%02d_%d", dt.Year(), dt.Month(), dt.Day(), dt.Unix())
	twse_cis_season_csv_download(dt_day+".html", "102", "1")
	return
	for _, this := range twse_no_download_list {

		//dtstr1 := this.Dt_id + " 00:00:00"
		//dt, _ := time.Parse("2006-01-02 15:04:05", dtstr1)
		//dt_day := fmt.Sprintf("%04d%02d%02d", dt.Year(), dt.Month(), dt.Day())
		pause_time := rand.Int31n(5) + 5
		fmt.Printf("pause time wait  %d sec download = %s\r\n", pause_time, this.Dt_id)
		//twse_cis_season_csv_download(this.Csv_filename, this.Dt_id, dt_day)

		return
		time.Sleep(time.Second * time.Duration(pause_time))
	}
}

// 下載台灣證卷交易所 每日收盤行情
/*
encodeURIComponent: 1
step: 1
firstin: 1
off: 1
isQuery: Y
TYPEK: sii
year: 102
season: 01
*/
func twse_cis_season_csv_download(csv_filename string, year string, season string) {

	/*
		cis_url := cj_func.G_save_twse_cis_season_url
		fmt.Println(cis_url)
		//os.Exit(1)
		data := url.Values{}
		data.Add("encodeURIComponent", "1")
		data.Add("step", "1")
		data.Add("firstin", "1")
		data.Add("off", "1")
		data.Add("isQuery", "Y")
		data.Add("TYPEK", "sii")
		data.Add("year", year)
		data.Add("season", season)
		body := cj_html.Twse_Post(cis_url, data)


		path_filename := cj_func.G_save_twse_cis_season_csv + "/" + csv_filename
		//big5ToUTF8 := traditionalchinese.Big5.NewDecoder()
		//big5Test := "\xb4\xfa\xb8\xd5" // 測試的 Big5 編碼
		//big5Test := string(body)
		//utf8, _, _ := transform.String(big5ToUTF8, big5Test)
		//fmt.Println(utf8) // 顯示「測試」
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
		fmt.Println(size)
		//cj_db.DB_twse_download_day_list_download_update(dt_date, size)
		// create from a file
	*/
	path_filename := cj_func.G_save_twse_cis_season_csv + "/20220830_1661854457.html"
	f, err := os.Open(path_filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".hasBorder").Each(func(index int, ele *goquery.Selection) {

		ele.Find("tr").Each(func(index_tr int, ele_tr *goquery.Selection) {
			fmt.Println(index_tr)
			fmt.Println(ele_tr.Text())
			ele_tr.Find("td").Each(func(index_td int, ele_td *goquery.Selection) {
				fmt.Println(index_td)
				fmt.Println(ele_td.Text())
			})
		})

		fmt.Println(index)
	})

}
