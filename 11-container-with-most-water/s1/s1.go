package main

import "fmt"

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := 1; j <= len(height)-1; j++ {
			length := j - i
			weight := min(height[i], height[j])
			area := length * weight
			if area > max {
				max = area
			}
		}
	}
	return max
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
