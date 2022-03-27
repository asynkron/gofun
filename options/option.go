package options

type Option[T any] struct {
	hasValue bool
	value    *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{true, &value}
}

func None[T any]() Option[T] {
	return Option[T]{false, nil}
}

func (o *Option[T]) IsSome() bool {
	if o == nil {
		return false
	}

	return o.hasValue
}

func (o *Option[T]) IsNone() bool {
	if o == nil {
		return true
	}
	return !o.hasValue
}
