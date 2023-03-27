package longestcycle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// You are given a directed graph of n nodes numbered from 0 to n - 1, where
// each node has at most one outgoing edge.

// The graph is represented with a given 0-indexed array edges of size n,
// indicating that there is a directed edge from node i to node edges[i].
// If there is no outgoing edge from node i, then edges[i] == -1.

// Return the length of the longest cycle in the graph. If no cycle exists, return -1.

// A cycle is a path that starts and ends at the same node.
func longestCycleBeta(edges []int) int {
	result := -1
	for i := range edges {
		cycleLength := findCycle(i, edges)
		if cycleLength > result {
			result = cycleLength
		}
	}
	return result
}

func findCycle(currentNode int, edges []int) int {
	root := currentNode
	visited := make([]bool, len(edges))
	// visited[currentNode] = true
	// currentNode = edges[currentNode]
	cycleLength := 0
	for {
		// a cycle occurred
		// if nextNode == index {
		// 	return cycleLength
		// }
		currentNode = edges[currentNode]
		// no connected nodes
		if currentNode == -1 {
			break
		}
		if visited[currentNode] {
			break
		}
		visited[currentNode] = true
		cycleLength++
		if currentNode == root {
			return cycleLength
		}
	}
	return -1
}

func TestFindCycle(t *testing.T) {
	var cases = []struct {
		index    int
		edges    []int
		expected int
	}{
		// {
		// 	index:    0,
		// 	edges:    []int{3, 3, 4, 2, 3},
		// 	expected: -1,
		// },
		// {
		// 	index:    3,
		// 	edges:    []int{3, 3, 4, 2, 3},
		// 	expected: 3,
		// },
		// {
		// 	index:    4,
		// 	edges:    []int{3, 3, 4, 2, 3},
		// 	expected: 3,
		// },
		// {
		// 	index:    1,
		// 	edges:    []int{3, 3, 4, 2, 3},
		// 	expected: -1,
		// },
		{
			index:    1,
			edges:    []int{2, -1, 3, 1},
			expected: -1,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := findCycle(c.index, c.edges)
			require.Equal(t, c.expected, result)
		})
	}
}

func TestLongestCycleBeta(t *testing.T) {
	var cases = []struct {
		edges    []int
		expected int
	}{
		{
			// 0 -> 3 , 1 -> 3, 2 -> 4, 3 -> 2, 4 -> 3
			edges:    []int{3, 3, 4, 2, 3},
			expected: 3,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := longestCycleBeta(c.edges)
			require.Equal(t, c.expected, result)
		})
	}
}

func longestCycle(edges []int) int {
	n := len(edges)
	nodeTraversalOrder := make([]int, n)
	visited := func(i int) bool {
		return nodeTraversalOrder[i] > 0
	}
	traversalNumber := 1
	res := -1
	for i := range edges {
		t0 := traversalNumber // start order of current DFS
		if visited(i) {
			continue
		}
		nodeTraversalOrder[i] = traversalNumber
		traversalNumber++
		// Iterate until end or reaching a previously visited node
		first, second := i, edges[i]
		for second != -1 && !visited(second) {
			nodeTraversalOrder[second] = traversalNumber
			first = second
			second = edges[second]
			traversalNumber++
		}
		if second != -1 &&
			nodeTraversalOrder[second] >= t0 { // if time of second is > start time of DFS, it's a cycle
			res = max(res, nodeTraversalOrder[first]-nodeTraversalOrder[second]+1)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLongestCycle(t *testing.T) {
	var cases = []struct {
		edges    []int
		expected int
	}{
		{
			// 0 -> 3 , 1 -> 3, 2 -> 4, 3 -> 2, 4 -> 3
			edges:    []int{3, 3, 4, 2, 3},
			expected: 3,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := longestCycle(c.edges)
			require.Equal(t, c.expected, result)
		})
	}
}
