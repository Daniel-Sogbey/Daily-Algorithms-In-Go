package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
	letters := "xyz"
	key := 2
	fmt.Println(caeserCipherEncryptor(letters, key))
	fmt.Println(CaeserCipherEncryptorPractice(letters, key))
}

// solution 1 : O(n) Time | O(n) Space
func caeserCipherEncryptor(str string, key int) string {
	var newLetters []string
	var newKey int = key % 26

	for _, letter := range str {
		newLetters = append(newLetters, string(getNewLetter(letter, newKey)))
	}

	return strings.Join(newLetters, "")
}

func getNewLetter(r rune, i int) rune {
	newLetterCode := int(r) + i

	if newLetterCode <= 122 {
		return rune(newLetterCode)
	} else {
		return rune(96 + newLetterCode%122)
	}

}
