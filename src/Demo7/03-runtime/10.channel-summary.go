package main

import "fmt"

func main() {
	/*
	   	1. 当管道写满了，写阻塞

	      2. 当缓冲区读完了，读阻塞

	      3. 如果管道没有使用make分配空间，管道默认是nil，从nil的管道读取数据、写入数据，都会阻塞（注意，不会崩溃）

	      4. 从一个已经close的管道读取数据时，会返回零值（不会崩溃）  《=== 验证

	      5. 向一个已经close的管道写数据时，会崩溃   《== 验证

	      6. 关闭一个已经close的管道，程序会崩溃  《== 验证

	      7. 关闭管道的动作，一定要在写端，不应该放在读端，否则写的继续写会崩溃

	      8. 读和写的次数，一定对等，否则：

	         1. 在多个go程中：资源泄露
	         2. 在主go程中，程序崩溃（deadlock)

	*/

	/*
		## ==三个场景使用make==

		切片，map，管道

		切片：

		```go
		var names []string
		names = make([]string, 长度, 容量)
		```

		字典/map/哈希表：

		```go
		var idNames map[int]string
		idNames = make(map[int]string, 容量)
		```

		管道：

		```go
		var numChan chan int
		numChan = make(chan int, 容量)
		```



	*/

	//判断管道是否已经关闭
	numChan := make(chan int, 10)

	//写
	go func() {
		for i := 0; i < 50; i++ {
			numChan <- i
			fmt.Println("写入数据：", i)
		}
		close(numChan) //写完之后记得关
	}()

	//读数据
	//for v := range numChan {
	//	fmt.Println("v:", v)
	//}

	for {
		v, ok := <-numChan
		if !ok {
			fmt.Println("管道已经关闭，准备退出!")
			break
		}

		fmt.Println("v", v)
	}

	fmt.Println("Over!")
}
