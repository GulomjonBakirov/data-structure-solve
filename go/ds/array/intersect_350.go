package array

func Intersect(nums1 []int, nums2 []int) []int {
	result := []int{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {

			if nums1[i] == nums2[j] {
				result = append(result, nums1[i])
				nums1 = append(nums1[:i], nums1[i+1:]...)
				nums2 = append(nums2[:j], nums2[j+1:]...)
				i -= 1
				j -= 1
				break
			}
		}
	}

	return result
}
