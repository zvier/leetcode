package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{Head: head}
}

func (this *Solution) GetRandom() int {
	rand.Seed(time.Now().UTC().UnixNano())
	index := 1
	pNode := this.Head
	result := this.Head.Val
	for pNode != nil {
		psudo := rand.Intn(index)
		if psudo == 0 {
			result = pNode.Val
		}
		index++
		pNode = pNode.Next
	}
	return result
}

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	obj := Constructor(node1)
	param_1 := obj.GetRandom()
	fmt.Println(param_1)
}
