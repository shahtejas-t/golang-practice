package main

import "fmt"

var pl = fmt.Println

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		secondNum := target - num
		if index, found := numMap[secondNum]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return nil
}

func main() {
	var nums = []int{2, 0, 5, 7}
	var target = 9
	pl(twoSum(nums, target))
}
