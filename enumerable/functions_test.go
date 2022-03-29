package enumerable

import (
	"fmt"
	"testing"
)

func TestToMapOfSlice(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10}
	e := FromSlice(values)
	m := ToMapOfSlice(e, func(i int) int {
		return i
	})
	if len(m) != 10 {
		t.Errorf("Expected 10, got %d", len(m))
	}
}

func TestMin(t *testing.T) {
	values := []int{1, 2, -3, 4, 5, 6, 7, 8, 9, 99, 10, 10, 10}
	e := FromSlice(values)
	min := Min(e)
	if min != -3 {
		t.Errorf("Expected -3, but got %d", min)
	}
}

func TestSum(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	e := FromSlice(values)
	sum := Sum(e)
	if sum != 85 {
		t.Errorf("Expected 85, got %d", sum)
	}
}

func TestMax(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 99, 10, 10, 10}
	e := FromSlice(values)
	max := Max(e)
	if max != 99 {
		t.Errorf("Expected 99, but got %d", max)
	}
}

func TestAggregate(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	e := FromSlice(values)
	sum := Aggregate(e, 0, func(a, b int) int {
		return a + b
	})
	if sum != 85 {
		t.Errorf("Expected 85, got %d", sum)
	}
}

func TestChunk(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	e := FromSlice(values)
	chunks := Chunk(e, 3)
	count := Count(chunks)
	if count != 5 {
		t.Errorf("Expected 5 chunks, got %d", count)
	}
}

func TestAvg(t *testing.T) {
	values := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	e := FromSlice(values)
	avg := Avg(e)
	if avg != 5.5 {
		t.Error("Expected 5.5, but got ", avg)
	}
}

func TestToSet(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	e := FromSlice(values)
	s := ToSet(e)
	if s.Size() != 10 {
		t.Error("Expected 10 but got ", s.Size())
	}
}

func TestExcept(t *testing.T) {
	values1 := []string{"a", "b", "c", "d", "a", "a", "a"}
	values2 := []string{"a", "y", "z"}
	e1 := FromSlice(values1)
	e2 := FromSlice(values2)
	c := Except(e1, e2)
	count := Count(c)
	if count != 3 { // b, c, d
		t.Errorf("Expected 4, got %d", count)
	}
}

//tests the functions in functions.go
func TestDistinct(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	d := Distinct(e)
	c := Count(d)
	if c != 4 {
		t.Errorf("Expected 4, got %d", c)
	}
}

func TestToSlice(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	s := ToSlice(e)
	//assert that values and s are equal
	for i := 0; i < len(values); i++ {
		if values[i] != s[i] {
			t.Errorf("Expected %s, got %s", values[i], s[i])
		}
	}
}

func TestCount(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	c := Count(e)
	if c != 7 {
		t.Errorf("Expected 7, got %d", c)
	}
}

func TestElementAtOrDefault(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	s := ElementAtOrDefault(e, 5, "default")
	if s != "a" {
		t.Errorf("Expected a, got %s", s)
	}
}

func TestFilter(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	f := Filter(e, func(s string) bool {
		return s == "a"
	})
	c := Count(f)
	if c != 4 {
		t.Errorf("Expected 3, got %d", c)
	}
}

func TestFirstOrDefault(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	s := FirstOrDefault(e, "default")
	if s != "a" {
		t.Errorf("Expected a, got %s", s)
	}
}

func TestLastOrDefault(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "x"}
	e := FromSlice(values)
	s := LastOrDefault(e, "default")
	if s != "x" {
		t.Errorf("Expected x, got %s", s)
	}
}

func TestLimit(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	l := Limit(e, 3)
	c := Count(l)
	if c != 3 {
		t.Errorf("Expected 3, got %d", c)
	}
}

func TestMap(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	m := Map(e, func(s string) string {
		return s + "!"
	})
	//assert that values and m are equal but m also has the ! appended
	i := 0
	m.Enumerate(func(s string) bool {
		expected := values[i] + "!"
		if expected != s {
			t.Errorf("Expected %s, got %s", expected, s)
		}
		i++
		return YieldContinue
	})
}

func TestAll(t *testing.T) {
	values := []string{"a", "a", "a", "a"}
	e := FromSlice(values)
	a := All(e, func(s string) bool {
		return s == "a"
	})
	if !a {
		t.Errorf("Expected true, got %t", a)
	}
}

func TestAny(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	a := Any(e, func(s string) bool {
		return s == "a"
	})
	if !a {
		t.Errorf("Expected true, got %t", a)
	}
}

func TestConcat(t *testing.T) {
	values1 := []string{"a", "b", "c", "d", "a", "a", "a"}
	values2 := []string{"x", "y", "z"}
	e1 := FromSlice(values1)
	e2 := FromSlice(values2)
	c := Concat(e1, e2)

	s := ToSlice(c)
	fmt.Printf("%v", s)

	f := FirstOrDefault(c, "default")
	if f != "a" {
		t.Errorf("Expected a got %s", f)
	}
	l := LastOrDefault(c, "default")
	if l != "z" {
		t.Errorf("Expected z got %s", l)
	}
}

func TestContains(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	c := Contains(e, "a")
	if !c {
		t.Errorf("Expected true, got %t", c)
	}
}

func TestDistinctBy(t *testing.T) {
	values := []string{"a", "b", "c", "d", "a", "a", "a"}
	e := FromSlice(values)
	d := DistinctBy(e, func(s string) string {
		return s
	})
	c := Count(d)
	if c != 4 {
		t.Errorf("Expected 4, got %d", c)
	}
}
