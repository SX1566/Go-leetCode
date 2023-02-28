package main

import "fmt"

func main() {
	var i interface{}
	describe1(i)

	i = 42
	describe1(i)

	i = 12301
	describe1(i)
}

func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
