package tree

func PreorderTraversal(root *TreeNode, result ...int) []int {

	if root == nil {
		return []int{}
	}

	right := PreorderTraversal(root.Right)
	left := PreorderTraversal(root.Left)

	return append(append([]int{root.Val}, left...), right...)
}
