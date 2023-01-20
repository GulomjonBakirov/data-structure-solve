package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(value int) *TreeNode {
	return &TreeNode{Val: value}
}

func InOrder(node *TreeNode) {
	if node != nil {
		InOrder(node.Left)
		fmt.Println(node.Val, " ")
		InOrder(node.Right)
	}
}

func PreOrder(node *TreeNode) {
	if node != nil {
		fmt.Println(node.Val, " ")
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}
