package problem130

import (
	"fmt"
	"testing"
)

func Test130(t *testing.T) {
	//testCase1 := [][]byte{
	//	{'X', 'X', 'X', 'X'},
	//	{'X', 'O', 'O', 'X'},
	//	{'X', 'X', 'O', 'X'},
	//	{'X', 'O', 'X', 'X'},
	//}
	//
	//solve(testCase1)

	testCase2 := [][]byte{
		{'O', 'O'},
		{'O', 'O'},
	}

	solve(testCase2)
}

func solve(board [][]byte) {
	var dfs func(x, y int)

	dfs = func(x int, y int) {
		if x < 0 || x > len(board)-1 ||
			y < 0 || y > len(board[x])-1 ||
			board[x][y] == 'X' ||
			board[x][y] == 'A' {
			return
		}

		if board[x][y] == 'O' {
			board[x][y] = 'A'
		}

		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	for i := 0; i < len(board); i++ {
		dfs(i, 0)
		dfs(i, len(board[i])-1)
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(0, j)
		dfs(len(board)-1, j)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	fmt.Println(board)
}
