package array

import "fmt"

func MaxSubArray(nums []int) int {
	maxSum := 0
	maxFar := nums[0]

	for _, el := range nums {
		maxSum = maxSum + el

		fmt.Println("Max sum before: ", maxSum)
		fmt.Println("Max far before: ", maxFar)
		fmt.Println("Element: ", el)

		if maxFar < maxSum {
			maxFar = maxSum
		}

		if maxSum < 0 {
			maxSum = 0
		}
		fmt.Println("Max sum after: ", maxSum)
		fmt.Println("Max far after: ", maxFar)

	}

	return maxFar
}
