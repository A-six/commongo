package datex

import "time"

//func Lately6DateExcludingTodayStr(t time.Time) (dates []string) {
//	for i := 0; i < 6; i++ {
//		t = t.Add(time.Hour * -24)
//		dates = append(dates, t.Format("2006-01-02"))
//	}
//	return
//}
//
//func Lately7DateStr(t time.Time) (dates []string) {
//	dates = append(dates, t.Format("2006-01-02")) //1
//	for i := 0; i < 6; i++ {
//		t = t.Add(time.Hour * -24)
//		dates = append(dates, t.Format("2006-01-02"))
//	}
//	return
//}
//
//func Lately7DateExcludingTodayStr(t time.Time) (dates []string) {
//	for i := 0; i < 7; i++ {
//		t = t.Add(time.Hour * -24)
//		dates = append(dates, t.Format("2006-01-02"))
//	}
//	return
//}
//
//func MonthDateStr(t time.Time) (dates []string) {
//	tMonth := t.Month()
//	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
//	dates = append(dates, t.Format("2006-01-02"))
//	for i := 0; i < 32; i++ {
//		t = t.Add(time.Hour * 24)
//		if t.Month() != tMonth {
//			return
//		}
//		dates = append(dates, t.Format("2006-01-02"))
//	}
//	return
//}
//
//func IncludingDate(stDate, edDate, date string) bool {
//	if date >= stDate && date <= edDate {
//		return true
//	}
//	return false
//}
//
//func ContainDates(dates []string, todayDate string) bool {
//	for _, date := range dates {
//		if todayDate == date {
//			return true
//		}
//	}
//	return false
//}
//
//func AnalysisStEdDate(stDate, edDate time.Time) (dates []string) {
//	stdate := edDate.Format("2006-01-02")
//	for {
//		dates = append(dates, stDate.Format("2006-01-02"))
//		stDate = stDate.Add(time.Hour * 24)
//
//		if stDate.Format("2006-01-02") > stdate {
//			return
//		}
//	}
//}

func DateStEdToDatas(stDate, edDate time.Time) (dates []string) {
	stdate := edDate.Format("2006-01-02")
	for {
		dates = append(dates, stDate.Format("2006-01-02"))
		stDate = stDate.Add(time.Hour * 24)

		if stDate.Format("2006-01-02") > stdate {
			return
		}
	}
}
