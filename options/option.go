package options

type Option[T any] struct {
	value *T
}

func Some[T any](value *T) *Option[T] {
	if value == nil {
		return nil
	}
	return &Option[T]{value}
}

func None[T any]() *Option[T] {
	return nil
}

func IsSome[T any](o *Option[T]) bool {
	if o == nil {
		return false
	}

	if o.value == nil {
		return false
	}

	return true
}

func IsNone[T any](o *Option[T]) bool {
	return !IsSome(o)
}

func Map[T any, U any](o *Option[T], f func(T) U) *Option[U] {
	if o == nil {
		return nil
	}

	res := f(*o.value)
	return Some[U](&res)
}

func Match[T any](o *Option[T], f func(T), n func()) {
	if IsSome(o) {
		f(*o.value)
	} else {
		n()
	}
}

func GetOrDefault[T any](o *Option[T], defaultValue T) T {
	if IsSome(o) {
		v := o.value
		return *v
	}

	return defaultValue
}
