package problem1254

func closedIsland(grid [][]int) int {
	var dfs func(x, y int) bool
	dfs = func(x, y int) (island bool) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			return false
		}
		if grid[x][y] != 0 {
			return true
		}

		grid[x][y] = -1
		ret1, ret2, ret3, ret4 := dfs(x-1, y), dfs(x+1, y), dfs(x, y-1), dfs(x, y+1)
		return ret1 && ret2 && ret3 && ret4
	}

	result := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 && dfs(i, j) {
				result++
			}
		}
	}
	return result
}
