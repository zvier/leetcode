package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	return checkEqualTree(t1, t2) || checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func checkEqualTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return checkEqualTree(t1.Left, t2.Left) && checkEqualTree(t1.Right, t2.Right)
}

func main() {
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node1 := &TreeNode{1, node2, node3}

	node := &TreeNode{2, nil, nil}
	ret := checkSubTree(node1, node)
	fmt.Println(ret)
}
