package gotimes

import (
	"time"
)

type gotimes struct {
	today   time.Time
	holiday []time.Time
}

type GoTimes struct {
	Time time.Time
	*gotimes
}

var weekend = map[int]bool{
	int(time.Saturday): true,
	int(time.Sunday):   true,
}

func newTime() goTimesImpl {
	return &gotimes{
		today: time.Now(),
	}
}

func (gt *gotimes) isWeekend(d time.Time) (w bool) {
	return weekend[int(d.Weekday())]
}

func (gt *gotimes) SetToday(today time.Time) *gotimes {
	gt.today = today
	return gt
}

func (gt *gotimes) SetHolidays(h []time.Time) *gotimes {
	gt.holiday = append(gt.holiday, h...)
	return gt
}

func (gt *gotimes) AddWeekDay(sla int) (t time.Time) {
	for i := 1; i <= sla; i++ {
		nextDay := gt.today.AddDate(0, 0, i)
		if gt.isWeekend(nextDay) {
			sla++
		} else {
			if gt.isOffDay(nextDay) {
				sla++
			}
			t = gt.today.AddDate(0, 0, i)
		}
	}
	return
}

func (gt *gotimes) isOffDay(d time.Time) (h bool) {
	for _, dayOff := range gt.holiday {
		if dayOff.Equal(d) {
			return true
		}
	}
	return
}

func AddWeekDay(day int, holidays []time.Time) time.Time {
	return newTime().SetToday(time.Now()).SetHolidays(holidays).AddWeekDay(day)
}
