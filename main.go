package main

import (
	"cj_db"
	"flag"
	"fmt"

	"cj_func"
	"cj_html"
)

func main() {

	filenamePtr := flag.String("ini", "sk88.ini", "a string")

	//numbPtr := flag.Int("numb", 42, "an int")
	//boolPtr := flag.Bool("sk88day", false, "a bool")

	//var svar string
	//flag.StringVar(&svar, "svar", "bar", "a string var")
	sk88day := flag.Bool("sk88day", false, "a bool")
	sk88week := flag.Bool("sk88week", false, "a bool")
	sk88month := flag.Bool("sk88month", false, "a bool")

	maday := flag.Bool("maday", false, "a bool")
	maweek := flag.Bool("maweek", false, "a bool")
	mamonth := flag.Bool("mamonth", false, "a bool")

	daylimit := flag.Bool("daylimit", false, "a bool")

	twse_day_csv := flag.Bool("twsedaycsv", false, "a bool")
	twse_day_csv_insert := flag.Bool("twsedayinsert", false, "a bool") // 新增資料到資料庫

	mr_increase := flag.Bool("mr_increase", false, "a bool")
	mr_increase_year := flag.Bool("mr_year", false, "a bool")

	futures_1min := flag.Bool("futures_1min", false, "a bool")

	// twse 綜合損益表 consolidated income statement 抓取
	twse_consolidated_income_statement := flag.Bool("twse_cis", false, "a bool")

	flag.Parse()
	fmt.Println("tail:", flag.Args())
	cj_func.Load_file(*filenamePtr)

	cj_db.G_db_dsn = cj_func.G_db_dsn
	cj_html.User_agent = cj_func.G_user_agent

	// 抓取每日股票
	if *sk88day {
		Sk88_load_day()
	}

	// 抓取每星期報表
	if *sk88week {
		Sk88_load_week()
	}

	// 抓取每月報表
	if *sk88month {
		Sk88_load_month()
	}

	// ma 統計 (天)
	if *maday {
		Stock_day_update_ma()
	}

	// ma 統計 (星期)
	if *maweek {
		Stock_week_update_ma()
	}

	// ma 統計 (月)
	if *mamonth {
		Stock_month_update_ma()
	}

	// 取得漲停跌停價
	if *daylimit {
		Stock_day_limit()
	}

	// 每日收盤行情 - TWSE 臺灣證券交易所
	if *twse_day_csv {
		Twse_day_csv_start()
	}

	// 統計財務 (月)
	if *mr_increase {
		Statistics_monthly_revenue_increase_start()
	}

	// 統計財務 (年)
	if *mr_increase_year {
		Statistics_monthly_revenue_year_start()
	}

	// 期貨五分鐘線
	if *futures_1min {
		Sk88_load_futures_min()
	}

	// 每日收盤行情 - TWSE 臺灣證券交易所
	if *twse_day_csv_insert {
		Twse_day_stock_csv_insert()
		Twse_stock_day_update_ma()
	}

	// 綜合損益表 - 公開資訊觀測站
	// 抓取綜合損益表
	if *twse_consolidated_income_statement {
		fmt.Println("hihihihi");
		Twse_cis_season_csv_start();
	}

	//fmt.Println("word:", *wordPtr)
	//fmt.Println("numb:", *numbPtr)
	//fmt.Println("fork:", *boolPtr)
	//fmt.Println("svar:", svar)

}
