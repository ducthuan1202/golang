package timeformat

import (
	"time"
	"strings"
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

const (
	Year string = "2006" 		// YYYY
	Month string = "01"			// MM
	Day string = "02"			// DD
	Hour string = "03"			// h
	HourUp string = "15"		// H
	Minute string = "04"		// i
	Second string = "05"		// s
	Flag string = "PM"			// F
	
	DayOfWeek string = "Monday"		// J
	DayOfWeekLess string = "Mon"	// j
	MonthOfYear string = "January"	// O
	MonthOfYearLess string = "Jan"	// o
)

func convert(strFormat string) string {
	syntaxReplace := strings.NewReplacer(
		"YYYY", Year, 
		"MM", Month, 
		"DD", Day, 
		"h", Hour,
		"H", HourUp,
		"i", Minute,
		"s", Second,
		"F", Flag,
		"J", DayOfWeek,
		"j", DayOfWeekLess,
		"O", MonthOfYear,
		"o", MonthOfYearLess,
	)
	return syntaxReplace.Replace(strFormat)
}


func TimeToString(dt time.Time, strFormat string) string {
	return dt.Format(convert(strFormat))	
}


func StringToTime(strTime string, strFormat string) (time.Time, error) {	
	return time.Parse(convert(strFormat), strTime)	
}