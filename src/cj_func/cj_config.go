package cj_func

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var G_uuid string
var G_db_dsn string
var G_user_agent string
var G_save_twse_day_init_date string
var G_save_twse_day_csv string
var G_save_twse_day_url string

var G_save_twse_cis_season_init_year string
var G_save_twse_cis_season_csv string
var G_save_twse_cis_season_url string

var g_host string
var g_port string
var g_username string
var g_password string
var g_database string

func Load_file(filename string) {
	fmt.Println("Go MySQL Tutorial")
	cfg, err := ini.Load(filename)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	G_user_agent = cfg.Section("global").Key("user_agent").String()
	G_uuid = cfg.Section("global").Key("uuid").String()

	G_save_twse_cis_season_csv = cfg.Section("twse_day_csv").Key("save_path").String()
	G_save_twse_cis_season_url = cfg.Section("twse_day_csv").Key("twse_url").String()
	G_save_twse_cis_season_init_year = cfg.Section("twse_day_csv").Key("init_year").String()

	// 綜合損益表 - 公開資訊觀測站 相關預設參數
	G_save_twse_day_csv = cfg.Section("twse_cis_csv").Key("save_path").String()
	G_save_twse_day_url = cfg.Section("twse_cis_csv").Key("twse_cis_url").String()
	G_save_twse_day_init_date = cfg.Section("twse_cis_csv").Key("init_year").String()

	g_host = cfg.Section("database").Key("host").String()
	g_port = cfg.Section("database").Key("port").String()

	if g_port == "" {
		g_port = "3306"
	}
	g_username = cfg.Section("database").Key("username").String()
	g_password = cfg.Section("database").Key("password").String()
	g_database = cfg.Section("database").Key("database").String()

	fmt.Println("host = ", g_host)
	fmt.Println("port = " + g_port)

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	G_db_dsn = g_username + ":" + g_password + "@tcp(" + g_host + ":" + g_port + ")/" + g_database
	fmt.Println("dsn = ", G_db_dsn)

}
