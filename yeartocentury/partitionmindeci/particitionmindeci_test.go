package partitionmindeci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func findMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minPartitions(n string) int {
	largestDigit := 0
	for _, v := range n {
		val := int(v - '0')
		fmt.Printf("%v\n", val)
		largestDigit = findMax(largestDigit, val)
	}
	return largestDigit
}

func TestMinPartitions(t *testing.T) {
	var cases = []struct {
		input    string
		expected int
	}{
		{
			input:    "32",
			expected: 3,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := minPartitions(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
