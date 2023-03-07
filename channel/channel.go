package main

import (
	"fmt"
	"time"
)

func channelDemo() {
	//c := make(chan int)
	// 读取数据
	//go func() {
	//	for {
	//		n := <-c
	//		fmt.Print(n)
	//	}
	//}()

	// 函数调用
	//go worker(1, c)
	//
	//// 往 channel 发数据
	//c <- 1
	//c <- 2

	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

// channel 作为参数
func worker(id int, c chan int) {
	for {
		//n := <-c
		//fmt.Print(reflect.TypeOf(<-c))
		fmt.Printf("游客 %d 接受了 %c\n", id, <-c)
	}
}

func main() {
	channelDemo()
}
