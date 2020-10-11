package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}
	if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
		return root
	}
	if p.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}

func main() {
	node3 := &TreeNode{3, nil, nil}
	node5 := &TreeNode{5, nil, nil}
	node0 := &TreeNode{0, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node9 := &TreeNode{9, nil, nil}
	node4 := &TreeNode{4, node3, node5}
	node2 := &TreeNode{2, node0, node4}
	node8 := &TreeNode{8, node7, node9}
	node6 := &TreeNode{6, node2, node8}
	ret := lowestCommonAncestor(node6, node2, node4)
	fmt.Println(ret)
}
