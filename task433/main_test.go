package task433

import (
	"strconv"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)
func compress(chars []byte) int {
    if len(chars) == 1 {
        return 1
    }
    left := 0
	right :=0
	for right < len(chars) {
		count :=0
		ch := chars[right]
		for right < len(chars) && chars[right] == ch {
			right++
			count++
		}
		chars[left] = ch
		left++
		if count > 1 {
			for _, v := range strconv.Itoa(count) {
				chars[left] = byte(v)
				left++
			}
		}
	}
    return left
}

func TestCompress(t *testing.T) {
	cases := []struct {
		name string
		input []byte
		expected int
	}{
		{
			name: "case1",
			input: []byte{'a','a','b','b','c','c','c'},
			expected: 6,
		},
		{
			name: "case2",
			input: []byte{'a'},
			expected: 1,
		},
		{
			name: "case3",
			input: []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'},
			expected: 4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := compress(c.input)
			td.Cmp(t, result, c.expected)
		})
	}
}