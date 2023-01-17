package main

func main() {
	p1 := testPtr1()
	print(*p1)

}

func testPtr1() *string {
	city := "北京"
	ptr := &city
	return ptr

}
