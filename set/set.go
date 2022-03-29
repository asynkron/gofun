package set

type Set[T comparable] interface {
	//all methods from MutableSet

	Contains(e T) bool
	Size() int
	IsEmpty() bool
	ToSlice() []T
	IsSubset(other Set[T]) bool
	IsSuperset(other Set[T]) bool
	IsProperSuperset(other Set[T]) bool
	IsProperSubset(other Set[T]) bool
	Equals(other Set[T]) bool
	String() string
}
