package interval

import "testing"

func assertEqual[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestNewEndpoint(t *testing.T) {
	t.Run("open", func(t *testing.T) {
		assertEqual(
			t,
			Endpoint[Int]{
				Value:   Int(1),
				Closed:  false,
				Bounded: true,
			},
			NewOpen(Int(1)),
		)
	})
	t.Run("closed", func(t *testing.T) {
		assertEqual(
			t,
			Endpoint[Int]{
				Value:   Int(2),
				Closed:  true,
				Bounded: true,
			},
			NewClosed(Int(2)),
		)
	})
	t.Run("unbounded", func(t *testing.T) {
		assertEqual(
			t,
			Endpoint[Int]{
				Bounded: false,
			},
			NewUnbounded[Int](),
		)
	})
}

func TestEqualAndBothClosed(t *testing.T) {
	unbounded := NewUnbounded[Int]()
	v1 := Int(1)
	v2 := Int(2)
	cases := []struct {
		name  string
		e, e2 Endpoint[Int]
		want  bool
	}{
		{
			name: "unbounded, unbounded",
			e:    unbounded,
			e2:   unbounded,
			want: false,
		},
		{
			name: "unbounded, open",
			e:    unbounded,
			e2:   NewOpen(v1),
			want: false,
		},
		{
			name: "unbounded, closed",
			e:    unbounded,
			e2:   NewClosed(v1),
			want: false,
		},
		{
			name: "open, unbounded",
			e:    NewOpen(v1),
			e2:   unbounded,
			want: false,
		},
		{
			name: "open, open",
			e:    NewOpen(v1),
			e2:   NewOpen(v1),
			want: false,
		},
		{
			name: "open, closed",
			e:    NewOpen(v1),
			e2:   NewClosed(v1),
			want: false,
		},
		{
			name: "closed, unbounded",
			e:    NewClosed(v1),
			e2:   unbounded,
			want: false,
		},
		{
			name: "closed, open",
			e:    NewClosed(v1),
			e2:   NewOpen(v1),
			want: false,
		},
		{
			name: "closed, closed, different values",
			e:    NewClosed(v1),
			e2:   NewClosed(v2),
			want: false,
		},
		{
			name: "closed, closed",
			e:    NewClosed(v1),
			e2:   NewClosed(v1),
			want: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.e.equalAndBothClosed(c.e2))
		})
	}
}
