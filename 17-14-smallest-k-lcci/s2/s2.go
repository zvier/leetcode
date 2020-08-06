package main

import "fmt"

func smallestK(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	heapify(arr, k)
	for i := k; i <= len(arr)-1; i++ {
		if arr[i] < arr[0] {
			arr[0] = arr[i]
			down(arr, 0, k)
		}
	}

	return arr[:k]
}

func heapify(arr []int, k int) {
	for i := k/2 - 1; i >= 0; i-- {
		down(arr, i, k)
	}
}

func down(arr []int, i, k int) bool {
	start := i
	for {
		left := 2*i + 1
		if left >= k || left < 0 {
			break
		}
		max := left
		right := left + 1
		if right < k && arr[right] > arr[left] {
			max = right
		}

		if arr[i] > arr[max] {
			break
		}
		arr[i], arr[max] = arr[max], arr[i]
		i = max
	}

	return i > start
}
func main() {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}
	result := smallestK(arr, 4)
	fmt.Println(result)
}
