package main

import "fmt"

func subarraySum(nums []int, k int) int {
	cnt, pre := 0, 0
	mp := make(map[int]int, len(nums))
	mp[0] = 1
	for i := 0; i <= len(nums)-1; i++ {
		pre += nums[i]
		if n, ok := mp[pre-k]; ok {
			cnt += n
		}
		mp[pre] += 1
	}
	return cnt
}

func main() {
	nums := []int{1, 1, 1}
	ret := subarraySum(nums, 2)
	fmt.Println(ret)
}
