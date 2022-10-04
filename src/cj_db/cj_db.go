package cj_db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var G_db *sql.DB
var G_db_dsn string

func Connect() *sql.DB {
	fmt.Println("Go MySQL Connect")

	db, err := sql.Open("mysql", G_db_dsn)

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatal(err)
	}
	G_db = db
	fmt.Println("connect db success")
	return G_db
}

func Close() {
	G_db.Close()
}

func Query(sql string, args ...interface{}) {
	results, err := G_db.Query(sql, args...)

	if err != nil {
		fmt.Println(err)
	} else {
		results.Close()
	}
}

func Exec(sql string, args ...interface{}) {
	results, err := G_db.Query(sql, args...)

	if err != nil {
		fmt.Println(err)
	} else {
		results.Close()
	}
}

func Hello_show() int {
	fmt.Println("Hello, 世界")
	//G_db = 123555
	return 1
}

func Change_date(add_day int) string {
	dt := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	//fmt.Println(dt.GoString())
	dt2 := dt.AddDate(0, 0, add_day)
	dt_date := dt2.Format("2006-01-02")
	//fmt.Println("MM-DD-YYYY : ", dt.Format("2006-01-02"))
	//fmt.Println("MM-DD-YYYY : ", dt_date)
	return dt_date
}

func Change_datetime(add_day int, add_time int) string {
	dt := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	//fmt.Println(dt.GoString())
	dt2 := dt.AddDate(0, 0, add_day)
	dt3 := dt2.Add(time.Duration(add_time) * time.Second)
	dt_date := dt3.Format("2006-01-02")
	dt_datetime := fmt.Sprintf("%s %02d:%02d:%02d", dt_date, dt3.Hour(), dt3.Minute(), dt3.Second())

	//dt_date := dt3.Format(time.RFC3339)
	fmt.Println(dt_datetime)
	//fmt.Println("MM-DD-YYYY : ", dt.Format("2006-01-02"))
	//fmt.Println("MM-DD-YYYY : ", dt_date)
	return dt_datetime
}

// 取得相差天數
func DBF_diffDay(dt_start string) int {

	var diffday int

	sql := "SELECT DATEDIFF(NOW(), \"" + dt_start + "\") as diffday"
	row := G_db.QueryRow(sql)
	err := row.Scan(&diffday)
	if err != nil {
		fmt.Println(sql)
		log.Fatal(err)
	}
	return diffday
}

// 取得相差天數
func DBF_diffMonth(dt_start string) int {

	var diffMonth int

	row := G_db.QueryRow("SELECT TIMESTAMPDIFF(MONTH, ?, NOW())", dt_start)
	err := row.Scan(&diffMonth)
	if err != nil {
		log.Fatal(err)
	}
	return diffMonth
}

// 取得相差天數
func DBF_diffTwoMonth(dt_start string, dt_end string) int {

	var diffMonth int

	row := G_db.QueryRow("SELECT TIMESTAMPDIFF(MONTH, ?, ?)", dt_start, dt_end)
	err := row.Scan(&diffMonth)
	if err != nil {
		log.Fatal(err)
	}
	return diffMonth
}
