package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 3, 5, 5, 7}
	fmt.Println(removeDuplicateFromSortedArray(nums))
}

func removeDuplicateFromSortedArray(nums []int) int {
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow = slow + 1
			nums[slow] = nums[fast]
		}
	}
	fmt.Println(nums)
	return slow + 1
}
