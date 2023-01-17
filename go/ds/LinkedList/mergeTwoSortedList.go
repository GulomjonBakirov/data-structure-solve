package linkedlist

func MergeTwoLinkedList(list1 *Node, list2 *Node) *Node {
	result := &Node{0, nil}
	curr := result
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}

	if p1 != nil {
		curr.Next = p1
	} else {
		curr.Next = p2
	}

	return result.Next
}
