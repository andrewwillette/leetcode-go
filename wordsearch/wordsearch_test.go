package wordsearch

import "testing"

func exist(board [][]byte, word string) bool {
	// 1. find the first character
	// 2. find the next character
	// 3. repeat step 2 until the end of the word
	// 4. if the word is found, return true
	// 5. if the word is not found, return false
	// 6. repeat step 1 until the end of the board
	// 7. if the word is not found, return false
	// 8. if the word is found, return true
	// 1. find the first character
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 2. find the next character
			if board[i][j] == word[0] {
				// 3. repeat step 2 until the end of the word
				if dfs(board, i, j, word, 0) {
					// 4. if the word is found, return true
					return true
				}
			}
		}
	}
	// 5. if the word is not found, return false
	return false
}

func dfs(board [][]byte, i, j int, word string, k int) bool {
	// 4. if the word is found, return true
	if k == len(word) {
		return true
	}
	// 5. if the word is not found, return false
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] != word[k] {
		return false
	}
	// 6. repeat step 1 until the end of the board
	// 7. if the word is not found, return false
	// 8. if the word is found, return true
	board[i][j], k = '*', k+1
	if dfs(board, i+1, j, word, k) || dfs(board, i-1, j, word, k) || dfs(board, i, j+1, word, k) || dfs(board, i, j-1, word, k) {
		return true
	}
	board[i][j] = word[k-1]
	return false
}

// write a test for the exist function
func TestExist(t *testing.T) {
	var cases = []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			board: [][]byte{
				{'C', 'A', 'A'},
				{'A', 'A', 'A'},
				{'B', 'C', 'D'},
			},
			word:     "AAB",
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			result := exist(c.board, c.word)
			if c.expected != result {
				t.Errorf("expected: %v, got: %v", c.expected, result)
			}
		})
	}
}
