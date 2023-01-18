package linkedlist

func ReverseList(head *Node) *Node {
	if head == nil {
		return &Node{}
	}

	current := head
	var prev *Node
	next := head.Next

	for current.Next != nil {
		next = current.Next // 2 , 3
		current.Next = prev // null, 1
		prev = current      // 1, 2
		current = next      // 2, 3
	}

	current.Next = prev

	return current
}
