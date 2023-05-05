package interval

import (
	"testing"
)

func TestInt(t *testing.T) {
	t.Run("NewEndpoint", func(t *testing.T) {
		testNewEndpoint(t, Int(1))
	})
	t.Run("EqualAndBothClosed", func(t *testing.T) {
		testEqualAndBothClosed(t, Int(1), Int(2))
	})
	t.Run("NewInterval", func(t *testing.T) {
		testNewInterval(t, Int(1), Int(2))
	})
	t.Run("IsEmpty", func(t *testing.T) {
		testIsEmpty(t, Int(1), Int(2))
	})
	t.Run("IsEntire", func(t *testing.T) {
		testIsEntire(t, Int(1))
	})
	t.Run("Contains", func(t *testing.T) {
		testContains(t, Int(1), Int(2), Int(3))
	})
	t.Run("CompareInterval", func(t *testing.T) {
		testCompareInterval(t, Int(1), Int(2), Int(3), Int(4))
	})
}
