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
	var newHead, tail *ListNode
	currentSum := 0
	for head != nil {
		currentSum += head.Val
		if head.Val == 0 && currentSum != 0 {
			newNode := &ListNode{Val: currentSum}
			if newHead == nil {
				newHead = newNode
				tail = newNode
			} else {
				tail.Next = newNode
				tail = newNode
			}
			currentSum = 0
		}
		head = head.Next
	}
	return newHead
}

func TestMergeNodes(t *testing.T) {
	var cases = []struct {
		input    *ListNode
		expected *ListNode
	}{
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
