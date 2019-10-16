package gigasecond

import (
	"fmt"
	"time"
)

//This function take a date of type Time
//and add 1 Gigasecond
func AddGigasecond(t time.Time) time.Time {
	duration, _ := time.ParseDuration(fmt.Sprintf("%1.fs", 1e9))
	return t.Add(duration)
}
