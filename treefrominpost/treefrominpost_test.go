package treefrominpost

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/**
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	lenpost := len(postorder)
	root := TreeNode{Val: postorder[lenpost-1]}
	mid := findIndex(inorder, postorder[lenpost-1])
	root.Left = buildTree(inorder[:mid], postorder[:mid])
	root.Right = buildTree(inorder[mid+1:], postorder[mid:lenpost-1])
	return &root
}

func findIndex(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

// TODO: Implement this
func arrayFromTreeNode(*TreeNode) []int {
	return []int{}
}

func TestBuildTree(t *testing.T) {

	var cases = []struct {
		inorder   []int
		postorder []int
		expected  []int
	}{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			expected:  []int{3, 9, 20, -1, -1, 15, 7},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.inorder), func(t *testing.T) {
			result := buildTree(c.inorder, c.postorder)

			require.Equal(t, c.expected, result)
		})
	}
}
