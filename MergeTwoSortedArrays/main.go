package main

import (
	"fmt"
)

func main() {
	a := []int{1, 4, 7, 20}
	b := []int{3, 5, 6}
	fmt.Println(MergeTwoSortedArrays(a, b))
}

func MergeTwoSortedArrays(a []int, b []int) []int {
	i := 0
	j := 0
	ans := []int{}

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ans = append(ans, a[i])
			i++
		} else {
			ans = append(ans, b[j])
			j++
		}
	}

	for i < len(a) {
		ans = append(ans, a[i])
		i++
	}

	for j < len(b) {
		ans = append(ans, b[j])
		j++
	}

	return ans
}
