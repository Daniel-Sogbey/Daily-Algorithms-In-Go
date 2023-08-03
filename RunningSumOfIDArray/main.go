package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	outPut := []int{nums[0]}
	runningSum := nums[0]

	for i := 1; i < len(nums); i++ {
		runningSum = nums[i] + runningSum
		fmt.Println(runningSum)
		outPut = append(outPut, runningSum)
	}

	return outPut

}
