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
