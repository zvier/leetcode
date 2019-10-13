package s1

func sortColors(nums []int) {
	counter := make(map[int]int, 3)
	for _, v := range nums {
		counter[v]++
	}

	inum := 0
	for _, v := range [...]int{0, 1, 2} {
		c := counter[v]
		for ; c > 0; c-- {
			nums[inum] = v
			inum++
		}
	}
}
