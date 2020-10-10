package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	quicksort(head, nil)
	return head
}

func quicksort(begin, end *ListNode) {
	if begin == end {
		return
	}
	pivot := partition(begin, end)
	quicksort(begin, pivot)
	quicksort(pivot.Next, end)
}

func partition(begin, end *ListNode) *ListNode {
	cursor, next := begin, begin
	for cursor != end {
		if cursor.Val < begin.Val {
			next = next.Next
			swap(cursor, next)
		}
		cursor = cursor.Next
	}
	swap(begin, next)
	return next
}

func swap(a, b *ListNode) {
	a.Val, b.Val = b.Val, a.Val
}

func main() {
	node3 := &ListNode{3, nil}
	node1 := &ListNode{1, node3}
	node2 := &ListNode{2, node1}
	node4 := &ListNode{4, node2}
	ret := sortList(node4)
	for cur := ret; cur != nil; cur = cur.Next {
		fmt.Println(cur)
	}
}
