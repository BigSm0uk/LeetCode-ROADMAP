package task165

import (
	"strconv"
	"strings"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func compareVersion(version1 string, version2 string) int {
    i, j := 0,0
    
    v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")

    for i < len(v1) && j < len(v2) {
        v1num,_ := strconv.Atoi(v1[i])
        v2num,_ := strconv.Atoi(v2[j])

        if v1num > v2num {
            return 1
        }
        if v1num < v2num {
            return -1
        }
        i++
        j++
    }
    for i < len(v1) {
        v1num,_ := strconv.Atoi(v1[i])
        if v1num != 0 {
            return 1
        } 
        i++
    }
    for j < len(v2) {
        v2num,_ := strconv.Atoi(v2[j])
        if v2num != 0 {
            return -1
        } 
        j++
    }
    return 0
}
func TestCompareVersion(t *testing.T) {
	cases := []struct {
		name     string
		version1 string
		version2 string
		expected int
	}{
		{
			name:     "case1",
			version1: "1.2",
			version2: "1.10",
			expected: -1,
		},
		{
			name:     "case2",
			version1: "1.01",
			version2: "1.001",
			expected: 0,
		},
		{
			name:     "case3",
			version1: "1.0",
			version2: "1.0.0.0",
			expected: 0,
		},
		{
			name:     "case4",
			version1: "1.0.1",
			version2: "1",
			expected: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := compareVersion(c.version1, c.version2)
			td.Cmp(t, result, c.expected)
		})
	}
}
