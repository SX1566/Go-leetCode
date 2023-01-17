package main

import "fmt"

// 函数返回值在参数列表之后，多个参数用 ,分割
func test(a int, b int, c string) (int, string) {
	return a * b, c

}

// 如果返回值只有一个参数并且没有名字，那么不需要加()
func test2(a, b int, c string) (res int, str string, bl bool) {
	//直接使用返回值的变量名字
	res = a * b
	str = c
	bl = false

	//当返回值有名字的时候直接用return
	return

}

// 如果返回值只有一个参数并且没有名字，那么不需要加()
func test3() int {
	return 100
}

func main() {
	v1, s1 := test(1, 30, "321312")
	fmt.Println(v1)
	fmt.Println(s1)

	v2, s2, _ := test2(1, 30, "321312")
	fmt.Println(v2)
	fmt.Println(s2)
}
