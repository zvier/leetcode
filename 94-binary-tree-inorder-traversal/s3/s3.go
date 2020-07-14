package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		if top.Right != nil {
			node = top.Right
		}
	}
	return result
}

func main() {
	n2 := &TreeNode{3, nil, nil}
	n1 := &TreeNode{2, n2, nil}
	n0 := &TreeNode{1, nil, n1}
	res := inorderTraversal(n0)
	fmt.Println(res)
}
