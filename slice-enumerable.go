package go_generic_linq

type SliceEnumerable[T any] struct {
	items []T
}

func (e *SliceEnumerable[T]) Filter(predicate func(T) bool) Enumerable[T] {
	return Filter[T](e, predicate)
}

func (e *SliceEnumerable[T]) FirstOrDefault(defaultValue T) T {
	//if has items, return first item
	if len(e.items) > 0 {
		return e.items[0]
	}
	//otherwise return default value
	return defaultValue
}

func (e *SliceEnumerable[T]) LastOrDefault(defaultValue T) T {
	//if has items, return last item
	if len(e.items) > 0 {
		return e.items[len(e.items)-1]
	}
	//otherwise return default value
	return defaultValue
}

func (e *SliceEnumerable[T]) ToSlice() []T {
	return append([]T{}, e.items...)
}

func (e *SliceEnumerable[T]) Count() int {
	return len(e.items)
}

func (e *SliceEnumerable[T]) ElementAtOrDefault(index int, defaultValue T) T {
	if len(e.items) > index {
		return e.items[index]
	}
	return defaultValue
}

func (e *SliceEnumerable[T]) Enumerate(yield Yield[T]) {
	for _, item := range e.items {
		yield(item)
	}
}
