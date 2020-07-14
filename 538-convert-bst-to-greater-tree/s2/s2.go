package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	node := root
	stack := []*TreeNode{}
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		sum += top.Val
		top.Val = sum
		node = top.Left
	}
	return root
}

func main() {
	node13 := &TreeNode{13, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node5 := &TreeNode{5, node2, node13}
	root := convertBST(node5)
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Right)
}
