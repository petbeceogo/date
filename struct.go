package date

import "time"

type (
	Date struct {
		Year  int
		Month int
		Day   int
	}
)

func (d *Date) StartTime() time.Time {
	return NewStart(d.Year, d.Month, d.Day)
}

func (d *Date) EndTime() time.Time {
	return NewEnd(d.Year, d.Month, d.Day)
}

func (d *Date) Same(other Date) bool {
	return d.Year == other.Year && d.Month == other.Month && d.Day == other.Day
}
