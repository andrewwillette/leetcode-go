package constructbintreefrompreorder

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

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}
	for i := 1; i < len(preorder); i++ {
		addNode(root, preorder[i])
	}
	return root
}

func addNode(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			addNode(root.Left, val)
		}
	}
	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			addNode(root.Right, val)
		}
	}
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	var cases = []struct {
		input    []int
		expected *TreeNode
	}{
		{
			input: []int{8, 5, 1, 7, 10, 12},
			expected: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 10,
					Right: &TreeNode{
						Val: 12,
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := bstFromPreorder(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}

func bstFromPreorderBeta(preorder []int) *TreeNode {
	fmt.Printf("bstFromPreorderBeta called with\n%v\n", preorder)
	if len(preorder) < 1 {
		return nil
	}
	bounds := 0
	for bounds < len(preorder) && preorder[bounds] <= preorder[0] {
		bounds++
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorderBeta(preorder[1:bounds]),
		Right: bstFromPreorderBeta(preorder[bounds:]),
	}
}

func TestBstFromPreorderBeta(t *testing.T) {
	bstFromPreorderBeta([]int{8, 5, 1, 7, 10, 12})
}
