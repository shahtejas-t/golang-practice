package main

import (
	"fmt"
	"sync"
)

var pl = fmt.Println

func nnums(n int, chnn chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i++ {
		chnn <- i
	}
	close(chnn)
}

func main() {
	chnn := make(chan int)
	var wg sync.WaitGroup

	wg.Wait()
	wg.Add(1)

	go nnums(10, chnn, &wg)

	for n := range chnn {
		pl(n)
	}

}
