package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	mp := make(map[int]int, len(nums2))
	stack := make([]int, 0, len(nums2))

	// range nums2
	for i := 0; i <= len(nums2)-1; i++ {
		// up stack, if current greater than top, pop top and store
		// because it next greater is current
		for len(stack) != 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			mp[nums2[stack[len(stack)-1]]] = i
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}

	// range nums1
	for i := 0; i <= len(nums1)-1; i++ {
		if index, ok := mp[nums1[i]]; ok {
			result = append(result, nums2[index])
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	ret := nextGreaterElement(nums1, nums2)
	fmt.Println(ret)
}
