package main

import (
	linkedlist "leetcode-ds/LinkedList"
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
	ll := &linkedlist.ListNode{}
	ll2 := &linkedlist.ListNode{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(4)
	ll.PrintList()
	ll2.Append(1)
	ll2.Append(3)
	ll2.Append(4)
	ll2.PrintList()
	// result := stringmod.CanConstruct("aa", "aab")
	// fmt.Println("Solve: ", result)

}
