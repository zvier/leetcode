package main

import "fmt"

const MAX = 1<<32 - 1

func coinChange(coins []int, amount int) int {
	cnt := helpCoinChange(0, amount, coins)
	return cnt
}

func helpCoinChange(idxCoin, amount int, coins []int) int {
	// amount is zero, no need any more, return 0 coin
	if amount == 0 {
		return 0
	}

	if idxCoin > len(coins)-1 || amount < 0 {
		return -1
	}
	// for coins[i] and current amount, the maxCnt is amount/coins[i]
	maxCnt := amount / coins[idxCoin]
	minCost := MAX
	// try each cnt in [0, maxCnt] for coins[i]
	for x := 0; x <= maxCnt; x++ {
		if amount >= x*coins[idxCoin] {
			res := helpCoinChange(idxCoin+1, amount-x*coins[idxCoin], coins)
			if res != -1 {
				minCost = min(minCost, res+x)
			}
		}
	}
	if minCost != MAX {
		return minCost
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
