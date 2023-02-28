package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("%v (%v years)", s.Name, s.Age)
}

func main() {
	a := Student{"Arthur Dent", 42}
	z := Student{"Zap hod Bumblebee", 9001}

	fmt.Println(a, z)
}
