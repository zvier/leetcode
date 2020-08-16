package main

import "fmt"

func longestPalindrome(s string) string {
	ret := []byte{}

	bs := []byte(s)
	for i := 0; i <= len(bs)-1; i++ {
		s1 := palindrome(bs, i, i)
		s2 := palindrome(bs, i, i+1)

		subStr := s1
		if len(s1) < len(s2) {
			subStr = s2
		}
		if len(subStr) > len(ret) {
			ret = subStr
		}
	}
	return string(ret)
}

func palindrome(s []byte, l, r int) []byte {
	for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
		l, r = l-1, r+1
	}
	return s[l+1 : r]
}

func main() {
	s := "cbbd"
	ret := longestPalindrome(s)
	fmt.Println(ret)
}
