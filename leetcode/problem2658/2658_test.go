package problem2658

func findMaxFish(grid [][]int) int {
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dfs func(x, y int) (fish int)
	dfs = func(x, y int) (fish int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
			return 0
		}
		fish += grid[x][y]
		grid[x][y] = 0
		for _, direction := range directions {
			fish += dfs(x+direction[0], y+direction[1])
		}
		return fish
	}

	maxFish := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			maxFish = max(maxFish, dfs(i, j))
		}
	}

	return maxFish
}
