package task283

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	slow := 0
	for nums[slow] != 0 && slow < len(nums)-1 {
		slow++
	}
	for i := slow + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow], nums[i] = nums[i], nums[slow]
			slow++
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "case1",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "case2",
			input:    []int{0},
			expected: []int{0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			moveZeroes(c.input)
			td.Cmp(t, c.input, c.expected)
		})
	}
}
