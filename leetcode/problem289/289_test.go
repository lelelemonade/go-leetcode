package problem289

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func gameOfLife(board [][]int) {
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	evolve := func(x, y int) {
		liveCount := 0
		for _, direction := range directions {
			if direction[0]+x >= 0 &&
				direction[0]+x < len(board) &&
				direction[1]+y >= 0 &&
				direction[1]+y < len(board[0]) &&
				(board[direction[0]+x][direction[1]+y] ==1||board[direction[0]+x][direction[1]+y] ==-1) {
				liveCount++
			}
		}
		if board[x][y] == 1 && (liveCount < 2 || liveCount > 3) {
			board[x][y] = -1
		}

		if board[x][y] == 0 && liveCount == 3 {
			board[x][y] = 2
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			evolve(i, j)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}

func Test289(t *testing.T) {
	board:=[][]int{{0,1,0},{0,0,1},{1,1,1},{0,0,0}}
	gameOfLife(board)
	assert.Equal(t,[][]int{{0,0,0},{1,0,1},{0,1,1},{0,1,0}},board)
}