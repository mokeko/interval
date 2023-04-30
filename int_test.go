package interval

import (
	"testing"
)

func TestIntIsEmpty(t *testing.T) {
	cases := []struct {
		name     string
		interval Interval[Int]
		want     bool
	}{
		{
			name: "unbounded",
			interval: NewInterval(
				NewUnbounded[Int](),
				NewUnbounded[Int](),
			),
			want: false,
		},
		{
			name: "lower unbounded",
			interval: NewInterval(
				NewUnbounded[Int](),
				NewClosed(Int(1)),
			),
			want: false,
		},
		{
			name: "upper unbounded",
			interval: NewInterval(
				NewClosed(Int(1)),
				NewUnbounded[Int](),
			),
			want: false,
		},
		{
			name: "lower < upper",
			interval: NewInterval(
				NewOpen(Int(1)),
				NewOpen(Int(2)),
			),
			want: false,
		},
		{
			name: "lowe == upper, closed",
			interval: NewInterval(
				NewClosed(Int(1)),
				NewClosed(Int(1)),
			),
			want: false,
		},
		{
			name: "lower == upper, left open",
			interval: NewInterval(
				NewOpen(Int(1)),
				NewClosed(Int(1)),
			),
			want: true,
		},
		{
			name: "lower == upper, right open",
			interval: NewInterval(
				NewClosed(Int(1)),
				NewOpen(Int(1)),
			),
			want: true,
		},
		{
			name: "lower == upper, both open",
			interval: NewInterval(
				NewOpen(Int(1)),
				NewOpen(Int(1)),
			),
			want: true,
		},
		{
			name: "lower > upper",
			interval: NewInterval(
				NewClosed(Int(2)),
				NewClosed(Int(1)),
			),
			want: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.interval.IsEmpty())
		})
	}
}

func TestIntIsEntire(t *testing.T) {
	cases := []struct {
		name     string
		interval Interval[Int]
		want     bool
	}{
		{
			name: "unbounded",
			interval: NewInterval(
				NewUnbounded[Int](),
				NewUnbounded[Int](),
			),
			want: true,
		},
		{
			name: "lower unbounded",
			interval: NewInterval(
				NewUnbounded[Int](),
				NewClosed(Int(1)),
			),
			want: false,
		},
		{
			name: "upper unbounded",
			interval: NewInterval(
				NewClosed(Int(1)),
				NewUnbounded[Int](),
			),
			want: false,
		},
		{
			name: "bounded",
			interval: NewInterval(
				NewClosed(Int(1)),
				NewClosed(Int(2)),
			),
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.interval.IsEntire())
		})
	}
}

func TestIntOverlap(t *testing.T) {
	cases := []struct {
		name string
		i    Interval[Int]
		i2   Interval[Int]
		want bool
	}{
		{
			name: "unbounded",
			i: NewInterval(
				NewUnbounded[Int](),
				NewUnbounded[Int](),
			),
			i2: NewInterval(
				NewUnbounded[Int](),
				NewUnbounded[Int](),
			),
			want: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.i.Overlap(c.i2))
		})
	}
}
