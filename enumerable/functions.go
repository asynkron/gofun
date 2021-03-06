package enumerable

import (
	"fmt"
	"github.com/asynkron/gofun/set"
	"golang.org/x/exp/constraints"
)

var (
	emptySequenceError = fmt.Errorf("sequence is empty")
)

func FromSlice[T any](items []T) Enumerable[T] {
	return &SliceEnumerable[T]{items}
}

func Min[T constraints.Ordered](enum Enumerable[T]) T {

	min := FirstOrDefault(enum, *new(T))
	enum.Enumerate(func(item T) bool {
		if item < min {
			min = item
		}
		return YieldContinue
	})

	return min
}

func Max[T constraints.Ordered](enum Enumerable[T]) T {

	max := FirstOrDefault(enum, *new(T))
	enum.Enumerate(func(item T) bool {
		if item > max {
			max = item
		}
		return YieldContinue
	})

	return max
}

func Sum[T constraints.Ordered](enum Enumerable[T]) T {

	sum := *new(T)
	enum.Enumerate(func(item T) bool {
		sum += item
		return YieldContinue
	})

	return sum
}

func Aggregate[T any](enum Enumerable[T], seed T, agg func(T, T) T) T {
	res := seed

	enum.Enumerate(func(item T) bool {
		res = agg(res, item)
		return YieldContinue
	})

	return res
}

func Chunk[T any](enum Enumerable[T], size int) Enumerable[[]T] {

	f := func(yield Yield[[]T]) {
		var chunk = make([]T, 0)
		enum.Enumerate(func(i T) bool {
			chunk = append(chunk, i)
			if len(chunk) == size {
				res := yield(chunk)
				chunk = make([]T, 0)
				return res
			}
			return YieldContinue
		})
		if len(chunk) > 0 {
			yield(chunk)
		}
	}

	return &FuncEnumerable[[]T]{f}
}

func Avg[T constraints.Signed | constraints.Unsigned | constraints.Float](enum Enumerable[T]) T {

	count := 0
	sum := *new(T)

	enum.Enumerate(func(item T) bool {
		sum += item
		count++
		return YieldContinue
	})

	return sum / T(count)
}

func All[T any](enum Enumerable[T], predicate func(T) bool) bool {
	var res = true
	enum.Enumerate(func(item T) bool {
		if !predicate(item) {
			res = false
			return YieldBreak
		}
		return YieldContinue
	})
	return res
}

func Any[T any](enum Enumerable[T], predicate func(T) bool) bool {
	var res = false
	enum.Enumerate(func(item T) bool {
		if predicate(item) {
			res = true
			return YieldBreak
		}
		return YieldContinue
	})
	return res
}

