package greatersumtree

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

func bstToGst(root *TreeNode) *TreeNode {
	var accumulationVal = 0
	var dfs func(root *TreeNode) *TreeNode

	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		dfs(root.Right)
		accumulationVal += root.Val
		root.Val = accumulationVal
		dfs(root.Left)

		return root
	}
	return dfs(root)
}

func TestBstToGst(t *testing.T) {
	var cases = []struct {
		input    *TreeNode
		expected *TreeNode
	}{
		{
			input: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 7,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			expected: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 36,
					Left: &TreeNode{
						Val: 36,
					},
					Right: &TreeNode{
						Val: 35,
						Right: &TreeNode{
							Val: 33,
						},
					},
				},
				Right: &TreeNode{
					Val: 21,
					Left: &TreeNode{
						Val: 26,
					},
					Right: &TreeNode{
						Val: 15,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
		{
			input: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val: 1,
				},
			},
			expected: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := bstToGst(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
