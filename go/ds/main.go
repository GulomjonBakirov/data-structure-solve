package main

import (
	"fmt"
	linkedlist "leetcode-ds/LinkedList"
	"leetcode-ds/array"
	"leetcode-ds/tree"
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
	ll := &linkedlist.Node{Val: 1}
	ll.Next = &linkedlist.Node{Val: 2}
	ll.Next.Next = &linkedlist.Node{Val: 3}
	ll.Next.Next.Next = &linkedlist.Node{Val: 4}
	fmt.Println(ll)
	ll2 := &linkedlist.Node{Val: 1}
	ll2.Next = &linkedlist.Node{Val: 3}
	ll2.Next.Next = &linkedlist.Node{Val: 4}
	// ll.Append(1)
	// ll.Append(2)
	// ll.Append(4)
	// ll2.Append(1)
	// ll2.Append(3)
	// ll2.Append(4)
	// result := linkedlist.MergeTwoLinkedList(ll, ll2)
	// result := linkedlist.RemoveElements(ll, 3)
	// for ll != nil {
	// 	fmt.Printf("%v -> ", ll.Val)
	// 	ll = ll.Next
	// }
	// fmt.Println("nil")

	root := tree.NewNode(1)
	root.Right = tree.NewNode(3)
	root.Right.Left = tree.NewNode(4)

	// fmt.Println("In order")
	// tree.InOrder(root)
	array.SortColors([]int{1, 2, 3, 1, 2})
	// fmt.Println(result)
}
