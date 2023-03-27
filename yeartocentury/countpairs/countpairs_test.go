package countpairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return *new(stack), 0
	}
	return s[:l-1], s[l-1]
}

func countPairsBeta(n int, edges [][]int) int64 {
	if n == 1 {
		return 0
	}
	if len(edges) == 0 {
		return int64((n / 2) * (n - 1))
	}
	nodesNotConnected := int64(0)
	connected := make(map[int][]int)
	// for each node, build array of nodes that it can traverse
	for i := 0; i < n; i++ {
		connected[i] = dfsBeta(i, edges)
	}
	// fmt.Printf("edges: %v\n", edges)
	// fmt.Printf("connected: %v\n", connected)
	for _, v := range connected {
		nodesNotConnected += int64(n) - int64(len(v))
	}
	return nodesNotConnected / 2
}

func countPairs(n int, edges [][]int) int64 {
	var ans int64 = 0
	adjList := make([][]int, n)
	visited := make([]bool, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			var len int = 0
			dfs(i, adjList, visited, &len)
			ans += int64(len) * int64(n-len)
		}
	}
	return ans / 2
}

func dfs(curr int, adjacentList [][]int, visited []bool, len *int) {
	if visited[curr] {
		return
	}
	visited[curr] = true
	*len++
	for _, next := range adjacentList[curr] {
		dfs(next, adjacentList, visited, len)
	}
}

// return all connected nodes for node
func dfsBeta(curr int, edges [][]int) []int {
	// fmt.Printf("dfs called with node: %v\n", node)
	var stack stack
	stack = stack.Push(curr)
	visited := []int{curr}
	// fmt.Printf("edges: %v\n", edges)
	nodeToFollow := stack[len(stack)-1]
	for {
		// fmt.Printf("stack: %v\n", stack)
		// fmt.Printf("visited: %v\n", visited)
		if len(stack) == 0 {
			break
		}
		// fmt.Printf("nodeToFollow: %v\n", nodeToFollow)
		for i := 0; i < len(edges); i++ {
			if edges[i][0] == nodeToFollow && !valueExistsInVisited(edges[i][1], visited) {
				connectedNode := edges[i][1]
				// fmt.Printf("connectedNode: %v\n", connectedNode)
				stack = stack.Push(connectedNode)
				visited = append(visited, connectedNode)
				nodeToFollow = connectedNode
				break
			} else if edges[i][1] == nodeToFollow && !valueExistsInVisited(edges[i][0], visited) {
				connectedNode := edges[i][0]
				// fmt.Printf("connectedNode: %v\n", connectedNode)
				stack = stack.Push(connectedNode)
				visited = append(visited, connectedNode)
				nodeToFollow = connectedNode
				break
			}
			if i == len(edges)-1 {
				stack, _ = stack.Pop()
				if len(stack) > 0 {
					nodeToFollow = stack[len(stack)-1]
				}
			}
		}
	}

	return visited
}

func valueExistsInVisited(val int, visited []int) bool {
	for i := 0; i < len(visited); i++ {
		if visited[i] == val {
			return true
		}
	}
	return false
}

func TestCountPairs(t *testing.T) {
	var cases = []struct {
		n        int
		edges    [][]int
		expected int64
	}{
		{
			n:        3,
			edges:    [][]int{{0, 1}, {0, 2}, {1, 2}},
			expected: 0,
		},
		{
			n:        7,
			edges:    [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}},
			expected: 14,
		},
		{
			n:        1,
			edges:    [][]int{},
			expected: 0,
		},
		{
			n:        5,
			edges:    [][]int{{1, 0}, {3, 1}, {0, 4}, {2, 1}},
			expected: 0,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := countPairs(c.n, c.edges)
			require.Equal(t, c.expected, result)
		})
	}
}

func TestDfs(t *testing.T) {
	var cases = []struct {
		n        int
		edges    [][]int
		expected []int
	}{
		{
			n:        0,
			edges:    [][]int{{1, 0}, {3, 1}, {0, 4}, {2, 1}},
			expected: []int{0, 1, 2, 3, 4},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := dfsBeta(c.n, c.edges)
			fmt.Printf("result: %v\n", result)
			// require.EqualValues(t, c.expected, result)
		})
	}
}
