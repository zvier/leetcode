package main

import "fmt"

const MAX = 1<<32 - 1

func coinChange(coins []int, amount int) int {
	// dp[i] is for the amount of i, dp[amount] is legal,
	// so the length of dp array shoud be amount+1
	dp := make([]int, amount+1)
	return helpCoinChange(coins, dp, amount)
}

func helpCoinChange(coins, dp []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	// to dp[i], if not zero, that's to say it already
	// compute, use it diretly and return
	if dp[amount] != 0 {
		return dp[amount]
	}
	min := MAX
	// range each coin to find to amount, which last coin is
	for _, coin := range coins {
		res := helpCoinChange(coins, dp, amount-coin)
		if res >= 0 && res < min {
			min = res + 1
		}
	}
	if min == MAX {
		min = -1
	}
	dp[amount] = min
	return dp[amount]
}

func main() {
	coins := []int{346, 29, 395, 188, 155, 109}
	amount := 9401
	ret := coinChange(coins, amount)
	fmt.Println(ret)
}
