package task13

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func romanToInt(s string) int {
	var m = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	prev := 0
	sb := []byte(s)
	out := 0
	for i := len(sb) - 1; i >= 0; i-- {
		cur := m[sb[i]]
		if cur >= prev {
			out += cur
		} else {
			out -= cur
		}
		prev = cur
	}
	return out

}

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, c := range cases {
		result := romanToInt(c.s)
		td.Cmp(t, result, c.expected)
	}
}
