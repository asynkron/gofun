package set

import (
	"fmt"
	"golang.org/x/exp/maps"
)

type ImmutableSet[T comparable] struct {
	m map[T]bool
}

func NewImmutable[T comparable](items ...T) *ImmutableSet[T] {
	d := make(map[T]bool)
	for _, item := range items {
		d[item] = true
	}
	return &ImmutableSet[T]{m: d}
}

func (s *ImmutableSet[T]) Add(e T) *ImmutableSet[T] {
	c := s.Clone()
	c.m[e] = true
	return c
}

func (s *ImmutableSet[T]) AddRange(e ...T) *ImmutableSet[T] {
	c := s.Clone()
	for _, i := range e {
		c.m[i] = true
	}

	return c
}

func (s *ImmutableSet[T]) Remove(e T) *ImmutableSet[T] {
	c := s.Clone()
	delete(c.m, e)
	return c
}

func (s *ImmutableSet[T]) Contains(e T) bool {
	_, ok := s.m[e]
	return ok
}

func (s *ImmutableSet[T]) Size() int {
	return len(s.m)
}

func (s *ImmutableSet[T]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *ImmutableSet[T]) ToSlice() []T {
	return maps.Keys(s.m)
}

func (s *ImmutableSet[T]) Clone() *ImmutableSet[T] {
	return &ImmutableSet[T]{m: maps.Clone(s.m)}
}

func (s *ImmutableSet[T]) Union(other Set[T]) *ImmutableSet[T] {
	if other == nil {
		return s.Clone()
	}

	union := s.Clone()
	for _, k := range other.ToSlice() {
		union = union.Add(k)
	}
	return union
}

func (s *ImmutableSet[T]) Except(other Set[T]) *ImmutableSet[T] {
	if other == nil {
		return s.Clone()
	}

	except := NewImmutable[T]()
	for k := range s.m {
		if !other.Contains(k) {
			except = except.Add(k)
		}
	}
	return except
}

func (s *ImmutableSet[T]) Intersect(other Set[T]) *ImmutableSet[T] {
	if other == nil {
		return NewImmutable[T]()
	}

	intersect := NewImmutable[T]()
	for k := range s.m {
		if other.Contains(k) {
			intersect = intersect.Add(k)
		}
	}
	return intersect
}

func (s *ImmutableSet[T]) ImmutableSet(other Set[T]) bool {
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

func (s *ImmutableSet[T]) IsSubset(other Set[T]) bool {
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

func (s *ImmutableSet[T]) IsSuperset(other Set[T]) bool {
	if other == nil {
		return false
	}

	return other.IsSubset(s)
}

func (s *ImmutableSet[T]) IsProperSuperset(other Set[T]) bool {
	if other == nil {
		return false
	}

	return s.Size() > other.Size() && s.IsSuperset(other)
}

func (s *ImmutableSet[T]) IsProperSubset(other Set[T]) bool {
	if other == nil {
		return false
	}

	return other.IsProperSuperset(s)
}

func (s *ImmutableSet[T]) Equals(other Set[T]) bool {
	if other == nil {
		return false
	}

	return s.IsSubset(other) && other.IsSubset(s)
}

func (s *ImmutableSet[T]) String() string {
	str := "{"
	for k := range s.m {
		str += fmt.Sprintf("%v, ", k)
	}
	str = str[:len(str)-2]
	str += "}"
	return str
}
