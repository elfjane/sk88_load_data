package cj_func

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func Change_date1(add_day int) string {
	dt := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	//fmt.Println(dt.GoString())
	dt2 := dt.AddDate(0, 0, add_day)
	dt_date := dt2.Format("2006-01-02")
	//fmt.Println("MM-DD-YYYY : ", dt.Format("2006-01-02"))
	//fmt.Println("MM-DD-YYYY : ", dt_date)
	return dt_date
}

// 置換 SQL 衝突字元
func Func_base_replace_string(str string) string {
	ret_str := strings.Replace(str, "'", "\\'", -1)
	ret_str = strings.TrimSpace(ret_str)
	return ret_str
}

// 置換 SQL 衝突字元包含逗點
func Func_base_replace_string_dot(str string) string {
	ret_str := strings.Replace(str, "'", "\\'", -1)
	ret_str = strings.Replace(ret_str, ",", "", -1)
	ret_str = strings.Replace(ret_str, "--", "-1", -1)
	ret_str = strings.TrimSpace(ret_str)
	return ret_str
}

func Func_read_file(filename string) []string {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	return text
}
