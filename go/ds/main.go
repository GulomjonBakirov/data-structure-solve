package main

import (
	"fmt"
	"leetcode-ds/array"
)

func main() {

	// result := array.MaxSubArray(nums)
	// array.ContainsDuplicate(nums)

	// result := array.TwoSum(nums, target)
	result := array.Intersect([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println("Solve: ", result)

}
