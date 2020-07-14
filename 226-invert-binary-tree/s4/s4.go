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
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
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
