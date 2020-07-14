package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	result := make([]int, 0)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if top != nil {
			stack = append(stack, top)
			stack = append(stack, nil)
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		} else {
			result = append(result, stack[len(stack)-1].Val)
			stack = stack[0 : len(stack)-1]
		}
	}
	return result
}

func main() {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, node3, nil}
	node1 := &TreeNode{1, nil, node2}
	ret := postorderTraversal(node1)
	fmt.Println(ret)
}
