package interval

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t1 := Time(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC))
	t2 := Time(time.Date(1, 1, 2, 0, 0, 0, 0, time.UTC))
	t3 := Time(time.Date(1, 1, 3, 0, 0, 0, 0, time.UTC))
	t4 := Time(time.Date(1, 1, 4, 0, 0, 0, 0, time.UTC))
	t.Run("NewEndpoint", func(t *testing.T) {
		testNewEndpoint(t, t1)
	})
	t.Run("EqualAndBothClosed", func(t *testing.T) {
		testEqualAndBothClosed(t, t1, t2)
	})
	t.Run("NewInterval", func(t *testing.T) {
		testNewInterval(t, t1, t2)
	})
	t.Run("IsEmpty", func(t *testing.T) {
		testIsEmpty(t, t1, t2)
	})
	t.Run("IsEntire", func(t *testing.T) {
		testIsEntire(t, t1)
	})
	t.Run("Overlap", func(t *testing.T) {
		testOverlap(t, t1, t2, t3, t4)
	})
}
