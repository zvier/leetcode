package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
func main() {
	node3 := &TreeNode{3, nil, nil}
	node1 := &TreeNode{1, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node2 := &TreeNode{2, node1, node3}
	node4 := &TreeNode{4, node2, node7}
	ret := insertIntoBST(node4, 5)
	fmt.Println(ret)
}
