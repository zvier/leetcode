package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, 0, len(nums))
	// init with 1
	for i := 0; i <= len(nums)-1; i++ {
		result = append(result, 1)
	}

	// multi from left part
	base := 1
	for i := 0; i <= len(nums)-1; i++ {
		result[i] = base * result[i]
		base = base * nums[i]
	}

	// multi from right part
	base = 1
	for j := len(nums) - 1; j >= 0; j-- {
		result[j] = base * result[j]
		base = base * nums[j]
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
