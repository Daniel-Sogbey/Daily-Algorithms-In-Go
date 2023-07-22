package main

// solution 4 : Time O(nlogn) : Space (O(1))
func twoNumberSum4(array []int, target int) []int {

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

func twoNumberSum5(array []int, target int) []int {

	hashMap := make(map[int]bool)

	for i := 0; i < len(array); i++ {
		var currentNumber int = array[i]
		var complement int = target - currentNumber

		if !hashMap[complement] {
			hashMap[currentNumber] = true
		} else {
			return []int{complement, currentNumber}
		}

	}

	return []int{}
}
