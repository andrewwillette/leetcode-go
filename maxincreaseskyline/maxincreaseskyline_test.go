package maxincreaseskyline

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/max-increase-to-keep-city-skyline/
func maxIncreaseKeepingSkyline(grid [][]int) int {
	result := 0
	for rowi, row := range grid {
		for columni, columnval := range row {
			maxrow := getRowMax(rowi, grid)
			maxcolumn := getColumnMax(columni, grid)
			skylineIncrease := minimum(maxrow, maxcolumn)
			result += skylineIncrease - columnval
		}
	}
	return result
}

func getRowMax(index int, grid [][]int) int {
	max := 0
	for _, val := range grid[index] {
		max = maximum(max, val)
	}
	return max
}

func getColumnMax(index int, grid [][]int) int {
	max := 0
	for _, val := range grid {
		for columni, columnval := range val {
			if columni == index {
				max = maximum(max, columnval)
			}
		}
	}
	return max
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	var cases = []struct {
		grid     [][]int
		expected int
	}{
		{
			grid:     [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}},
			expected: 35,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := maxIncreaseKeepingSkyline(c.grid)
			require.Equal(t, c.expected, result)
		})
	}
}
