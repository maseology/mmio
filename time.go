package mmio

import "time"

const unixToInternal int64 = 62135596800 // (1969*365 + 1969/4 - 1969/100 + 1969/400) * 24 * 60 * 60 // number of seconds between Year 1 and 1970 (62135596800 seconds)

var mdays = map[time.Month]int{
	time.January:   31,
	time.February:  28,
	time.March:     31,
	time.April:     30,
	time.May:       31,
	time.June:      30,
	time.July:      31,
	time.August:    31,
	time.September: 30,
	time.October:   31,
	time.November:  30,
	time.December:  31,
}

// Yr year
type Yr int

// Mo month
type Mo = time.Month

// MinMaxTime returns the limits of time.Time (see https://stackoverflow.com/questions/25065055/what-is-the-maximum-time-time-in-go/32620397#32620397)
func MinMaxTime() (_, _ time.Time) {
	return time.Unix(1<<63-1, 999999999), time.Unix(1<<63-1-unixToInternal, 999999999)
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// MonthDays returns the number of days in a given month
func MonthDays(year Yr, month Mo) int {
	if month == time.February && isLeap(int(year)) {
		return 29
	}
	return mdays[month]
}

// MMdate returns a string date formatted yymmdd
func MMdate(d time.Time) string {
	return d.Format("060102")
}

// MMtime returns a string date formatted yymmdd_hhmmss
func MMtime(d time.Time) string {
	return d.Format("060102-150405")
}
