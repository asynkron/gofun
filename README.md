# gofun

## Generic Enumerables

Lazy generic enumerables for Go.

```go

//create an enumerable of strings
v1 := enumerable.From("a", "b", "c", "d", "e", "f", "hello", "world", "hi", "hoi", "howdy")

//filter the enumerable to only include strings that start with "h"
v2 := enumerable.Filter(v1, func(s string) bool {
    return strings.HasPrefix(s, "h")
})

//map the enumerable of string to int, collect the string length as the new value
v3 := enumerable.Map[string, int](v2, func(s string) int {
    return len(s)
})

//skip the first 3 elements
v4 := enumerable.Skip(v3, 3)

//then take the next 3 elements
v5 := enumerable.Limit(v4, 3)

//iterate the enumerable and create a slice of ints
v6 := enumerable.ToSlice(v5)

```

Clone(ish) of .NET LINQ

## Generic Set

Typesafe Set type.

## Generic Option type
