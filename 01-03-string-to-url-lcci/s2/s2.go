package main

import "fmt"

const emptyChar = ' '

func replaceSpaces(S string, length int) string {
	bytes := []byte(S)
	right := len(bytes) - 1
	mid := length - 1
	for ; mid >= 0; mid-- {
		if bytes[mid] != emptyChar {
			bytes[mid], bytes[right] = bytes[right], bytes[mid]
			right--
		} else {
			bytes[right-2], bytes[right-1], bytes[right] = '%', '2', '0'
			right -= 3
		}
	}
	return string(bytes[right+1:])
}

func main() {
	s := "ds sdfs afs sdfa dfssf asdf             "
	r := replaceSpaces(s, 27)
	fmt.Println(r)
}
