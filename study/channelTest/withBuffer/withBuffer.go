package main

import (
	"fmt"
	"time"
)

func Send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Send ready", i)
		c <- i
		fmt.Println("Send", i)
	}
}

func Recv(c <-chan int) {
	for i := range c {
		fmt.Println("Received", i)
	}
}

func main() {
	c := make(chan int, 10)
	go Send(c)
	go Recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}
