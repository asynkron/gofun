package main

const (
	YieldContinue = true
	YieldBreak    = false
)

type Yield[T any] func(item T) bool

type Enumerable[T any] interface {
	Enumerate(yield Yield[T])
	Filter(predicate func(T) bool) Enumerable[T]
	FirstOrDefault(defaultValue T) T
	LastOrDefault(defaultValue T) T
	ToSlice() []T
	Count() int
	ElementAtOrDefault(index int, defaultValue T) T
}
