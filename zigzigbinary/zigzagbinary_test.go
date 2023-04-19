package zigzagbinary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// longestZigZag
// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
func longestZigZag(root *TreeNode) int {
	return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func dfs(root *TreeNode, isLeft bool, count int) int {
	if root == nil {
		return count
	}
	if isLeft {
		return max(dfs(root.Left, true, 0), dfs(root.Right, false, count+1))
	} else {
		return max(dfs(root.Left, true, count+1), dfs(root.Right, false, 0))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLongestZigZag(t *testing.T) {
	var cases = []struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val: 1,
								Right: &TreeNode{
									Val: 1,
								},
							},
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			expected: 3,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := longestZigZag(c.root)
			require.Equal(t, c.expected, result)
		})
	}
}
