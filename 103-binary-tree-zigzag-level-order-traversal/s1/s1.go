package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root, nil}
	// use to control direction
	leftToRight := true
	result := make([][]int, 0)
	line := make([]int, 0)
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front != nil {
			if leftToRight {
				line = append(line, front.Val)
			} else {
				line = append([]int{front.Val}, line...)
			}
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		} else {
			leftToRight = !leftToRight
			tmp := make([]int, len(line))
			copy(tmp, line)
			result = append(result, tmp)
			line = line[:0]
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		}
	}

	return result
}

func main() {
	node15 := &TreeNode{15, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node20 := &TreeNode{20, node15, node7}
	node9 := &TreeNode{9, nil, nil}
	node3 := &TreeNode{3, node9, node20}
	result := zigzagLevelOrder(node3)
	fmt.Println(result)
}
