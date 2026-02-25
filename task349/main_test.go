package task349

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	out := make([]int, 0)

	for _, num2 := range nums2 {
		m[num2] = struct{}{}
	}
	for _, num1 := range nums1 {
		if _, ok := m[num1]; ok {
			out = append(out, num1)
			delete(m, num1)
		}
	}
	return out
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "case1",
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			name:     "case2",
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{9, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := intersection(c.nums1, c.nums2)
			td.Cmp(t, result, td.Bag(td.Flatten(c.expected)))
		})
	}
}
