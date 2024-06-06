package main

import (
	"fmt"
)

func main() {

	var i = []int{1, 2, 3, 4, 5}
	var j = i[:3]
	j = append(j, 6)
	j = append(j, 7)
	j = append(j, 8)
	j = append(j, 9)
	fmt.Println(i)
	fmt.Println(j)
}
