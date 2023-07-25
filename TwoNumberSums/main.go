package main

import (
	"fmt"
	"sort"
)

/*
Two Number Sum •
Write a function that takes in a non-empty array of distinct integers and an integer representing a target sum. If any two numbers in the input array sum up to the target sum, the function should return them in an array, in any order. If no two numbers sum up to the target sum, the function should return an empty array.
Note that the target sum has to be obtained by summing two different integers in the array: vou can't add a single integer to itself in order to obtain the target sum.
You can assume that there will be at most one pair of numbers summing up to the target sum.
*/

func main() {
	// TwoNumberSum1()
	array := []int{-3, -1, -4, 5, 6, 7, 10, 11}
	target := 10
	fmt.Println(twoNumberSum1(array, target))
	fmt.Println(twoNumberSum2(array, target))
	fmt.Println(twoNumberSum3(array, target))
	fmt.Println(twoNumberSum4(array, target))
	fmt.Println(twoNumberSum5(array, target))
	fmt.Println(twoSum(array, target))
}

//Solution 1

// O(n^2) Time | O(1) space
func TwoNumberSum1(array []int, target int) []int {
	// Write your code here.
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i]+array[j] == target && array[i] != array[j] {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}

// Solution 2
// O(n) Time | O(n) space
func TwoNumberSum2(array []int, target int) []int {
	im := make(map[int]bool)

	for i := 0; i < len(array); i++ {
		complement := target - array[i]
		if !im[complement] {
			im[array[i]] = true
		} else {
			return []int{array[i], complement}
		}
	}

	return []int{}
}

// Solution 3
// O(nlogn) Time | O(1) space
func TwoNumberSum3(array []int, target int) []int {
	// Write your code here.

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

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]bool)

	for i, num := range nums {
		complement := target - num

		if !hashMap[complement] {
			hashMap[num] = true
		} else {
			return []int{i, indexOf(complement, nums)}
		}
	}

	return []int{}
}

func indexOf(value int, source []int) int {
	for i, num := range source {
		if num == value {
			return i
		}
	}

	return -1
}
