package mmio

import (
	"fmt"
	"time"
)

// Timer is a common timer used for profiling
type Timer struct{ t time.Time }

// NewTimer constructs a Timer
func NewTimer() Timer {
	var t Timer
	t.t = time.Now()
	return t
}

// TimerReset start a timer
func (t *Timer) TimerReset() {
	t.t = time.Now()
}

// TimerPrint reports lap time
func (t *Timer) TimerPrint(msg string) {
	if len(msg) == 0 {
		fmt.Println(time.Now().Sub(t.t))
	} else {
		fmt.Printf(" %s - %v\n", msg, time.Now().Sub(t.t))
	}
}
