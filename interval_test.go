package interval

import "testing"

func testNewInterval[T Ordered[T]](t *testing.T, v1, v2 T) {
	assertEqual(t, Interval[T]{
		Lower: OpenEp(v1),
		Upper: OpenEp(v2),
	}, New(OpenEp(v1), OpenEp(v2)))
}

// expect v1 < v2
func testIsEmpty[T Ordered[T]](t *testing.T, v1, v2 T) {
	unbounded := UnboundedEp[T]()
	cases := []struct {
		name     string
		interval Interval[T]
		want     bool
	}{
		{
			name:     "unbounded",
			interval: New(unbounded, unbounded),
			want:     false,
		},
		{
			name:     "lower unbounded",
			interval: New(unbounded, OpenEp(v1)),
			want:     false,
		},
		{
			name:     "upper unbounded",
			interval: New(OpenEp(v1), unbounded),
			want:     false,
		},
		{
			name:     "lower < upper",
			interval: New(OpenEp(v1), OpenEp(v2)),
			want:     false,
		},
		{
			name:     "lowe = upper, closed",
			interval: New(ClosedEp(v1), ClosedEp(v1)),
			want:     false,
		},
		{
			name:     "lower = upper, lower open",
			interval: New(OpenEp(v1), ClosedEp(v1)),
			want:     true,
		},
		{
			name:     "lower = upper, upper open",
			interval: New(ClosedEp(v1), OpenEp(v1)),
			want:     true,
		},
		{
			name:     "lower = upper, open",
			interval: New(OpenEp(v1), OpenEp(v1)),
			want:     true,
		},
		{
			name:     "lower > upper",
			interval: New(OpenEp(v2), OpenEp(v1)),
			want:     true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.interval.IsEmpty())
		})
	}
}

func testIsEntire[T Ordered[T]](t *testing.T, v T) {
	unbounded := UnboundedEp[T]()
	cases := []struct {
		name     string
		interval Interval[T]
		want     bool
	}{
		{
			name:     "unbounded",
			interval: New(unbounded, unbounded),
			want:     true,
		},
		{
			name:     "lower unbounded",
			interval: New(unbounded, OpenEp(v)),
			want:     false,
		},
		{
			name:     "upper unbounded",
			interval: New(OpenEp(v), unbounded),
			want:     false,
		},
		{
			name:     "bounded",
			interval: New(ClosedEp(v), ClosedEp(v)),
			want:     false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.interval.IsEntire())
		})
	}
}

