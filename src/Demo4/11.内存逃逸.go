package main

import "fmt"

func main() {
	p1 := testPtr1()
	fmt.Println(p1)

}

func testPtr1() *string {
	name := "aaa"
	p0 := &name
	fmt.Println("p0: ", p0)

	city := "北京"
	ptr := &city
	return ptr

}
