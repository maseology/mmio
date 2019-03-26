package mmio

import "time"

const unixToInternal int64 = 62135596800 // (1969*365 + 1969/4 - 1969/100 + 1969/400) * 24 * 60 * 60 // number of seconds between Year 1 and 1970 (62135596800 seconds)

// MinMaxTime returns the limits of time.Time (see https://stackoverflow.com/questions/25065055/what-is-the-maximum-time-time-in-go/32620397#32620397)
func MinMaxTime() (_, _ time.Time) {
	return time.Unix(1<<63-1, 999999999), time.Unix(1<<63-1-unixToInternal, 999999999)
}
