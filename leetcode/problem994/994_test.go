package problem994

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func orangesRotting(grid [][]int) int {
	queue := make([][]int, 0)
	freshOrangeCount := 0
	for i, v := range grid {
		for j := range v {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				freshOrangeCount++
			}
		}
	}

	if freshOrangeCount == 0 {
		return 0
	}

	result := 0

	for len(queue) != 0 {
		lenQueue := len(queue)

		for i := 0; i < lenQueue; i++ {
			x, y := queue[i][0], queue[i][1]

			if x-1 >= 0 && grid[x-1][y] == 1 {
				grid[x-1][y] = 2
				freshOrangeCount--
				queue = append(queue, []int{x - 1, y})
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				grid[x][y-1] = 2
				freshOrangeCount--
				queue = append(queue, []int{x, y - 1})
			}
			if x+1 < len(grid) && grid[x+1][y] == 1 {
				grid[x+1][y] = 2
				freshOrangeCount--
				queue = append(queue, []int{x + 1, y})
			}
			if y+1 < len(grid[0]) && grid[x][y+1] == 1 {
				grid[x][y+1] = 2
				freshOrangeCount--
				queue = append(queue, []int{x, y + 1})
			}
			grid[x][y] = 0
		}

		queue = queue[lenQueue:]
		result++
	}

	if freshOrangeCount == 0 {
		return result - 1
	}

	return -1
}

func Test994(t *testing.T) {
	assert.Equal(t, 4, orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}))
}
