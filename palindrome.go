package main

import (
	"fmt"
	"sync"
)

var pl = fmt.Println

func palindrome(text string, chnn chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	r := []rune(text)
	l := len(r)
	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		if r[i] != r[j] {
			chnn <- false
		}
	}
	chnn <- true
	close(chnn)
}

func main() {
	text := "hi"
	chnn := make(chan bool)
	var wg sync.WaitGroup
	wg.Wait()
	wg.Add(1)
	go palindrome(text, chnn, &wg)
	res := <-chnn
	pl(res)
}
