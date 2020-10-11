package main

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, 0, len(s))
	for i := 0; i <= len(s)-1; i++ {
		row := make([]bool, len(s))
		dp = append(dp, row)
	}

	maxLen := 0
	maxStr := ""
	for j := 0; j <= len(s)-1; j++ {
		for i := 0; i <= j; i++ {
			dp[i][j] = (s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]))
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				maxStr = string(s[i : j+1])
			}
		}
	}
	return maxStr
}

func main() {
	s := "cbbd"
	r := longestPalindrome(s)
	fmt.Println(r)
}
