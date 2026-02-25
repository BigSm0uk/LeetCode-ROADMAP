package task977

import (
	"math"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func sortedSquares(nums []int) []int {
	out := make([]int, len(nums))
	l,r, i := 0, len(nums)-1, len(nums)-1
	for l < r {
		if math.Abs(float64(nums[l])) > math.Abs(float64(nums[r])) {
			out[i] = nums[l]*nums[l]
			l++
		} else {
			out[i] = nums[r]*nums[r]
			r--
		}
		i--
	}
	if i == 0 {
		out[i] = nums[l]*nums[l]
	}
	return out
}

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "case1",
			input:    []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "case2",
			input:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := sortedSquares(c.input)
			td.Cmp(t, result, c.expected)
		})
	}
}
