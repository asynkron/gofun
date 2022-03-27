package main

import (
	"fmt"
	"strings"
)

func main() {
	v1 := From("hello", "foo", "bar", "hello", "hellbar", "helloworld", "a", "b", "c", "d", "e")
	v2 := Filter(v1, func(s string) bool {
		return strings.HasPrefix(s, "hell")
	})
	v3 := Map[string, int](v2, func(s string) int {
		return len(s)
	})

	v4 := Skip(v3, 0)
	v5 := Limit(v4, 2)
	v6 := ToSlice(v5)

	for _, v := range v6 {
		fmt.Printf("res %v\n", v)
	}

	m := ToMapOfSlice(v1, func(s string) int {
		return len(s)
	})

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}
