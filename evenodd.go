package main

import (
	"fmt"
)

var pl = fmt.Println

func printEven(n int, channel chan<- int) {
	for i := 2; i <= n; i = i + 2 {
		channel <- i
	}
	close(channel)
}

func printOdd(n int, channel chan<- int) {
	for i := 1; i <= n; i = i + 2 {
		channel <- i
	}
	close(channel)
}

func main() {
	n := 10
	evenChan := make(chan int)
	oddChan := make(chan int)

	go printEven(n, evenChan)
	go printOdd(n, oddChan)
	for {
		select {
		case evenno, ok := <-evenChan:
			if ok {
				pl(evenno)
			} else {
				evenChan = nil
			}
		case oddno, ok := <-oddChan:
			if ok {
				pl(oddno)
			} else {
				oddChan = nil
			}
		}
		if evenChan == nil && oddChan == nil {
			break
		}

	}
	// for value1 := range evenChan {
	// 	pl(value1)
	// }
	// for value2 := range oddChan {
	// 	pl(value2)
	// }
}
