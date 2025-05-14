package algorithm

import (
	"fmt"
	"testing"
)

func MaxOnesByMe(input []int) int {
	maxLen := 0

	leftCnt := 0
	rightCnt := 0

	for i, num := range input {
		if num == 0 {
			rightCnt = 0

			for j := i + 1; j < len(input); j++ {
				num = input[j]

				if num == 0 {
					break
				}

				rightCnt++
			}

			currLen := leftCnt + rightCnt
			if currLen > maxLen {
				maxLen = currLen
			}

			leftCnt = 0
			rightCnt = 0

			continue
		}

		leftCnt++
	}

	currLen := leftCnt + rightCnt
	if currLen > maxLen {
		maxLen = currLen
	}

	if maxLen == len(input) {
		maxLen--
	}

	return maxLen
}

func TestMaxOnesByMe(t *testing.T) {
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
			res := MaxOnesByMe(test.input)
			if res != test.result {
				t.Errorf("%v => %d, got %d", test.input, test.result, res)
			}
		})
	}
}
