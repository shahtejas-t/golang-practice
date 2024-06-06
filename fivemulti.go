package main

import (
	"fmt"
	"sync"
)

var pl = fmt.Println

func fivemulti(n int, ch chan int, wg *sync.WaitGroup, a bool) {
	defer wg.Done()
	if a {
		for i := 1; i <= n/2; i++ {
			ch <- (i * 10)
		}
		close(ch)
	} else {
		for i := 1; i <= n; i++ {
			if (i*5)%10 != 0 {
				ch <- (i * 5)
			}
		}
		close(ch)
	}
}

func main() {
	chn1 := make(chan int)
	chn2 := make(chan int)
	var wg sync.WaitGroup
	defer wg.Wait()
	n := 20

	wg.Add(1)
	go fivemulti(n, chn1, &wg, false)

	wg.Add(1)
	go fivemulti(n, chn2, &wg, true)

	// for num := range chn1 {
	// 	pl("chn1 : ", num)
	// }

	// for num := range chn2 {
	// 	pl("chn2 : ", num)
	// }

	for {
		select {
		case val1, ok := <-chn1:
			if ok {
				pl("chn1 : ", val1)
			} else {
				chn1 = nil
			}

		case val2, ok := <-chn2:
			if ok {
				pl("chn2 : ", val2)
			} else {
				chn2 = nil
			}
		default:
			if chn1 == nil && chn2 == nil {
				return
			}
		}
	}
}
