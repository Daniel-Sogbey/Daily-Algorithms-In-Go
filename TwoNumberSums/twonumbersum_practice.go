package main

import "sort"

// solution 1 : O(n*2) time | O(1) space
func twoNumberSum1(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		firstNumber := array[i]
		for j := 0; j < len(array); j++ {
			secondNumber := array[j]
			if firstNumber+secondNumber == target && firstNumber != secondNumber {
				return []int{firstNumber, secondNumber}
			}
		}
	}
	return []int{}
}

// solution 2 : O(n) space | O(n) time
func twoNumberSum2(array []int, target int) []int {
	hashMap := make(map[int]bool)

	for i := 0; i < len(array); i++ {
		currentNum := array[i]
		complement := target - currentNum

		if !hashMap[complement] {
			hashMap[currentNum] = true
		} else {
			return []int{complement, currentNum}
		}
	}

	return []int{}
}

// solution 3 : O(nlogn) Time : O(1) space
func twoNumberSum3(array []int, target int) []int {

	sort.Ints(array)

	left := 0

	right := len(array) - 1

	for left < right {

		currentSum := array[left] + array[right]

		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left = left + 1
		} else if currentSum > target {
			right = right - 1
		}
	}

	return []int{}
}
