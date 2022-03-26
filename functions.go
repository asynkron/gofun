package main

func FromSlice[T any](items []T) Enumerable[T] {
	return &SliceEnumerable[T]{items}
}

func Filter[T any](enum Enumerable[T], predicate func(T) bool) Enumerable[T] {

	f := func(yield Yield[T]) {
		enum.Enumerate(func(item T) bool {
			if predicate(item) {
				yield(item)
			}
			return YieldContinue
		})
	}

	return &FuncEnumerable[T]{f}
}

func ToSlice[T any](enum Enumerable[T]) []T {
	s := make([]T, 0)
	enum.Enumerate(func(item T) bool {
		s = append(s, item)
		return YieldContinue
	})
	return s
}

func FirstOrDefault[T any](enum Enumerable[T], defaultValue T) T {
	switch enum.(type) {
	case *SliceEnumerable[T]:
		if len(enum.(*SliceEnumerable[T]).items) > 0 {
			return enum.(*SliceEnumerable[T]).items[0]
		}
		return defaultValue
	default:
		var result T
		enum.Enumerate(func(item T) bool {
			result = item
			return YieldBreak
		})
		return result
	}
}

func LastOrDefault[T any](enum Enumerable[T], defaultValue T) T {

	switch enum.(type) {
	case *SliceEnumerable[T]:
		s := enum.(*SliceEnumerable[T])
		if len(s.items) > 0 {
			return s.items[len(s.items)-1]
		}
		return defaultValue
	default:
		var result T
		enum.Enumerate(func(item T) bool {
			result = item
			return YieldContinue
		})
		return result
	}
}

func ElementAtOrDefault[T any](enum Enumerable[T], index int, defaultValue T) T {
	switch enum.(type) {
	case *SliceEnumerable[T]:
		if len(enum.(*SliceEnumerable[T]).items) > index {
			return enum.(*SliceEnumerable[T]).items[index]
		}
		return defaultValue
	default:
		var result T
		enum.Enumerate(func(item T) bool {
			if index == 0 {
				result = item
				return YieldBreak
			}
			index--
			return YieldContinue
		})
		return result
	}
}

func Count[T any](enum Enumerable[T]) int {
	switch enum.(type) {
	case *SliceEnumerable[T]:
		return len(enum.(*SliceEnumerable[T]).items)
	default:
		var count int
		enum.Enumerate(func(item T) bool {
			count++
			return YieldContinue
		})
		return count
	}
}

func From[T any](items ...T) Enumerable[T] {
	return &SliceEnumerable[T]{items}
}

func Map[T any, U any](enum Enumerable[T], mapper func(T) U) Enumerable[U] {

	f := func(yield Yield[U]) {
		enum.Enumerate(func(item T) bool {
			yield(mapper(item))
			return YieldContinue
		})
	}

	return &FuncEnumerable[U]{f}
}

func Skip[T any](enum Enumerable[T], count int) Enumerable[T] {
	f := func(yield Yield[T]) {
		enum.Enumerate(func(item T) bool {
			if count > 0 {
				count--
				return YieldContinue
			}
			yield(item)
			return YieldContinue
		})
	}
	return &FuncEnumerable[T]{f}
}

func Limit[T any](enum Enumerable[T], count int) Enumerable[T] {
	f := func(yield Yield[T]) {
		enum.Enumerate(func(item T) bool {
			if count > 0 {
				yield(item)
				count--
				return YieldContinue
			}
			return YieldBreak
		})
	}
	return &FuncEnumerable[T]{f}
}
