package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}

	result := reverseList(node1)
	for cur := result; cur != nil; cur = cur.Next {
		fmt.Println(cur)
	}
}
