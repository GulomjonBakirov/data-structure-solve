package array

import "fmt"

func insertionSort(nums []int) []int {

	for j := 2; j < len(nums); j++ {
		key := nums[j]

		i := j - 1

		for i > 0 || nums[i] > key {
			nums[i+1] = nums[i]
			i = i - 1
		}
		nums[i+1] = key

	}

	return nums
}

func ContainsDuplicate(nums []int) {
	fmt.Println("Unsorted nums: ", nums)
	nums2 := insertionSort(nums)
	fmt.Println("Sorted nums: ", nums2)
}
