package set

import (
	"fmt"
	"golang.org/x/exp/maps"
)

type MutableSet[T comparable] struct {
	m map[T]bool
}

func New[T comparable](items ...T) *MutableSet[T] {
	d := make(map[T]bool)
	for _, item := range items {
		d[item] = true
	}
	return &MutableSet[T]{m: d}
}

func (s *MutableSet[T]) Add(e T) {
	s.m[e] = true
}

func (s *MutableSet[T]) AddRange(e ...T) {
	for _, i := range e {
		s.Add(i)
	}
}

func (s *MutableSet[T]) Remove(e T) {
	delete(s.m, e)
}

func (s *MutableSet[T]) Contains(e T) bool {
	_, ok := s.m[e]
	return ok
}

func (s *MutableSet[T]) Size() int {
	return len(s.m)
}

func (s *MutableSet[T]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *MutableSet[T]) Clear() {
	s.m = make(map[T]bool)
}

func (s *MutableSet[T]) TryAdd(e T) bool {
	if s.Contains(e) {
		return false
	}
	s.Add(e)
	return true
}

func (s *MutableSet[T]) TryRemove(e T) bool {
	if !s.Contains(e) {
		return false
	}
	s.Remove(e)
	return true
}

func (s *MutableSet[T]) ToSlice() []T {
	return maps.Keys(s.m)
}

func (s *MutableSet[T]) Clone() *MutableSet[T] {
	return &MutableSet[T]{m: maps.Clone(s.m)}
}

func (s *MutableSet[T]) Union(other Set[T]) *MutableSet[T] {
	if other == nil {
		return s.Clone()
	}

	union := s.Clone()

	for _, k := range other.ToSlice() {
		union.Add(k)
	}
	return union
}

func (s *MutableSet[T]) Except(other Set[T]) *MutableSet[T] {
	if other == nil {
		return s.Clone()
	}

	except := New[T]()
	for k := range s.m {
		if !other.Contains(k) {
			except.Add(k)
		}
	}
	return except
}

func (s *MutableSet[T]) Intersect(other Set[T]) *MutableSet[T] {
	if other == nil {
		return New[T]()
	}

	intersect := New[T]()
	for k := range s.m {
		if other.Contains(k) {
			intersect.Add(k)
		}
	}
	return intersect
}

func (s *MutableSet[T]) IsSubset(other Set[T]) bool {
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

func (s *MutableSet[T]) IsSuperset(other Set[T]) bool {
	if other == nil {
		return false
	}

	return other.IsSubset(s)
}

func (s *MutableSet[T]) IsProperSuperset(other Set[T]) bool {
	if other == nil {
		return false
	}

	return s.Size() > other.Size() && s.IsSuperset(other)
}

func (s *MutableSet[T]) IsProperSubset(other Set[T]) bool {
	if other == nil {
		return false
	}

	return other.IsProperSuperset(s)
}

func (s *MutableSet[T]) Equals(other Set[T]) bool {
	if other == nil {
		return false
	}

	return s.IsSubset(other) && other.IsSubset(s)
}

func (s *MutableSet[T]) String() string {
	str := "{"
	for k := range s.m {
		str += fmt.Sprintf("%v, ", k)
	}
	str = str[:len(str)-2]
	str += "}"
	return str
}
