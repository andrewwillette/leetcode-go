package simplifypath

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// simplifyPath
// https://leetcode.com/problems/simplify-path/
func simplifyPath(path string) string {
	simplifiedPath := make([]string, 0)
	dirs := strings.Split(path, "/")
	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		}
		if dir != ".." {
			simplifiedPath = append(simplifiedPath, dir)
		} else if len(simplifiedPath) > 0 {
			simplifiedPath = simplifiedPath[:len(simplifiedPath)-1]
		}

	}
	return "/" + strings.Join(simplifiedPath, "/")
}

func TestSimplifyPath(t *testing.T) {
	var cases = []struct {
		path     string
		expected string
	}{
		{
			path:     "/home/",
			expected: "/home",
		},
		{
			path:     "/../",
			expected: "/",
		},
		{
			path:     "/home//foo/",
			expected: "/home/foo",
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := simplifyPath(c.path)
			require.Equal(t, c.expected, result)
		})
	}
}
