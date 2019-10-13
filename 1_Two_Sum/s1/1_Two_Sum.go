package main

import "fmt"

func towSum(nums []int, target int) []int {
	for ia, numa := range nums {
		diff := target - numa
		for ib, numb := range nums[ia+1:] {
			if numb == diff {
				return []int{ia, ib + (ia + 1)}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 26
	result := towSum(nums, target)
	fmt.Println(result)
}