// expect v1 < v2
func testContainsPoint[T Ordered[T]](t *testing.T, v1, v2, v3 T) {
	unbounded := UnboundedEp[T]()
	cases := []struct {
		name     string
		interval Interval[T]
		point    T
		want     bool
	}{
		{
			name:     "empty",
			interval: New(OpenEp(v3), OpenEp(v1)),
			point:    v2,
			want:     false,
		},
		{
			name:     "unbounded",
			interval: New(unbounded, unbounded),
			point:    v1,
			want:     true,
		},
		{
			name:     "lower unbounded, point < upper",
			interval: New(unbounded, OpenEp(v2)),
			point:    v1,
			want:     true,
		},
		{
			name:     "lower unbounded, point = upper, closed",
			interval: New(unbounded, ClosedEp(v1)),
			point:    v1,
			want:     true,
		},
		{
			name:     "lower unbounded, point = upper, open",
			interval: New(unbounded, OpenEp(v1)),
			point:    v1,
			want:     false,
		},
		{
			name:     "lower unbounded, point > upper",
			interval: New(unbounded, OpenEp(v1)),
			point:    v2,
			want:     false,
		},
		{
			name:     "upper unbounded, point < lower",
			interval: New(OpenEp(v2), unbounded),
			point:    v1,
			want:     false,
		},
		{
			name:     "upper unbounded, point = lower, open",
			interval: New(OpenEp(v1), unbounded),
			point:    v1,
			want:     false,
		},
		{
			name:     "upper unbounded, point = lower, closed",
			interval: New(ClosedEp(v1), unbounded),
			point:    v1,
			want:     true,
		},
		{
			name:     "upper unbounded, point > lower",
			interval: New(OpenEp(v1), unbounded),
			point:    v2,
			want:     true,
		},
		{
			name:     "bounded, lower = point = upper",
			interval: New(ClosedEp(v1), ClosedEp(v1)),
			point:    v1,
			want:     true,
		},
		{
			name:     "bounded, point < lower",
			interval: New(OpenEp(v2), OpenEp(v3)),
			point:    v1,
			want:     false,
		},
		{
			name:     "bounded, point = lower, open",
			interval: New(OpenEp(v1), OpenEp(v2)),
			point:    v1,
			want:     false,
		},
		{
			name:     "bounded, point = lower, closed",
			interval: New(ClosedEp(v1), OpenEp(v2)),
			point:    v1,
			want:     true,
		},
		{
			name:     "bounded, lower < point < upper",
			interval: New(OpenEp(v1), OpenEp(v3)),
			point:    v2,
			want:     true,
		},
		{
			name:     "bounded, point = upper, open",
			interval: New(OpenEp(v1), OpenEp(v2)),
			point:    v2,
			want:     false,
		},
		{
			name:     "bounded, point = upper, closed",
			interval: New(OpenEp(v1), ClosedEp(v2)),
			point:    v2,
			want:     true,
		},
		{
			name:     "bounded, point > upper",
			interval: New(OpenEp(v1), OpenEp(v2)),
			point:    v3,
			want:     false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.interval.ContainsPoint(c.point))
		})
	}
}

