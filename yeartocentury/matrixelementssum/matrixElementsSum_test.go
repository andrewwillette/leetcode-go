package matrixelementssum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// After becoming famous, the CodeBots decided to move into a new building together.
// Each of the rooms has a different cost, and some of them are free, but there's
// a rumour that all the free rooms are haunted! Since the CodeBots are quite
// superstitious, they refuse to stay in any of the free rooms, or any of the rooms
// below any of the free rooms.
func matrixElementsSum(matrix [][]int) int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] == 0 {
				for k := i + 1; k < len(matrix); k++ {
					matrix[k][j] = 0
				}
			}
		}
	}
	matrixElementsSum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrixElementsSum += matrix[i][j]
		}
	}
	return matrixElementsSum
}

func TestMatrixElementsSum(t *testing.T) {
	var cases = []struct {
		input    [][]int
		expected int
	}{
		{
			// {{0, 1, 1, 2},
			// 	{0, 5, 0, 0},
			// 	{2, 0, 3, 3}},
			input: [][]int{{0, 1, 1, 2},
				{0, 5, 0, 0},
				{2, 0, 3, 3}},
			expected: 9,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := matrixElementsSum(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
