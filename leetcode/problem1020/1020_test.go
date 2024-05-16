package problem1020

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1020(t *testing.T) {
	testCase1 := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, 3, numEnclaves(testCase1))

	testCase2 := [][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, 0, numEnclaves(testCase2))
}

func numEnclaves(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		dfs(i, 0, grid)
		grid[i][0] = -1
		dfs(i, len(grid[0])-1, grid)
		grid[i][len(grid[0])-1] = -1
	}
	for j := 0; j < len(grid[0]); j++ {
		dfs(0, j, grid)
		grid[0][j] = -1
		dfs(len(grid)-1, j, grid)
		grid[len(grid)-1][j] = -1
	}

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func dfs(x int, y int, grid [][]int) {
	if x < 0 || x > len(grid)-1 ||
		y < 0 || y > len(grid[0])-1 ||
		grid[x][y] <= 0 {
		return
	}

	grid[x][y] = -1

	dfs(x-1, y, grid)
	dfs(x+1, y, grid)
	dfs(x, y-1, grid)
	dfs(x, y+1, grid)
}
