package interval

// Interval represents a interval consisting of two endpoints of type T.
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

// True if no points are contained in interval.
func (i Interval[T]) IsEmpty() bool {
	if !i.Lower.Bounded || !i.Upper.Bounded {
		return false
	}
	if i.Lower.Value.LessThan(i.Upper.Value) {
		return false
	}
	return !i.Lower.equalAndBothClosed(i.Upper)
}

// True if both endpoints are unbounded.
func (i Interval[T]) IsEntire() bool {
	return !i.Lower.Bounded && !i.Upper.Bounded
}

// True if interval contains given point.
func (i Interval[T]) ContainsPoint(p T) bool {
	if i.IsEmpty() {
		return false
	}
	if i.Lower.Bounded && (p.LessThan(i.Lower.Value) || (p.Equal(i.Lower.Value) && !i.Lower.Closed)) {
		return false
	}
	if i.Upper.Bounded && (i.Upper.Value.LessThan(p) || (p.Equal(i.Upper.Value) && !i.Upper.Closed)) {
		return false
	}
	return true
}

// True if receiver interval ends before argument interval starts.
func (i Interval[T]) Before(i2 Interval[T]) bool {
	if i.IsEmpty() || i2.IsEmpty() {
		return false
	}
	if i.Upper.Bounded && i2.Lower.Bounded {
		return i.Upper.Value.LessThan(i2.Lower.Value) || (i.Upper.Value.Equal(i2.Lower.Value) && (!i.Upper.Closed || !i2.Lower.Closed))
	}
	return false
}

// True if receiver interval starts after argument interval ends.
func (i Interval[T]) After(i2 Interval[T]) bool {
	if i.IsEmpty() || i2.IsEmpty() {
		return false
	}
	if i2.Upper.Bounded && i.Lower.Bounded {
		return i2.Upper.Value.LessThan(i.Lower.Value) || (i2.Upper.Value.Equal(i.Lower.Value) && (!i2.Upper.Closed || !i.Lower.Closed))
	}
	return false
}

// True if two interval share at least one point.
func (i Interval[T]) Overlaps(i2 Interval[T]) bool {
	// empty interval never overlaps
	if i.IsEmpty() || i2.IsEmpty() {
		return false
	}
	return !i.Before(i2) && !i.After(i2)
}
