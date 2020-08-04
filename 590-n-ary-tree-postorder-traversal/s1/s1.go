package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	for _, child := range root.Children {
		c := postorder(child)
		result = append(result, c...)
	}
	result = append(result, root.Val)
	return result
}

func main() {
	node5 := &Node{5, nil}
	node6 := &Node{6, nil}
	node3 := &Node{3, []*Node{node5, node6}}
	node2 := &Node{2, nil}
	node4 := &Node{4, nil}
	node1 := &Node{1, []*Node{node3, node2, node4}}
	ret := postorder(node1)
	fmt.Println(ret)
}
