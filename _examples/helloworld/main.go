package main

import (
	"fmt"
	"github.com/asynkron/gofun/enumerable"
	"github.com/asynkron/gofun/options"
	"strings"
)

func main() {

	//var n *options.Option[string] = nil
	//fmt.Println(options.IsSome(n))
	//s := options.Some("Hello World")
	//fmt.Println(options.IsSome(s))

	//how to handle?
	s2 := options.None[string]()
	fmt.Println(options.IsSome(s2))
	options.Match(s2,
		func(v string) {
			fmt.Println(v)
		},
		func() {
			fmt.Println("None")
		})

	v1 := enumerable.From("hello", "foo", "bar", "hello", "hellbar", "helloworld", "a", "b", "c", "d", "e")
	v22 := enumerable.Sum(v1)

	fmt.Println(v22)
	v2 := enumerable.Filter(v1, func(s string) bool {
		return strings.HasPrefix(s, "hell")
	})
	v3 := enumerable.Map[string, int](v2, func(s string) int {
		return len(s)
	})

	v4 := enumerable.Skip(v3, 0)
	v5 := enumerable.Limit(v4, 2)
	v6 := enumerable.ToSlice(v5)

	for _, v := range v6 {
		fmt.Printf("res %v\n", v)
	}

	m := enumerable.ToMapOfSlice(v1, func(s string) int {
		return len(s)
	})

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}
