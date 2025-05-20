package main

import (
	"fmt"
	"sync"
)

// helps you synchronize the completion of multiple goroutines by providing a
// mechanism to track when all of them have finished.
func display(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go display(&wg)
	fmt.Println("hello world")
	wg.Wait()
}
