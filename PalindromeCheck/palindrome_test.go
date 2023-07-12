package main

import "testing"

func TestPalindromeCheck(t *testing.T) {
	str := "abcdcba"

	if len(str) <= 0 {
		t.Errorf("Expected string of length greater than 0, but go %v", len(str))
	}

}
