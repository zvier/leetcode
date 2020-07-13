package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return isSub(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSub(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return isSub(head.Next, root.Left) || isSub(head.Next, root.Right)
}

func main() {
	hnode8 := &ListNode{8, nil}
	hnode2 := &ListNode{2, hnode8}
	hnode4 := &ListNode{4, hnode2}

	tnode3 := &TreeNode{3, nil, nil}
	tnode1 := &TreeNode{1, nil, nil}
	tnode8 := &TreeNode{8, tnode1, tnode3}
	tnode6 := &TreeNode{6, nil, nil}
	tnode2 := &TreeNode{2, tnode6, tnode8}
	tnoder4 := &TreeNode{4, tnode2, nil}
	tnode41 := &TreeNode{1, nil, nil}
	tnode32 := &TreeNode{2, tnode41, nil}
	tnodel4 := &TreeNode{4, nil, tnode32}

	root := &TreeNode{1, tnodel4, tnoder4}

	ret := isSubPath(hnode4, root)
	fmt.Println(ret)
}
