package gotimes_test

import (
	"gotimes"
	"testing"
)

func TestToday(t *testing.T) {
	t.Run("testAddWeekDay", func(t *testing.T) {
		expected := "2022-10-10"
		actual := gotimes.AddWeekDay(7).Format("2006-01-02")
		if actual != expected {
			t.Errorf("invalid addWeekDay\n\tExpected : %v\n\tActual : %v", expected, actual)
		}
	})

}
