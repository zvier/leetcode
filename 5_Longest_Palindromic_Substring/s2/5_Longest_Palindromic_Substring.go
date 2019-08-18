package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	// prepare bs base on s
	ss := []byte(s)
	bs := make([]byte, 0, len(s))
	bs = append(bs, []byte("^")...)
	for _, bchar := range ss {
		bs = append(bs, []byte("*")...)
		bs = append(bs, bchar)
	}
	bs = append(bs, []byte("*$")...)

	mx := 0
	id := 0
	maxLen := 0
	start := 2
	end := 2
	p := make([]int, len(bs))
	for i := 1; i < len(bs)-1; i++ {
		if i < mx {
			p[i] = min(p[2*id-i], mx-i)
		}
		for bs[i-p[i]-1] == bs[i+p[i]+1] {
			p[i]++
		}

		if i+p[i] > mx {
			id = i
			mx = i + p[i]
		}
		if p[i] > maxLen {
			start = i - p[i]
			end = i + p[i]
			maxLen = p[i]
		}
	}
	// i => ni = 2*i + 1 and ni = 2*i + 2
	// start: 2*i+1=>i 2*i+2=>i ni-1/2
	start = (start - 1) / 2
	// end: 2*i+1=>i-1 2*i+2=>i ni-2/2
	end = (end - 1 - 1) / 2
	return string(ss[start : end+1])
}

func main() {
	in := "cbbd"
	out := longestPalindrome(in)
	fmt.Println(out)
}
