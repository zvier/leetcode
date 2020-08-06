/**
* @file: s2.go
* @brief: use double pointer
* @author zvier, zvier60@gmail.com
* @version: 1.0.0
* @date: 2020-08-06
 */
package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxValue := 0
	for left < right {
		width := right - left
		high := min(height[left], height[right])
		area := width * high
		if area > maxValue {
			maxValue = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxValue
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(input)
	fmt.Println(res)
}
