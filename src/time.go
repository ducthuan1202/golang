package main

import (
	"fmt"
	"time"
	"packages/timeformat"
)

/*
* Thời gian tham chiếu cụ thể là
* Mon Jan 2 15:04:05 MST 2006
* Giá trị ở Unix là: 1136239445
* GMT-0700: 
* 01/02 03:04:05PM '06 -0700
*  1  2  3  4  5     6  -7
* ngày 2, tháng 1, năm 2006, giờ 3, phút 4, giây 5, timezone -6
*/

func main(){

	timeString := "2020-04-21 23:30:45"
	// chuyển đổi 1 chuỗi sang dạng time.Time
	df, _ := timeformat.StringToTime(timeString, "YYYY-MM-DD H:i:s")

	// định dạng time sang năm-tháng-ngày
	fmt.Println(timeformat.TimeToString(df, "YYYY-MM-DD, J"))

	// định dạng time sang giờ:phút:giây
	fmt.Println(timeformat.TimeToString(df, "h:i:sF"))

	// Ví dụ về format
	now := time.Now()

	// Lấy 1 chuỗi theo định dạng YYYY-MM-DD
	fmt.Println("YYYY-MM-DD \t", timeformat.TimeToString(now, "j YYYY-MM-DD"))

	// Lấy 1 chuỗi theo định dạng DD/MM/YYYY
	fmt.Println("DD/MM/YYYY \t", timeformat.TimeToString(now, "DD/MM/YYYY"))

	// Lấy 1 chuỗi theo định dạng YYYY-MM-DD H:i:s
	fmt.Println("YYYY-MM-DD H:i:s", timeformat.TimeToString(now, "YYYY-MM-DD H:i:s"))

	// Lấy 1 chuỗi theo định dạng h:i:s
	fmt.Println("h:i:s \t\t", timeformat.TimeToString(now, "h:i:s"))

	// Lấy 1 chuỗi theo định dạng h:i:sm
	fmt.Println("h:i:sF \t\t", timeformat.TimeToString(now, "h:i:sF"))

	// Lấy 1 chuỗi theo định dạng H:i:s
	fmt.Println("H:i:s \t\t", timeformat.TimeToString(now, "H:i:s"))

}

func initial (){

	type Status int
	const (
		Pending Status = 1 + iota
		Success
		Cancel
	)

	var myStt Status
	myStt = Success
	fmt.Println(myStt)

	fmt.Println(Pending, Success, Cancel)
}