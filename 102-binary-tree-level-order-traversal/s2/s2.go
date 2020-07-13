package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		line := make([]int, 0, len(queue))
		nextQueue := make([]*TreeNode, 0)
		for _, node := range queue {
			line = append(line, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
		result = append(result, line)
	}
	return result
}

func main() {
	node7 := &TreeNode{7, nil, nil}
	node15 := &TreeNode{15, nil, nil}
	node20 := &TreeNode{20, node15, node7}
	node9 := &TreeNode{9, nil, nil}
	node3 := &TreeNode{3, node9, node20}
	ret := levelOrder(node3)
	fmt.Println(ret)
}
