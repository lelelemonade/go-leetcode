package problem1034

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	vis := make([][]bool, len(grid))
	for i := range vis {
		vis[i] = make([]bool, len(grid[0]))
	}
	originalColor := grid[row][col]
	borders := [][]int{}
	isBoundary := func(r, c int) bool {
        if r == 0 || r == len(grid)-1 || c == 0 || c == len(grid[0])-1 {
            return true
        }
        if grid[r-1][c] != originalColor || grid[r+1][c] != originalColor || grid[r][c-1] != originalColor || grid[r][c+1] != originalColor {
            return true
        }
        return false
    }
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || vis[x][y] || grid[x][y] != originalColor {
			return
		}

		vis[x][y] = true

		dfs(x-1,y)
		dfs(x+1,y)
		dfs(x,y-1)
		dfs(x,y+1)

        if isBoundary(x, y) {
            borders = append(borders, []int{x, y})
        }
	}

	dfs(row,col)
	for _, border := range borders {
        grid[border[0]][border[1]] = color
    }

	return grid
}
