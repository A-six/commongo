package datex

import (
	"fmt"
	"testing"
	"time"
)

func TestMonthDateStr(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2021-01-01")
	t2, _ := time.Parse("2006-01-02", "2021-02-01")
	fmt.Println(MonthDateStr(t1))
	fmt.Println(MonthDateStr(t2))
}

func TestMonthDateStr2(t *testing.T) {
	t2011, _ := time.Parse("2006-01-02", "2011-02-01")
	t2012, _ := time.Parse("2006-01-02", "2012-02-01")
	t2013, _ := time.Parse("2006-01-02", "2013-02-01")
	t2014, _ := time.Parse("2006-01-02", "2014-02-01")
	t2015, _ := time.Parse("2006-01-02", "2015-02-01")
	t2016, _ := time.Parse("2006-01-02", "2016-02-01")
	t2017, _ := time.Parse("2006-01-02", "2017-02-01")
	t2018, _ := time.Parse("2006-01-02", "2018-02-01")
	t2019, _ := time.Parse("2006-01-02", "2019-02-01")
	t2020, _ := time.Parse("2006-01-02", "2020-02-01")
	t2021, _ := time.Parse("2006-01-02", "2021-02-01")
	fmt.Println(MonthDateStr(t2011))
	fmt.Println(MonthDateStr(t2012))
	fmt.Println(MonthDateStr(t2013))
	fmt.Println(MonthDateStr(t2014))
	fmt.Println(MonthDateStr(t2015))
	fmt.Println(MonthDateStr(t2016))
	fmt.Println(MonthDateStr(t2017))
	fmt.Println(MonthDateStr(t2018))
	fmt.Println(MonthDateStr(t2019))
	fmt.Println(MonthDateStr(t2020))
	fmt.Println(MonthDateStr(t2021))
}
