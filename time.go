package interval

import "time"

var _ Ordered[Time] = Time{}

// Time is a wrapper of time.Time.
// It implements the Ordered interface.
type Time time.Time

// Equal checks if t is equal to t2.
func (t Time) Equal(t2 Time) bool {
	return time.Time(t).Equal(time.Time(t2))
}

// LessThan checks if t is less than t2.
func (t Time) LessThan(t2 Time) bool {
	return time.Time(t).Before(time.Time(t2))
}
