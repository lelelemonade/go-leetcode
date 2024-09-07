package problem498

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findDiagonalOrder(mat [][]int) []int {
    x,y:=0,0
	result:=[]int{}
	upRight:=true
	for len(result) < len(mat)*len(mat[0]) {
		if upRight{
			for x>0&&y<len(mat[0])-1{
				result = append(result, mat[x][y])
				x--
				y++
			}
			result = append(result, mat[x][y])
			// turn direction
			upRight=false
			if y+1<len(mat[0]){
				y++
			}else{
				x++
			}
		}else{
			for y>0&&x<len(mat)-1{
				result = append(result, mat[x][y])
				x++
				y--
			}
			// turn direction
			upRight=true
			result = append(result, mat[x][y])
			if x+1<len(mat){
				x++
			}else{
				y++
			}
		}
		
	}
	return result
}

func Test498(t *testing.T) {
	assert.Equal(t,[]int{1,2,4,7,5,3,6,8,9},findDiagonalOrder(
		[][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		}))

	assert.Equal(t,[]int{2,3},findDiagonalOrder(
		[][]int{
			{2,3},
		}))
}
