package clock

import "fmt"

const (
	maxMinutes = 60
	maxHour    = 24
)

// Clock implements simple clock with hours and minutes.
type Clock struct {
	h int
	m int
}

// New returns new clock instance.
func New(hour, minute int) Clock {
	m := minute % maxMinutes
	hour += (minute / maxMinutes)

	// On rolling negative minutes adjust to it's previous hour
	if m < 0 {
		hour--
		m += 60
	}

	h := hour % maxHour
	if h < 0 {
		h = maxHour + h
	}
	return Clock{h, m}
}

// String returns string format of Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add method adds time to clock.
func (c Clock) Add(a int) Clock {
	temp := New(c.h, c.m+a)
	c.h = temp.h
	c.m = temp.m
	return c
}
