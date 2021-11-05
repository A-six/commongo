package datex

import (
	"fmt"
	"testing"
	"time"
)

//func TestLately7DateStr(t *testing.T) {
//	fmt.Println(Lately7DateStr(time.Now()))
//	fmt.Println(Lately7DateExcludingTodayStr(time.Now()))
//	fmt.Println(MonthDateStr(time.Now()))
//}
//
//func TestMonthDateStr(t *testing.T) {
//	fmt.Println(MonthDateStr(time.Now()))
//}
//
//func TestIncludingDate(t *testing.T) {
//	fmt.Println(IncludingDate("2021-01-01", "2021-01-31", "2021-01-01"))
//	fmt.Println(IncludingDate("2021-02-01", "2021-02-31", "2021-02-02"))
//	fmt.Println(IncludingDate("2021-03-01", "2021-03-31", "2021-03-03"))
//	fmt.Println(IncludingDate("2021-04-01", "2021-04-31", "2021-04-04"))
//	fmt.Println(IncludingDate("2021-05-01", "2021-05-31", "2021-05-05"))
//	fmt.Println(IncludingDate("2021-06-01", "2021-06-31", "2021-06-06"))
//	fmt.Println(IncludingDate("2021-07-01", "2021-07-31", "2021-07-07"))
//}

func TestDateStEdToDatas(t *testing.T) {
	st1, _ := time.Parse("2006-01-02", "2021-01-01")
	ed1, _ := time.Parse("2006-01-02", "2021-12-15")

	fmt.Println(DateStEdToDatas(st1, ed1))
}
