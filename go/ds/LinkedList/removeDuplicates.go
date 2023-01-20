package linkedlist

import "fmt"

func DeleteDuplicates(head *Node) *Node {
	fmt.Println(head)
	if head == nil {
		return &Node{}
	}
	curr := head
	next := head.Next

	for curr != nil && next != nil {
		if curr.Val == next.Val {
			curr.Next = next.Next
			next = next.Next
		} else {
			curr = curr.Next
			next = next.Next
		}
	}

	return head
}
