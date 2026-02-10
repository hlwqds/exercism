package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	switch wSched {
	case Teenth:
		t := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
		teenthStart := t.Weekday()
		return 13 + int(wDay+7-teenthStart)%7
	case Last:
		t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
		monthEnd := t.Weekday()
		return t.Day() - int(monthEnd+7-wDay)%7
	default:
		t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		monthStart := t.Weekday()
		return int(wDay+7-monthStart)%7 + int(wSched)*7 + 1
	}
}
