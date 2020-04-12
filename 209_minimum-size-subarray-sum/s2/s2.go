package s2

func minSubArrayLen(s int, nums []int) int {
	max := len(nums) + 1
	minLen := max

	for i := 0; i <= len(nums)-1; i++ {
		sum := 0
		for j := i; j <= len(nums)-1; j++ {
			sum += nums[j]
			if sum >= s {
				minLen = min(minLen, j-i+1)
				break
			}
		}
	}

	if minLen == max {
		return 0
	}

	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
