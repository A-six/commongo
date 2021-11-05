package datex

import (
	"fmt"
	"testing"
	"time"
)

func TestWeekMondaySt(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2021-11-01")
	t2, _ := time.Parse("2006-01-02", "2021-11-02")
	t3, _ := time.Parse("2006-01-02", "2021-11-03")
	t4, _ := time.Parse("2006-01-02", "2021-11-04")
	t5, _ := time.Parse("2006-01-02", "2021-11-05")
	t6, _ := time.Parse("2006-01-02", "2021-11-06")
	t7, _ := time.Parse("2006-01-02", "2021-11-07")
	t8, _ := time.Parse("2006-01-02", "2021-11-08")
	t9, _ := time.Parse("2006-01-02", "2021-11-09")

	fmt.Println(WeekMondaySt(t1))
	fmt.Println(WeekMondaySt(t2))
	fmt.Println(WeekMondaySt(t3))
	fmt.Println(WeekMondaySt(t4))
	fmt.Println(WeekMondaySt(t5))
	fmt.Println(WeekMondaySt(t6))
	fmt.Println(WeekMondaySt(t7))
	fmt.Println(WeekMondaySt(t8))
	fmt.Println(WeekMondaySt(t9))
}
