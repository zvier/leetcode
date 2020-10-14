package main

import "fmt"

type TreeNode struct {
    Value int
    Right *TreeNode
    Left *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    stack := []*TreeNode{root}
    ans := 0 
    for len(stack) > 0 {
        cnt := len(stack)
        for cnt > 0 {
            front := stack[0]
            stack = stack[1:]
            if front.Left != nil { 
                stack = append(stack, front.Left)
            }
            if front.Right != nil {
                stack = append(stack, front.Right)
            }
            cnt--
        }
        ans++
    }
    return ans
}

func main() {
    node7 := &TreeNode{7, nil, nil}    
    node15 := &TreeNode{15, nil, nil}    
    node20 := &TreeNode{20, node15, node7}    
    node9 := &TreeNode{9, nil, nil}    
    node3 := &TreeNode{3, node9, node20}    
    ret := maxDepth(node3)
    fmt.Println(ret)
}
