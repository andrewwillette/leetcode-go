package symmetrictree

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

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func TestIsBinaryComplete(t *testing.T) {
	var cases = []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3,
					Left: &TreeNode{Val: 6},
				},
			},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := isSymmetric(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
