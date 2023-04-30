package interval

import (
	"testing"
)

func TestInt(t *testing.T) {
	t.Run("NewEndpoint", func(t *testing.T) {
		testNewEndpoint(t, Int(1))
	})
	t.Run("EndpointEqualAndBothClosed", func(t *testing.T) {
		testEndpointEqualAndBothClosed(t, Int(1), Int(2))
	})
	t.Run("NewInterval", func(t *testing.T) {
		testNewInterval(t, Int(1), Int(2))
	})

}

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
			name: "lowe = upper, closed",
			interval: NewInterval(
				NewClosed(Int(1)),
				NewClosed(Int(1)),
			),
			want: false,
		},
		{
			name: "lower = upper, lower open",
			interval: NewInterval(
				NewOpen(Int(1)),
				NewClosed(Int(1)),
			),
			want: true,
		},
		{
			name: "lower = upper, upper open",
			interval: NewInterval(
				NewClosed(Int(1)),
				NewOpen(Int(1)),
			),
			want: true,
		},
		{
			name: "lower = upper, both open",
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
