package problem1926

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func nearestExit(maze [][]byte, entrance []int) int {
	queue := make([][]int, 0)
	queue = append(queue, entrance)
	result := 0
	visited := make([][]bool, len(maze))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(maze[0]))
	}

	for len(queue) != 0 {
		lenQueue := len(queue)
		for i := 0; i < lenQueue; i++ {
			x, y := queue[i][0], queue[i][1]
			if (x == 0 || y == 0 || x == len(maze)-1 || y == len(maze[0])-1) &&
				!(x == entrance[0] && y == entrance[1]) {
				return result
			}
			if x-1 >= 0 && maze[x-1][y] == '.' && !visited[x-1][y] {
				queue = append(queue, []int{x - 1, y})
				visited[x-1][y] = true
			}
			if y-1 >= 0 && maze[x][y-1] == '.' && !visited[x][y-1] {
				queue = append(queue, []int{x, y - 1})
				visited[x][y-1] = true
			}
			if x+1 < len(maze) && maze[x+1][y] == '.' && !visited[x+1][y] {
				queue = append(queue, []int{x + 1, y})
				visited[x+1][y] = true
			}
			if y+1 < len(maze[0]) && maze[x][y+1] == '.' && !visited[x][y+1] {
				queue = append(queue, []int{x, y + 1})
				visited[x][y+1] = true
			}
			maze[x][y] = '+'
		}
		result++
		queue = queue[lenQueue:]
	}

	return -1
}

func Test1926(t *testing.T) {
	// assert.Equal(t, 1, nearestExit([][]byte{
	// 	{'+', '+', '.', '+'},
	// 	{'.', '.', '.', '+'},
	// 	{'+', '+', '+', '.'},
	// }, []int{1, 2}))
	assert.Equal(t, 2, nearestExit([][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'},
	}, []int{1, 0}))
}
