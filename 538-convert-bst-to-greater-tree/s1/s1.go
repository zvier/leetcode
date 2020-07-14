package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	transfer(root, 0)
	return root
}

func transfer(root *TreeNode, total int) int {
	if root == nil {
		return total
	}
	total = transfer(root.Right, total)
	root.Val += total
	total = transfer(root.Left, root.Val)
	return total
}

func main() {
	node13 := &TreeNode{13, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node5 := &TreeNode{5, node2, node13}
	root := convertBST(node5)
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Right)
}
