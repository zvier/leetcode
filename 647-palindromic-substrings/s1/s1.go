package main

import "fmt"

func countSubstrings(s string) int {
	n := len(s)
	cnt := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			cnt++
		}
	}
	return cnt
}

func main() {
	s := "abc"
	r := countSubstrings(s)
	fmt.Println(r)
}
