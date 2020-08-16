package main

import "fmt"

func largestRectangleArea(heights []int) int {
	ret := 0
	stack := make([]int, 0, len(heights))
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	for i := 0; i <= len(heights)-1; i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// top is the heigt, extent it left and right
			// right can reach to i-1, because heights[top] > heights[i]
			right := i - 1
			// because left element may have already pop
			// so left need to use top in stack add 1
			left := stack[len(stack)-1] + 1
			area := (right - left + 1) * heights[top]
			if area > ret {
				ret = area
			}
		}
		stack = append(stack, i)
	}
	return ret
}

func main() {
	//nums := []int{2, 1, 5, 6, 2, 3}
	nums := []int{6, 7, 5, 2, 4, 5, 9, 3}
	ret := largestRectangleArea(nums)
	fmt.Println(ret)
}
