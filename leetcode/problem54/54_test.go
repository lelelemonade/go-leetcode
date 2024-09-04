package problem54

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func spiralOrder(matrix [][]int) []int {
    result := make([]int,0)

	isNotValid := func(x,y int)bool{
		return x<0||x>=len(matrix)||y<0||y>=len(matrix[0])||matrix[x][y]<(-100)
	}

	x,y:=0,0
	direction:=Right

	for {
		if isNotValid(x,y){
			break
		}
		result = append(result, matrix[x][y])
		matrix[x][y]=-101
		switch direction{
		case Up:
			x-=1
			if isNotValid(x,y){
				direction=Right
				x+=1
				y+=1
			}
		case Right:
			y+=1
			if isNotValid(x,y){
				direction=Down
				y-=1
				x+=1
			}
		case Down:
			x+=1
			if isNotValid(x,y){
				direction=Left
				x-=1
				y-=1
			}
		case Left:
			y-=1
			if isNotValid(x,y){
				direction=Up
				y+=1
				x-=1
			}
		}
	}

	return result
}

func Test54(t *testing.T) {
	assert.Equal(t,[]int{1,2,3,6,9,8,7,4,5},spiralOrder(
		[][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		},
	))
}