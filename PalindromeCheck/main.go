package main

import "fmt"

func main() {
	str := "abcdcba"
	fmt.Println(IsPalindrome2(str))
	fmt.Println(isPalindromePractice(str))
}

func IsPalindrome1(str string) bool {
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

func IsPalindrome2(str string) bool {
	// Write your code here.

	newStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])

		// fmt.Println(string(str[i]), string(newStr[i]))
	}

	return newStr == str

	// return false
}
