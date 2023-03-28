package minpathsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func minPathSum(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinPathSum(t *testing.T) {
	var cases = []struct {
		grid     [][]int
		expected int
	}{
		{
			grid:     [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expected: 7,
		},
		// {
		// 	grid:     [][]int{{1, 2, 3}, {4, 5, 6}},
		// 	expected: 12,
		// },
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := minPathSum(c.grid)
			require.Equal(t, c.expected, result)
		})
	}
}
