package main

import "fmt"

func main() {

	fmt.Println(NumberOfSteps(5))

}

func NumberOfSteps(num int) int {
	var step = 0

	for {
		if num == 0 {
			break
		} else if num%2 == 0 {
			num /= 2
			step++
		} else {
			num--
			step++
		}

	}

	return step

}
