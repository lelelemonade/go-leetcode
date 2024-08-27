package problem1514

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
    edgeMap:=make(map[int]map[int]float64)

	for i,iv:=range edges{
		if edgeMap[iv[0]] == nil {
			edgeMap[iv[0]]=make(map[int]float64)
		}
		edgeMap[iv[0]][iv[1]]=succProb[i]
		if edgeMap[iv[1]]==nil{			
			edgeMap[iv[1]]=make(map[int]float64)
		}
		edgeMap[iv[1]][iv[0]]=succProb[i]
	}

	visited:=make([]float64,n)
	queue:=make([]direction,0)
	queue=append(queue, direction{start_node,start_node,1})
	for len(queue)!=0{
		queueLen:=len(queue)
		for i:=0;i<queueLen;i++{
			currentNode:=queue[i]
			if visited[currentNode.end]==0||visited[currentNode.end]<currentNode.prob{
				visited[currentNode.end]=currentNode.prob
				for k,v:=range edgeMap[currentNode.end]{
					if v*currentNode.prob>visited[k]{
						queue=append(queue, direction{currentNode.end,k,v*currentNode.prob})
					}
				}
			}
		}
		queue=queue[queueLen:]
	}
	return visited[end_node]
}

type direction struct{
	start int
	end int
	prob float64
}

func Test1514(t *testing.T) {
	assert.Equal(t,0.25,maxProbability(3,[][]int{{0,1},{1,2},{0,2}},[]float64{0.5,0.5,0.2},0,2))
	assert.Equal(t,0.3,maxProbability(3,[][]int{{0,1},{1,2},{0,2}},[]float64{0.5,0.5,0.3},0,2))
}