package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return helper(root.Left, min, root) && helper(root.Right, root, max)
}

func main() {
	node3 := &TreeNode{3, nil, nil}
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, node1, node3}
	r := isValidBST(node2)
	fmt.Println(r)
}
