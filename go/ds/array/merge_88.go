package array

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 {
		nums1 = nums2
	} else {
		nums1 = nums1[:m]
		nums1 = append(nums1, nums2...)
	}
	var i, key, j int

	for i = 1; i < len(nums1); i++ {
		key = nums1[i]
		j = i - 1

		for j >= 0 && nums1[j] > key {
			nums1[j+1] = nums1[j]
			j = j - 1
		}

		nums1[j+1] = key
	}

	return nums1
}
