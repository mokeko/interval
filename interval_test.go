package interval

import "testing"

func testNewInterval[T Ordered[T]](t *testing.T, v1, v2 T) {
	assertEqual(t, Interval[T]{
		Lower: NewOpen(v1),
		Upper: NewOpen(v2),
	}, NewInterval(NewOpen(v1), NewOpen(v2)))
}

// expect v1 < v2
func testIsEmpty[T Ordered[T]](t *testing.T, v1, v2 T) {
	unbounded := NewUnbounded[T]()
	cases := []struct {
		name     string
		interval Interval[T]
		want     bool
	}{
		{
			name:     "unbounded",
			interval: NewInterval(unbounded, unbounded),
			want:     false,
		},
		{
			name:     "lower unbounded",
			interval: NewInterval(unbounded, NewOpen(v1)),
			want:     false,
		},
		{
			name:     "upper unbounded",
			interval: NewInterval(NewOpen(v1), unbounded),
			want:     false,
		},
		{
			name:     "lower < upper",
			interval: NewInterval(NewOpen(v1), NewOpen(v2)),
			want:     false,
		},
		{
			name:     "lowe = upper, closed",
			interval: NewInterval(NewClosed(v1), NewClosed(v1)),
			want:     false,
		},
		{
			name:     "lower = upper, lower open",
			interval: NewInterval(NewOpen(v1), NewClosed(v1)),
			want:     true,
		},
		{
			name:     "lower = upper, upper open",
			interval: NewInterval(NewClosed(v1), NewOpen(v1)),
			want:     true,
		},
		{
			name:     "lower = upper, open",
			interval: NewInterval(NewOpen(v1), NewOpen(v1)),
			want:     true,
		},
		{
			name:     "lower > upper",
			interval: NewInterval(NewOpen(v2), NewOpen(v1)),
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
	unbounded := NewUnbounded[T]()
	cases := []struct {
		name     string
		interval Interval[T]
		want     bool
	}{
		{
			name:     "unbounded",
			interval: NewInterval(unbounded, unbounded),
			want:     true,
		},
		{
			name:     "lower unbounded",
			interval: NewInterval(unbounded, NewOpen(v)),
			want:     false,
		},
		{
			name:     "upper unbounded",
			interval: NewInterval(NewOpen(v), unbounded),
			want:     false,
		},
		{
			name:     "bounded",
			interval: NewInterval(NewClosed(v), NewClosed(v)),
			want:     false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.interval.IsEntire())
		})
	}
}

