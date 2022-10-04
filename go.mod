module sk88_load_data.so

go 1.17

require github.com/go-sql-driver/mysql v1.6.0

require (
	cj_db v1.0.0
	cj_func v1.0.0
	cj_html v1.0.0
	github.com/xuri/excelize/v2 v2.5.0
	golang.org/x/text v0.3.7
)

replace cj_html v1.0.0 => ./src/cj_html

replace cj_db v1.0.0 => ./src/cj_db

replace cj_func v1.0.0 => ./src/cj_func

require (
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/richardlehane/mscfb v1.0.3 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/xuri/efp v0.0.0-20210322160811-ab561f5b45e3 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
	gopkg.in/ini.v1 v1.63.2 // indirect
)
