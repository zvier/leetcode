package main

import "fmt"

func climbStairs(n int) int {
	dp := []int{1, 2}
	if n <= 2 {
		return n
	}
	for ; n > 2; n-- {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
	}
	return dp[1]
}

func main() {
	res := climbStairs(4)
	fmt.Println(res)
}
