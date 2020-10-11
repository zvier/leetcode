package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	counter := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		counter[num]++
	}
	result := make([]int, 0, len(nums2))
	for _, num := range nums2 {
		if cnt, ok := counter[num]; ok && cnt > 0 {
			result = append(result, num)
			counter[num]--
		}
	}
	return result
}

func main() {
	//nums1 := []int{1, 2, 2, 1}
	//nums2 := []int{2, 2}
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	ret := intersect(nums1, nums2)
	fmt.Println(ret)
}
