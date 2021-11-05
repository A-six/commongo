package datex

import "time"

func WeekMondaySt(t time.Time) (dates []string) {
	if t.Weekday() != time.Monday {
		for i := 0; i < 6; i++ {
			t = t.Add(time.Hour * -24)
			if t.Weekday() == time.Monday {
				break
			}
		}
	}
	dates = append(dates, t.Format("2006-01-02"))
	for i := 0; i < 6; i++ {
		t = t.Add(time.Hour * 24)
		dates = append(dates, t.Format("2006-01-02"))
	}
	return
}
