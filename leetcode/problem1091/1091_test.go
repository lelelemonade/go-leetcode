package problem1091

import "math"

type validNode struct {
	x      int
	y      int
	length int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n:=len(grid)
	directions:=[][]int{
		{1,1},
		{1,-1},
		{-1,1},
		{-1,-1},
		{1,0},
		{0,1},
		{-1,0},
		{0,-1},
	}
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j:=0;j<n;j++{
			dist[i][j]=math.MaxInt
		}
	}
	queue := make([]validNode, 0)
	if grid[0][0] == 1|| grid[n-1][n-1] == 1 {
		return -1
	}
	queue = append(queue, validNode{0, 0, 1})
	dist[0][0]=1
	for len(queue) != 0 {
		first:=queue[0]
		queue=queue[1:]
		for _,direction := range directions{
			x,y:=first.x+direction[0],first.y+direction[1]
			if x<0||y<0||x>=n||y>=n||grid[x][y]==1{
				continue
			}
			if dist[x][y] > first.length+1 {
				dist[x][y] = first.length + 1
				queue = append(queue, validNode{x, y, first.length + 1})
			}
		}
	}
	if dist[n-1][n-1]==math.MaxInt{
		return -1
	}else{
		return dist[n-1][n-1]
	}
}
