package s3

func minSubArrayLen(s int, nums []int) int {
	max := len(nums) + 1
	minLen := max

	sums := make([]int, 0, len(nums))
	sum := 0
	for _, num := range nums {
		sum += num
		sums = append(sums, sum)
	}

	for i := 0; i <= len(nums)-1; i++ {
		beforeSum := 0
		if i >= 1 {
			beforeSum = sums[i-1]
		}

		left, right := i, len(nums)-1
		for left <= right {
			mid := left + (right-left)/2
			sum := sums[mid] - beforeSum
			if sum >= s {
				minLen = min(minLen, mid-i+1)
				right = mid - 1
			} else {
				left = mid + 1
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
