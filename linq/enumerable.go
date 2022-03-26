package linq

const (
	YieldContinue = true
	YieldBreak    = false
)

type Yield[T any] func(item T) bool

type Enumerable[T any] interface {
	Enumerate(yield Yield[T])
}
