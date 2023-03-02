package main

import "fmt"

func main() {
	//只会读前几个字符

	ch := make(chan int, 6)
	ch <- 1
	ch <- 2
	ch <- 2
	ch <- 2
	ch <- 2
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
