package gutime

import (
	"strconv"
	"strings"
	"time"
)

//)

//func GetTime() string {

//	t := time.Now()
//	out := (t.Format("20060102150405"))
//	return out
//}

//func GetTimeLongVersion() string {

//	t := time.Now()
//	out := (t.Format("2006/01/02 15:04:05"))
//	return out
//}

//func GetDDay() string {

//	t := time.Now()
//	//      out := (t.Format("20180101"))
//	out := (t.Format("20060102"))
//	return out
//}

func GetHourMinSec(cleaned bool) string {

	t := time.Now()
	out := (t.Format("15:04:05"))

	if cleaned == true {
		out = strings.Replace(out, ":", "", -1)
	}

	return out
}

func GetHourMinSecAsInt() int {

	tmp := GetHourMinSec(true)
	out, _ := strconv.Atoi(tmp)
	return out
}

//func DateConvertor(s string) string {

//	year := (s[0:4])
//	month := (s[5:7])
//	day := (s[6:8])

//	hours := (s[8:10])
//	minutes := (s[10:12])
//	seconds := (s[12:14])

//	out := year + "/" + month + "/" + day + " - " + hours + ":" + minutes + ":" + seconds

//	return out
//}

//func GetTodayAndYesterday(toRevert bool) []string {
//	today := GetDDay()

//	i, _ := strconv.Atoi((GetDDay()))
//	d := int64(i) - 1
//	yesterday := IntToString(d)

//	var days []string
//	days = append(days, yesterday)
//	days = append(days, today)

//	if toRevert == true {
//		for i, j := 0, len(days)-1; i < j; i, j = i+1, j-1 {
//			days[i], days[j] = days[j], days[i]
//		}
//	}

//	return days
//}

/////////////////// To move //////////////////

//func IntToString(target int64) string {
//	return strconv.FormatInt(target, 10) // , 64)
//	//      (target, 'f', 6, 64)
//}

/////////////////// To move //////////////////

//func GetLastSevenDays() []string {

//	var days []string
//	today, _ := strconv.Atoi(GetDDay())

//	for i := 0; i < 7; i++ {
//		d := int64(today) - int64(i)
//		days = append(days, IntToString(d))
//	}

//	for i, j := 0, len(days)-1; i < j; i, j = i+1, j-1 {
//		days[i], days[j] = days[j], days[i]
//	}

//	return days
//}
