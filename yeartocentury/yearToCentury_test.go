package yeartocentury

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func yearToCentury(year int) int {
	var century float64
	century = float64(year) / 100.0
	if century == float64(int64(century)) {
		return int(century)
	} else {
		return int(century) + 1
	}
}

func TestYearToCentury(t *testing.T) {
	var cases = []struct {
		input    int
		expected int
	}{
		{
			input:    1705,
			expected: 18,
		},
		{
			input:    1700,
			expected: 17,
		},
	}
	for _, c := range cases {
		result := yearToCentury(c.input)
		require.Equal(t, c.expected, result)
	}

}
