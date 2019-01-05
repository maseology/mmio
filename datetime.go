package mmio

import (
	"fmt"
	. "time"
)

// VBdotNetToUNIX converts the VB.NET binary datetime format,
// which is the number of seconds since 0001-01-01 00:00:00
// to UNIX format, which is the number of nanoseconds since
// 1970-01-01 00:00:00
func VBdotNetToUNIX(vb int64) Time {
	t := Unix(vb/10000000-62135596800, 0)
	fmt.Println(t.Location)
	return t
}
