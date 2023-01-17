package linkedlist

func RemoveElements(head *Node, val int) *Node {
	curr := head

	var prev *Node

	for curr != nil {
		if curr.Val == val {
			if prev == nil {
				head = curr.Next
			} else {
				prev.Next = curr.Next
			}
			curr = curr.Next
		} else {
			prev = curr
			curr = curr.Next
		}

	}
	return head
}
