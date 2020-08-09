package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

func main() {
	node31 := &TreeNode{3, nil, nil}
	node41 := &TreeNode{4, nil, nil}
	node42 := &TreeNode{4, nil, nil}
	node32 := &TreeNode{3, nil, nil}
	node21 := &TreeNode{2, node31, node41}
	node22 := &TreeNode{2, node42, node32}
	node1 := &TreeNode{1, node21, node22}
	ret := isSymmetric(node1)
	fmt.Println(ret)
}
