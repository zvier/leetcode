package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	track := make([]int, 0, len(nums))
	bt(nums, track, &res)
	return res
}

func bt(nums, track []int, res *[][]int) {
	// end bt if meet
	if len(track) == len(nums) {
		item := make([]int, len(nums))
		copy(item, track)
		*res = append(*res, item)
		return
	}

	mp := make(map[int]bool, len(track))
	for _, t := range track {
		mp[t] = true
	}
	for _, num := range nums {
		if mp[num] {
			continue
		}
		track = append(track, num)
		bt(nums, track, res)
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
