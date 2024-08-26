package problem2017

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func gridGame(grid [][]int) int64 {
    if len(grid[0])==1{
        return 0
    }
	firstTail:=0
	secondHead:=0
	for i:=1;i<len(grid[0]);i++{
		firstTail+=grid[0][i]
	}
	result :=math.MaxInt

	for i:=1;i<len(grid[0]);i++{
        result=min(result,max(secondHead,firstTail))
		secondHead+=grid[1][i-1]
		firstTail-=grid[0][i]
		result=min(result,max(secondHead,firstTail))
	}

	return int64(result)
}

func Test2017(t *testing.T) {
	assert.Equal(t,int64(4),gridGame([][]int{{2,5,4},{1,5,1}}))
	assert.Equal(t,int64(5),gridGame([][]int{{1,3,1,5},{1,3,3,1}}))
	assert.Equal(t,int64(63),gridGame([][]int{{20,3,20,17,2,12,15,17,4,15},{20,10,13,14,15,5,2,3,14,3}}))
}