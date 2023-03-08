package main

import "fmt"

// 非阻塞通道操作
// 常规的通过通道发送和接收数据是阻塞的。 然而，我们可以使用带一个 default 子句的 select 来实现 非阻塞 的发送、接收，
// 甚至是非阻塞的多路 select。
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "h1"
	select {
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case meg := <-messages:
		fmt.Println("received message:", meg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
