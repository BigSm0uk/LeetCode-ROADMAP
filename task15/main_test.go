package task15

import (
	"sort"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func threeSum(nums []int) [][]int {
	out := make([][]int, 0)
	target := 0
	m := map[[3]int]struct{}{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				key := [3]int{nums[i], nums[left], nums[right]}
				if _, ok := m[key]; !ok {
					out = append(out, key[:])
					m[key] = struct{}{}
				}

			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return out
}

func TestThreeSum(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:     "case1",
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "case2",
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "case3",
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := threeSum(c.input)
			td.Cmp(t, result, td.Bag(td.Flatten(c.expected)))
		})
	}
}
