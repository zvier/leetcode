package main

import "fmt"

const emptyChar = ' '

func replaceSpaces(S string, length int) string {
	result := make([]byte, 0, length)
	for i := 0; i <= length-1; i++ {
		if S[i] == emptyChar {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, S[i])
		}
	}
	return string(result)
}

func main() {
	s := "Mr John Smith    "
	r := replaceSpaces(s, 13)
	fmt.Println(r)
}
