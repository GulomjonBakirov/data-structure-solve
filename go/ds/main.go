package main

import (
	"fmt"
	"leetcode-ds/array"
)

func main() {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1

	// result := array.MaxSubArray(nums)
	// array.ContainsDuplicate(nums)

	// result := array.TwoSum(nums, target)
	result := array.Merge(nums1, m, nums2, n)
	fmt.Println("Solve: ", result)

}
