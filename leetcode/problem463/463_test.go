package problem463

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test463(t *testing.T) {
	testCase1 := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}

	assert.Equal(t, 16, islandPerimeter(testCase1))
}

func islandPerimeter(grid [][]int) int {
	visited := make(map[Point]struct{})

	perimeter := 0

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x > len(grid)-1 ||
			y < 0 || y > len(grid[0])-1 ||
			grid[x][y] == 0 {
			return
		}
		point := Point{
			x: x,
			y: y,
		}

		if _, exist := visited[point]; !exist {
			visited[point] = struct{}{}
			//left
			if (x-1 >= 0 && grid[x-1][y] == 0) || x-1 < 0 {
				perimeter++
			}
			dfs(x-1, y)
			//right
			if (x+1 < len(grid) && grid[x+1][y] == 0) || x+1 >= len(grid) {
				perimeter++
			}
			dfs(x+1, y)

			if y-1 >= 0 && grid[x][y-1] == 0 || y-1 < 0 {
				perimeter++
			}
			dfs(x, y-1)

			if y+1 < len(grid[x]) && grid[x][y+1] == 0 || y+1 >= len(grid[x]) {
				perimeter++
			}
			dfs(x, y+1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				return perimeter
			}
		}
	}

	return 0
}

type Point struct {
	x int
	y int
}
