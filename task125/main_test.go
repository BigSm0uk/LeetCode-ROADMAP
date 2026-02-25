package task125

import (
	"strings"
	"testing"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.Trim(strings.ToLower(s), " ")

	l, r := 0, len(s)-1
	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "case1",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "case2",
			input:    "race a car",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := isPalindrome(c.input)
			if result != c.expected {
				t.Errorf("Expected %v, got %v", c.expected, result)
			}
		})
	}
}
