package problem59

import (
	"fmt"
	"testing"
)

func Test59(t *testing.T) {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	i, j := 0, 0
	direction := 0
	iter := 1

	for i >= 0 && i < n && j >= 0 && j < n && result[i][j] == 0 {
		result[i][j] = iter
		iter++

		if direction == 0 {
			j++
		} else if direction == 1 {
			i++
		} else if direction == 2 {
			j--
		} else if direction == 3 {
			i--
		}

		if direction == 0 && (j >= n-1 || result[i][j+1] != 0) {
			direction = 1
		} else if direction == 1 && (i >= n-1 || result[i+1][j] != 0) {
			direction = 2
		} else if direction == 2 && (j <= 0 || result[i][j-1] != 0) {
			direction = 3
		} else if direction == 3 && (i <= 0 || result[i-1][j] != 0) {
			direction = 0
		}
	}

	return result
}
