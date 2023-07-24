package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string = "AAAAAAAAAAAAABBCCCCDD"

	fmt.Println(RunLengthEncoding(str))
}

func RunLengthEncoding(str string) string {
	var encodedListCharacters []string
	currentRunLength := 1

	for i := 1; i < len(str); i++ {
		// fmt.Println(string(str[i]))
		currentCharacter := str[i]
		previousCharacter := str[i-1]

		if string(currentCharacter) != string(previousCharacter) || currentRunLength == 9 {
			encodedListCharacters = append(encodedListCharacters, strconv.Itoa(currentRunLength))
			encodedListCharacters = append(encodedListCharacters, string(previousCharacter))

			currentRunLength = 0
		}

		currentRunLength = currentRunLength + 1

	}

	encodedListCharacters = append(encodedListCharacters, strconv.Itoa(currentRunLength))
	encodedListCharacters = append(encodedListCharacters, string(str[len(str)-1]))

	return strings.Join(encodedListCharacters, "")

}
