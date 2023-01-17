package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type ListNode struct {
	head *Node
}

func (ll *ListNode) Append(data int) {
	node := &Node{Val: data}

	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.Next != nil {
			current = current.Next
		}

		current.Next = node
	}

}

func (ll *ListNode) Prepend(data int) {
	node := &Node{Val: data, Next: ll.head}
	ll.head = node
}

func (ll *ListNode) Delete(data int) {
	if ll.head == nil {
		return
	}

	if ll.head.Val == data {
		ll.head = ll.head.Next
		return
	}

	current := ll.head

	for current.Next != nil {
		if current.Next.Val == data {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
}

func (ll *ListNode) PrintList() {
	current := ll.head

	for current != nil {
		fmt.Printf("%v -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
