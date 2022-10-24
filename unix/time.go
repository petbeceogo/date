package unix

import "time"

func MilliToTime(unix int) time.Time {
	return time.UnixMilli(int64(unix))
}
