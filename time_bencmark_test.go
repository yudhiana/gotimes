package gotimes_test

import (
	"gotimes"
	"testing"
)

func BenchmarkAddWeekDay(b *testing.B) {
	gotimes.AddWeekDay(7, nil, nil)
}
