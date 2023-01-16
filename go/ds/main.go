package main

import (
	"fmt"
	stringmod "leetcode-ds/string_mod"
)

func main() {
	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }
	// result := array.MaxSubArray(nums)
	// array.ContainsDuplicate(nums)

	// result := array.TwoSum(nums, target)
	// result := array.Intersect([]int{1, 2, 2, 1}, []int{2, 2})
	// result := array.GetRow(3)
	// result := stringmod.FirstUniqChar("ssaalom")
	result := stringmod.CanConstruct("aa", "aab")
	fmt.Println("Solve: ", result)

}
