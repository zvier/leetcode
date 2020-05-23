package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return _climbStairs(n, 1, 2)
}

func _climbStairs(n, a, b int) int {
	if n == 2 {
		return b
	}
	return _climbStairs(n-1, b, a+b)
}

func main() {
	res := climbStairs(4)
	fmt.Println(res)
}
