package array

func TwoSum(nums []int, target int) []int {
	result := []int{}

	for i := 0; i < len(nums); i++ {
		res := target - nums[i]

		for j := i + 1; j < len(nums); j++ {
			if res == nums[j] {
				result = append(result, i, j)
				break
			}
		}

	}
	return result

}
