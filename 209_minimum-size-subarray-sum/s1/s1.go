package s1

func minSubArrayLen(s int, nums []int) int {
	max := len(nums) + 1
	slow, fast := 0, 0
	minLen, sum := max, 0

	for ; fast <= len(nums)-1; fast++ {
		sum += nums[fast]
		for sum >= s {
			curLen := fast - slow + 1
			minLen = min(minLen, curLen)
			sum -= nums[slow]
			slow++
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
