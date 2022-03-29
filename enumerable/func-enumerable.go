package enumerable

type FuncEnumerable[T any] struct {
	enumerateFun func(yield Yield[T])
}

func (e *FuncEnumerable[T]) Enumerate(yield Yield[T]) {
	e.enumerateFun(yield)
}

func (e *FuncEnumerable[T]) GetEnumerator() Enumerator[T] {
	return &FuncEnumerator[T]{
		fun: func() (T, bool) {
			return *new(T), true
		},
	}
}
