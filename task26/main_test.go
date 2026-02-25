package task26

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[cur] {
			continue
		}
		nums[cur+1] = nums[i]
		cur++
	}
	return cur + 1
}

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "case1",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "case2",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := removeDuplicates(c.input)
			td.Cmp(t, k, c.expected)
		})
	}
}
