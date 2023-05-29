package main

func main() {
	myClient := &Myclient{}
	myClient.InitClient("127.0.0.1:8800")
	var resp string
	err := myClient.HelloWorld("test", &resp)
	if err != nil {
		println("HelloWorld err:", err)
		return
	}
	println(resp, err)

}
