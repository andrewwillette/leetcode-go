package mergenodes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-nodes-in-between-zeros/
func mergeNodes(head *ListNode) *ListNode {
	result := new(ListNode)
	currentNode := result
	summation := 0
	node := head.Next
	for {
		if node.Val == 0 {
			result.Val = summation
			result.Next = new(ListNode)
			result = result.Next
			node = node.Next
			continue
		} else {
			summation += result.Val
			node = node.Next
		}
		if node.Next == nil {
			break
		}
	}
	return result
}

func TestMergeNodes(t *testing.T) {
	var cases = []struct {
		input    *ListNode
		expected *ListNode
	}{
		// Input: head = [0,3,1,0,4,5,2,0]
		// Output: [4,11]
		{
			input: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 2,
										Next: &ListNode{
											Val: 0,
										},
									},
								},
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 11,
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := mergeNodes(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
