package interval

import "time"

var _ Ordered[Time] = Time{}

// Time is a wrapper of time.Time.
// It implements the Ordered interface.
type Time time.Time

func (t Time) Equal(t2 Time) bool {
	return time.Time(t).Equal(time.Time(t2))
}

func (t Time) LessThan(t2 Time) bool {
	return time.Time(t).Before(time.Time(t2))
}
