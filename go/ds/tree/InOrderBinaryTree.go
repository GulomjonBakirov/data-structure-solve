package tree

import "fmt"

func InOrderBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	fmt.Println("Val: ", root.Val)

	left := InOrderBinaryTree(root.Left)
	right := InOrderBinaryTree(root.Right)

	return append(append([]int{root.Val}, right...), left...)
}
