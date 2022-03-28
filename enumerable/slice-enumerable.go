package enumerable

type SliceEnumerable[T any] struct {
	items []T
}

func (e *SliceEnumerable[T]) Enumerate(yield Yield[T]) {
	for _, item := range e.items {
		res := yield(item)
		if res == YieldBreak {
			break
		}
	}
}
