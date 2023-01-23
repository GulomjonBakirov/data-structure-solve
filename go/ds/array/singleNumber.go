package array

func SingleNumber(nums []int) int {
	duplicates := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if duplicates[nums[i]] != 0 {
			duplicates[nums[i]] += 1
		} else {
			duplicates[nums[i]] = 1
		}

	}

	for _, el := range nums {
		if duplicates[el] == 1 {
			return duplicates[el]
		}
	}

	return 1
}
