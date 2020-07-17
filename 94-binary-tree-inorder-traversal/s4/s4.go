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
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		// 1. pop a node from the stack
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if top != nil {
			// 2. right node enqueue
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			// 3. cur node re-enqueue
			stack = append(stack, top)
			stack = append(stack, nil)
			// 4. left node enqueue
			if top.Left != nil {
				stack = append(stack, top.Left)
			}

		} else {
			top = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			result = append(result, top.Val)
		}
	}

	return result
}

func main() {
	n3 := &TreeNode{3, nil, nil}
	n2 := &TreeNode{2, n3, nil}
	n1 := &TreeNode{1, nil, n2}
	res := inorderTraversal(n1)
	fmt.Println(res)
}
