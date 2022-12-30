package main

import "fmt"

func main() {
	name := "111"

	//需要换行，原生输出字符串使用反引号``
	usage := `./a.out <option>
		-h help
		-a xxxx`
	fmt.Println(name)
	fmt.Println(usage)

	//2. 长度，访问
	//name.length C++
	//len()  GO
	length := len(name)
	fmt.Println(length)

	for i := 0; i < len(name); i++ {
		fmt.Printf("i: %d, v v: %c\n", i, name[i])
	}

	//拼接
	i, j := "hello", "world"
	fmt.Println(i + j)

}
