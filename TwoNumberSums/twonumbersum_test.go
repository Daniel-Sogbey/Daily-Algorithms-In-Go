package main

import (
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	array := []int{-3, -1, -4, 5, 6, 7, 10, 11}
	target := 10

	expectedOutput := []int{-1, 11}
	outPut := TwoNumberSum1(array, target)

	if len(outPut) != len(expectedOutput) {
		t.Errorf("Expect the output to be of length %v, but got %v", len(expectedOutput), len(outPut))
	}
}

type testCase struct {
	input  []int
	output []int
	target int
	isErr  bool
}

func TestTwoSum(t *testing.T) {

	testCases := []testCase{
		{
			input:  []int{-3, -1, -4, 5, 6, 7, 10, 11},
			output: []int{-1, 11},
			target: 10,
			isErr:  false,
		},
		{
			input:  []int{-3, 1, 4, 5, 6, 7, 10, 11},
			output: []int{1, 7},
			target: 8,
			isErr:  false,
		},
	}

	for _, tt := range testCases {
		output := twoNumberSum1(tt.input, tt.target)

		if !check(output, tt.input) {
			t.Errorf("Expected %d, got %d", tt.output, output)
		}
	}
}

func check(a []int, b []int) bool {
	i := 0

	for j := 0; j < len(b); j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}
