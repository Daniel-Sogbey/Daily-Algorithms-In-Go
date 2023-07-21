package main

import "strings"

func CaeserCipherEncryptorPractice(str string, key int) string {
	// factor 26 letters of the English alphabet
	// str = "xyz"
	newLetters := []string{}
	newKey := key % 26

	for _, letter := range str {

		newLetters = append(newLetters, string(getNewRunes(letter, newKey)))
	}

	return strings.Join(newLetters, "")
}

func getNewRunes(r rune, newKey int) rune {
	newLetterCode := int(r) + newKey
	if newLetterCode <= 122 {
		return rune(newLetterCode)
	} else {
		return rune(96 + newLetterCode%122)
	}
}
