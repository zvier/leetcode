package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func main() {
	res := climbStairs(4)
	fmt.Println(res)
}
