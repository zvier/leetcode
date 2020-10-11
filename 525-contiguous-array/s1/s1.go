package main

import "fmt"

func findMaxLength(nums []int) int {
	mp := make(map[int]int, len(nums))
	mp[0] = -1
	count := 0
	maxLen := 0
	for i := 0; i <= len(nums)-1; i++ {
		count = count + 2*nums[i] - 1
		if index, ok := mp[count]; ok {
			l := i - index
			if l > maxLen {
				maxLen = l
			}
		} else {
			mp[count] = i
		}
	}
	return maxLen
}

func main() {
	nums := []int{0, 1}
	ret := findMaxLength(nums)
	fmt.Println(ret)
}
