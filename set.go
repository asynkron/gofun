package main

import "fmt"

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

func (s *Set[T]) TryAdd(e T) bool {
	if s.Contains(e) {
		return false
	}
	s.Add(e)
	return true
}

func (s *Set[T]) TryRemove(e T) bool {
	if !s.Contains(e) {
		return false
	}
	s.Remove(e)
	return true
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.m))
	for k := range s.m {
		slice = append(slice, k)
	}
	return slice
}

func (s *Set[T]) Clone() *Set[T] {
	set := NewSet[T]()
	for k := range s.m {
		set.Add(k)
	}
	return set
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	if other == nil {
		return s.Clone()
	}

	union := s.Clone()
	for k := range other.m {
		union.Add(k)
	}
	return union
}

func (s *Set[T]) Except(other *Set[T]) *Set[T] {
	if other == nil {
		return s.Clone()
	}

	except := NewSet[T]()
	for k := range s.m {
		if !other.Contains(k) {
			except.Add(k)
		}
	}
	return except
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	if other == nil {
		return NewSet[T]()
	}

	intersect := NewSet[T]()
	for k := range s.m {
		if other.Contains(k) {
			intersect.Add(k)
		}
	}
	return intersect
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if other == nil {
		return false
	}

	for k := range s.m {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	if other == nil {
		return false
	}

	return other.IsSubset(s)
}

func (s *Set[T]) IsProperSuperset(other *Set[T]) bool {
	if other == nil {
		return false
	}

	return s.Size() > other.Size() && s.IsSuperset(other)
}

func (s *Set[T]) IsProperSubset(other *Set[T]) bool {
	if other == nil {
		return false
	}

	return other.IsProperSuperset(s)
}

func (s *Set[T]) Equals(other *Set[T]) bool {
	if other == nil {
		return false
	}

	return s.IsSubset(other) && other.IsSubset(s)
}

func (s *Set[T]) String() string {
	str := "{"
	for k := range s.m {
		str += fmt.Sprintf("%v, ", k)
	}
	str = str[:len(str)-2]
	str += "}"
	return str
}
