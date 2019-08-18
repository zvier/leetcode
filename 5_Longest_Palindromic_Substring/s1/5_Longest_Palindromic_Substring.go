package main

import "fmt"

func palindRadix(s []byte, i int) int {
	left := i
	right := i
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++
	}
	radix := right - i - 1
	return radix
}

func longestPalindromic(s string) string {
	bs := make([]byte, 0, len(s))
	for _, bchar := range []byte(s) {
		bs = append(bs, []byte("*")...)
		bs = append(bs, bchar)
	}
	bs = append(bs, []byte("*")...)

	maxPalindRadix := 0
	start := 0
	end := 0
	for i, _ := range bs {
		radix := palindRadix(bs, i)
		if radix > maxPalindRadix {
			maxPalindRadix = radix
			start = i - radix
			end = i + radix
		}
	}
	result := []byte(s)[start/2 : end/2]
	return string(result)
}

func main() {
	in := "babad"
	out := longestPalindromic(in)
	fmt.Println(out)
}
