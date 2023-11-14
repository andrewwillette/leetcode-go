package uniquepaths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/unique-paths/
func uniquePaths(m int, n int) int {
	var memo = make(map[*int]int, m)
	// var dfs func(
	return 0
}

func TestUniquePaths(t *testing.T) {
	var cases = []struct {
		m        int
		n        int
		expected int
	}{
		{},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := uniquePaths(c.m, c.n)
			require.Equal(t, c.expected, result)
		})
	}
}
