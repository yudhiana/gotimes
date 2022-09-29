package gotimes

import "time"

type gotimesImpl interface {
	SetToday(today time.Time) *Gotimes
	SetHolidays(h []time.Time) *Gotimes
	AddWeekDay(sla int) (t time.Time)
	IsWeekend() (w bool)
}
