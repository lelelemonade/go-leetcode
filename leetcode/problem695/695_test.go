package problem695

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test695(t *testing.T) {
	testCase1 := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	assert.Equal(t, 6, maxAreaOfIsland(testCase1))
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	mark := -1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			islandArea := dfs(i, j, grid, mark)
			if islandArea != 0 {
				mark--
				if maxArea < islandArea {
					maxArea = islandArea
				}
			}
		}
	}

	return maxArea
}

func dfs(x int, y int, grid [][]int, mark int) (count int) {
	if x < 0 || x > len(grid)-1 ||
		y < 0 || y > len(grid[0])-1 ||
		grid[x][y] <= 0 {
		return 0
	}

	grid[x][y] = mark

	return dfs(x-1, y, grid, mark) +
		dfs(x, y-1, grid, mark) +
		dfs(x+1, y, grid, mark) +
		dfs(x, y+1, grid, mark) +
		1
}
