package task205

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func isIsomorphic(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)

	if len(sb) != len(tb) {
		return false
	}
	ms := make(map[byte]byte, len(sb))
	mt := make(map[byte]byte, len(sb))

	for i := range sb {
		sch := sb[i]
		tch := tb[i]

		if v, ok := ms[sch]; ok {
			if v != tch {
				return false
			}
		} else {
			ms[sch] = tch
		}
		if v, ok := mt[tch]; ok {
			if v != sch {
				return false
			}
		} else {
			mt[tch] = sch
		}
	}
	return true
}

func TestIsIsomorphic(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}
	for _, c := range cases {
		result := isIsomorphic(c.s, c.t)
		td.Cmp(t, result, c.expected)
	}
}
