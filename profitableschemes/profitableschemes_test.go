package profitableschemes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// profitableSchemes
// https://leetcode.com/problems/profitable-schemes/
// dfs with memoization
// dp[i][j][k] - the number of schemes after i-th crime, used j people and gain at least k profit
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	MOD := 1000000007
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		// dp[i] is the number of crimes
		dp[i] = make([][]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, minProfit+1)
		}
	}
	dp[0][0][0] = 1
	for i := 1; i <= len(group); i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				dp[i%2][j][k] = dp[(i-1)%2][j][k] % MOD
				if j-group[(i-1)] >= 0 {
					preProfit := max(k-profit[i-1], 0)
					fmt.Printf("preprofit %d\n", preProfit)
					dp[i%2][j][k] += dp[(i-1)%2][j-group[(i-1)]][preProfit] % MOD
				}
			}
		}
	}
	res := 0
	for j := 0; j <= n; j++ {
		res += dp[len(group)%2][j][minProfit]
	}
	return res % MOD
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestProfitableSchemes(t *testing.T) {
	var cases = []struct {
		n         int
		minProfit int
		group     []int
		profit    []int
		expected  int
	}{
		{
			n:         5,
			minProfit: 3,
			group:     []int{2, 2},
			profit:    []int{2, 3},
			expected:  2,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := profitableSchemes(c.n, c.minProfit, c.group, c.profit)
			require.Equal(t, c.expected, result)
		})
	}
}
