package main

import "fmt"

func countSubstrings(s string) int {
	dp := make([][]bool, 0, len(s))
	for i := 0; i <= len(s)-1; i++ {
		row := make([]bool, len(s))
		dp = append(dp, row)
	}

	cnt := 0
	for j := 0; j <= len(s)-1; j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	s := "abc"
	r := countSubstrings(s)
	fmt.Println(r)
}
