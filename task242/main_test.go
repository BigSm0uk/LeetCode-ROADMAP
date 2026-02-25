package task242

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int)
	for k, _ := range s {
		sMap[s[k]] += 1
		sMap[t[k]] -= 1
	}
	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}
	return true
}
func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}
	for _, c := range cases {
		result := isAnagram(c.s, c.t)
		td.Cmp(t, result, c.expected)
	}
}
