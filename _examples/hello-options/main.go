package main

import (
	"fmt"
	"github.com/asynkron/gofun/options"
)

func main() {
	s := options.None[string]()
	fmt.Println(options.IsSome(s))
	options.Match(s,
		func(v string) {
			fmt.Println(v)
		},
		func() {
			fmt.Println("None")
		})
}
