package main

// https://gobyexample-cn.github.io/channel-synchronization
import (
	"fmt"
	"time"
)

func worker1(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker1(done)

	<-done
}
