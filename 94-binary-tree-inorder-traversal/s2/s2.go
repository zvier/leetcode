package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	l := inorderTraversal(root.Left)
	r := inorderTraversal(root.Right)
	result = append(result, l...)
	result = append(result, root.Val)
	result = append(result, r...)
	return result
}

func main() {
	n2 := &TreeNode{3, nil, nil}
	n1 := &TreeNode{2, n2, nil}
	n0 := &TreeNode{1, nil, n1}
	res := inorderTraversal(n0)
	fmt.Println(res)
}
