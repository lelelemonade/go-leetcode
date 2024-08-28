package problem1905

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	visited:=make([][]bool,len(grid2))
	for j:=0;j<len(grid2);j++{
		visited[j]=make([]bool, len(grid2[0]))
	}

	var dfs func(i,j,mark int)

	dfs =func(i,j,mark int){
		if i<0||i>len(grid2)-1||j<0||j>len(grid2[0])-1||visited[i][j]||grid2[i][j]==0{
			return
		}
		visited[i][j]=true
		grid2[i][j]=mark
		dfs(i+1,j,mark)
		dfs(i-1,j,mark)
		dfs(i,j+1,mark)
		dfs(i,j-1,mark)
	}

	islandMark:=-1

    for i:=0;i<len(grid2);i++{
		for j:=0;j<len(grid2[0]);j++{
			if visited[i][j]||grid2[i][j]==0{
				continue
			}
			dfs(i,j,islandMark)
			islandMark--
		}
	}
	resultMap:=make(map[int]struct{})
	for i:=-1;i>islandMark;i--{
		resultMap[i]=struct{}{}
	}
	for i:=0;i<len(grid2);i++{
		for j:=0;j<len(grid2[0]);j++{
			if grid2[i][j]<0&&grid1[i][j]==0{
				delete(resultMap,grid2[i][j])
			}
		}
	}
	return len(resultMap)
}

func Test1905(t *testing.T) {
	assert.Equal(t,3 ,countSubIslands(
		[][]int{{1,1,1,0,0},{0,1,1,1,1},{0,0,0,0,0},{1,0,0,0,0},{1,1,0,1,1}},
		[][]int{{1,1,1,0,0},{0,0,1,1,1},{0,1,0,0,0},{1,0,1,1,0},{0,1,0,1,0}},
				))
}