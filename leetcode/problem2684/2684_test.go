package problem2684

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxMoves(grid [][]int) int {
	memo := make([][]int, len(grid))
    for i := range memo {
        memo[i] = make([]int, len(grid[0]))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
	var dfs func(x, y, lastValue int) int
	dfs = func(x, y, lastValue int) (moves int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || lastValue >= grid[x][y] {
			return 0
		}

		if memo[x][y] != -1 {
            return memo[x][y]
        }

		moves = max(dfs(x-1, y+1, grid[x][y]), dfs(x, y+1, grid[x][y]), dfs(x+1, y+1, grid[x][y])) + 1
		memo[x][y]=moves
		return 
	}
	result := 0

	for i:=0;i<len(grid);i++{
		result=max(result,dfs(i-1,1,grid[i][0]))
		result=max(result,dfs(i,1,grid[i][0]))
		result=max(result,dfs(i+1,1,grid[i][0]))
	}

	return result
}

func Test2684(t *testing.T) {
	assert.Equal(t,3,maxMoves([][]int{
		{187,167,209,251,152,236,263,128,135},
		{267,249,251,285,73,204,70,207,74},
		{189,159,235,66,84,89,153,111,189},
		{120,81,210,7,2,231,92,128,218},
		{193,131,244,293,284,175,226,205,245},
	}))
	// assert.Equal(t,3,maxMoves([][]int{{2,4,3,5},{5,4,9,3},{3,4,2,11},{10,9,13,15}}))
}