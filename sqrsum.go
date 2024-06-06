package main

import (
	"fmt"
	"sync"
)

var pl = fmt.Println

func sumsqr(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	ch <- sum
}

func main() {
	n := 6
	ch := make(chan int)
	var wg sync.WaitGroup

	defer wg.Wait()

	wg.Add(1)
	go sumsqr(n, ch, &wg)
	pl(<-ch)
}
