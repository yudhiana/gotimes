package gotimes

import "time"

type GoTimesImpl interface {
	SetToday(today time.Time) *gotimes
	SetHolidays(h []time.Time) *gotimes
	AddWeekDay(sla int) (t time.Time)
}
