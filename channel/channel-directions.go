package main

import "fmt"

// ping 函数定义了一个只能发送数据的（只写）通道。 尝试从这个通道接收数据会是一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数接收两个通道，pings 仅用于接收数据（只读），pongs 仅用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// 建立通道缓冲为1
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	fmt.Println("ping")
	ping(pings, "passed message")

	fmt.Println("pong")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
