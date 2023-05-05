package interval

var _ Ordered[Int] = Int(0)

// Int is a wrapper of int.
// It implements the Ordered interface.
type Int int

// Equal checks if i is equal to i2.
func (i Int) Equal(i2 Int) bool {
	return i == i2
}

// LessThan checks if i is less than i2.
func (i Int) LessThan(i2 Int) bool {
	return i < i2
}
