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

func (e *SliceEnumerable[T]) GetEnumerator() Enumerator[T] {
	var index = 0
	return &FuncEnumerator[T]{
		fun: func() (T, bool) {
			if index < len(e.items) {
				index++
				return e.items[index-1], true
			}
			return *new(T), false
		},
	}
}
