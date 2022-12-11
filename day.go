package date

func DayRange(
	year int,
	month int,
) (firstDay int, latsDay int) {
	firstDayTime := NewStart(year, month, 1)
	lastDayTime := firstDayTime.AddDate(0, 1, -1)

	return 1, lastDayTime.Day()
}
