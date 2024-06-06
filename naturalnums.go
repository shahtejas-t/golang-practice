package main

import (
	"fmt"
	"sync"
)

var pl = fmt.Println

func naturalNums(n int, wg *sync.WaitGroup, channel chan<- int) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		channel <- i
	}
	// close(channel)
}

func main() {
	var wg sync.WaitGroup
	numChan := make(chan int)
	wg.Add(1)
	go naturalNums(5, &wg, numChan)

	go func() {
		wg.Wait()
		close(numChan)
	}()

	for n := range numChan {
		pl(n)
	}
}
