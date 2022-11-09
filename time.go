package date

import (
	"time"

	"github.com/petbeceogo/date/dateloc"
)

func Now() time.Time {
	return time.Now().In(dateloc.Current())
}

func NowAfter(d time.Duration) time.Time {
	return Now().Add(d)
}

func NewStart(
	year int,
	month int,
	day int,
) time.Time {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, dateloc.Current())

	return t
}

func NewEnd(
	year int,
	month int,
	day int,
) time.Time {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, dateloc.Current())
	t = t.Add((time.Duration(24) * time.Hour) - (time.Duration(1) * time.Nanosecond))

	return t
}
