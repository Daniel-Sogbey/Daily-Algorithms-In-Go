package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 3
	fmt.Println(removeElement(nums, target))
}

func removeElement(nums []int, target int) int {
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if target != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	fmt.Println(nums)

	return slow
}
