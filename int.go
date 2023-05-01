package interval

var _ Ordered[Int] = Int(0)

// Int is a wrapper of int.
// It implements the Ordered interface.
type Int int

func (i Int) Equal(i2 Int) bool {
	return i == i2
}

func (i Int) LessThan(i2 Int) bool {
	return i < i2
}
