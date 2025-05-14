package algorithm

import (
	"fmt"
	"testing"
)

func MaxOnesByInterviewer(input []int) int {
	maxLen := 0
	curLen := 0
	prevLen := 0

	for _, num := range input {
		if num == 1 {
			curLen++
			maxLen = max(maxLen, curLen+prevLen)
		} else {
			prevLen = curLen
			curLen = 0
		}
	}

	if maxLen == len(input) {
		maxLen--
	}

	return maxLen
}

func TestMaxOnesByInterviewer(t *testing.T) {
	var tests = []struct {
		input  []int
		result int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{1, 1, 0, 1}, 3},
		{[]int{1, 1, 0, 0, 1}, 2},
		{[]int{1, 0, 1, 0, 1, 0, 1}, 2},
		{[]int{1, 1, 0, 0, 1, 1, 1, 1, 1}, 5},
		{[]int{1, 1, 0, 0, 0, 1, 1}, 2},
		{[]int{1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1}, 6},
		{[]int{1, 1, 1}, 2},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v,%v", test.input, test.result)
		t.Run(testName, func(t *testing.T) {
			res := MaxOnesByInterviewer(test.input)
			if res != test.result {
				t.Errorf("%v => %d, got %d", test.input, test.result, res)
			}
		})
	}
}
