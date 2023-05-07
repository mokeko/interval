package interval

import "testing"

func assertEqual(t *testing.T, want, got any) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func testNewEndpoint[T Ordered[T]](t *testing.T, v T) {
	t.Run("open", func(t *testing.T) {
		assertEqual(
			t,
			Endpoint[T]{
				Value:     v,
				Closed:    false,
				Unbounded: false,
			},
			OpenEp(v),
		)
	})
	t.Run("closed", func(t *testing.T) {
		assertEqual(
			t,
			Endpoint[T]{
				Value:     v,
				Closed:    true,
				Unbounded: false,
			},
			ClosedEp(v),
		)
	})
	t.Run("unbounded", func(t *testing.T) {
		assertEqual(
			t,
			Endpoint[T]{
				Unbounded: true,
			},
			UnboundedEp[T](),
		)
	})
}

func testEqualAndBothClosed[T Ordered[T]](t *testing.T, v1, v2 T) {
	if v1.Equal(v2) {
		t.Fatalf("v1 and v2 must not be equal. v1: %v, v2: %v", v1, v2)
	}

	unbounded := UnboundedEp[T]()
	cases := []struct {
		name  string
		e, e2 Endpoint[T]
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
			e2:   OpenEp(v1),
			want: false,
		},
		{
			name: "unbounded, closed",
			e:    unbounded,
			e2:   ClosedEp(v1),
			want: false,
		},
		{
			name: "open, unbounded",
			e:    OpenEp(v1),
			e2:   unbounded,
			want: false,
		},
		{
			name: "open, open",
			e:    OpenEp(v1),
			e2:   OpenEp(v1),
			want: false,
		},
		{
			name: "open, closed",
			e:    OpenEp(v1),
			e2:   ClosedEp(v1),
			want: false,
		},
		{
			name: "closed, unbounded",
			e:    ClosedEp(v1),
			e2:   unbounded,
			want: false,
		},
		{
			name: "closed, open",
			e:    ClosedEp(v1),
			e2:   OpenEp(v1),
			want: false,
		},
		{
			name: "closed, closed, different values",
			e:    ClosedEp(v1),
			e2:   ClosedEp(v2),
			want: false,
		},
		{
			name: "closed, closed",
			e:    ClosedEp(v1),
			e2:   ClosedEp(v1),
			want: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertEqual(t, c.want, c.e.equalAndBothClosed(c.e2))
		})
	}
}

func TestBounded(t *testing.T) {
	assertEqual(t, true, Endpoint[Int]{}.Bounded())
	assertEqual(t, false, Endpoint[Int]{Unbounded: true}.Bounded())
}
