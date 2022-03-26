package go_generic_linq

type FuncEnumerable[T any] struct {
	enumerateFun func(yield Yield[T])
}

func (e *FuncEnumerable[T]) Filter(predicate func(T) bool) Enumerable[T] {
	return Filter[T](e, predicate)
}

func (e *FuncEnumerable[T]) FirstOrDefault(defaultValue T) T {
	return FirstOrDefault[T](e, defaultValue)
}

func (e *FuncEnumerable[T]) LastOrDefault(defaultValue T) T {
	return LastOrDefault[T](e, defaultValue)
}

func (e *FuncEnumerable[T]) ToSlice() []T {
	return ToSlice[T](e)
}

func (e *FuncEnumerable[T]) Count() int {
	return Count[T](e)
}

func (e *FuncEnumerable[T]) ElementAtOrDefault(index int, defaultValue T) T {
	return ElementAtOrDefault[T](e, index, defaultValue)
}

func (e *FuncEnumerable[T]) Enumerate(yield Yield[T]) {
	e.enumerateFun(yield)
}
