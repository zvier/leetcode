package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	node4 := &ListNode{-4, nil}
	node0 := &ListNode{0, node4}
	node2 := &ListNode{2, node0}
	node3 := &ListNode{3, node2}
	node4.Next = node2
	ret := hasCycle(node3)
	fmt.Println(ret)
}
