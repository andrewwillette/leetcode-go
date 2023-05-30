package countpoints

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/
func countPoints(points [][]int, queries [][]int) []int {
	result := []int{}
	for _, circle := range queries {
		pointsInCircle := 0
		for _, point := range points {
			if isPointInCircle(point, circle) {
				pointsInCircle++
			}
		}
		result = append(result, pointsInCircle)
	}
	return result
}

// point is [x,y]
// circle is [x, y, rad]
func isPointInCircle(point []int, circle []int) bool {
	result := math.Sqrt(float64(math.Pow(float64(point[0]-circle[0]), 2) + math.Pow(float64(point[1]-circle[1]), 2)))
	return result <= float64(circle[2])
}

func TestCountPoints(t *testing.T) {
	var cases = []struct {
		points   [][]int
		queries  [][]int
		expected []int
	}{
		{
			points:   [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
			queries:  [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
			expected: []int{3, 2, 2},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.points), func(t *testing.T) {
			result := countPoints(c.points, c.queries)
			require.Equal(t, c.expected, result)
		})
	}
}
