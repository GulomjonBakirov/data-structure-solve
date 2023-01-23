package array

func SortColors(nums []int) {
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				temp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = temp
				sorted = false
			}
		}
	}
}
