package main

import (
	"fmt"
	"sync"
)

var pl = fmt.Println

func generateTable(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		if i%5 == 0 {
			ch <- i
		}
	}
	close(ch)
}

func main() {
	n := 20
	g1Chn := make(chan int)
	g2Chn := make(chan int)
	var wg sync.WaitGroup

	// Spawn goroutines
	wg.Add(2) // Increment the WaitGroup counter for two goroutines
	go generateTable(n, g1Chn, &wg)
	go generateTable(n, g2Chn, &wg)

	// Wait for both goroutines to finish
	go func() {
		wg.Wait()
		close(g1Chn)
		close(g2Chn)
	}()

	// Process values from channels
	for {
		select {
		case value1, ok := <-g1Chn:
			if !ok {
				pl("g1 channel closed")
				g1Chn = nil
			} else {
				pl("g1:", value1)
			}
		case value2, ok := <-g2Chn:
			if !ok {
				pl("g2 channel closed")
				g2Chn = nil
			} else {
				pl("g2:", value2)
			}
		default:
			// Exit the loop when both channels are closed
			if g1Chn == nil && g2Chn == nil {
				return
			}
		}
	}
}
