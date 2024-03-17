package problem37

import (
	"fmt"
	"testing"
)

func TestByte(t *testing.T) {
	fmt.Println(byte(4 + 48))
	fmt.Println('4')
	fmt.Println('.')
}

func Benchmark37(b *testing.B) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	for i := 0; i < b.N; i++ {
		solveSudoku(board)
	}
}

func Test37(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku(board [][]byte) {
	dfs(board)
}

func dfs(board [][]byte) (success bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}

			availableAnswer := findAvailableAnswer(board, i, j)
			if len(availableAnswer) == 0 {
				return false
			}
			for _, v := range availableAnswer {
				board[i][j] = byte(v + 48)
				if dfs(board) {
					return true
				}
				board[i][j] = '.'
			}
			return false
		}
	}

	return true
}

func findAvailableAnswer(board [][]byte, x int, y int) []int {
	answer := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}

	for i := 0; i < 9; i++ {
		delete(answer, int(board[x][i])-'0')
		delete(answer, int(board[i][y])-'0')
	}

	for i := (x / 3) * 3; i < (x/3)*3+3; i++ {
		for j := (y / 3) * 3; j < (y/3)*3+3; j++ {
			delete(answer, int(board[i][j])-'0')
		}
	}

	availableAnswer := []int{}
	for k := range answer {
		availableAnswer = append(availableAnswer, k)
	}

	return availableAnswer
}
