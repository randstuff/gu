package gutime

import (
	"fmt"
	"testing"
)

// testing "hack" to see the output
//

func TestGetTime(t *testing.T) {
	fmt.Println(GetTime())
}

func TestGetDDay(t *testing.T) {
	fmt.Println(GetDDay())
}

func TestGetHourMinSec(t *testing.T) {
	fmt.Println(GetHourMinSec(true))
	fmt.Println(GetHourMinSec(false))
}

func TestGetHourMinSecAsInt(t *testing.T) {
	fmt.Println(GetHourMinSecAsInt())
}
func TestGetDate(t *testing.T) {
	fmt.Println(GetTimeLongVersion())
}

func TestDateConvertor(t *testing.T) {
	out := GetTime()
	fmt.Println(DateConvertor(out))
}

func TestGetTodayAndYesterday(t *testing.T) {
	fmt.Println(GetTodayAndYesterday(true))
	fmt.Println(GetTodayAndYesterday(false))
}

func TestGetLastSevenDays(t *testing.T) {
	fmt.Println(GetLastSevenDays())
}
