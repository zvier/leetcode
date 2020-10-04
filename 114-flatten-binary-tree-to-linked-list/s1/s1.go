package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var pre *TreeNode
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil {
			pre.Left, pre.Right = nil, top
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		pre = top
	}
}

func main() {
	node6 := &TreeNode{6, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node5 := &TreeNode{5, nil, node6}
	node2 := &TreeNode{2, node3, node4}
	node1 := &TreeNode{1, node2, node5}

	flatten(node1)
	for cur := node1; cur != nil; cur = cur.Right {
		fmt.Println(cur.Val)
	}
}
