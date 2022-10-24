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
