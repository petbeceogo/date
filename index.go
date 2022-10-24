package date

import (
	"github.com/petbeceogo/date/dateloc"
)

func Init(loc *string) error {
	return dateloc.Init(loc)
}
