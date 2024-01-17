package containermostwater

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	// memo :=
	maxArea := 0
	x1, x2 := 0, len(height)-1
	for {
		if x1 == x2 {
			break
		}
		area := getArea(x1, x2, height[x1], height[x2])
		if area > maxArea {
			maxArea = area
		}
		if height[x1] > height[x2] {
			x2--
		} else {
			x1++
		}
	}
	return maxArea
}

func getArea(x1, x2, y1, y2 int) int {
	x := x2 - x1
	y := min(y2, y1)
	result := x * y
	return result
}

func TestMaxArea(t *testing.T) {
	var cases = []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := maxArea(c.height)
			require.Equal(t, c.expected, result)
		})
	}
}
