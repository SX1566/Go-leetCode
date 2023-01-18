package main

import (
	ADD "Demo5-import/add"
	//直接使用
	. "Demo5-import/sub"
	"fmt"
)

func main() {
	fmt.Println(ADD.Add(1, 6))
	fmt.Println(Sub(0, 10000))
}
