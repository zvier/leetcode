package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	p1 *TreeNode
	p2 *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []pair{{p, q}}
	for len(queue) != 0 {
		p := queue[0].p1
		q := queue[0].p2
		queue = queue[1:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, pair{p.Left, q.Left})
		queue = append(queue, pair{p.Right, q.Right})
	}
	return true
}

func main() {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node1 := &TreeNode{1, node2, node3}
	ret := isSameTree(node1, node1)
	fmt.Println(ret)
}
