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

// Reset start a timer
func (t *Timer) Reset() {
	t.t = time.Now()
}

// Print reports current time
func (t *Timer) Print(msg string) {
	if len(msg) == 0 {
		fmt.Println(time.Now().Sub(t.t))
	} else {
		fmt.Printf(" %s - %v\n", msg, time.Now().Sub(t.t))
	}
}

// Lap reports lap time (resets timer)
func (t *Timer) Lap(msg string) {
	t.Print(msg)
	t.Reset()
}
