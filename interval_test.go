package interval

import "testing"

func TestNewIntInterval(t *testing.T) {
	assertEqual(t, Interval[Int]{
		Lower: NewOpen(Int(1)),
		Upper: NewClosed(Int(2)),
	}, NewInterval(NewOpen(Int(1)), NewClosed(Int(2))))

	assertEqual(t, Interval[Int]{
		Lower: NewUnbounded[Int](),
		Upper: NewOpen(Int(4)),
	}, NewInterval(NewUnbounded[Int](), NewOpen(Int(4))))
}
