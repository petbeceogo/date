package date_test

import (
	"testing"
	"time"

	"github.com/devjoeween/smfp/v1"
	"github.com/petbeceogo/date"
)

func TestBetween(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Seoul")
	from := time.Date(2022, time.December, 28, 0, 0, 0, 0, loc)
	to := time.Date(2023, time.January, 5, 0, 0, 0, 0, loc)

	dates := date.Between(from, to)

	if len(dates) != 9 {
		t.Error("dates length should be 9")
		return
	}

	expected := []date.Date{
		{2022, 12, 28},
		{2022, 12, 29},
		{2022, 12, 30},
		{2022, 12, 31},
		{2023, 1, 1},
		{2023, 1, 2},
		{2023, 1, 3},
		{2023, 1, 4},
		{2023, 1, 5},
	}
	success := smfp.Every(func(item date.Date, index int) bool {
		exp := expected[index]

		return exp.Year == item.Year && exp.Month == item.Month && exp.Day == item.Day
	}).Execute(dates).(bool)

	if !success {
		t.Fail()
		return
	}
}
