package main

import "fmt"

func towSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for i, num := range nums {
		mp[num] = i
	}

	for ia, num := range nums {
		diff := target - num
		ib, ok := mp[diff]
		if ok && ib != ia {
			return []int{ia, ib}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := towSum(nums, target)
	fmt.Println(result)
}
