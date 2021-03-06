package set

import "testing"

func TestSet_IsSuperset(t *testing.T) {
	s := New[string]("a", "b", "c", "d")

	if !s.IsSuperset(New[string]()) {
		t.Errorf("Expected true, got false")
	}

	if !s.IsSuperset(New[string]("a")) {
		t.Errorf("Expected true, got false")
	}

	if !s.IsSuperset(New("a", "b")) {
		t.Errorf("Expected true, got false")
	}

	if !s.IsSuperset(New("a", "b", "c")) {
		t.Errorf("Expected true, got false")
	}

	if !s.IsSuperset(New("a", "b", "c", "d")) {
		t.Errorf("Expected true, got false")
	}

	if s.IsSuperset(New("a", "b", "c", "d", "e")) {
		t.Errorf("Expected false, got true")
	}
}

func TestSet_Remove(t *testing.T) {
	s := New[string]("a", "b", "c", "d")

	s.Remove("a")
	if s.Contains("a") {
		t.Errorf("Expected false, got true")
	}

	s.Remove("b")
	if s.Contains("b") {
		t.Errorf("Expected false, got true")
	}

	s.Remove("c")
	if s.Contains("c") {
		t.Errorf("Expected false, got true")
	}

	s.Remove("d")
	if s.Contains("d") {
		t.Errorf("Expected false, got true")
	}
}

func TestSet_Size(t *testing.T) {
	s := New[string]("a", "b", "c", "d")

	if s.Size() != 4 {
		t.Errorf("Expected 4, got %d", s.Size())
	}

	s.Remove("a")
	if s.Size() != 3 {
		t.Errorf("Expected 3, got %d", s.Size())
	}

	s.Remove("b")
	if s.Size() != 2 {
		t.Errorf("Expected 2, got %d", s.Size())
	}

	s.Remove("c")
	if s.Size() != 1 {
		t.Errorf("Expected 1, got %d", s.Size())
	}

	s.Remove("d")
	if s.Size() != 0 {
		t.Errorf("Expected 0, got %d", s.Size())
	}
}

func TestSet_String(t *testing.T) {
	s := New[string]("a", "b", "c", "d")

	if s.String() != "{a, b, c, d}" {
		t.Errorf("Expected {a, b, c, d}, got %s", s.String())
	}
}

func TestSet_ToSlice(t *testing.T) {
	s := New[string]("a", "b", "c", "d")
	slice := s.ToSlice()

	if len(slice) != 4 {
		t.Errorf("Expected 4, got %d", len(slice))
	}
}

func TestSet_TryAdd(t *testing.T) {
	s := New[string]("a", "b", "c", "d")

	if s.TryAdd("a") {
		t.Errorf("Expected false, got true")
	}

	if !s.TryAdd("e") {
		t.Errorf("Expected true, got false")
	}
}

func TestSet_TryRemove(t *testing.T) {
	s := New[string]("a", "b", "c", "d")

	if !s.TryRemove("a") {
		t.Errorf("Expected true, got false")
	}

	if s.TryRemove("e") {
		t.Errorf("Expected false, got true")
	}
}

func TestSet_Union(t *testing.T) {
	s := New[string]("a", "b", "c", "d")

	s2 := New[string]("a", "b", "c", "e")

	s3 := s.Union(s2)
	if s3.Size() != 5 {
		t.Errorf("Expected 5, got %d", s3.Size())
	}

	if !s3.Contains("a") {
		t.Errorf("Expected true, got false")
	}

	if !s3.Contains("b") {
		t.Errorf("Expected true, got false")
	}

	if !s3.Contains("c") {
		t.Errorf("Expected true, got false")
	}

	if !s3.Contains("d") {
		t.Errorf("Expected true, got false")
	}

	if !s3.Contains("e") {
		t.Errorf("Expected true, got false")
	}
}

func TestNew(t *testing.T) {
	s := New[string]()
	if s.Size() != 0 {
		t.Errorf("New set should be empty")
	}
}

func TestSet_Add(t *testing.T) {
	s := New[string]("a", "b", "c")
	if s.Size() != 3 {
		t.Errorf("MutableSet should have 3 elements")
	}
}

func TestSet_Clear(t *testing.T) {
	s := New[string]("a", "b", "c")
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("MutableSet should be empty")
	}
}

func TestSet_Clone(t *testing.T) {
	s := New[string]("a", "b", "c")
	s2 := s.Clone()
	s.Clear()
	if s2.Size() != 3 {
		t.Errorf("MutableSet should have 3 elements")
	}
}

func TestSet_Contains(t *testing.T) {
	s := New[string]("a", "b", "c")
	if !s.Contains("a") {
		t.Errorf("MutableSet should contain 'a'")
	}
	if s.Contains("d") {
		t.Errorf("MutableSet should not contain 'd'")
	}
}

func TestSet_Equals(t *testing.T) {
	s := New[string]("a", "b", "c")
	s2 := New[string]("a", "b", "c")
	if !s.Equals(s2) {
		t.Errorf("Sets should be equal")
	}
	s2.Clear()
	s2.Add("d")
	if s.Equals(s2) {
		t.Errorf("Sets should not be equal")
	}
}

func TestSet_Except(t *testing.T) {
	s := New[string]("a", "b", "c")
	s2 := New[string]("a", "b", "c")
	s3 := s.Except(s2)
	if s3.Size() != 0 {
		t.Errorf("MutableSet should be empty")
	}

	s.Add("e")
	s4 := s.Except(s2)
	if s4.Size() != 1 {
		t.Errorf("MutableSet should contain 1 item")
	}

}

func TestSet_Intersect(t *testing.T) {
	s := New[string]("a", "b", "c", "d")
	s2 := New[string]("a", "b", "c", "e")
	s3 := s.Intersect(s2)
	if s3.Size() != 3 {
		t.Errorf("MutableSet should have 3 elements")
	}

}

func TestSet_IsEmpty(t *testing.T) {
	s := New[string]()
	if !s.IsEmpty() {
		t.Errorf("MutableSet should be empty")
	}
	s.Add("a")
	if s.IsEmpty() {
		t.Errorf("MutableSet should not be empty")
	}
}

func TestSet_IsProperSubset(t *testing.T) {
	s := New[string]("a", "b", "c")
	s2 := New[string]("a", "b", "c")
	s3 := s.IsProperSubset(s2)
	if s3 {
		t.Errorf("MutableSet should not be a proper subset")
	}
	s2.Add("d")
	s3 = s.IsProperSubset(s2)
	if !s3 {
		t.Errorf("MutableSet should be a proper subset")
	}
}

func TestSet_IsProperSuperset(t *testing.T) {
	s := New[string]("a", "b", "c")
	s2 := New[string]("a", "b", "c")
	s3 := s.IsProperSuperset(s2)
	if s3 {
		t.Errorf("MutableSet should not be a proper superset")
	}
	s.Add("d")
	s3 = s.IsProperSuperset(s2)
	if !s3 {
		t.Errorf("MutableSet should be a proper superset")
	}
}

func TestSet_IsSubset(t *testing.T) {
	s := New[string]("a", "b", "c")
	s2 := New[string]("a", "b", "c")
	s3 := s.IsSubset(s2)
	if !s3 {
		t.Errorf("MutableSet should be a subset")
	}

}
