package main

import (
	"fmt"
	"sync"
)

func printnums(n int, chnn chan int, wg *sync.WaitGroup, even bool) {
	defer wg.Done()

	for i := 1; i <= n; i++ {
		if even && i%2 == 0 {
			chnn <- i
		} else if !even && i%2 != 0 {
			chnn <- i
		}
	}
	close(chnn)
}

func main() {
	echn := make(chan int)
	ochn := make(chan int)
	var wg sync.WaitGroup
	wg.Wait()
	wg.Add(1)
	go printnums(10, echn, &wg, true)
	wg.Add(1)
	go printnums(10, ochn, &wg, false)
	for {
		select {
		case even, ok := <-echn:
			if ok {
				fmt.Println("EVEN CHANNEL : ", even)
			} else {
				echn = nil
			}
		case odd, ok := <-ochn:
			if ok {
				fmt.Println("ODD CHANNEL : ", odd)
			} else {
				ochn = nil
			}
		}
		if echn == nil && ochn == nil {
			break
		}
	}
}
