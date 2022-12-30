package main

import "fmt"

func main() {

	name := "name"
	ptr := &name
	fmt.Println("name ptr", ptr, " ==>指针")

	name2Ptr := new(string)
	*name2Ptr = "name2"
	ptr2 := &name2Ptr
	fmt.Println(ptr2)

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
