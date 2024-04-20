package problem79__test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test79(t *testing.T) {
	testCase1 := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	assert.True(t, exist(testCase1, "ABCCED"))

	testCase2 := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	assert.True(t, exist(testCase2, "SEE"))

	testCase3 := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	assert.False(t, exist(testCase3, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, i int, j int, idx int, word string) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if idx < 0 || idx >= len(word) {
		return false
	}

	if board[i][j] != word[idx] {
		return false
	}

	if idx == len(word)-1 && board[i][j] == word[idx] {
		return true
	}

	temp := board[i][j]
	board[i][j] = '!'

	result := dfs(board, i-1, j, idx+1, word) ||
		dfs(board, i, j-1, idx+1, word) ||
		dfs(board, i+1, j, idx+1, word) ||
		dfs(board, i, j+1, idx+1, word)

	board[i][j] = temp

	return result
}
