package main

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 0

	for i := 0; i <= len(nums)-1; i++ {
		preMax := 0
		for j := 0; j <= i-1; j++ {
			if nums[i] > nums[j] && dp[j] > preMax {
				preMax = dp[j]
			}
		}

		dp[i] = preMax + 1
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(nums)
	fmt.Println(result)
}
