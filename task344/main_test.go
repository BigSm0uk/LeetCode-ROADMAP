package task344

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func reverseString(s []byte)  {
	if len(s) == 1 {
		return
	}
    l, r := 0, len(s)-1
    for l < r {
        s[l], s[r] = s[r], s[l]

        l++
        r--
    }
}

func TestReverseString(t *testing.T) {
    cases := []struct {
        name string
        input []byte
        expected []byte
    }{
        {
            name: "case1",
            input: []byte{'h','e','l','l','o'},
            expected: []byte{'o','l','l','e','h'},
        },
        {
            name: "case2",
            input: []byte{'H','a','n','n','a', 'h'},
            expected: []byte{'h','a','n','n','a', 'H'},
        },
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            reverseString(c.input)
            td.Cmp(t, c.input, c.expected)
        })
    }
}