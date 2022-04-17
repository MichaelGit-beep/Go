package main

import "fmt"

type ints interface{}

func main() {
	a := ints(2)

	fmt.Printf("%v", a)
}
