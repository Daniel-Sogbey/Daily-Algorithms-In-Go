package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	s := "ace"
	t := "abcde"
	fmt.Println(IsSubsequence(t, s))
}

func IsSubsequence(t string, s string) bool {
	slow := 0
	fast := 0

	for slow < len(s) && fast < len(t) {
		if s[slow] == t[fast] {
			slow++
		}
		fast++
	}
	return slow == len(s)
}
