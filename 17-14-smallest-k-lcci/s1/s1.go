package s1

func smallestK(arr []int, k int) []int {
	quicksort(arr, 0, len(arr)-1)
	return arr[0:k]
}

func quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}
	j := partition(arr, low, high)
	quicksort(arr, low, j-1)
	quicksort(arr, j+1, high)
}

func partition(arr []int, low, high int) int {
	l, r := low, high
	for {
		for arr[l] <= arr[low] && l < high {
			l++
		}
		for arr[r] >= arr[low] && r > low {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[low], arr[r] = arr[r], arr[low]

	return r
}
