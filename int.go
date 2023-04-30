package interval

var _ Ordered[Int] = Int(0)

type Int int

func (i Int) Equal(i2 Int) bool {
	return i == i2
}

func (i Int) LessThan(i2 Int) bool {
	return i < i2
}
