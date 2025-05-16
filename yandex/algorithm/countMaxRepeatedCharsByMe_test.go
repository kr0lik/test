package algorithm

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
)

var ErrInvalidInput = errors.New("invalid input")

func CountMaxRepeatedCharsByMe(input string) (int, error) {
	re := regexp.MustCompile("[A-Z]+")

	if !re.MatchString(input) {
		return 0, ErrInvalidInput
	}

	if len(input) < 2 {
		return len(input), nil
	}

	maxLen := 0
	lastChar := ""
	currentLen := 0

	for _, s := range input {
		if string(s) != lastChar {
			lastChar = string(s)
			currentLen = 0
		}

		currentLen++

		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	if currentLen > maxLen {
		maxLen = currentLen
	}

	return maxLen, nil
}

func TestCountMaxRepeatedCharsByMe(t *testing.T) {
	var tests = []struct {
		input  string
		result int
	}{
		{"12313", 0},
		{"", 0},
		{"A", 1},
		{"AB", 1},
		{"AAAB", 3},
		{"AABBBCC", 3},
		{"ABABAAAB", 3},
		{"AABBBCCCD", 3},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v,%v", test.input, test.result)
		t.Run(testName, func(t *testing.T) {
			res, _ := CountMaxRepeatedCharsByMe(test.input)
			if res != test.result {
				t.Errorf("%v => %d, got %d", test.input, test.result, res)
			}
		})
	}
}
