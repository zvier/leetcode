package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 1. outqueue the first node
		node := queue[0]
		queue = queue[1:]

		// 2. swap the node's left and right
		node.Left, node.Right = node.Right, node.Left

		// 3. left and right enqueue if not nil
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

func main() {
	node1 := &TreeNode{1, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node6 := &TreeNode{6, nil, nil}
	node9 := &TreeNode{9, nil, nil}
	node2 := &TreeNode{2, node1, node3}
	node7 := &TreeNode{7, node6, node9}
	node4 := &TreeNode{4, node2, node7}
	root := invertTree(node4)
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Right)
}
