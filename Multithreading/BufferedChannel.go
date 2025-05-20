package main

import (
	"fmt"
	"time"
)

// allows sending data to the channel without immediately blocking the sender, as
// long as there is space available in the buffer.

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("producer ", i)
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
func consumer(ch chan int) {
	for num := range ch {
		fmt.Println("consumer ", num)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ch := make(chan int, 3)
	go producer(ch)
	go consumer(ch)
	time.Sleep(10 * time.Second)
}
