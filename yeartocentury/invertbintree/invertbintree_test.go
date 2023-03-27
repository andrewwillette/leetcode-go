package invertbintree

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

func invertTree(root *TreeNode) *TreeNode {
	invertNode(root)
	return root
}

func invertNode(root *TreeNode) {
	if root != nil {
		tmp := root.Left
		root.Left = root.Right
		root.Right = tmp
		invertTree(root.Left)
		invertTree(root.Right)
	}
}

func TestIsBinaryComplete(t *testing.T) {
	var cases = []struct {
		input    *TreeNode
		expected *TreeNode
	}{
		{
			input: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := invertTree(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
