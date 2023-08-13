package main

import "testing"

func TestMain(t *testing.T) {
	type TestCase struct {
		n     int
		wants []string
	}

	var testCases = []TestCase{
		{n: 2, wants: []string{"1", "Fizz"}},
		{n: 3, wants: []string{"1", "2", "Fizz"}},
		{n: 5, wants: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{n: 15, wants: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, tc := range testCases {
		got := fizzBuzz(tc.n)

		result := compareSlices(tc.wants, got)

		if !result {
			t.Errorf("Test case failed : Expect value %v but got value %v", tc.wants, got)
		}

	}

}

func compareSlices(slice1 []string, slice2 []string) bool {

	indexOfSlice1 := 0
	indexOfSlice2 := 0

	for indexOfSlice1 < len(slice1) && indexOfSlice2 < len(slice2) {
		if slice1[indexOfSlice1] == slice2[indexOfSlice2] {
			indexOfSlice1++
			indexOfSlice2++
		} else {
			return false
		}
	}

	return true
}
