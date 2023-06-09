// Package interval provides generic interval types and operations.
package interval

// Interval represents a interval consisting of two endpoints.
// The zero value of Interval is an empty interval.
type Interval[T Ordered[T]] struct {
	Lower Endpoint[T]
	Upper Endpoint[T]
}

// New returns an interval with given endpoints.
func New[T Ordered[T]](lower, upper Endpoint[T]) Interval[T] {
	return Interval[T]{
		Lower: lower,
		Upper: upper,
	}
}

// IsEmpty returns true if no points are contained in interval.
func (i Interval[T]) IsEmpty() bool {
	if i.Lower.Unbounded || i.Upper.Unbounded {
		return false
	}
	if i.Lower.Value.LessThan(i.Upper.Value) {
		return false
	}
	return !i.Lower.equalAndBothClosed(i.Upper)
}

// IsEntire returns true if both endpoints are unbounded.
func (i Interval[T]) IsEntire() bool {
	return i.Lower.Unbounded && i.Upper.Unbounded
}

// Contains returns true if interval contains the point with given value.
func (i Interval[T]) Contains(p T) bool {
	if i.IsEmpty() {
		return false
	}
	if i.Lower.Bounded() && (p.LessThan(i.Lower.Value) || (p.Equal(i.Lower.Value) && !i.Lower.Closed)) {
		return false
	}
	if i.Upper.Bounded() && (i.Upper.Value.LessThan(p) || (p.Equal(i.Upper.Value) && !i.Upper.Closed)) {
		return false
	}
	return true
}

// Before returns true if interval ends before other interval starts.
func (i Interval[T]) Before(i2 Interval[T]) bool {
	if i.IsEmpty() || i2.IsEmpty() {
		return false
	}
	if i.Upper.Bounded() && i2.Lower.Bounded() {
		return i.Upper.Value.LessThan(i2.Lower.Value) || (i.Upper.Value.Equal(i2.Lower.Value) && (!i.Upper.Closed || !i2.Lower.Closed))
	}
	return false
}

// After returns true if interval starts after other interval ends.
func (i Interval[T]) After(i2 Interval[T]) bool {
	if i.IsEmpty() || i2.IsEmpty() {
		return false
	}
	if i2.Upper.Bounded() && i.Lower.Bounded() {
		return i2.Upper.Value.LessThan(i.Lower.Value) || (i2.Upper.Value.Equal(i.Lower.Value) && (!i2.Upper.Closed || !i.Lower.Closed))
	}
	return false
}

// Overlap returns true if interval shares at least one point with other interval.
func (i Interval[T]) Overlaps(i2 Interval[T]) bool {
	// empty interval never overlaps
	if i.IsEmpty() || i2.IsEmpty() {
		return false
	}
	return !i.Before(i2) && !i.After(i2)
}
