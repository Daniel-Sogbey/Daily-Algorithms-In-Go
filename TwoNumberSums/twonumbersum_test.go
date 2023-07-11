package main

import "testing"

func TestTwoNumberSum(t *testing.T) {
	array := []int{-3, -1, -4, 5, 6, 7, 10, 11}
	target := 10

	expectedOutput := []int{-1, 11}
	outPut := TwoNumberSum1(array, target)

	if len(outPut) != len(expectedOutput) {
		t.Errorf("Expect the output to be of length %v, but got %v", len(expectedOutput), len(outPut))
	}
}
