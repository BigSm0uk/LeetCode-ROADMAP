package task1

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
func twoSumWithHashMap(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{m[target-v], i}
		}
		m[v] = i
	}
	return []int{}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := generateLargeArray(100000)
	target := 99998
	b.ResetTimer()
	for b.Loop() {
		twoSum(nums, target)
	}
}
func BenchmarkTwoSumWithHashMap(b *testing.B) {
	nums := generateLargeArray(100000)
	target := 99998
	b.ResetTimer()
	for b.Loop() {
		twoSumWithHashMap(nums, target)
	}
}
func generateLargeArray(size int) []int {
	arr := make([]int, size)
	for i := range size {
		arr[i] = i * 2 // например: 0, 2, 4, 6, ...
	}
	return arr
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
			expected: []int{0, 1},
		},
		{
			name:     "case2",
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "case3",
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := twoSum(c.input, c.target)
			td.Cmp(t, result, c.expected)
		})
	}
}
