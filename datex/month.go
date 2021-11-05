package datex

import "time"

func MonthDateStr(t time.Time) (dates []string) {
	tMonth := t.Month()
	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	dates = append(dates, t.Format("2006-01-02"))
	for i := 0; i < 32; i++ {
		t = t.Add(time.Hour * 24)
		if t.Month() != tMonth {
			return
		}
		dates = append(dates, t.Format("2006-01-02"))
	}
	return
}
