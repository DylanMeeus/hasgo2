package pkg

type Maybe[T any] struct {
	Value   T
	Present bool
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{Present: false}
}

func Just[T any](t T) Maybe[T] {
	return Maybe[T]{
		Value:   t,
		Present: true,
	}
}
