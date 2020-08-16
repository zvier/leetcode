package main

import "fmt"

func trap(height []int) int {
	stack := make([]int, 0, len(height))
	ret := 0
	for i := 0; i <= len(height)-1; i++ {
		// decrease stack
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// height[i] > height[st.top()], pop the top
			top := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			// caclu the distance, ready to get water
			distance := i - stack[len(stack)-1] - 1
			m := min(height[i], height[stack[len(stack)-1]])
			ret = ret + distance*(m-top)
		}
		stack = append(stack, i)
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(nums)
	fmt.Println(res)
}
