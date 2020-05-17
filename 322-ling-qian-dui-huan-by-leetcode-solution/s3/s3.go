package main

import "fmt"

const MAX = 1<<32 - 1

func coinChange(coins []int, amount int) int {
	// dp[i] is for the amount of i, dp[amount] is legal,
	// so the length of dp array shoud be amount+1
	dp := make([]int, amount+1)
	for i := 1; i <= len(dp)-1; i++ {
		dp[i] = MAX
	}
	for _, coin := range coins {
		// because a new coin shoud be select one or more
		// the target can start a least coin and end amount
		// add one coin each time and update dp
		for target := coin; target <= amount; target++ {
			dp[target] = min(dp[target], dp[target-coin]+1)
		}
	}
	if dp[amount] != MAX {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	coins := []int{346, 29, 395, 188, 155, 109}
	amount := 9401
	ret := coinChange(coins, amount)
	fmt.Println(ret)
}
