package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	result := []int{}

	for len(stack) > 0 {
		// 1. pop a node from the stack
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		// 2. process the node
		result = append(result, node.Val)

		// 3. enqueue the children node
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

func main() {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, node3, nil}
	node1 := &TreeNode{1, nil, node2}
	result := preorderTraversal(node1)
	fmt.Println(result)
}
