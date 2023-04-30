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

func testOverlap[T Ordered[T]](t *testing.T) {
	unbounded := NewUnbounded[T]()
	cases := []struct {
		name string
		i    Interval[T]
		i2   Interval[T]
		want bool
	}{
		{
			name: "unbounded",
			i:    NewInterval(unbounded, unbounded),
			i2:   NewInterval(unbounded, unbounded),
			want: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.i.Overlap(c.i2))
		})
	}
}
