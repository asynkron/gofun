package enumerable

import "testing"

func TestSliceEnumerableCanEnumerateAllElements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	enumerable := FromSlice(slice)
	enumerator := enumerable.GetEnumerator()

	for i := 0; i < len(slice); i++ {
		item, ok := enumerator.MoveNext()
		if !ok {
			t.Fail()
		}
		if item != i+1 {
			t.Fail()
		}
	}
}
