package linq

type FuncEnumerable[T any] struct {
	enumerateFun func(yield Yield[T])
}

func (e *FuncEnumerable[T]) Enumerate(yield Yield[T]) {
	e.enumerateFun(yield)
}
