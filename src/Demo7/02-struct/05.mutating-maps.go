package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["A"] = 48
	fmt.Println(m["A"])

	m["A"] = 22
	fmt.Println(m["A"])

	delete(m, "A")
	fmt.Println(m["A"])

	v, ok := m["A"]
	fmt.Println("The value:", v, "Present?", ok)
}
