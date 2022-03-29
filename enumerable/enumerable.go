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

type Enumerator[T any] interface {
	MoveNext() (T, bool)
}

type FuncEnumerator[T any] struct {
	fun func() (T, bool)
}

func (fe *FuncEnumerator[T]) MoveNext() (T, bool) {
	return fe.fun()
}