// expect v1 < v2 < v3 < v4
func testOverlap[T Ordered[T]](t *testing.T, v1, v2, v3, v4 T) {
	unbounded := NewUnbounded[T]()
	cases := []struct {
		name string
		i    Interval[T]
		i2   Interval[T]
		want bool
	}{
		{
			name: "both are empty",
			i:    NewInterval(NewOpen(v1), NewOpen(v1)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v1)),
			want: false,
		},
		{
			name: "i is empty",
			i:    NewInterval(NewOpen(v1), NewOpen(v1)),
			i2:   NewInterval(unbounded, unbounded),
			want: false,
		},
		{
			name: "i2 is empty",
			i:    NewInterval(unbounded, unbounded),
			i2:   NewInterval(NewOpen(v1), NewOpen(v1)),
			want: false,
		},
		// in the following cases, the intervals are not empty
		{
			name: "both are unbounded",
			i:    NewInterval(unbounded, unbounded),
			i2:   NewInterval(unbounded, unbounded),
			want: true,
		},
		{
			name: "i is unbounded",
			i:    NewInterval(unbounded, unbounded),
			i2:   NewInterval(NewOpen(v1), NewOpen(v2)),
			want: true,
		},
		{
			name: "i2 is unbounded",
			i:    NewInterval(NewOpen(v1), NewOpen(v2)),
			i2:   NewInterval(unbounded, unbounded),
			want: true,
		},
		{
			name: "both are lower unbounded",
			i:    NewInterval(unbounded, NewOpen(v1)),
			i2:   NewInterval(unbounded, NewOpen(v2)),
			want: true,
		},
		{
			name: "both are upper unbounded",
			i:    NewInterval(NewOpen(v1), unbounded),
			i2:   NewInterval(NewOpen(v2), unbounded),
			want: true,
		},
		// i.Lower is unbounded
		{
			name: "i.Lower is unbounded, i.Upper < i2.Lower",
			i:    NewInterval(unbounded, NewOpen(v1)),
			i2:   NewInterval(NewOpen(v2), NewOpen(v3)),
			want: false,
		},
		{
			name: "i.Lower is unbounded, i.Upper = i2.Lower, no contact",
			i:    NewInterval(unbounded, NewOpen(v1)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v2)),
			want: false,
		},
		{
			name: "i.Lower is unbounded, i.Upper = i2.Lower, contact",
			i:    NewInterval(unbounded, NewClosed(v1)),
			i2:   NewInterval(NewClosed(v1), NewOpen(v2)),
			want: true,
		},
		{
			name: "i.Lower is unbounded, i.Upper > i2.Lower",
			i:    NewInterval(unbounded, NewOpen(v2)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v3)),
			want: true,
		},
		{
			name: "i.Lower is unbounded, i.Upper > i2.Upper",
			i:    NewInterval(unbounded, NewOpen(v3)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v2)),
			want: true,
		},
		// i.Upper is unbounded
		{
			name: "i.Upper is unbounded, i.Lower > i2.Upper",
			i:    NewInterval(NewOpen(v3), unbounded),
			i2:   NewInterval(NewOpen(v1), NewOpen(v2)),
			want: false,
		},
		{
			name: "i.Upper is unbounded, i.Lower = i2.Upper, no contact",
			i:    NewInterval(NewOpen(v2), unbounded),
			i2:   NewInterval(NewOpen(v1), NewOpen(v2)),
			want: false,
		},
		{
			name: "i.Upper is unbounded, i.Lower = i2.Upper, contact",
			i:    NewInterval(NewClosed(v2), unbounded),
			i2:   NewInterval(NewOpen(v1), NewClosed(v2)),
			want: true,
		},
		{
			name: "i.Upper is unbounded, i.Lower< i2.Upper",
			i:    NewInterval(NewOpen(v2), unbounded),
			i2:   NewInterval(NewOpen(v1), NewOpen(v3)),
			want: true,
		},
		{
			name: "i.Upper is unbounded, i.Lower < i2.Lower",
			i:    NewInterval(NewOpen(v1), unbounded),
			i2:   NewInterval(NewOpen(v2), NewOpen(v3)),
			want: true,
		},
		// i2.Lower is unbounded
		{
			name: "i2.Lower is unbounded, i2.Upper < i.Lower",
			i:    NewInterval(NewOpen(v2), NewOpen(v3)),
			i2:   NewInterval(unbounded, NewOpen(v1)),
			want: false,
		},
		{
			name: "i2.Lower is unbounded, i2.Upper = i.Lower, no contact",
			i:    NewInterval(NewOpen(v1), NewOpen(v2)),
			i2:   NewInterval(unbounded, NewOpen(v1)),
			want: false,
		},
		{
			name: "i2.Lower is unbounded, i2.Upper = i.Lower, contact",
			i:    NewInterval(NewClosed(v1), NewOpen(v2)),
			i2:   NewInterval(unbounded, NewClosed(v1)),
			want: true,
		},
		{
			name: "i2.Lower is unbounded, i2.Upper > i.Lower",
			i:    NewInterval(NewOpen(v1), NewOpen(v3)),
			i2:   NewInterval(unbounded, NewOpen(v2)),
			want: true,
		},
		{
			name: "i2.Lower is unbounded, i2.Upper > i.Upper",
			i:    NewInterval(NewOpen(v1), NewOpen(v2)),
			i2:   NewInterval(unbounded, NewOpen(v3)),
			want: true,
		},
		// i2.Upper is unbounded
		{
			name: "i2.Upper is unbounded, i2.Lower > i.Lower",
			i:    NewInterval(NewOpen(v1), NewOpen(v2)),
			i2:   NewInterval(NewOpen(v3), unbounded),
			want: false,
		},
		{
			name: "i2.Upper is unbounded, i2.Lower = i.Upper, no contact",
			i:    NewInterval(NewOpen(v1), NewOpen(v2)),
			i2:   NewInterval(NewOpen(v2), unbounded),
			want: false,
		},
		{
			name: "i2.Upper is unbounded, i2.Lower = i.Upper, contact",
			i:    NewInterval(NewOpen(v1), NewClosed(v2)),
			i2:   NewInterval(NewClosed(v2), unbounded),
			want: true,
		},
		{
			name: "i2.Upper is unbounded, i2.Lower < i.Upper",
			i:    NewInterval(NewOpen(v1), NewOpen(v3)),
			i2:   NewInterval(NewOpen(v2), unbounded),
			want: true,
		},
		{
			name: "i2.Upper is unbounded, i2.Lower < i.Lower",
			i:    NewInterval(NewOpen(v2), NewOpen(v3)),
			i2:   NewInterval(NewOpen(v1), unbounded),
			want: true,
		},
		// in the following cases, both intervals are bounded
		{
			name: "i.Lower = i.Upper = i2.Lower = i2.Upper",
			i:    NewInterval(NewClosed(v1), NewClosed(v1)),
			i2:   NewInterval(NewClosed(v1), NewClosed(v1)),
			want: true,
		},
		{
			name: "i.Lower < i.Upper < i2.Lower < i2.Upper",
			i:    NewInterval(NewOpen(v1), NewOpen(v2)),
			i2:   NewInterval(NewOpen(v3), NewOpen(v4)),
			want: false,
		},
		{
			name: "i.Lower < i.Upper = i2.Lower < i2.Upper, no contact",
			i:    NewInterval(NewOpen(v1), NewOpen(v2)),
			i2:   NewInterval(NewOpen(v2), NewOpen(v3)),
			want: false,
		},
		{
			name: "i.Lower < i.Upper = i2.Lower < i2.Upper, contact",
			i:    NewInterval(NewOpen(v1), NewClosed(v2)),
			i2:   NewInterval(NewClosed(v2), NewOpen(v3)),
			want: true,
		},
		{
			name: "i.Lower < i2.Lower < i.Upper < i2.Upper",
			i:    NewInterval(NewOpen(v1), NewOpen(v3)),
			i2:   NewInterval(NewOpen(v2), NewOpen(v4)),
			want: true,
		},
		{
			name: "i.Lower < i2.Lower < i.Upper = i2.Upper",
			i:    NewInterval(NewOpen(v1), NewOpen(v3)),
			i2:   NewInterval(NewOpen(v2), NewOpen(v3)),
			want: true,
		},
		{
			name: "i.Lower < i2.Lower < i2.Upper < i.Upper",
			i:    NewInterval(NewOpen(v1), NewOpen(v4)),
			i2:   NewInterval(NewOpen(v2), NewOpen(v3)),
			want: true,
		},
		{
			name: "i.Lower = i2.Lower",
			i:    NewInterval(NewOpen(v1), NewOpen(v2)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v4)),
			want: true,
		},
		{
			name: "i2.Lower < i.Lower < i.Upper < i2.Upper",
			i:    NewInterval(NewOpen(v2), NewOpen(v3)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v4)),
			want: true,
		},
		{
			name: "i2.Lower < i.Lower < i2.Upper < i.Upper",
			i:    NewInterval(NewOpen(v2), NewOpen(v4)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v3)),
			want: true,
		},
		{
			name: "i2.Lower < i.Lower = i2.Upper = i.Upper, no contact",
			i:    NewInterval(NewClosed(v2), NewClosed(v2)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v2)),
			want: false,
		},
		{
			name: "i2.Lower < i.Lower = i2.Upper = i.Upper, contact",
			i:    NewInterval(NewClosed(v2), NewClosed(v2)),
			i2:   NewInterval(NewOpen(v1), NewClosed(v2)),
			want: true,
		},
		{
			name: "i2.Lower < i.Lower = i2.Upper < i.Upper, no contact",
			i:    NewInterval(NewOpen(v2), NewOpen(v3)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v2)),
			want: false,
		},
		{
			name: "i2.Lower < i.Lower = i2.Upper < i.Upper, contact",
			i:    NewInterval(NewClosed(v2), NewOpen(v3)),
			i2:   NewInterval(NewOpen(v1), NewClosed(v2)),
			want: true,
		},
		{
			name: "i2.Lower < i2.Upper < i.Lower < i.Upper",
			i:    NewInterval(NewOpen(v3), NewOpen(v4)),
			i2:   NewInterval(NewOpen(v1), NewOpen(v2)),
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.i.Overlap(c.i2))
		})
	}
}
