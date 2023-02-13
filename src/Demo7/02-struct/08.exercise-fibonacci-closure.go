package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func(int) int {
	return func(x int) int {
		if x == 0 {
			return 0
		} else if x <= 2 {
			return 1
		} else if x > 2 {
			a, b := 1, 1
			result := 0
			for i := 3; i <= x; i++ {
				result = a + b
				a, b = b, result
			}
			return result
		} else {
			return -1
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f(i))
	}
}
