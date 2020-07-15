package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	l := searchBST(root.Left, val)
	if l != nil {
		return l
	}
	r := searchBST(root.Right, val)
	return r
}

func main() {
	node3 := &TreeNode{3, nil, nil}
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, node1, node3}
	node7 := &TreeNode{7, nil, nil}
	node4 := &TreeNode{4, node2, node7}
	r := searchBST(node4, 4)
	fmt.Println(r)
}
