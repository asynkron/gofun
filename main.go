package main

import "fmt"

func main() {
	v1 := From("hello", "foo", "bar", "hello")
	v2 := Filter(v1, func(s string) bool {
		return s == "hello"
	})
	v3 := Map[string, int](v2, func(s string) int {
		return len(s)
	})
	res := v3.ToSlice()

	fmt.Printf("%v", res)
}