func Filter[T any](enum Enumerable[T], predicate func(T) bool) Enumerable[T] {
	f := func(yield Yield[T]) {
		enum.Enumerate(func(item T) bool {
			if predicate(item) {
				res := yield(item)
				return res
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
	res, err := First(enum)
	if err != nil {
		return defaultValue
	}
	return res
}

func First[T any](enum Enumerable[T]) (T, error) {

	switch e := enum.(type) {
	case *SliceEnumerable[T]:
		if len(e.items) > 0 {
			return e.items[0], nil
		}
		return *new(T), emptySequenceError
	default:
		var result T
		var found = false
		enum.Enumerate(func(item T) bool {
			result = item
			found = true
			return YieldBreak
		})
		if found {
			return result, nil
		}
		return *new(T), emptySequenceError
	}
}

func Last[T any](enum Enumerable[T]) (T, error) {
	switch e := enum.(type) {
	case *SliceEnumerable[T]:
		if len(e.items) > 0 {
			return e.items[len(e.items)-1], nil
		}
		return *new(T), emptySequenceError
	default:
		var result T
		var found = false
		enum.Enumerate(func(item T) bool {
			result = item
			found = true
			return YieldContinue
		})
		if found {
			return result, nil
		}
		return *new(T), emptySequenceError
	}
}

func LastOrDefault[T any](enum Enumerable[T], defaultValue T) T {
	res, err := Last(enum)
	if err != nil {
		return defaultValue
	}
	return res
}

func ElementAtOrDefault[T any](enum Enumerable[T], index int, defaultValue T) T {
	res, err := ElementAt(enum, index)
	if err != nil {
		return defaultValue
	}
	return res
}

func ElementAt[T any](enum Enumerable[T], index int) (T, error) {
	switch e := enum.(type) {
	case *SliceEnumerable[T]:
		if len(e.items) > index {
			return e.items[index], nil
		}
		return *new(T), emptySequenceError
	default:
		var result T
		var found = false
		enum.Enumerate(func(item T) bool {
			if index == 0 {
				result = item
				found = true
				return YieldBreak
			}
			index--
			return YieldContinue
		})

		if found {
			return result, nil
		}
		return *new(T), emptySequenceError
	}
}

func Count[T any](enum Enumerable[T]) int {
	switch e := enum.(type) {
	case *SliceEnumerable[T]:
		return len(e.items)
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
			res := yield(mapper(item))
			return res
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
			res := yield(item)
			return res
		})
	}
	return &FuncEnumerable[T]{f}
}

func Limit[T any](enum Enumerable[T], count int) Enumerable[T] {
	f := func(yield Yield[T]) {
		enum.Enumerate(func(item T) bool {
			if count > 0 {
				res := yield(item)
				count--
				return res
			}
			return YieldBreak
		})
	}
	return &FuncEnumerable[T]{f}
}

func ToSet[T comparable](enum Enumerable[T]) *set.MutableSet[T] {
	s := set.New[T]()
	enum.Enumerate(func(item T) bool {
		s.Add(item)
		return YieldContinue
	})
	return s
}

func Except[T comparable](enum Enumerable[T], except Enumerable[T]) Enumerable[T] {
	s := ToSet(except)
	f := func(yield Yield[T]) {
		enum.Enumerate(func(item T) bool {
			if !s.Contains(item) {
				res := yield(item)
				return res
			}
			return YieldContinue
		})
	}
	return &FuncEnumerable[T]{f}
}

func DistinctBy[T any, U comparable](enum Enumerable[T], keySelector func(T) U) Enumerable[T] {
	f := func(yield Yield[T]) {
		seen := set.New[U]()
		enum.Enumerate(func(item T) bool {
			key := keySelector(item)
			if seen.TryAdd(key) {
				res := yield(item)
				return res
			}
			return YieldContinue
		})
	}
	return &FuncEnumerable[T]{f}
}

func Distinct[T comparable](enum Enumerable[T]) Enumerable[T] {
	f := func(yield Yield[T]) {
		seen := set.New[T]()
		enum.Enumerate(func(item T) bool {
			if seen.TryAdd(item) {
				res := yield(item)
				return res
			}
			return YieldContinue
		})
	}
	return &FuncEnumerable[T]{f}
}

func ToMapOfSlice[T any, U comparable](enum Enumerable[T], mapper func(T) U) map[U][]T {
	m := make(map[U][]T)
	enum.Enumerate(func(item T) bool {
		key := mapper(item)
		m[key] = append(m[key], item)
		return YieldContinue
	})
	return m
}

func Concat[T any](enum Enumerable[T], other Enumerable[T]) Enumerable[T] {
	f := func(yield Yield[T]) {
		b := YieldContinue
		enum.Enumerate(func(item T) bool {
			b = yield(item)
			return b
		})
		//did we break in the first loop?
		if b == YieldBreak {
			return
		}
		other.Enumerate(func(item T) bool {
			res := yield(item)
			return res
		})
	}
	return &FuncEnumerable[T]{f}
}

func Contains[T comparable](enum Enumerable[T], item T) bool {
	var result bool
	enum.Enumerate(func(i T) bool {
		if i == item {
			result = true
			return YieldBreak
		}
		return YieldContinue
	})
	return result
}
