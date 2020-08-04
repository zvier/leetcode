/**
* @file: s2.go
* @brief: postorder use stack
* @author zvier, zvier6@gmail.com
* @version: 1.0.0
* @date: 2020-08-05
 */
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
	stack := []*Node{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top == nil {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, top.Val)
			continue
		}
		if top.Children == nil {
			result = append(result, top.Val)
			continue
		}
		stack = append(stack, top)
		stack = append(stack, nil)
		for i := len(top.Children) - 1; i >= 0; i-- {
			child := top.Children[i]
			stack = append(stack, child)
		}
	}
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
