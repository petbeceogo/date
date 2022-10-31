package dateform

import (
	"math"
	"strconv"
	"time"
)

// TODO: regex 유사한 패턴 기반 동작으로 변경
func Date(t time.Time) string {
	yStamp := strconv.Itoa(t.Year())
	mStamp := toDigit(strconv.Itoa(int(t.Month())), 2)
	dStamp := toDigit(strconv.Itoa(t.Day()), 2)

	return yStamp + "." + mStamp + "." + dStamp
}

func KRMonthDate(t time.Time) string {
	mStamp := strconv.Itoa(int(t.Month()))
	dStamp := strconv.Itoa(t.Day())

	return mStamp + "월 " + dStamp + "일"
}

func KRTime(t time.Time) string {
	h := t.Hour()
	var hStamp string
	if h > 12 {
		hStamp = "오후 " + strconv.Itoa(h-12) + "시"
	} else {
		hStamp = "오전 " + strconv.Itoa(h) + "시"
	}

	m := t.Minute()
	var mStamp string
	if m > 0 {
		mStamp = " " + strconv.Itoa(m) + "분"
	} else {
		mStamp = ""
	}

	return hStamp + mStamp
}

func TimeHourMin(t time.Time) string {
	stamp := toDigit(strconv.Itoa(t.Hour()), 2) +
		":" +
		toDigit(strconv.Itoa(t.Minute()), 2)

	return stamp
}

func DurationHourMin(d time.Duration) string {
	hour := d.Hours()
	h := int(math.Floor(hour))
	m := int(math.Floor(d.Minutes())) - h*60
	stamp := toDigit(strconv.Itoa(h), 2) +
		"h " +
		toDigit(strconv.Itoa(m), 2) +
		"m"

	return stamp
}

func DurationMinSecond(d time.Duration) string {
	minutes := d.Minutes()
	m := int(math.Floor(minutes))
	s := int(math.Floor(d.Seconds())) - m*60
	stamp := toDigit(strconv.Itoa(m), 2) + "m " + toDigit(strconv.Itoa(s), 2) + "s"

	return stamp
}
