package enumerable

const (
	YieldContinue = true
	YieldBreak    = false
)

type Yield[T any] func(item T) bool

type Enumerable[T any] interface {
	Enumerate(yield Yield[T])

	//Cannot make Map/Select as methods cannot have type parameters
}
