package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}

	fmt.Println(reverseString(s))
}

// func reverseString(s []byte) {
// left := 0
// right := len(s) - 1

// for left < right {
// s[left], s[right] = s[right], s[left]
// right--
// left++
// }

// }

func reverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}
