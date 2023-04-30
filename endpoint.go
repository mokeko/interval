package interval

type Ordered[T any] interface {
	Equal(T) bool
	LessThan(T) bool
}

type Endpoint[T Ordered[T]] struct {
	Value   T
	Closed  bool
	Bounded bool
}

func NewOpen[T Ordered[T]](v T) Endpoint[T] {
	return Endpoint[T]{
		Value:   v,
		Closed:  false,
		Bounded: true,
	}
}

func NewClosed[T Ordered[T]](v T) Endpoint[T] {
	return Endpoint[T]{
		Value:   v,
		Closed:  true,
		Bounded: true,
	}
}

func NewUnbounded[T Ordered[T]]() Endpoint[T] {
	return Endpoint[T]{
		Bounded: false,
	}
}
