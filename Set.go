package main

type Set[T comparable] struct {
	m map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	d := make(map[T]bool)
	return &Set[T]{m: d}
}

func (s *Set[T]) Add(e T) {
	s.m[e] = true
}

func (s *Set[T]) Remove(e T) {
	delete(s.m, e)
}

func (s *Set[T]) Contains(e T) bool {
	_, ok := s.m[e]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]bool)
}
