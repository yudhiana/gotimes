package gotimes_test

import (
	"encoding/json"
	"gotimes"
	"io"
	"os"
	"testing"
	"time"
)

func TestGotimes(t *testing.T) {
	t.Run("testAddWeekDay", func(t *testing.T) {
		fJSON, _ := os.Open("example_day_off_national_INA.json")
		defer fJSON.Close()

		byteValue, _ := io.ReadAll(fJSON)
		var holidaysData []map[string]interface{}
		var holidays []time.Time

		json.Unmarshal(byteValue, &holidaysData)
		for _, dayOff := range holidaysData {
			dof, _ := time.Parse("2006-01-02 15:04:05", dayOff["calendar_date"].(string))
			holidays = append(holidays, dof)
		}

		expected := "2022-10-10"
		today := time.Now()
		actual := gotimes.AddWeekDay(7, &today, holidays).Format("2006-01-02")
		if actual != expected {
			t.Errorf("invalid addWeekDay\n\tExpected : %v\n\tActual : %v", expected, actual)
		}

		actualUnsetToday := gotimes.AddWeekDay(7, nil, holidays).Format("2006-01-02")
		if actualUnsetToday != expected {
			t.Errorf("invalid addWeekDay\n\tExpected : %v\n\tActual : %v", expected, actualUnsetToday)
		}

		actualUnsetHoliday := gotimes.AddWeekDay(7, &today, nil).Format("2006-01-02")
		if actualUnsetHoliday != expected {
			t.Errorf("invalid addWeekDay\n\tExpected : %v\n\tActual : %v", expected, actualUnsetHoliday)
		}

		actualUnsetTodayAndHolidays := gotimes.AddWeekDay(7, nil, nil).Format("2006-01-02")
		if actualUnsetTodayAndHolidays != expected {
			t.Errorf("invalid addWeekDay\n\tExpected : %v\n\tActual : %v", expected, actualUnsetTodayAndHolidays)
		}
	})

}
