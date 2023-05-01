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

// expect v1 < v2 < v3 < v4
func testOverlap[T Ordered[T]](t *testing.T, v1, v2, v3, v4 T) {
	unbounded := UnboundedEp[T]()
	cases := []struct {
		name string
		i    Interval[T]
		i2   Interval[T]
		want bool
	}{
		{
			name: "both are empty",
			i:    New(OpenEp(v1), OpenEp(v1)),
			i2:   New(OpenEp(v1), OpenEp(v1)),
			want: false,
		},
		{
			name: "i is empty",
			i:    New(OpenEp(v1), OpenEp(v1)),
			i2:   New(unbounded, unbounded),
			want: false,
		},
		{
			name: "i2 is empty",
			i:    New(unbounded, unbounded),
			i2:   New(OpenEp(v1), OpenEp(v1)),
			want: false,
		},
		// in the following cases, the intervals are not empty
		{
			name: "both are unbounded",
			i:    New(unbounded, unbounded),
			i2:   New(unbounded, unbounded),
			want: true,
		},
		{
			name: "i is unbounded",
			i:    New(unbounded, unbounded),
			i2:   New(OpenEp(v1), OpenEp(v2)),
			want: true,
		},
		{
			name: "i2 is unbounded",
			i:    New(OpenEp(v1), OpenEp(v2)),
			i2:   New(unbounded, unbounded),
			want: true,
		},
		{
			name: "both are lower unbounded",
			i:    New(unbounded, OpenEp(v1)),
			i2:   New(unbounded, OpenEp(v2)),
			want: true,
		},
		{
			name: "both are upper unbounded",
			i:    New(OpenEp(v1), unbounded),
			i2:   New(OpenEp(v2), unbounded),
			want: true,
		},
		// i.Lower is unbounded
		{
			name: "i.Lower is unbounded, i.Upper < i2.Lower",
			i:    New(unbounded, OpenEp(v1)),
			i2:   New(OpenEp(v2), OpenEp(v3)),
			want: false,
		},
		{
			name: "i.Lower is unbounded, i.Upper = i2.Lower, no contact",
			i:    New(unbounded, OpenEp(v1)),
			i2:   New(OpenEp(v1), OpenEp(v2)),
			want: false,
		},
		{
			name: "i.Lower is unbounded, i.Upper = i2.Lower, contact",
			i:    New(unbounded, ClosedEp(v1)),
			i2:   New(ClosedEp(v1), OpenEp(v2)),
			want: true,
		},
		{
			name: "i.Lower is unbounded, i.Upper > i2.Lower",
			i:    New(unbounded, OpenEp(v2)),
			i2:   New(OpenEp(v1), OpenEp(v3)),
			want: true,
		},
		{
			name: "i.Lower is unbounded, i.Upper > i2.Upper",
			i:    New(unbounded, OpenEp(v3)),
			i2:   New(OpenEp(v1), OpenEp(v2)),
			want: true,
		},
		// i.Upper is unbounded
		{
			name: "i.Upper is unbounded, i.Lower > i2.Upper",
			i:    New(OpenEp(v3), unbounded),
			i2:   New(OpenEp(v1), OpenEp(v2)),
			want: false,
		},
		{
			name: "i.Upper is unbounded, i.Lower = i2.Upper, no contact",
			i:    New(OpenEp(v2), unbounded),
			i2:   New(OpenEp(v1), OpenEp(v2)),
			want: false,
		},
		{
			name: "i.Upper is unbounded, i.Lower = i2.Upper, contact",
			i:    New(ClosedEp(v2), unbounded),
			i2:   New(OpenEp(v1), ClosedEp(v2)),
			want: true,
		},
		{
			name: "i.Upper is unbounded, i.Lower< i2.Upper",
			i:    New(OpenEp(v2), unbounded),
			i2:   New(OpenEp(v1), OpenEp(v3)),
			want: true,
		},
		{
			name: "i.Upper is unbounded, i.Lower < i2.Lower",
			i:    New(OpenEp(v1), unbounded),
			i2:   New(OpenEp(v2), OpenEp(v3)),
			want: true,
		},
		// i2.Lower is unbounded
		{
			name: "i2.Lower is unbounded, i2.Upper < i.Lower",
			i:    New(OpenEp(v2), OpenEp(v3)),
			i2:   New(unbounded, OpenEp(v1)),
			want: false,
		},
		{
			name: "i2.Lower is unbounded, i2.Upper = i.Lower, no contact",
			i:    New(OpenEp(v1), OpenEp(v2)),
			i2:   New(unbounded, OpenEp(v1)),
			want: false,
		},
		{
			name: "i2.Lower is unbounded, i2.Upper = i.Lower, contact",
			i:    New(ClosedEp(v1), OpenEp(v2)),
			i2:   New(unbounded, ClosedEp(v1)),
			want: true,
		},
		{
			name: "i2.Lower is unbounded, i2.Upper > i.Lower",
			i:    New(OpenEp(v1), OpenEp(v3)),
			i2:   New(unbounded, OpenEp(v2)),
			want: true,
		},
		{
			name: "i2.Lower is unbounded, i2.Upper > i.Upper",
			i:    New(OpenEp(v1), OpenEp(v2)),
			i2:   New(unbounded, OpenEp(v3)),
			want: true,
		},
		// i2.Upper is unbounded
		{
			name: "i2.Upper is unbounded, i2.Lower > i.Lower",
			i:    New(OpenEp(v1), OpenEp(v2)),
			i2:   New(OpenEp(v3), unbounded),
			want: false,
		},
		{
			name: "i2.Upper is unbounded, i2.Lower = i.Upper, no contact",
			i:    New(OpenEp(v1), OpenEp(v2)),
			i2:   New(OpenEp(v2), unbounded),
			want: false,
		},
		{
			name: "i2.Upper is unbounded, i2.Lower = i.Upper, contact",
			i:    New(OpenEp(v1), ClosedEp(v2)),
			i2:   New(ClosedEp(v2), unbounded),
			want: true,
		},
		{
			name: "i2.Upper is unbounded, i2.Lower < i.Upper",
			i:    New(OpenEp(v1), OpenEp(v3)),
			i2:   New(OpenEp(v2), unbounded),
			want: true,
		},
		{
			name: "i2.Upper is unbounded, i2.Lower < i.Lower",
			i:    New(OpenEp(v2), OpenEp(v3)),
			i2:   New(OpenEp(v1), unbounded),
			want: true,
		},
		// in the following cases, both intervals are bounded
		{
			name: "i.Lower = i.Upper = i2.Lower = i2.Upper",
			i:    New(ClosedEp(v1), ClosedEp(v1)),
			i2:   New(ClosedEp(v1), ClosedEp(v1)),
			want: true,
		},
		{
			name: "i.Lower < i.Upper < i2.Lower < i2.Upper",
			i:    New(OpenEp(v1), OpenEp(v2)),
			i2:   New(OpenEp(v3), OpenEp(v4)),
			want: false,
		},
		{
			name: "i.Lower < i.Upper = i2.Lower < i2.Upper, no contact",
			i:    New(OpenEp(v1), OpenEp(v2)),
			i2:   New(OpenEp(v2), OpenEp(v3)),
			want: false,
		},
		{
			name: "i.Lower < i.Upper = i2.Lower < i2.Upper, contact",
			i:    New(OpenEp(v1), ClosedEp(v2)),
			i2:   New(ClosedEp(v2), OpenEp(v3)),
			want: true,
		},
		{
			name: "i.Lower < i2.Lower < i.Upper < i2.Upper",
			i:    New(OpenEp(v1), OpenEp(v3)),
			i2:   New(OpenEp(v2), OpenEp(v4)),
			want: true,
		},
		{
			name: "i.Lower < i2.Lower < i.Upper = i2.Upper",
			i:    New(OpenEp(v1), OpenEp(v3)),
			i2:   New(OpenEp(v2), OpenEp(v3)),
			want: true,
		},
		{
			name: "i.Lower < i2.Lower < i2.Upper < i.Upper",
			i:    New(OpenEp(v1), OpenEp(v4)),
			i2:   New(OpenEp(v2), OpenEp(v3)),
			want: true,
		},
		{
			name: "i.Lower = i2.Lower",
			i:    New(OpenEp(v1), OpenEp(v2)),
			i2:   New(OpenEp(v1), OpenEp(v4)),
			want: true,
		},
		{
			name: "i2.Lower < i.Lower < i.Upper < i2.Upper",
			i:    New(OpenEp(v2), OpenEp(v3)),
			i2:   New(OpenEp(v1), OpenEp(v4)),
			want: true,
		},
		{
			name: "i2.Lower < i.Lower < i2.Upper < i.Upper",
			i:    New(OpenEp(v2), OpenEp(v4)),
			i2:   New(OpenEp(v1), OpenEp(v3)),
			want: true,
		},
		{
			name: "i2.Lower < i.Lower = i2.Upper = i.Upper, no contact",
			i:    New(ClosedEp(v2), ClosedEp(v2)),
			i2:   New(OpenEp(v1), OpenEp(v2)),
			want: false,
		},
		{
			name: "i2.Lower < i.Lower = i2.Upper = i.Upper, contact",
			i:    New(ClosedEp(v2), ClosedEp(v2)),
			i2:   New(OpenEp(v1), ClosedEp(v2)),
			want: true,
		},
		{
			name: "i2.Lower < i.Lower = i2.Upper < i.Upper, no contact",
			i:    New(OpenEp(v2), OpenEp(v3)),
			i2:   New(OpenEp(v1), OpenEp(v2)),
			want: false,
		},
		{
			name: "i2.Lower < i.Lower = i2.Upper < i.Upper, contact",
			i:    New(ClosedEp(v2), OpenEp(v3)),
			i2:   New(OpenEp(v1), ClosedEp(v2)),
			want: true,
		},
		{
			name: "i2.Lower < i2.Upper < i.Lower < i.Upper",
			i:    New(OpenEp(v3), OpenEp(v4)),
			i2:   New(OpenEp(v1), OpenEp(v2)),
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.i.Overlap(c.i2))
		})
	}
}
