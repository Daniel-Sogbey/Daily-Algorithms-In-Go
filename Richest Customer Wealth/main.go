package main

import "fmt"

func main() {
	accounts := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println(maximumWealth(accounts))
}

func maximumWealth(accounts [][]int) int {
	highestSum := 0
	currentSum := 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			currentSum += accounts[i][j]
		}

		if currentSum > highestSum {
			highestSum = currentSum
		}
		currentSum = 0
	}
	return highestSum
}
