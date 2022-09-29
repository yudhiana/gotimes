package gotimes

import (
	"encoding/json"
	"fmt"
	"time"
)

type Gotimes struct {
	Today   time.Time
	Holiday []time.Time
	Error   error
}

var weekend = map[int]bool{
	int(time.Saturday): true,
	int(time.Sunday):   true,
}

func newTime() gotimesImpl {
	return &Gotimes{
		Today: time.Now(),
	}
}

func (gt *Gotimes) IsWeekend() (w bool) {
	return weekend[int(gt.Today.Weekday())]
}

func (gt *Gotimes) SetToday(today time.Time) *Gotimes {
	gt.Today = today
	return gt
}

func (gt *Gotimes) SetHolidays(h []time.Time) *Gotimes {
	gt.Holiday = append(gt.Holiday, h...)
	return gt
}

func (gt *Gotimes) Println() {
	gbyte, _ := json.Marshal(*gt)
	fmt.Println(string(gbyte))
}

func (gt *Gotimes) IsOffDay(d time.Time) (h bool) {
	for _, dayOff := range gt.Holiday {
		if dayOff.Equal(d) {
			return true
		}
	}
	return
}

func (gt *Gotimes) AddDate(years, month, days int) time.Time {
	return gt.Today.AddDate(years, month, days)
}

func (gt *Gotimes) Add(d time.Duration) time.Time {
	return gt.Today.Add(d)
}

func (gt *Gotimes) AddWeekDay(sla int) (t time.Time) {
	for i := 1; i <= sla; i++ {
		nextDay := gt.Today.AddDate(0, 0, i)
		if weekend[int(nextDay.Weekday())] {
			sla++
		} else {
			if gt.IsOffDay(nextDay) {
				sla++
			}
			t = gt.Today.AddDate(0, 0, i)
		}
	}
	return
}

func AddWeekDay(d int) (t time.Time) {
	return newTime().AddWeekDay(d)
}
