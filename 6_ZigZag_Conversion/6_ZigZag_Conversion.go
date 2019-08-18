package main

import "fmt"

func convert(s string, numRows int) string {
	bs := []byte(s)
	ret := make([]byte, 0, len(bs))
	for i := 0; i < numRows; i++ {
		gap1 := 2 * (numRows - i - 1)
		gap2 := 2 * i
		if gap2 == 0 && gap1 != 0 {
			gap2 = gap1
		}
		if gap1 == 0 && gap2 != 0 {
			gap1 = gap2
		}
		if gap1 == 0 && gap2 == 0 {
			gap1 = 1
			gap2 = 1
		}

		cnt := 0
		for j := i; j <= len(bs)-1; cnt++ {
			ret = append(ret, bs[j])
			if cnt%2 == 0 {
				j = j + gap1
			} else {
				j = j + gap2
			}
		}
	}
	return string(ret)
}

func main() {
	in := "ABC"
	out := convert(in, 1)
	fmt.Println(out)
}
