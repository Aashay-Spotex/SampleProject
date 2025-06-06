package main

import "fmt"

// Example of Unbuffered Channel the sender will block until the receiver is
// ready to receive the data

func printDisplay(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go printDisplay(ch)
	for num := range ch {
		fmt.Println(num)
	}
	fmt.Println("channel Done")
}
