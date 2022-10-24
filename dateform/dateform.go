package dateform

import (
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
