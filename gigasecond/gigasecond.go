// Package gigasecond implements simple time arithmetic by adding
// gigaseconds which 1e9 to given time stamp.
package gigasecond

// import path for the time package from the standard library
import "time"

// constant to maintain gigasecond value.
const gigaSecond = 1e9

// AddGigasecond takes time stamp as an argument and returns new time stamp
// by adding gigasecond to its value.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Duration(gigaSecond) * time.Second)
	return t
}
