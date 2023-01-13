package main

import "fmt"

// 函数返回值在参数列表之后
func test(a int, b int, c string) (int, string) {
	fmt.Println("11")
	return a * b, c

}

func main() {
	test(1, 2, "3")
	fmt.Println("123")
}
