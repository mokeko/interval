package interval

// Ordered is an interface for types that have a order relation.
type Ordered[T any] interface {
	Equal(T) bool
	LessThan(T) bool
}

// Endpoint represents an endpoint of an interval.
// It contains a value and a flag indicating whether it is closed and bounded.
// The type T must have a order relation.
type Endpoint[T Ordered[T]] struct {
	Value     T
	Closed    bool
	Unbounded bool
}

// OpenEp returns an open endpoint.
func OpenEp[T Ordered[T]](v T) Endpoint[T] {
	return Endpoint[T]{
		Value:  v,
		Closed: false,
	}
}

// ClosedEp returns a closed endpoint.
func ClosedEp[T Ordered[T]](v T) Endpoint[T] {
	return Endpoint[T]{
		Value:  v,
		Closed: true,
	}
}

// UnboundedEp returns an unbounded endpoint.
func UnboundedEp[T Ordered[T]]() Endpoint[T] {
	return Endpoint[T]{
		Unbounded: true,
	}
}

// Bounded is just a negation of Unbounded.
func (e Endpoint[T]) Bounded() bool {
	return !e.Unbounded
}

func (e Endpoint[T]) equalAndBothClosed(e2 Endpoint[T]) bool {
	return e.Value.Equal(e2.Value) && e.Closed && e2.Closed
}
