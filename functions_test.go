package main

import "testing"

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

}

func TestElementAtOrDefault(t *testing.T) {

}

func TestFilter(t *testing.T) {

}

func TestFirstOrDefault(t *testing.T) {

}

func TestLastOrDefault(t *testing.T) {

}

func TestLimit(t *testing.T) {

}

func TestMap(t *testing.T) {

}

func TestAll(t *testing.T) {

}

func TestAny(t *testing.T) {

}

func TestConcat(t *testing.T) {

}

func TestContains(t *testing.T) {

}

func TestDistinctBy(t *testing.T) {

}

func TestExcept(t *testing.T) {

}
