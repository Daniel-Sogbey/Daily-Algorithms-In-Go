package main

import "fmt"

func main() {
	str := "abcdcba"
	fmt.Println(IsPalindrome(str))
}

func IsPalindrome(str string) bool {
	// Write your code here.
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
