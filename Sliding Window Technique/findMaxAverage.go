//Leetcode Problem: 643. Maximum Average Subarray I

package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("Hello, World!")
	k := 4
	nums := []int{1, 12, -5, -6, 50, 3}
	answer := findMaxAverageBad(nums, k)
	fmt.Println(answer)
}

func findMaxAverageBad(nums []int, k int) float64 {
	max := math.Inf(-1)

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := 0; j < k; j++ {
			if !(i+(k-1) > len(nums)-1) {
				sum += nums[i+j]
			} else {
				i = len(nums)
				j = 4
			}
		}

		if sum != 0 {
			max = math.Max(max, float64(sum)/float64(k))
		}
	}

	if max == math.Inf(-1) {
		return 0.0
	}

	return max
}
