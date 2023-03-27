package isbalancedtree

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

func isBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := dfs(root.Left)
	isRightBalanced, rightHeight := dfs(root.Right)
	diff := abs(leftHeight - rightHeight)
	if isLeftBalanced && isRightBalanced && diff <= 1 {
		return true, 1 + max(leftHeight, rightHeight)
	}
	return false, -1
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestIsBalanced(t *testing.T) {
	var cases = []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 20,
					},
				},
			},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := isBalanced(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
