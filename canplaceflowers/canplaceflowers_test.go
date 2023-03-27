package canplaceflowers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, plt := range flowerbed {
		if plt == 1 {
			continue
		}
		if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}

func TestCanPlaceFlowers(t *testing.T) {
	var cases = []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := canPlaceFlowers(c.flowerbed, c.n)
			require.Equal(t, c.expected, result)
		})
	}
}
