package main

import (
	"fmt"
	"sync"
)

// Ensures only one routine will work on piece of code at a time
var counter int
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Print("final counter: ", counter)
}
