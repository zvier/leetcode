package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node1 := &TreeNode{1, node2, node3}
	ret := isSameTree(node1, node1)
	fmt.Println(ret)
}
