package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, 0, len(s))
	for i := 0; i <= len(s)-1; i++ {
		row := make([]int, len(s))
		row[i] = 1
		dp = append(dp, row)
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j <= len(s)-1; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "bbbab"
	ret := longestPalindromeSubseq(s)
	fmt.Println(ret)
}
