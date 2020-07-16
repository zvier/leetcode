package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		rightMinNode := getRightMinNode(root.Right)
		root.Val = rightMinNode.Val
		root.Right = deleteNode(root.Right, rightMinNode.Val)
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getRightMinNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func main() {
	node7 := &TreeNode{7, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, node2, node4}
	node6 := &TreeNode{6, nil, node7}
	node5 := &TreeNode{5, node3, node6}
	r := deleteNode(node5, 3)
	fmt.Println(r.Left)
	fmt.Println(r.Right)
}
