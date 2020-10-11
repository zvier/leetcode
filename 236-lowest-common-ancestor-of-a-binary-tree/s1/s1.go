package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func main() {
	node4 := &TreeNode{4, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node2 := &TreeNode{2, node7, node4}
	node6 := &TreeNode{6, nil, nil}
	node0 := &TreeNode{0, nil, nil}
	node8 := &TreeNode{8, nil, nil}
	node5 := &TreeNode{5, node6, node2}
	node1 := &TreeNode{1, node0, node8}
	node3 := &TreeNode{3, node5, node1}
	ret := lowestCommonAncestor(node3, node5, node1)
	fmt.Println(ret)
}
