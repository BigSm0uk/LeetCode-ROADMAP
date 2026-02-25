package task167

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{1, 2}
	}
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{left + 1, right + 1}
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{
			name:     "case1",
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "case2",
			input:    []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			name:     "case3",
			input:    []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := twoSum(c.input, c.target)
			td.Cmp(t, result, c.expected)
		})
	}
}