// expect v1 < v2 < v3 < v4
func testCompareInterval[T Ordered[T]](t *testing.T, v1, v2, v3, v4 T) {
	unbounded := UnboundedEp[T]()
	cases := []struct {
		name     string
		i        Interval[T]
		i2       Interval[T]
		before   bool
		after    bool
		overlaps bool
	}{
		{
			name:     "both are empty",
			i:        New(OpenEp(v1), OpenEp(v1)),
			i2:       New(OpenEp(v1), OpenEp(v1)),
			before:   false,
			after:    false,
			overlaps: false,
		},
		{
			name:     "i is empty",
			i:        New(OpenEp(v1), OpenEp(v1)),
			i2:       New(unbounded, unbounded),
			before:   false,
			after:    false,
			overlaps: false,
		},
		{
			name:     "i2 is empty",
			i:        New(unbounded, unbounded),
			i2:       New(OpenEp(v1), OpenEp(v1)),
			before:   false,
			after:    false,
			overlaps: false,
		},
		// in the following cases, the intervals are not empty
		{
			name:     "both are unbounded",
			i:        New(unbounded, unbounded),
			i2:       New(unbounded, unbounded),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i is unbounded",
			i:        New(unbounded, unbounded),
			i2:       New(OpenEp(v1), OpenEp(v2)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2 is unbounded",
			i:        New(OpenEp(v1), OpenEp(v2)),
			i2:       New(unbounded, unbounded),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "both are lower unbounded",
			i:        New(unbounded, OpenEp(v1)),
			i2:       New(unbounded, OpenEp(v2)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "both are upper unbounded",
			i:        New(OpenEp(v1), unbounded),
			i2:       New(OpenEp(v2), unbounded),
			before:   false,
			after:    false,
			overlaps: true,
		},
		// i.Lower is unbounded
		{
			name:     "i.Lower is unbounded, i.Upper < i2.Lower",
			i:        New(unbounded, OpenEp(v1)),
			i2:       New(OpenEp(v2), OpenEp(v3)),
			before:   true,
			after:    false,
			overlaps: false,
		},
		{
			name:     "i.Lower is unbounded, i.Upper = i2.Lower, no contact",
			i:        New(unbounded, OpenEp(v1)),
			i2:       New(OpenEp(v1), OpenEp(v2)),
			before:   true,
			after:    false,
			overlaps: false,
		},
		{
			name:     "i.Lower is unbounded, i.Upper = i2.Lower, contact",
			i:        New(unbounded, ClosedEp(v1)),
			i2:       New(ClosedEp(v1), OpenEp(v2)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Lower is unbounded, i.Upper > i2.Lower",
			i:        New(unbounded, OpenEp(v2)),
			i2:       New(OpenEp(v1), OpenEp(v3)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Lower is unbounded, i.Upper > i2.Upper",
			i:        New(unbounded, OpenEp(v3)),
			i2:       New(OpenEp(v1), OpenEp(v2)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		// i.Upper is unbounded
		{
			name:     "i.Upper is unbounded, i.Lower > i2.Upper",
			i:        New(OpenEp(v3), unbounded),
			i2:       New(OpenEp(v1), OpenEp(v2)),
			before:   false,
			after:    true,
			overlaps: false,
		},
		{
			name:     "i.Upper is unbounded, i.Lower = i2.Upper, no contact",
			i:        New(OpenEp(v2), unbounded),
			i2:       New(OpenEp(v1), OpenEp(v2)),
			before:   false,
			after:    true,
			overlaps: false,
		},
		{
			name:     "i.Upper is unbounded, i.Lower = i2.Upper, contact",
			i:        New(ClosedEp(v2), unbounded),
			i2:       New(OpenEp(v1), ClosedEp(v2)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Upper is unbounded, i.Lower< i2.Upper",
			i:        New(OpenEp(v2), unbounded),
			i2:       New(OpenEp(v1), OpenEp(v3)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Upper is unbounded, i.Lower < i2.Lower",
			i:        New(OpenEp(v1), unbounded),
			i2:       New(OpenEp(v2), OpenEp(v3)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		// i2.Lower is unbounded
		{
			name:     "i2.Lower is unbounded, i2.Upper < i.Lower",
			i:        New(OpenEp(v2), OpenEp(v3)),
			i2:       New(unbounded, OpenEp(v1)),
			before:   false,
			after:    true,
			overlaps: false,
		},
		{
			name:     "i2.Lower is unbounded, i2.Upper = i.Lower, no contact",
			i:        New(OpenEp(v1), OpenEp(v2)),
			i2:       New(unbounded, OpenEp(v1)),
			before:   false,
			after:    true,
			overlaps: false,
		},
		{
			name:     "i2.Lower is unbounded, i2.Upper = i.Lower, contact",
			i:        New(ClosedEp(v1), OpenEp(v2)),
			i2:       New(unbounded, ClosedEp(v1)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Lower is unbounded, i2.Upper > i.Lower",
			i:        New(OpenEp(v1), OpenEp(v3)),
			i2:       New(unbounded, OpenEp(v2)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Lower is unbounded, i2.Upper > i.Upper",
			i:        New(OpenEp(v1), OpenEp(v2)),
			i2:       New(unbounded, OpenEp(v3)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		// i2.Upper is unbounded
		{
			name:     "i2.Upper is unbounded, i2.Lower > i.Lower",
			i:        New(OpenEp(v1), OpenEp(v2)),
			i2:       New(OpenEp(v3), unbounded),
			before:   true,
			after:    false,
			overlaps: false,
		},
		{
			name:     "i2.Upper is unbounded, i2.Lower = i.Upper, no contact",
			i:        New(OpenEp(v1), OpenEp(v2)),
			i2:       New(OpenEp(v2), unbounded),
			before:   true,
			after:    false,
			overlaps: false,
		},
		{
			name:     "i2.Upper is unbounded, i2.Lower = i.Upper, contact",
			i:        New(OpenEp(v1), ClosedEp(v2)),
			i2:       New(ClosedEp(v2), unbounded),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Upper is unbounded, i2.Lower < i.Upper",
			i:        New(OpenEp(v1), OpenEp(v3)),
			i2:       New(OpenEp(v2), unbounded),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Upper is unbounded, i2.Lower < i.Lower",
			i:        New(OpenEp(v2), OpenEp(v3)),
			i2:       New(OpenEp(v1), unbounded),
			before:   false,
			after:    false,
			overlaps: true,
		},
		// in the following cases, both intervals are bounded
		{
			name:     "i.Lower = i.Upper = i2.Lower = i2.Upper",
			i:        New(ClosedEp(v1), ClosedEp(v1)),
			i2:       New(ClosedEp(v1), ClosedEp(v1)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Lower < i.Upper < i2.Lower < i2.Upper",
			i:        New(OpenEp(v1), OpenEp(v2)),
			i2:       New(OpenEp(v3), OpenEp(v4)),
			before:   true,
			after:    false,
			overlaps: false,
		},
		{
			name:     "i.Lower < i.Upper = i2.Lower < i2.Upper, no contact",
			i:        New(OpenEp(v1), OpenEp(v2)),
			i2:       New(OpenEp(v2), OpenEp(v3)),
			before:   true,
			after:    false,
			overlaps: false,
		},
		{
			name:     "i.Lower < i.Upper = i2.Lower < i2.Upper, contact",
			i:        New(OpenEp(v1), ClosedEp(v2)),
			i2:       New(ClosedEp(v2), OpenEp(v3)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Lower < i2.Lower < i.Upper < i2.Upper",
			i:        New(OpenEp(v1), OpenEp(v3)),
			i2:       New(OpenEp(v2), OpenEp(v4)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Lower < i2.Lower < i.Upper = i2.Upper",
			i:        New(OpenEp(v1), OpenEp(v3)),
			i2:       New(OpenEp(v2), OpenEp(v3)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Lower < i2.Lower < i2.Upper < i.Upper",
			i:        New(OpenEp(v1), OpenEp(v4)),
			i2:       New(OpenEp(v2), OpenEp(v3)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i.Lower = i2.Lower",
			i:        New(OpenEp(v1), OpenEp(v2)),
			i2:       New(OpenEp(v1), OpenEp(v4)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Lower < i.Lower < i.Upper < i2.Upper",
			i:        New(OpenEp(v2), OpenEp(v3)),
			i2:       New(OpenEp(v1), OpenEp(v4)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Lower < i.Lower < i2.Upper < i.Upper",
			i:        New(OpenEp(v2), OpenEp(v4)),
			i2:       New(OpenEp(v1), OpenEp(v3)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Lower < i.Lower = i2.Upper = i.Upper, no contact",
			i:        New(ClosedEp(v2), ClosedEp(v2)),
			i2:       New(OpenEp(v1), OpenEp(v2)),
			before:   false,
			after:    true,
			overlaps: false,
		},
		{
			name:     "i2.Lower < i.Lower = i2.Upper = i.Upper, contact",
			i:        New(ClosedEp(v2), ClosedEp(v2)),
			i2:       New(OpenEp(v1), ClosedEp(v2)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Lower < i.Lower = i2.Upper < i.Upper, no contact",
			i:        New(OpenEp(v2), OpenEp(v3)),
			i2:       New(OpenEp(v1), OpenEp(v2)),
			before:   false,
			after:    true,
			overlaps: false,
		},
		{
			name:     "i2.Lower < i.Lower = i2.Upper < i.Upper, contact",
			i:        New(ClosedEp(v2), OpenEp(v3)),
			i2:       New(OpenEp(v1), ClosedEp(v2)),
			before:   false,
			after:    false,
			overlaps: true,
		},
		{
			name:     "i2.Lower < i2.Upper < i.Lower < i.Upper",
			i:        New(OpenEp(v3), OpenEp(v4)),
			i2:       New(OpenEp(v1), OpenEp(v2)),
			before:   false,
			after:    true,
			overlaps: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Run("overlaps", func(t *testing.T) {
				assertEqual(t, c.overlaps, c.i.Overlap(c.i2))
			})
			t.Run("before", func(t *testing.T) {
				assertEqual(t, c.before, c.i.Before(c.i2))
			})
			t.Run("after", func(t *testing.T) {
				assertEqual(t, c.after, c.i.After(c.i2))
			})
		})
	}
}
