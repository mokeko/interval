package interval

type Interval[T Ordered[T]] struct {
	Lower Endpoint[T]
	Upper Endpoint[T]
}

func NewInterval[T Ordered[T]](lower, upper Endpoint[T]) Interval[T] {
	return Interval[T]{
		Lower: lower,
		Upper: upper,
	}
}

func (i Interval[T]) IsEmpty() bool {
	if !i.Lower.Bounded || !i.Upper.Bounded {
		return false
	}
	if i.Lower.Value.LessThan(i.Upper.Value) {
		return false
	}
	if i.Lower.Value.Equal(i.Upper.Value) {
		return !(i.Lower.Closed && i.Upper.Closed)
	}
	// i.Lower.Value > i.Upper.Value
	return true
}

func (i Interval[T]) IsEntire() bool {
	return !i.Lower.Bounded && !i.Upper.Bounded
}

func (i Interval[T]) Overlap(i2 Interval[T]) bool {
	// empty interval never overlaps
	if i.IsEmpty() || i2.IsEmpty() {
		return false
	}
	// entire interval overlaps with any non-empty interval
	if i.IsEntire() || i2.IsEntire() {
		return true
	}
	// At most 2 of the 4 endpoints are unbounded
	// If both intervals are unbounded in same direction, they overlap
	if (!i.Lower.Bounded && !i2.Lower.Bounded) || (!i.Upper.Bounded && !i2.Upper.Bounded) {
		return true
	}
	if !i.Lower.Bounded || !i2.Upper.Bounded {
		// i.Upper and i2.Lower are bounded
		return i2.Lower.Value.LessThan(i.Upper.Value) || (i2.Lower.Value.Equal(i.Upper.Value) && i2.Lower.Closed && i.Upper.Closed)
	}
	if !i.Upper.Bounded || !i2.Lower.Bounded {
		// i.Lower and i2.Upper are bounded
		return i.Lower.Value.LessThan(i2.Upper.Value) || (i.Lower.Value.Equal(i2.Upper.Value) && i.Lower.Closed && i2.Upper.Closed)
	}
	// both intervals are bounded in both directions
	if i.Lower.Value.Equal(i2.Upper.Value) {
		return i.Lower.Closed && i2.Upper.Closed
	}
	if i.Upper.Value.Equal(i2.Lower.Value) {
		return i.Upper.Closed && i2.Lower.Closed
	}
	return i.Lower.Value.LessThan(i2.Upper.Value) && i2.Lower.Value.LessThan(i.Upper.Value)
}
