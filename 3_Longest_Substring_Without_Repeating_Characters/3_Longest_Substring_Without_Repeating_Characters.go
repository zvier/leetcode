package leetcode

func lengthOfLongestSubstring(s string) int {
	dp := make(map[rune]int)
	start := 0
	length := 0

	for i, v := range s {
		if iv, ok := dp[v]; ok && iv >= start {
			start = iv + 1
		}
		curlen := i - start + 1
		if curlen > length {
			length = curlen
		}
		dp[v] = i
	}
	return length
}
