package main

import (
	"fmt"
)

var pl = fmt.Println

func rsum(nums []int) []int {
	sums := make([]int, len(nums))
	if len(nums) > 0 {
		sums[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			sums[i] = sums[i-1] + nums[i]
		}
	}
	return sums
}

func main() {
	var num = []int{1, 2, 3, 4, 5}
	sums := rsum(num)
	pl(sums)

}
