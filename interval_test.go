package interval

import "testing"

func testNewInterval[T Ordered[T]](t *testing.T, v1, v2 T) {
	assertEqual(t, Interval[T]{
		Lower: NewOpen(v1),
		Upper: NewOpen(v2),
	}, NewInterval(NewOpen(v1), NewOpen(v2)))
}
