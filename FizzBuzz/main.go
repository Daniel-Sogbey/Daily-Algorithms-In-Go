package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	answers := []string{}

	for i := 1; i <= n; i++ {
		answers = append(answers, strconv.Itoa(i))
	}

	for i, answer := range answers {
		currentAnswer, _ := strconv.Atoi(answer)
		fmt.Println(currentAnswer)
		if currentAnswer%3 == 0 {
			answers[i] = "Fizz"
		}

		if currentAnswer%5 == 0 {
			answers[i] = "Buzz"
		}

		if currentAnswer%3 == 0 && currentAnswer%5 == 0 {
			answers[i] = "FizzBuzz"
		}
	}

	return answers
}
